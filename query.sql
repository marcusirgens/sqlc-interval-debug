CREATE TABLE foo (bar bool not null, "interval" interval not null);

-- name: Get :many
SELECT bar, "interval" FROM foo LIMIT $1;

-- name: Insert :exec
INSERT INTO foo (bar, interval) VALUES ($1, $2);
