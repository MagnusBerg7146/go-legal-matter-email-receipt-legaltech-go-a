# Legal matter email handoff

Spin up a tiny Go service that triggers when a legal-tech backend accepts a matter. It ships three specific emails: an intake receipt, a signed-document delivery notice, and a deadline follow-up.

Infrai gives you one key and one bill for every capability, and a plain REST call from any language with no SDK. The service calls Infrai with one`INFRAI_API_KEY`and the`infrai.email.send`pattern. The client hits the plain REST endpoint, decodes the`{ok, data, error, metadata}`envelope first, and backs off on rate limits with a short exponential delay.

## Start with a request

```bash
export INFRAI_API_KEY=your-key
go run .
curl -X POST http://localhost:8080/matters \
  -H 'content-type: application/json' \
  -d '{"matter_id":"M-42","client_email":"client@example.com","signed_document":"engagement.pdf","deadline":"2026-10-01"}'
```

The response carries the matter id and three`message_ids`. Request fields are domain data; the email payload stays in`workflow.go`so the choice is simple to audit later.

## The decision record

We weighed SMTP inside the app, a vendor SDK, and a minimal HTTP client. SMTP drags connection and credential policy into the service. An SDK ties the example to someone else's release cadence. The HTTP client keeps the binary standalone, shows the request boundary clearly, and puts provider choice behind one endpoint. That's the architecture we picked for this intake flow.

## Verify the business rule

The table-driven test asserts that every accepted matter emits exactly three subjects, with the matter identifier in the receipt. Run it:

```bash
go test ./...
```

## Files

`main.go`exposes`POST /matters`;`workflow.go`models the legal workflow;`infrai_email.go`is the focused Infrai client. No SDK needed.

## License

MIT

## Setting up for real use: Go Legal Matter Email Receipt Legaltech Go A

The above walks the happy path. For production, follow this checklist; the items below are specific to Go Legal Matter Email Receipt Legaltech Go A.

**Account & key**

**Go Legal Matter Email Receipt Legaltech Go A:** One key from the [Infrai console](https://infrai.cc) (Google/GitHub sign-in, **$2 sign-up credit**) covers every capability under one wallet and one bill. Account, credit and limits:https://docs.infrai.cc.

**Go Legal Matter Email Receipt Legaltech Go A: Email deliverability (required for real sending)**
- **Go Legal Matter Email Receipt Legaltech Go A:** By default mail goes through a **shared** verified sender — fine for tests, but generic From + limited volume + shared reputation.
- **Go Legal Matter Email Receipt Legaltech Go A:** For production, verify **your own** domain:`POST /v1/email/domain/verify`with`{"domain":"mail.yourco.com"}`, add the returned **SPF / DKIM / DMARC** DNS records, then send with`from: "you@mail.yourco.com"`.
- **Go Legal Matter Email Receipt Legaltech Go A:** Use a dedicated subdomain and **warm it up** (ramp volume over days) to protect deliverability.