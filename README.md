# 金庫 - Kinko Vault

A minimal CSV based database for login information and secrets with zero external dependencies. Meant to be deployed using Docker Compose and accessed over a secure VPN.

## What is the motivation behind this project?

According to my belief security stems from not messing up what we could. So much brilliance went into creating secure VPN protocols, that it would be a shame to not rely upon them. This small web app is a containerized server side rendered database which functions without Javascript code on the client side for the sake of simplicity.

## How to start it?

Customize the included compose.yaml file:

```yaml
services:
  kinko:
    image: ghcr.io/kerezsiz42/kinko:latest
    environment:
      TOKEN: "generate a long string for a token"
      DATABASE_TYPE: "csv"
      CSV_DATABASE_PATH: "/path/to/database.csv"
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
    ports:
      - 8080:8080
```

## How to recover from disasters?

## How to deploy it?

## How to keep it secure?

## How to look at logs?

```sh
sudo less /var/lib/docker/containers/${CONTAINER_ID}/${CONTAINER_ID}-json.log
```

## TODO

- implement csv database functionality
- implement import/export functionality
- create pipeline
- use golangci-lint
- use grype scan
- write unit tests for database layer
- write e2e tests
