
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Users
(
    id serial NOT NULL,
    username character varying(80) NOT NULL,
    hash character varying(60) NOT NULL,
    email character varying(100) NOT NULL,
    api_key character varying(64) NOT NULL,
    privilege integer NOT NULL DEFAULT 0,
    problems_solved integer NOT NULL DEFAULT 0,
    activated boolean NOT NULL DEFAULT false,
    email_update boolean NOT NULL DEFAULT false,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),

    CONSTRAINT id_users_pkey PRIMARY KEY (id),
    CONSTRAINT username_unique UNIQUE (username),
    CONSTRAINT email_unique UNIQUE (email)
);

CREATE INDEX user_username_index ON Users (username);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX user_username_index;
DROP TABLE Users;
