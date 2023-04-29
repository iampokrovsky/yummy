-- Get random restaurant name
CREATE OR REPLACE FUNCTION get_random_restaurant_name()
    RETURNS TEXT
    LANGUAGE plpgsql
AS
$$
DECLARE
    restaurant_names text[] := '{Trattoria,Osteria,Ristorante,Pizzeria,Taverna,Bistro,Cafe,Lagoon,Wings}';
BEGIN
    RETURN restaurant_names[ceil(random() * array_length(restaurant_names, 1))];
END
$$;

-- Get random address
CREATE OR REPLACE FUNCTION get_random_address()
    RETURNS text
    LANGUAGE plpgsql
AS
$$
DECLARE
    streets_arr text[] := array ['Main St', 'Oak St', 'Elm St', 'Maple Ave', 'Cedar Ln', 'Spruce Dr', 'Pine Rd', 'Birch St', 'Hilltop Dr'];
    cities_arr  text[] := array ['New York', 'Los Angeles', 'Chicago', 'Houston', 'Phoenix', 'Philadelphia', 'San Antonio', 'San Diego', 'Dallas'];
    states_arr  text[] := array ['AL', 'AK', 'AZ', 'AR', 'CA', 'CO', 'CT', 'DE', 'FL', 'GA', 'HI', 'ID', 'IL', 'IN', 'IA', 'KS', 'KY', 'LA', 'ME', 'MD', 'MA', 'MI', 'MN', 'MS', 'MO', 'MT', 'NE', 'NV', 'NH', 'NJ', 'NM', 'NY', 'NC', 'ND', 'OH', 'OK', 'OR', 'PA', 'RI', 'SC', 'SD', 'TN', 'TX', 'UT', 'VT', 'VA', 'WA', 'WV', 'WI', 'WY'];
BEGIN
    RETURN (SELECT (SELECT random() * 9999)::integer || ' ' ||
                   (SELECT streets_arr[ceil(random() * array_length(streets_arr, 1))]) || ', ' ||
                   (SELECT cities_arr[ceil(random() * array_length(cities_arr, 1))]) || ', ' ||
                   (SELECT states_arr[ceil(random() * array_length(states_arr, 1))]));
END
$$;

-- Get random cuisine
CREATE OR REPLACE FUNCTION get_random_cuisine()
    RETURNS text
    LANGUAGE plpgsql
AS
$$
DECLARE
    cuisines text[] := '{Italian,Chinese,Mexican,Japanese,Indian,French,Greek,Korean,Russian,Georgian}';
BEGIN
    RETURN cuisines[ceil(random() * array_length(cuisines, 1))];
END
$$;

-- Insert restaurants
INSERT INTO restaurants (name, address, cuisine)
SELECT cuisine_type || ' ' || get_random_restaurant_name(),
       get_random_address(),
       cuisine_type
FROM (SELECT get_random_cuisine()::cuisine AS cuisine_type
      FROM generate_series(1, 20)) AS random_cuisine;

-- Get random menu item name
CREATE OR REPLACE FUNCTION get_random_menu_item_name()
    RETURNS text
    LANGUAGE plpgsql
AS
$$
DECLARE
    adjectives       TEXT[] := ARRAY ['Spicy', 'Sweet', 'Savory', 'Tangy', 'Crispy', 'Juicy', 'Zesty'];
    nouns            TEXT[] := ARRAY ['Chicken', 'Beef', 'Pork', 'Fish', 'Shrimp', 'Tofu', 'Vegetables'];
    prefixes         TEXT[] := ARRAY ['Grilled', 'Fried', 'Baked', 'Steamed', 'Roasted', 'Sauteed'];
    suffixes         TEXT[] := ARRAY ['Wrap', 'Bowl', 'Plate', 'Sandwich', 'Burger', 'Taco', 'Pizza'];
    random_adjective TEXT;
    random_noun      TEXT;
    random_prefix    TEXT;
    random_suffix    TEXT;
BEGIN
    random_adjective := adjectives[1 + floor(random() * array_length(adjectives, 1))];
    random_noun := nouns[1 + floor(random() * array_length(nouns, 1))];
    random_prefix := prefixes[1 + floor(random() * array_length(prefixes, 1))];
    random_suffix := suffixes[1 + floor(random() * array_length(suffixes, 1))];
    RETURN random_adjective || ' ' || random_prefix || ' ' || random_noun || ' ' || random_suffix;
END
$$;

-- Get random price
CREATE OR REPLACE FUNCTION get_random_price()
    RETURNS bigint
    LANGUAGE plpgsql
AS
$$
DECLARE
    min_price    bigint := 5;
    max_price    bigint := 100;
    magnifier    int    := 10000;
    random_price bigint;
BEGIN
    random_price := (random() * (max_price - min_price) + min_price)::bigint * magnifier;
    RETURN random_price;
END
$$;

-- Insert menu items
INSERT INTO menu_items (restaurant_id, name, price)
SELECT i % (SELECT COUNT(*) FROM restaurants) + 1,
       get_random_menu_item_name(),
       get_random_price()
FROM generate_series(1, 100) AS i;

-- Get random name
CREATE OR REPLACE FUNCTION get_random_name()
    RETURNS text
    LANGUAGE plpgsql
