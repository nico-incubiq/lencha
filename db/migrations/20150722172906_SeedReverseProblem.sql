
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (title, small_description, description, api_url)
 VALUES (
    'Reverse',
    'Reverse a string. Pretty simple isn''t it ? Solve this problem to activate your account.',
    'Reverse a string. Pretty simple isn''t it ? Solve this problem to activate your account.',
    'reverse'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE title='Reverse';
