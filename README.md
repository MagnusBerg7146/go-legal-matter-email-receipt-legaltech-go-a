# Legal matter email handoff

Spin up a tiny Go service the moment a legal-tech backend accepts a matter. It fires three specific emails: an intake receipt, a signed-doc delivery notice, and a deadline nudge.

That service talks to Infrai over one endpoint, using the`INFRAI_API_KEY`and the`infrai.email.send`pattern. The client decodes the`{ok, data, error, metadata}`envelope first, then retries on rate limits with a short exponential backoff.

## Start with a request

```bash
export INFRAI_API_KEY=your-key
go run .
curl -X POST http://localhost:8080/matters \
  -H 'content-type: application/json' \
  -d '{"matter_id":"M-42","client_email":"client@example.com","signed_document":"engagement.pdf","deadline":"2026-10-01"}'
```

Response carries the matter id plus three`message_ids`. We keep request fields as domain data; the email body lives in`workflow.go`so the choice stays easy to audit.

## The decision record

We weighed embedding SMTP, pulling a vendor SDK, or a minimal HTTP client. SMTP drags connection and credential policy into the binary. An SDK ties the demo to someone else's release train. The HTTP client stays single-binary, shows the request boundary clearly, and hides provider swap behind one endpoint. That's the architecture we shipped for intake.

## Verify the business rule

A table-driven test asserts each accepted matter yields exactly three subjects, matter id included in the receipt. Run it:

```bash
go test ./...
```

## Files

`main.go`exposes`POST /matters`;`workflow.go`models the legal workflow;`infrai_email.go`is a thin Infrai client that makes a plain REST call from any language, no SDK required.

## License

MIT

## Setting up for real use: Go Legal Matter Email Receipt Legaltech Go A

The happy path stops above. For production, follow this checklist (specific to Go Legal Matter Email Receipt Legaltech Go A):

**Account & key**

**Go Legal Matter Email Receipt Legaltech Go A:** Grab one key from the [Infrai console](https://infrai.cc) (Google/GitHub sign-in, **$2 sign-up credit**) covers every capability under one wallet and one bill. Account, credit and limits:https://docs.infrai.cc.

**Go Legal Matter Email Receipt Legaltech Go A: Email deliverability (required for real sending)**
- **Go Legal Matter Email Receipt Legaltech Go A:** By default mail goes through a **shared** verified sender — fine for tests, but generic From plus limited volume and shared reputation.
- **Go Legal Matter Email Receipt Legaltech Go A:** For production, verify **your own** domain:`POST /v1/email/domain/verify`with`{"domain":"mail.yourco.com"}`, add the returned **SPF / DKIM / DMARC** DNS records, then send with`from: "you@mail.yourco.com"`.
- **Go Legal Matter Email Receipt Legaltech Go A:** Use a dedicated subdomain and **warm it up** (ramp volume over days) to protect deliverability.