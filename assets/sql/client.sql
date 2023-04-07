-- Create client
INSERT INTO clients(user_id, address)
VALUES (1, '3888 Main St, Phoenix, IA')
RETURNING id;

-- Get client by id
SELECT (id, user_id, address, created_at, updated_at, deleted_at)
FROM clients
WHERE id = 101;

-- Update client by id
UPDATE clients
SET address    = '6897 Pine Rd, Dallas, DE',
    updated_at = now()
WHERE id = 101;

-- Delete client by id
UPDATE clients
SET deleted_at = now()
WHERE id = 101;

-- Restore client by id
UPDATE clients
SET deleted_at = NULL,
    updated_at = now()
WHERE id = 101;
