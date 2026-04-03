-- name: FeedList :many
select feeds.name, feeds.url, users.name as username
from feeds
inner join users on feeds.user_id = users.id
order by feeds.updated_at asc;