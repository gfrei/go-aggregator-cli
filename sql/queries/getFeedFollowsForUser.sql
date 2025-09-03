-- name: GetFeedFollowsForUser :many
-- It should return all the feed follows for a given user, 
-- and include the names of the feeds and user in the result.

SELECT feed_follows.*, users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
INNER JOIN users
ON feed_follows.user_id = users.id
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;