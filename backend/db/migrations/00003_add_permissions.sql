-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS permissions(
    id bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);


INSERT INTO permissions(code)
VALUES
    ('movies:read'),
    ('movies:write'),
    ('clips:read'),
    ('clips:write');


-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS users_permissions;
DROP TABLE IF EXISTS permissions;
