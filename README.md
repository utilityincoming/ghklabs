# ghklabs

Webhook delivery for outbound event notifications.

`webhook.Dispatcher` registers subscriber endpoints, publishes events, and POSTs a JSON body with an HMAC-SHA256 signature. Failed attempts stay pending and become due again after backoff. A delivery is marked failed after the attempt limit.

Signature header: `X-GHK-Signature: t=<unix>,v1=<hex>`, where `v1` is HMAC-SHA256 of `<unix>.<body>` using the endpoint secret. Receivers check it with `webhook.Verify`.

```go
d := webhook.NewDispatcher(webhook.Options{})
_ = d.Register(webhook.Endpoint{
    ID: "billing", URL: "https://example.com/hooks", Secret: os.Getenv("HOOK_SECRET"),
    Types: []string{"invoice.paid"},
})
ids, _ := d.Publish(webhook.Event{ID: "evt_1", Type: "invoice.paid", Data: json.RawMessage(`{"id":"in_1"}`)})
_ = d.DeliverOnce(ctx, ids[0])
```
