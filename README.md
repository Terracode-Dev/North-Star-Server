# HR Management System

A Go-based HR Management System with a clean architecture approach, using MariaDB as the database.

## Project Structure

```
.
├── cmd
│   └── main.go           # Application entry point
├── internal
│   ├── config           # Configuration management
│   ├── database         # Database operations and models
│   ├── middleware       # HTTP middleware
│   ├── modules          # Business logic modules
│   ├── pkg             # Shared packages
│   └── server          # HTTP server setup
├── migration           # Database migrations
└── sqlc.yaml          # SQLC configuration
```

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Goose (for migrations)
- SQLC (for SQL generation)

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd <project-directory>
```

2. Start the database:
```bash
docker-compose up -d
```

3. Run the migrations:
```bash
make mig-up
```

4. Start the server:
```bash
make run-server
```

## Database Management

The project uses Goose for database migrations and SQLC for type-safe SQL queries.

### Available Make Commands

- `make mig-up` - Apply all pending migrations
- `make mig-down` - Roll back the last migration
- `make mig-reset` - Reset all migrations
- `make mig-status` - Check migration status
- `make create-migrate n=migration_name` - Create a new migration
- `make sql-gen` - Generate SQLC code
- `make run-server` - Start the application

### Creating New Migrations

To create a new migration:
```bash
make create-migrate n=add_new_table
```

### Database Configuration

The MariaDB instance is configured with the following default settings:
- Port: 3306
- Database: hr_management
- User: hrapp
- Password: hrpassword

You can modify these settings in the `docker-compose.yml` file.

## Development

1. Generate SQL code after modifying queries:
```bash
make sql-gen
```

2. Apply new migrations:
```bash
make mig-up
```

3. Start the development server:
```bash
make run-server
```

## Project Layout

- `cmd/`: Contains the main application entry points
- `internal/`: Application-specific code
  - `config/`: Configuration management
  - `database/`: Database operations and models
  - `middleware/`: HTTP middleware components
  - `modules/`: Business logic modules (HR, etc.)
  - `pkg/`: Shared utilities and helpers
  - `server/`: HTTP server setup and routing
- `migration/`: Database migration files
- `Makefile`: Build and development commands
- `sqlc.yaml`: SQLC configuration

## Contributing

1. Create a new branch for your feature
2. Make your changes
3. Submit a pull request

## License

[Add your license information here]
