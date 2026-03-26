# Reviews Plugin - Cosmo gRPC Service Example

This repository contains a simple Cosmo gRPC service plugin that showcases how to design APIs with GraphQL Federation but implement them using gRPC methods instead of traditional resolvers.

## What is this demo about?

This demo illustrates a key pattern in Cosmo gRPC service development:
- **Design with GraphQL**: Define your API using GraphQL schema
- **Implement with gRPC**: Instead of writing GraphQL resolvers, implement gRPC service methods
- **Bridge the gap**: The Cosmo router connects GraphQL operations to your gRPC implementations
- **Test-Driven Development**: Test your gRPC service implementation with gRPC client and server without external dependencies

The `reviews` subgraph is one of two subgraphs that make up the federated graph. It extends the `Product` entity with a `reviews` field and exposes a `review(id: ID!)` query. The `products` subgraph owns the `Product` entity; this subgraph joins via the `id` key to attach review data to each product.

The plugin demonstrates:
- How GraphQL types and operations map to gRPC service methods
- Proper structure for a Cosmo gRPC service plugin
- How to test your gRPC service implementation with gRPC client and server without external dependencies
- `QueryReview` — resolves a single review by ID
- `LookupProductById` — entity lookup used by the router to attach reviews to products

## Getting Started

Plugin structure:

   ```
    plugins/reviews/
    ├── go.mod                # Go module file with dependencies
    ├── go.sum                # Go checksums file
    ├── src/
    │   ├── main.go           # Main plugin implementation
    │   ├── main_test.go      # Tests for the plugin
    │   └── schema.graphql    # GraphQL schema defining the API
    ├── generated/            # Generated code (created during build)
    └── bin/                  # Compiled binaries (created during build)
        └── plugin            # The compiled plugin binary
   ```

## 🔧 Customizing Your Plugin

- Change the GraphQL schema in `src/schema.graphql` and regenerate the code with `make generate`.
- Implement the changes in `src/main.go` and test your implementation with `make test`.
- Build the plugin with `make build`.

## 📚 Learn More

For more information about Cosmo and building router plugins:
- [Cosmo Documentation](https://cosmo-docs.wundergraph.com/)
- [Cosmo Router Plugins Guide](https://cosmo-docs.wundergraph.com/connect/plugins)

---

<p align="center">Made with ❤️ by <a href="https://wundergraph.com">WunderGraph</a></p>