-- name: CreateUser :one
INSERT INTO users (user_name, email, password_hash, bio)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetUserById :one
SELECT
    id,
    user_name,
    password_hash,
    email,
    bio,
    create_at,
    update_at
FROM users
WHERE id = $1;