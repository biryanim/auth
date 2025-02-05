-- +goose Up
create table users (
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) unique not null,
    email varchar(255) unique not null,
    role smallint not null default 0,
    password text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose Down
drop table users;
