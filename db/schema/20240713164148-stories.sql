-- migrate:up
create table stories (
    id SERIAL primary key,
    user_id SERIAL references users (id),
    title TEXT not null,
    type TYPE not null default 'feature',
    status STATUS not null default 'unstarted',
    points POINTS not null default 'none',
    description TEXT not null
)
;

-- migrate:down
drop table stories;

