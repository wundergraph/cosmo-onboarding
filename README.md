<p align="center">
<img width="350" src="./docs/assets/logo.png"/>
</p>

<div align="center">
<h5>WunderGraph Cosmo - The GraphQL Federation Platform</h5>
<h6><i>Reach for the stars, ignite your cosmo!</i></h6>
</div>

<p align="center">
  <a href="https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding"><strong>Quickstart</strong></a> ·
  <a href="https://github.com/wundergraph/cosmo/tree/main/examples"><strong>Examples</strong></a> ·
  <a href="https://cosmo-docs.wundergraph.com"><strong>Docs</strong></a> ·
  <a href="https://cosmo-docs.wundergraph.com/cli"><strong>CLI</strong></a> ·
  <a href="https://wundergraph.com/discord"><strong>Community</strong></a> ·
  <a href="https://github.com/wundergraph/cosmo/releases"><strong>Changelog</strong></a> ·
  <a href="https://wundergraph.com/jobs"><strong>Hiring</strong></a>
</p>

<p align="center">
  <a href="https://human-oss.dev"><img src="https://human-oss.dev/badge.svg" alt="Open Source AI Manifesto"></a>
</p>

## Overview

WunderGraph Cosmo is a comprehensive Lifecycle API Management platform tailored for Federated GraphQL. It encompasses everything from Schema Registry, composition checks, and analytics, to metrics, tracing, and routing. Whether you’re looking to deploy 100% on-prem or prefer a [Managed Service](https://cosmo.wundergraph.com/login), Cosmo offers flexibility without vendor lock-in, all under the Apache 2.0 license.

## Onboarding

This repository hosts a configuration for [Cosmo Router](https://github.com/wundergraph/cosmo/tree/main/router) and sample applications, which provide a way to get yourself familiar with GraphQL federation and the Cosmo Cloud platform.

## Quickstart

By default, the router runs on port `3002`.

1. `docker run --rm -p 3002:3002 cosmo-demo`
2. Execute example query via `cURL`:

```shell
curl -s -X POST http://localhost:3002/graphql -H 'Content-Type: application/json' -d '{"query":"query GetProductWithReviews($id: ID\u0021) { product(id: $id) { id title price { currency amount } reviews { id author rating contents } } }","variables":{"id":"1"}}'
```

## Development

Make sure you have [make](https://www.gnu.org/software/make/) and [pnpm](https://pnpm.io/) installed.

1. Run `pnpm install` to pull down the dependencies
2. Run `make start` to build and run the router image.

Some other make tasks:

* `make build` - build plugins and router configuration
* `make compose` - generate router execution config
* `make docker` - build a docker image
