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

## Usage for testing without coverage.

1. Execute command.
   ```sh
   go test ./..
   ```

## Usage for testing with coverage.

1. Execute command.
   ```sh
   go test ./... -cover
   ```

## Default port services

- **Redis UI:** 9000
- **API:** 8080

## Diagrams

- [Diagram Flow](./docs/Diagram%20Flow.jpg)
- [Architecture Sync](./docs/Architecture%20Synchronous.jpg)
- [Diagram Flow Async](./docs/Flow%20Async.jpg)
- [Architecture Disaster Recovery (ADR)](<./docs/Architecture%20Disaster%20Recovery%20(ADR).jpg>)
- [Architecture for stats](./docs/Architecture%20for%20Stats.jpg)

## Explain

- [tecnologias.md](./docs/tecnologias.md)

## Evidence

- [Testing Stress](./docs/Prueba%20Estress.jpg)
- RPS: 3090
- RPM Projection: 185400 (TPS x 60)

- [Infrastructure K8s](./docs/Evidence%20Infrastructure.png)
- HPA (Horizontal Pod Autoscaler)
  -> Min: 40 pods
  -> Max: 300 pods
  -> Metrics for HPA
  --> CPU: 80%
  --> RAM: 80%.
