-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         bigserial primary key,
    email      text      not null unique,
    username   text      not null unique,
    password   text      not null,
    role       text      not null,

    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp          default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
