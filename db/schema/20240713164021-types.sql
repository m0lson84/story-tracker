-- migrate:up
create type POINTS as enum (
    'none', 'one', 'two', 'three', 'five', 'eight', 'thirteen'

)
;

create type STATUS as enum (
    'unstarted', 'started', 'finished', 'delivered', 'rejected', 'approved'
)
;

create type TYPE as enum (
'bug', 'chore', 'feature'
)
;

-- migrate:down
drop type POINTS;
drop type STATUS;
drop type TYPE;

