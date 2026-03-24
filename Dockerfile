FROM ghcr.io/wundergraph/cosmo/router:latest

ENV DEV_MODE=true
ENV LISTEN_ADDR=0.0.0.0:3002

# Copy the router configuration
COPY config.yaml /app/config.yaml
COPY config.json /app/config.json

# Copy the plugins to the container
COPY ./plugins /app/plugins

# Set the working directory
WORKDIR /app

# The entrypoint is already set in the base image
