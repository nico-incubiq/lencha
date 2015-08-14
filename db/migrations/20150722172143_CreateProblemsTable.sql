
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Problems
(
    id serial NOT NULL,
    title character varying(80) NOT NULL,
    small_description character varying(400) NOT NULL,
    description character varying(1000) NOT NULL,
    api_url character varying(200) NOT NULL,
    solved_total integer NOT NULL DEFAULT 0,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),

    CONSTRAINT id_problems_pkey PRIMARY KEY (id),
    CONSTRAINT title_unique UNIQUE (title)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Problems;
