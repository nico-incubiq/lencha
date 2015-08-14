
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (title, small_description, description, api_url)
 VALUES (
    'Maze',
    'Find the only solution of a maze',
    'Find the only solution of a maze',
    'maze'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE title='Reverse';
