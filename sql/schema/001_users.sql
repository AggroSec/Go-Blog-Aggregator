-- +goose Up
create table users (
    id UUID primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name text unique not null
);

-- +goose Down
drop table users;