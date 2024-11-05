-- name: CreateMessage :one
INSERT INTO messages (message, sender_id, conversation_id) VALUES ($1, $2, $3) returning *;