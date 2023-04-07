-- Create courier
INSERT INTO couriers(user_id, transport)
VALUES (150, 'car'::transport)
RETURNING id;

-- Get courier by id
SELECT (id, user_id, transport, created_at, updated_at, deleted_at)
FROM couriers
WHERE id = 51;

-- Update courier by id
UPDATE couriers
SET transport  = 'bike'::transport,
    updated_at = now()
WHERE id = 51;

-- Delete courier by id
UPDATE couriers
SET deleted_at = now()
WHERE id = 51;

-- Restore courier by id
UPDATE couriers
SET deleted_at = NULL,
    updated_at = now()
WHERE id = 51;