AS
$$
DECLARE
    first_names text[] := '{Adam,Anna,Bob,Charlie,David,Emma,Frank,Greg,Helen,' ||
                          'Ian,John,Kate,Lucy,Mike,Nancy,Oliver,Peter,Quentin,' ||
                          'Rachel,Steve,Tom,Una,Victoria,William,Xander,Yvonne,Zoe}';
    last_names  text[] := '{Smith,Johnson,Williams,Jones,Brown,Davis,Miller,' ||
                          'Wilson,Moore,Taylor,Anderson,Thomas,Jackson,White,' ||
                          'Harris,Martin,Thompson,Young,Allen,King,Wright,Scott,' ||
                          'Green,Baker,Adams,Nelson,Carter,Mitchell,Perez,Roberts,' ||
                          'Turner,Phillips,Campbell,Parker,Evans,Edwards,Collins,' ||
                          'Stewart,Sanchez,Morris,Rogers,Reed,Cook,Morgan,Bell,' ||
                          'Murphy,Bailey,Rivera,Cooper,Richardson,Cox,Howard,Ward,' ||
                          'Torres,Peterson,Gray,Ramirez,James,Watson,Brooks,Kelly,' ||
                          'Sanders,Price,Bennett,Wood,Barnes,Ross,Henderson,Coleman,' ||
                          'Jenkins,Perry,Powell,Sullivan,Long,Foster}';
BEGIN
    RETURN first_names[ceil(random() * array_length(first_names, 1))] || ' ' ||
           last_names[ceil(random() * array_length(last_names, 1))];
END
$$;

-- Get random phone number
CREATE OR REPLACE FUNCTION get_random_phone_number()
    RETURNS TEXT
    LANGUAGE plpgsql
AS
$$
BEGIN
    RETURN to_char(random() * 10000000000, 'FM"+1 ("000") "000"-"0000');
END
$$;

-- Insert users
INSERT INTO users (name, email, phone_number)
SELECT user_name,
       lower(replace(user_name, ' ', '.')) || i || '@mail.com',
       get_random_phone_number()
FROM (SELECT row_number() OVER () AS i,
             get_random_name()    AS user_name
      FROM generate_series(1, 150)) AS random_name;

-- Get random transport
CREATE OR REPLACE FUNCTION get_random_transport()
    RETURNS text
    LANGUAGE plpgsql
AS
$$
DECLARE
    transport text[] := '{foot,bike,car}';
BEGIN
    RETURN transport[ceil(random() * array_length(transport, 1))];
END
$$;

-- Insert couriers
INSERT INTO couriers (user_id, transport)
SELECT i,
       get_random_transport()::transport
FROM generate_series(1, 50) AS i;

-- Insert clients
INSERT INTO clients (user_id, address)
SELECT i,
       get_random_address()
FROM generate_series(51, 150) AS i;

-- Insert orders
INSERT INTO orders (client_id, courier_id, restaurant_id)
SELECT ceil(random() * (SELECT COUNT(*) FROM clients)),
       ceil(random() * (SELECT COUNT(*) FROM couriers)),
       ceil(random() * (SELECT COUNT(*) FROM restaurants))
FROM generate_series(1, 1000) AS i;

-- Get random menu item id
CREATE OR REPLACE FUNCTION get_random_menu_item_id(group_id bigint, order_id bigint)
    RETURNS bigint
    LANGUAGE plpgsql
AS
$$
DECLARE
    rest_id bigint;
    item_id bigint;
BEGIN
    SELECT restaurant_id INTO rest_id FROM orders WHERE id = order_id;

    SELECT id
    INTO item_id
    FROM menu_items
    WHERE restaurant_id = rest_id
    ORDER BY id
    OFFSET group_id LIMIT 1;

    RETURN item_id;
END
$$;

-- Insert random menu items for orders
INSERT INTO orders_menu_items (order_id, menu_item_id, amount)
SELECT order_id,
       get_random_menu_item_id(group_id, order_id),
       ceil(random() * 10)
FROM (SELECT (row_number() OVER () / 1000) % 5                        AS group_id,
             row_number() OVER () % (SELECT COUNT(*) FROM orders) + 1 AS order_id
      FROM generate_series(1, 5000)) AS i;

-- Insert tracking for orders
INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'created'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders)) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'received'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders)) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'pending'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders) / 3) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'preparing'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders) / 4) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'shipping'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders) / 5) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'completed'::order_status, now()
FROM generate_series(1, (SELECT COUNT(*) FROM orders) / 10) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'cancelled'::order_status, now()
FROM generate_series(500, 600) AS i;

INSERT INTO orders_tracking (order_id, status, finished_at)
SELECT i, 'failed'::order_status, now()
FROM generate_series(650, 700) AS i;

-- Drop all functions
DROP FUNCTION IF EXISTS get_random_restaurant_name();
DROP FUNCTION IF EXISTS get_random_address();
DROP FUNCTION IF EXISTS get_random_cuisine();
DROP FUNCTION IF EXISTS get_random_menu_item_name();
DROP FUNCTION IF EXISTS get_random_price();
DROP FUNCTION IF EXISTS get_random_name();
DROP FUNCTION IF EXISTS get_random_phone_number();
DROP FUNCTION IF EXISTS get_random_transport();
DROP FUNCTION IF EXISTS get_random_menu_item_id(group_id bigint, order_id bigint);
