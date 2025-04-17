ENV_FILE := .env

.PHONY: kustomize-dev wire update-submodule get-local-env get-dev-env get-prd-env test build-docker build-docker-compose up down

# wire dependency injection
wire:
	./bin/wire ./infura/di/wire.go

update-submodule:
	go get github.com/runetale/client-go 

test: dev
	go test -v ./...

build-docker:
	docker build --no-cache -t runetale-handshake-server -f Dockerfile .

build-docker-compose:
	docker-compose -f docker-compose.yml build

dev: build-docker-compose
	docker-compose up

down:
	docker-compose rm -fsv runetale-handshake-server 

nix-build:
	nix build .#runetale-handshake-server --no-link --print-out-paths --print-build-logs --extra-experimental-features flakes
  