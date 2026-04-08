-- name: GetPostsForUser :many

select posts.title, posts.url, posts.description, posts.published_at from posts
inner join feeds on posts.feed_id = feeds.id
where feeds.user_id = $1
order by posts.published_at desc
limit $2;