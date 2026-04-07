-- name: FeedLookup :one
select *
from feeds
where url = $1;