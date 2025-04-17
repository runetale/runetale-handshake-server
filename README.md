## runetale-handshake-server
This repository enables Runetale to establish peer-to-peer handshakes between connected nodes, serving as a core component for initiating secure P2P communication.

# enviroment variables
To configure the application, copy the `.env.example` file and rename it to `.env`:

```bash
cp .env.example .env

Run

```sh
make dev
```

for nix
```sh
make nix-build
```

### grpc
use grpcurl for debugging.

get a list of available rpcs.
⚠ When using grpcurl with TLS, remove `-plaintext`.
`grpcurl -plaintext 127.0.0.1:10000 list`

for health check, use the following.
`grpc-health-probe -addr=127.0.0.1:10000`

### dependency injection

remember to run di when you create an interactor or repository.
`make wire`
