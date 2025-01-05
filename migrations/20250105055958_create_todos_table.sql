-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
	id serial NOT NULl,
	uuid character(40) NOT NULL,
	title character(30) NOT NULL,
	is_completed boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id"),
	UNIQUE("uuid"),
	UNIQUE("title")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
