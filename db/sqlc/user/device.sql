-- name: CreateDevice :one
INSERT INTO devices (
    operating_system,
    mac_address,
    client_ip,
    user_agent,
    is_banned
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetDevice :one
SELECT * FROM devices
WHERE id = $1 LIMIT 1;

-- name: UpdateDevice :exec
UPDATE devices
SET is_banned = COALESCE(sqlc.narg(is_banned), is_banned)
WHERE id = $1;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE id = $1;