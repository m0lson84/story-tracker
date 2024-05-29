-- migrate:up
create table users(
    id SERIAL primary key,
    username TEXT not null,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP
)
;

-- migrate:down
drop table users;

