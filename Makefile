run:
	ENV=prod go run main.go

dev:
	ENV=dev go run main.go

test:
	go test -v ./...
	go run cmd/test-main.go

create-migrate:
	migrate create -ext sql -dir ./db/migrations $(name)

prod-db-migrate:
	ENV=prod go run cmd/migrations/main.go

dev-db-migrate:
	ENV=dev go run cmd/migrations/main.go

dev-db-rollback-all:
	ENV=dev go run cmd/dev_rollback_all/main.go