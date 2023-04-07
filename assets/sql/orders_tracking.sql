-- Create a new order state
INSERT INTO orders_tracking (order_id, status)
VALUES (400, 'pending')
RETURNING id;

-- Get current order state
SELECT (id, order_id, status, started_at, finished_at)
FROM orders_tracking
WHERE order_id = 400
  AND finished_at IS NULL;

-- Finish current order state
UPDATE orders_tracking
SET finished_at = now()
WHERE id = 400;

-- Delete order state
DELETE
FROM orders_tracking
WHERE id = 1400;


