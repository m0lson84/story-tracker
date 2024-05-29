-- name: CreateStory :one
insert into stories (user_id, title, type, status, points, description)
values ($1, $2, $3, $4, $5, $6)
returning *
;

-- name: DeleteStory :exec
delete from stories
where id = $1
;


-- name: GetStory :one
select id, user_id, title, type, status, points, description
from stories
where id = $1
limit 1
;

-- name: ListStories :many
select id, user_id, title, type, status, points, description
from stories
order by id
;

-- name: UpdateStory :exec
update stories
set title = coalesce(sqlc.narg(title), title),
    user_id = coalesce(sqlc.narg(user_id), user_id),
    type = coalesce(sqlc.narg(type), type),
    status = coalesce(sqlc.narg(status), status),
    points = coalesce(sqlc.narg(points), points),
    description = coalesce(sqlc.narg(description), description)
where id = sqlc.arg(id)
;

