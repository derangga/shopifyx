# What service is this?
Shopifyx service.

## Development Normal
- Isi env sesuai dengan yang di env.example
- `go mod download` kalau belum.
- `go get -u -d github.com/golang-migrate/migrate/cmd/migrate` untuk install golang migrate
- Jalankan. `go run main.go`.

## Create migration
- `migrate create -ext sql -dir db/migrations -seq add_user_table`
- `migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up`
- `migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations down`

## Specs
- [https://jsonapi.org/](https://jsonapi.org/)
- Auth Bearer token

## Tech Stack
- Golang
- Gin
- PostgreSQL
- JWT

## Dependency