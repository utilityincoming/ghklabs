package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Handler serves the publish and drain API.
// Tokens authorizes Bearer GHK_API_TOKEN. Drain also accepts each extra token,
// which is how Vercel Cron's CRON_SECRET is allowed in.
type Handler struct {
	Service *Service
	Tokens  []string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case (r.URL.Path == "/" || r.URL.Path == "/healthz") && r.Method == http.MethodGet:
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.URL.Path == "/v1/endpoints" && r.Method == http.MethodPost:
		if !h.allow(w, r, false) {
			return
		}
		h.register(w, r)
	case r.URL.Path == "/v1/events" && r.Method == http.MethodPost:
		if !h.allow(w, r, false) {
			return
		}
		h.publish(w, r)
	case (r.URL.Path == "/v1/drain" || r.URL.Path == "/api" || r.URL.Path == "/api/drain") && (r.Method == http.MethodPost || r.Method == http.MethodGet):
		if !h.allow(w, r, true) {
			return
		}
		h.drain(w, r)
	case strings.HasPrefix(r.URL.Path, "/v1/deliveries/") && r.Method == http.MethodGet:
		if !h.allow(w, r, false) {
			return
		}
		h.get(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h Handler) allow(w http.ResponseWriter, r *http.Request, drain bool) bool {
	configured := false
	got, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	for i, token := range h.Tokens {
		if token == "" {
			continue
		}
		configured = true
		if !ok || got == "" {
			continue
		}
		if i > 0 && !drain {
			continue
		}
		if got == token {
			return true
		}
	}
	if !configured {
		http.Error(w, "webhook api token is not configured", http.StatusServiceUnavailable)
		return false
	}
	http.Error(w, "unauthorized", http.StatusUnauthorized)
	return false
}

func (h Handler) register(w http.ResponseWriter, r *http.Request) {
	var ep Endpoint
	if err := decodeJSON(r, &ep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.Register(r.Context(), ep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"id": ep.ID})
}

func (h Handler) publish(w http.ResponseWriter, r *http.Request) {
	var ev Event
	if err := decodeJSON(r, &ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ids, err := h.Service.Publish(r.Context(), ev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ids == nil {
		ids = []string{}
	}
	// First attempt happens in this request. Later retries wait for drain.
	result, err := h.Service.DeliverIDs(r.Context(), ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]any{
		"delivery_ids": ids,
		"drain":        result,
	})
}

func (h Handler) drain(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Drain(r.Context(), 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h Handler) get(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/deliveries/")
	dlv, ok, err := h.Service.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"id":          dlv.ID,
		"event_id":    dlv.Event.ID,
		"event_type":  dlv.Event.Type,
		"endpoint_id": dlv.EndpointID,
		"status":      dlv.Status,
		"attempts":    dlv.Attempts,
		"next_at":     dlv.NextAt,
	})
}

func decodeJSON(r *http.Request, dest any) error {
	dec := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return err
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
