-- +goose Up
ALTER TABLE users ADD COLUMN specialization jsonb;

-- +goose Down
ALTER TABLE users DROP COLUMN specialization;
