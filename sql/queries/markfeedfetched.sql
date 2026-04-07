-- name: MarkFeedFetched :exec
update feeds
set 
    last_fetched_at = now(),
    updated_at = now()
where id = $1;