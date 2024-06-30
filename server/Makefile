dev: build
	@docker-compose up -d db
	@docker-compose up server

build:
	@go build -o bin/ecommerce_go cmd/main.go

test:
	@go test -v ./...

clean:
	@docker-compose down
	@rm -f bin/ecommerce_go

run: build
	@./bin/ecommerce_go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

