
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE ProblemsSolved
(
    id serial NOT NULL,
    problem_id integer REFERENCES Problems (id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_id integer REFERENCES Users (id) ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT id_problems_solved_pkey PRIMARY KEY (id)
);

CREATE INDEX problems_solved_index ON ProblemsSolved (user_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX problems_solved_index;
DROP TABLE ProblemsSolved;
