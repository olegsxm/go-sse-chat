migrate_new:
	goose --dir=./db/migrations create $(name) sql

migrate_up:
	goose postgres "postgres://puser:ppassword@localhost:5432/db" --dir=./db/migrations  up && sqlc generate

sqlc:
	sqlc generate