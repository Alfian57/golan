include .env
export

DATABASE_URL = "mysql://$$DB_USERNAME:$$DB_PASSWORD@tcp($$DB_HOST:$$DB_PORT)/$$DB_NAME?query"

run:
	air

migrate-create:
	@read -p "Enter migration name (use underscore): " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-up:
	migrate -database $(DATABASE_URL) -path migrations up
	
migrate-down:
	migrate -database $(DATABASE_URL) -path migrations down

migrate-force:
	@read -p "Enter migration version: " version; \
	migrate -database $(DATABASE_URL) -path migrations force $$version