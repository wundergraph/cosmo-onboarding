.PHONY: build compose docker format lint

make: build

# Runs build recursively
build:
	@pnpm build

# Generate router execution config (required before building Docker image or running the router)
# Usage: make compose
compose:
	@pnpm compose

# Build linux plugin binary + Docker image (run `make compose` first if config.json is missing)
# Usage: make docker
docker:
	@$(MAKE) -C plugins/products build-linux
	@$(MAKE) -C plugins/reviews build-linux
	docker build -t cosmo-demo .

# Build linux plugin binary + Docker image (run `make compose` first if config.json is missing)
# for local development
# Usage: make docker-local
docker-local:
	@$(MAKE) -C plugins/products build-linux
	@$(MAKE) -C plugins/reviews build-linux
	docker build -f Dockerfile.local -t cosmo-demo-local .

# Test all plugins
test:
	@pnpm test

start:
	@pnpm start

# Format Go source in all plugins
format:
	@$(MAKE) -C plugins/products format
	@$(MAKE) -C plugins/reviews format

# Lint Go source in all plugins
lint:
	@$(MAKE) -C plugins/products lint
	@$(MAKE) -C plugins/reviews lint
