-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users
(
    id         bigserial primary key,
    email      varchar(255) not null unique,
    username   varchar(30)  not null unique,
    password   varchar(60)  not null,
    role       user_role    not null,

    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now(),
    deleted_at timestamp             default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE user_role;
-- +goose StatementEnd
