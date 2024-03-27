.PHONY: run
run:
	go run cmd/main.go

.PHONY: start_docker
start_docker:
	docker run --name=todo-db -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p 5436:5432 -d --rm  -v todo-db-data:/var/lib/postgresql/data postgres

.PHONY: migrate_up
migrate_up:
	migrate -path ./schema -database 'postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5436/postgres?sslmode=disable' up

.PHONY: migrate_down
migrate_down:
	migrate -path ./schema -database 'postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5436/postgres?sslmode=disable' down

.PHONY: migrate_create
migrate_create:
	migrate create -ext sql -dir ./schema -seq init