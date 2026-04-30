#!/usr/bin/env sh

set -o xtrace

fly deploy -a cosmo-onboarding-router -c cosmo-onboarding-router.fly.toml --image ghcr.io/wundergraph/cosmo/router:latest
