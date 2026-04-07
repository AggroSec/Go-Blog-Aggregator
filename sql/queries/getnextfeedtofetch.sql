-- name: GetNextFeedToFetch :one
select *
from feeds
order by last_fetched_at nulls first, created_at asc
limit 1;