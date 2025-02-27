run dev: ## Run gin web server with air hot reload
	air

create-migration: ## Create new migration file. It takes parameter `file` as filename. Usage: `make create-migration file=add_column_time`
	migrate create -ext sql -dir db/migrations -seq $(file)

up-migration: ## Run migration. It takes parameter `dsn` as database string connection. Usage: `make up-migration dsn="postgres://postgres:postgres@localhost:5432/hrisdb?sslmode=disable"`
	migrate -path db/migrations -database $(dsn) up ${ver}

down-migration: ## Rollback migration. Takes the parameter ver to drop how many times as specified.
	migrate -path db/migrations -database $(dsn) down ${count}

status-migration: ## Check migration status
	migrate -path db/migrations -database $(dsn) version

force-migration: ## Fix dirty migration, specify the version
	migrate -path db/migrations -database $(dsn) force $(ver)
