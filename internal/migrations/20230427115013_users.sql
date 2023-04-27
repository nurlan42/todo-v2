-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id int not null,
    login text,
    password text,
    created_at timestamp,
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table users
-- +goose StatementEnd
