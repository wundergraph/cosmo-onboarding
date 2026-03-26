.PHONY: build compose docker format lint

make: build

# Runs build recursively
build:
	@pnpm build

# Runs build for CI (serial mode)
build-ci:
	@pnpm build:ci

# Generate router execution config (required before building Docker image or running the router)
# Usage: make compose
compose:
	@pnpm compose

# Build linux plugin binary + Docker image (run `make compose` first if config.json is missing)
# for local development
# Usage: make docker-local
docker-local:
	@pnpm -r build-linux
	docker build -f Dockerfile.local -t cosmo-demo-local .

# Generates protobuf and client code
generate:
	@pnpm generate

# Test all plugins
test:
	@pnpm test

start:
	@pnpm start

# Format Go source in all plugins
format:
	@pnpm -r format

# Lint Go source in all plugins
lint:
	@pnpm -r lint
