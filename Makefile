include .env
export $(shell sed 's/=.*//' .env)

dev: ## Run gin web server with air hot reload
	air

create-migration: ## Create new migration file. Usage: `make create-migration file=add_column_time`
	migrate create -ext sql -dir db/migrations -seq $(file)

up-migration: ## Run migration. Usage: `make up-migration`
	migrate -path db/migrations -database $(DATABASE_URL) up ${ver}

down-migration: ## Rollback migration
	migrate -path db/migrations -database $(DATABASE_URL) down ${count}

status-migration: ## Check migration status
	migrate -path db/migrations -database $(DATABASE_URL) version

force-migration: ## Fix dirty migration, specify the version
	migrate -path db/migrations -database $(DATABASE_URL) force $(ver)
