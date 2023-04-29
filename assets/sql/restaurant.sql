-- Create a _restaurant
INSERT INTO restaurants (name, address, cuisine)
VALUES ('Chiho', '5568 Pine Rd, Houston, NM', 'Italian'::cuisine)
RETURNING id;

-- Get all restaurants
SELECT id, name, address, cuisine, created_at, updated_at, deleted_at
FROM restaurants;

-- Get _restaurant by id
SELECT id, name, address, cuisine, created_at, updated_at, deleted_at
FROM restaurants
WHERE id = 1;

-- Get _restaurant by name
SELECT id, name, address, cuisine, created_at, updated_at, deleted_at
FROM restaurants
WHERE name ILIKE '%Chiho%';

-- Get _restaurant by cuisine
SELECT id, name, address, cuisine, created_at, updated_at, deleted_at
FROM restaurants
WHERE cuisine = 'Italian';

-- Update _restaurant by id
UPDATE restaurants
SET name       = 'Unicorn',
    address    = '135 Elm St, San Diego, OH',
    cuisine    = 'Greek',
    updated_at = now()
WHERE id = 1;

-- Delete _restaurant by id
UPDATE restaurants
SET deleted_at = now()
WHERE id = 1;

-- Restore _restaurant by id
UPDATE users
SET deleted_at = NULL,
    updated_at = now()
WHERE id = 151;

