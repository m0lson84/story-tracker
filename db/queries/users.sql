-- name: CreateUser :one
insert into users (username)
values ($1)
returning *
;

-- name: DeleteUser :exec
delete from users
where id = $1
;

-- name: GetUser :one
select id, username, created_at, updated_at
from users
where id = $1
limit 1
;

-- name: ListUsers :many
select id, username
from users
order by username
;


-- name: UpdateUser :exec
update users
set username = $2
where id = $1
returning id, username
;

