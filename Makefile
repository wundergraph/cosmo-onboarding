.PHONY: build compose docker

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

start:
	@pnpm start
