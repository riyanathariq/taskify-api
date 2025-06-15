-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, username, password, name, created_at, updated_at)
VALUES (
           'af850194-8216-4fed-8501-948216ffed43',
           'riyanathariq',
           '$2a$10$LdR4hZgxIWRi3/OM504BOOgVddYq0xPvq5rAJD.9zsXbX0dcycKOu',
           'Riyan Athariq',
           CURRENT_TIMESTAMP,
           CURRENT_TIMESTAMP
       );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE username = 'riyanathariq';
-- +goose StatementEnd
