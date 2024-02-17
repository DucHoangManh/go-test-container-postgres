DB_DSN = denishoang:pgpassword@localhost:5432/products?sslmode=disable
.PHONY: migrations/new
migrations/new:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: migrations/up
migrations/up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./migrations -database="postgres://${DB_DSN}" up

.PHONY: migrations/down
migrations/down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./migrations -database="postgres://${DB_DSN}" down
