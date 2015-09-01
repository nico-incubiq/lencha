
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE problems
SET name = 'Tic-Tac-Toe',
small_description = 'Beat the server at a Tic-tac-toe game.',
description = '<h2>TITLE</h2>',
api_url = 'tictactoe'
WHERE api_url = 'shortestpath';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
UPDATE problems
SET name = '',
small_description = '',
description = '',
api_url = ''
WHERE api_url = 'tictactoe';

