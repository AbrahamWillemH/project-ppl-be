# Gin Web Server with PostgreSQL

This is a simple **Go (Gin) web server** with PostgreSQL database integration. The project uses **Air** for live reloading, **pgx** for database connection, **dotenv** for environment variables, and **golang-migrate** for handling database migrations.

## 🚀 Getting Started

### 1️⃣ Install Dependencies
Make sure you have **Go** installed. Then, install the required dependencies:

```sh
# Install Air for live reloading
go install github.com/cosmtrek/air@latest

# Install PostgreSQL driver (pgx)
go get github.com/jackc/pgx/v5

# Install dotenv to load environment variables
go get github.com/joho/godotenv

# Install golang-migrate for database migrations
brew install golang-migrate  # macOS (Homebrew)
sudo apt install golang-migrate -y  # Linux
# For Windows, download from: https://github.com/golang-migrate/migrate/releases
```

### 2️⃣ Setup Environment Variables
Create a `.env` file in the project root:

```ini
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

### 3️⃣ Run Database Migrations
Use **golang-migrate** to manage database schema:

```sh
# Create a new migration file (example: create_users_table)
migrate create -ext sql -dir migrations -seq create_users_table

# Apply migrations
migrate -path migrations -database "$DATABASE_URL" up

# Rollback last migration
migrate -path migrations -database "$DATABASE_URL" down 1
```

### 4️⃣ Run the Server with Air (Live Reload)
Start the server with **Air** for auto-reloading:

```sh
air
```

If you're not using **Air**, you can run the server normally:

```sh
go run main.go
```

## 📌 Notes
- Ensure PostgreSQL is running before starting the server.
- Modify the `DATABASE_URL` in `.env` to match your database credentials.

Happy coding! 🚀
