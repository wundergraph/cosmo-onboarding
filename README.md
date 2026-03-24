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

1. `docker run --rm -p 3002:3002 <TBD>` 🚧 official image pending
2. Execute example query via `cURL`:

```shell
curl -s -X POST http://localhost:3002/graphql -H 'Content-Type: application/json' -d '{"query":"query GetProductWithReviews($id: ID\u0021) { product(id: $id) { id title price { currency amount } reviews { id author rating contents } } }","variables":{"id":"1"}}'
```

## Development

Make sure you have [make](https://www.gnu.org/software/make/) and [pnpm](https://pnpm.io/) installed.

This project uses `pnpm@10.29.3` (defined in `packageManager`). The easiest way to get the right version is via [Corepack](https://nodejs.org/api/corepack.html), which ships with Node.js:

```shell
corepack enable
corepack install
```

> [!NOTE]
> The image is built with name `cosmo-demo-local`, it is intended to run against local instance of Cosmo platform.
> See more information on [how to develop Cosmo platform locally](https://github.com/wundergraph/cosmo/blob/main/CONTRIBUTING.md#local-development)

1. Run `pnpm install` to pull down the dependencies
2. Run `make start` to build and run the router image.
3. Visit [http://localhost:3002](http://localhost:3002) to interact with the playground
4. (optional) Execute example query via `cURL`:

```shell
curl -s -X POST http://localhost:3002/graphql -H 'Content-Type: application/json' -d '{"query":"query GetProductWithReviews($id: ID\u0021) { product(id: $id) { id title price { currency amount } reviews { id author rating contents } } }","variables":{"id":"1"}}'
```

### Local development with Cosmo

1. Build the image `make docker-local`
2. Have Cosmo platform running locally by [following the contribution documentation](https://github.com/wundergraph/cosmo/blob/main/CONTRIBUTING.md#local-development)
3. Create federated graph ([documentation](https://cosmo-docs.wundergraph.com/cli/federated-graph/create)): `pnpm wgc federated-graph create --routing-url http://localhost:3002/graphql onboarding`
4. Create subgraphs ([documentation](https://cosmo-docs.wundergraph.com/cli/subgraph/create)):

```shell
pnpm wgc subgraph create --routing-url http://localhost:3002/graphql products
pnpm wgc subgraph create --routing-url http://localhost:3002/graphql reviews
```

5. Publish the subgraphs ([documentation](https://cosmo-docs.wundergraph.com/cli/subgraph/publish)):

```shell
pnpm wgc subgraph publish products --schema <path-to-cloned-repository>/plugins/products/src/schema.graphql
pnpm wgc subgraph publish reviews --schema <path-to-cloned-repository>/plugins/reviews/src/schema.graphql
```

6. Run this image alongside local instance(s) of Cosmo services with provided `<token>` ([documentation](https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding#create-a-router-token)):

```shell
docker run --rm -p 3002:3002 -e GRAPH_API_TOKEN=<token> -e LOG_LEVEL=info cosmo-demo
```

Some other make tasks:

* `make build` - build plugins and router configuration
* `make compose` - generate router execution config
* `make docker-local` - build a development version of the docker image
