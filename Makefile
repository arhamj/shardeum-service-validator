.PHONY: build
build:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o service_validator ./cmd

.PHONY: scp
scp:
		scp ./service_validator local-data-rpc:/home/shardeumcoredev/rpc_local_data/go-service-validator
		scp ./scripts/start_script.sh local-data-rpc:/home/shardeumcoredev/rpc_local_data/go-service-validator
		scp ./config/config.json local-data-rpc:/home/shardeumcoredev/rpc_local_data/go-service-validator