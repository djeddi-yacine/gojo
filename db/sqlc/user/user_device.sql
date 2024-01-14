-- name: CreateUserDevice :exec
INSERT INTO user_devices (
    user_id,
    device_id
) VALUES (
  $1, $2
)
ON CONFLICT (user_id,device_id) 
DO NOTHING;

-- name: ListUserDevices :many
SELECT device_id FROM user_devices
WHERE user_id = $1
ORDER BY id;