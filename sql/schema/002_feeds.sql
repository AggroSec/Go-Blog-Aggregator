-- +goose Up

create table feeds (
    id UUID primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name text not null,
    url text unique not null,
    user_id UUID not null,
    foreign key (user_id) references users(id) on delete cascade
);

-- +goose Down

drop table feeds;