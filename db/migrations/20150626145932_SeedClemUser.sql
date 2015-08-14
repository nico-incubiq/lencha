
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (username, hash, email, privilege, api_key, activated, created_at)
 VALUES ('Clem', '$2a$10$ni0si13xp.G0y9CCJwJGxeLlbj6tI7WDwSWwJfWw9QMSXP/uGpsbi', 'clement.laisne@gmail.com', 1, '1d5b6006e16c6e921d2ee85fe3629af1d32c39254a5c00c666529e72be27c533', true, now());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users WHERE username='Clem';
