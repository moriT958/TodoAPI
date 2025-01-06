-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_todos_uuid ON todos("uuid");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_todos_uuid;
-- +goose StatementEnd
