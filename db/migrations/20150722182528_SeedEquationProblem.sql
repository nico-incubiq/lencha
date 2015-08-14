
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (title, small_description, description, api_url)
 VALUES (
    'Equation',
    'Solve a simple equation of the form : a*x^2 + b*x + c = 0',
    'Solve a simple equation of the form : a*x^2 + b*x + c = 0',
    'equation'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE title='Reverse';
