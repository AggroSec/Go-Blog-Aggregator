-- name: UnfollowFeed :exec

delete from feed_follows
where user_id = (select users.id from users where users.name = $1)
and feed_id = (select feeds.id from feeds where feeds.url = $2);