-- +goose Up

CREATE TABLE users (
    id UNIQUEIDENTIFIER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at DATETIME NOT NULL,
    
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;