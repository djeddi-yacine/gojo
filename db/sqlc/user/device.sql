-- name: CreateDevice :one
INSERT INTO devices (
    id,
    device_name,
    device_hash,
    operating_system,
    mac_address,
    client_ip,
    user_agent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (device_hash)
DO UPDATE SET user_agent = excluded.user_agent
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