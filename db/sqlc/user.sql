-- name: CreateUser :one
INSERT INTO users (
          username, 
          email, 
          hashed_password, 
          full_name
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($2, hashed_password),
  password_changed_at = COALESCE($3, password_changed_at),
  full_name = COALESCE($4, full_name),
  email = COALESCE($5, email),
  is_email_verified = COALESCE($6, is_email_verified)
WHERE
  username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;