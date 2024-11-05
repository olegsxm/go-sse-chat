-- name: CreateParticipant :exec
INSERT INTO participants (conversation_id, user_id) VALUES ($1, $2);
