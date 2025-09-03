-- name: GetFeedByName :one
SELECT *
FROM feeds
WHERE name = $1;