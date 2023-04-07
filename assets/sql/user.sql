-- Create user
INSERT INTO users(name, email, phone_number)
VALUES ('Regina Johnson', 'regina.johnson1997@gmail.com', '+1 (999) 999-9999')
RETURNING id;

-- Get user by id
SELECT (id, name, email, phone_number, created_at, updated_at, deleted_at)
FROM users
WHERE id = 151;

-- Get user by phone
SELECT (id, name, email, phone_number, created_at, updated_at, deleted_at)
FROM users
WHERE phone_number = '+1 (999) 999-9999';

-- Get user by email
SELECT (id, name, email, phone_number, created_at, updated_at, deleted_at)
FROM users
WHERE email = 'regina.johnson1997@gmail.com';

-- Update user by id
UPDATE users
SET name         = 'Oleg Green',
    email        = 'oleg.green1961@yandex.ru',
    phone_number = '+7 (955) 550-0335',
    updated_at   = now()
WHERE id = 151;

-- Delete user by id
UPDATE users
SET deleted_at = now()
WHERE id = 151;

-- Restore user by id
UPDATE users
SET deleted_at = NULL,
    updated_at = now()
WHERE id = 151;
