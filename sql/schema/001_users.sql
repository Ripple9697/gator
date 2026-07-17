-- +goose Up
CREATE TABLE users (
  id TEXT PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT not NULL UNIQUE
);

-- +goose Down
DROP TABLE users;

