-- Create a menu item
INSERT INTO menu_items (restaurant_id, name, price)
VALUES (1, 'Fresh Baked Bread', 100000)
RETURNING id;

-- Get menu item by id
SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at
FROM menu_items
WHERE id = 1;

-- Get menu items by _restaurant id
SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at
FROM menu_items
WHERE restaurant_id = 1;

-- Get menu items by name
SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at
FROM menu_items
WHERE name ILIKE '%Crispy%';

-- Update menu items by id
UPDATE menu_items
SET name       = 'Horny Sausage',
    price      = 143000,
    updated_at = now()
WHERE id = 20;

-- Delete menu item by id
UPDATE menu_items
SET deleted_at = now()
WHERE id = 20;

-- Restore menu item by id
UPDATE menu_items
SET deleted_at = NULL,
    updated_at = now()
WHERE id = 20;
