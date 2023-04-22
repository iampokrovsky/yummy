-- Create an order
INSERT INTO orders (client_id, courier_id, restaurant_id)
VALUES (3, 5, 9)
RETURNING id;

-- Get order by client id
SELECT client_id, courier_id, restaurant_id
FROM orders
WHERE client_id = 3;

-- Get order by _restaurant id
SELECT client_id, courier_id, restaurant_id
FROM orders
WHERE restaurant_id = 3;

-- Get order by courier id
SELECT client_id, courier_id, restaurant_id
FROM orders
WHERE courier_id = 3;

-- Update order's courier id
UPDATE orders
SET courier_id = 10,
    updated_at = now()
WHERE id = 20;

-- Delete order by id
UPDATE orders
SET deleted_at = now()
WHERE id = 100;
