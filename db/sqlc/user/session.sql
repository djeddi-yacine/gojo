-- name: CreateSession :one
INSERT INTO sessions (
  id,
  username,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: UpdateSession :one
UPDATE sessions
SET is_blocked = $2
WHERE username = $1
RETURNING username;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1 ;

-- name: RefreshSessions :exec
DELETE FROM sessions AS s1
WHERE s1.username = $1
AND s1.is_blocked = true
AND (s1.expires_at < NOW()
     OR s1.expires_at != (SELECT MAX(expires_at) FROM sessions AS s2
                         WHERE s2.username = $1 AND s2.is_blocked = true)
);
