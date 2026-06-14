# mock-transaction

Mock UPI payment upstream service for local development of the Shield fraud detection system.

Listens on **port 9001** and responds to `POST /upi/pay` with a hardcoded success payload. Logs the `X-Risk-Score` header set by the Kong shield-check plugin on every allowed request.

## Run

```bash
go run .
# [MOCK-TXN] Starting on :9001
```

## Response

```json
{ "status": "payment_accepted", "transactionId": "TXN-MOCK-001" }
```

## Part of

See [shield](https://github.com/krishnaMohan2501/shield) for full local setup instructions.
