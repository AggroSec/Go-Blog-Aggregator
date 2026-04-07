-- name: FollowingLookup :many
select feeds.name as feed_name, feeds.url as feed_url
from feeds
inner join feed_follows on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1;