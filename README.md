# api-shortener-url

[![codecov](https://codecov.io/github/challenge-mercadolibre-cl/api-shortener-url/branch/main/graph/badge.svg?token=KZFA9Z94CG)](https://codecov.io/github/challenge-mercadolibre-cl/api-shortener-url)
[![Go](https://github.com/challenge-mercadolibre-cl/api-shortener-url/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/challenge-mercadolibre-cl/api-shortener-url/actions/workflows/go.yml)

> Service to shortener url.

# Requirements

- [Docker-compose](https://docs.docker.com/compose/)
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/)

## Usage on Development

1. Copy .env.example to .env
2. Execute command.
   ```sh
   go run app/main.go
   ```
3. Happy Coding.

## Usage with docker-compose

1. Copy .env.example to .env
2. Execute command.
   ```sh
   docker-compose up -d
   ```
3. Happy Coding.

## Default port services

- **Redis UI:** 9000
- **API:** 8080
