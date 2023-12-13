-- name: CreateAnimeMovieLogoImage :one
INSERT INTO anime_movie_logo_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieLogoImage :one
SELECT * FROM anime_movie_logo_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeMovieLogoImageByAnimeID :one
SELECT * FROM anime_movie_logo_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeMovieLogoImages :many
SELECT image_id
FROM anime_movie_logo_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMovieLogoImage :exec
DELETE FROM anime_movie_logo_images
WHERE anime_id = $1 AND image_id = $2;