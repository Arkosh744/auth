-- +goose Up
create table accesses
(
    id               bigserial primary key,
    endpoint_address text      not null,
    role             text      not null,
    created_at       timestamp not null default now(),
    updated_at       timestamp not null default now()
);

insert into accesses (endpoint_address, role)
values ('/chat_v1.ChatV1/CreateChat', 'admin');

-- +goose Down
drop table accesses;