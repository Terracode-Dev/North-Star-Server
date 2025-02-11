mig-up:
	goose up

mig-down:
	goose down

mig-reset:
	goose reset

mig-status:
	goose status

create-migrate:
	goose create $(n) sql

sql-gen:
	sqlc generate

run-server:
	go run cmd/main.go
