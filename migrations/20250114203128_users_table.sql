-- +goose Up
CREATE TYPE role_type AS ENUM (
    'UNKNOWN_ROLE_TYPE',
    'user',
    'admin'
);

create table users (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) unique not null,
    is_admin role_type not null default 'UNKNOWN_ROLE_TYPE',
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose Down
drop table users;
drop type role_type;
