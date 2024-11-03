-- name: GetUser :one
SELECT * FROM users where id = $1;

-- name: GetUserByLogin :one
SELECT * FROM users where login = $1;

-- name: CreateUser :one
INSERT INTO users (login, password, salt) VALUES ($1, $2, $3) returning *;