-- Create an order's menu item
INSERT INTO orders_menu_items (order_id, menu_item_id, amount)
VALUES (900, 27, 10)
RETURNING id;

-- Get order's menu items by order id
SELECT (id, order_id, menu_item_id, amount, created_at, updated_at)
FROM orders_menu_items
WHERE order_id = 900;

-- Update amount of order's menu item by id
UPDATE orders_menu_items
SET amount     = 10,
    updated_at = now()
WHERE id = 899;

-- Delete order's menu item by id
UPDATE orders_menu_items
SET deleted_at = now()
WHERE id = 899;

-- Restore order's menu item by id
UPDATE orders_menu_items
SET deleted_at = NULL,
    amount     = 1,
    updated_at = now()
WHERE id = 899;
