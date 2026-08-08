# Diffie-Hellman

A toy client-server implementation of the Diffie-Hellman key exchange, written to pick up Go.

There's a server (`dhk-server`) and a client (`dhk-client`). The flow:

1. Client hits `GET /publicKey` on the server, which generates a fresh 256-bit prime `P` and picks `G`, and sends both back along with an id.
2. Client generates its own private value, computes its public value (`G^private mod P`), and sends it to the server via `POST /exchange`.
3. Server generates its own private value, computes its public value, derives the shared secret from the client's value, and returns its public value.
4. Client uses the server's returned value to derive the same shared secret on its end.

Both sides end up with the same `SharedSecret`, never sent over the wire.

## Running it

```bash
# terminal 1
cd dhk-server
go run .

# terminal 2
cd dhk-client
go run .
```

Server listens on `:8080` by default (`-addr` flag to change it). Client is hardcoded to talk to `http://localhost:8080`.

## Known limitations

This is a learning exercise, not something to trust with real secrets:

- `G` is hardcoded to `5` instead of being derived as an actual primitive root of `P` — verifying primitive roots properly was more number theory than I wanted to get into for this pass.
- No session management. Each exchange is one-shot; there's no key confirmation step and nothing persists beyond the single request/response.

Work in progress, mainly used as a Go sandbox.
