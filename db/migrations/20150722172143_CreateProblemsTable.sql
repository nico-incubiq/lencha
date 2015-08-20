
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Problems
(
    id serial NOT NULL,
    name character varying(80) NOT NULL,
    small_description character varying(400) NOT NULL,
    description character varying(10000) NOT NULL,
    api_url character varying(200) NOT NULL,
    solved_total integer NOT NULL DEFAULT 0,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),

    CONSTRAINT id_problems_pkey PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Problems;
