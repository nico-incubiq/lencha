
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (name, small_description, description, api_url)
 VALUES (
    'Shortest path',
    'Several cities are connected, find the shortest path from city A to city B',
    'Several cities are connected, find the shortest path from city A to city B',
    'shortestpath'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE name='Shortest path';
