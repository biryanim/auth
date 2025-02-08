-- +goose Up
create table accesses(
    id serial primary key,
    endpoint_address text not null,
    role smallint not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table accesses;
