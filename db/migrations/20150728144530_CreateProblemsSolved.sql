
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE problems_solved
(
    id serial NOT NULL,
    problem_id integer REFERENCES Problems (id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_id integer REFERENCES Users (id) ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT id_problems_solved_pkey PRIMARY KEY (id),
    CONSTRAINT user_id_problem_id_unique UNIQUE (problem_id, user_id)
);

CREATE INDEX problems_solved_index ON problems_solved (user_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX problems_solved_index;
DROP TABLE problems_solved;
