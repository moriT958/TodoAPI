-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
	id serial NOT NULl,
	uuid varchar(40) NOT NULL,
	title varchar(30) NOT NULL,
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
