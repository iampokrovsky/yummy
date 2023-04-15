-- +goose Up
CREATE TYPE transport AS ENUM (
    'foot',
    'bike',
    'car'
    );

CREATE TYPE order_status AS ENUM (
    'created',
    'received',
    'pending',
    'preparing',
    'shipping',
    'completed',
    'cancelled',
    'failed'
    );

CREATE TYPE cuisine AS ENUM (
    'Italian',
    'Chinese',
    'Mexican',
    'Japanese',
    'Indian',
    'Thai',
    'French',
    'Greek',
    'Korean',
    'Russian',
    'Georgian'
    );

CREATE TABLE restaurants
(
    id         bigserial PRIMARY KEY,
    name       text        NOT NULL,
    address    text        NOT NULL,
    cuisine    cuisine,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE menu_items
(
    id            bigserial PRIMARY KEY,
    restaurant_id bigint      NOT NULL,
    name          text        NOT NULL,
    price         bigint      NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT (now()),
    updated_at    timestamptz,
    deleted_at    timestamptz
);

CREATE TABLE users
(
    id           bigserial PRIMARY KEY,
    name         text        NOT NULL,
    email        text UNIQUE NOT NULL,
    phone_number text UNIQUE NOT NULL,
    created_at   timestamptz NOT NULL DEFAULT (now()),
    updated_at   timestamptz,
    deleted_at   timestamptz
);

CREATE TABLE couriers
(
    id         bigserial PRIMARY KEY,
    user_id    bigint UNIQUE NOT NULL,
    transport  transport     NOT NULL,
    created_at timestamptz   NOT NULL DEFAULT (now()),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE clients
(
    id         bigserial PRIMARY KEY,
    user_id    bigint UNIQUE NOT NULL,
    address    text          NOT NULL,
    created_at timestamptz   NOT NULL DEFAULT (now()),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE orders
(
    id            bigserial PRIMARY KEY,
    client_id     bigint      NOT NULL,
    courier_id    bigint      NOT NULL,
    restaurant_id bigint      NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT (now()),
    updated_at    timestamptz,
    deleted_at    timestamptz
);

CREATE TABLE orders_menu_items
(
    id           bigserial PRIMARY KEY,
    order_id     bigint      NOT NULL,
    menu_item_id bigint      NOT NULL,
    amount       int         NOT NULL,
    created_at   timestamptz NOT NULL DEFAULT (now()),
    updated_at   timestamptz,
    deleted_at   timestamptz
);

CREATE TABLE orders_tracking
(
    id          bigserial PRIMARY KEY,
    order_id    bigint       NOT NULL,
    status      order_status NOT NULL,
    started_at  timestamptz  NOT NULL DEFAULT (now()),
    finished_at timestamptz
);

CREATE EXTENSION pg_trgm;

CREATE INDEX ON restaurants USING gin (name gin_trgm_ops);

CREATE INDEX ON restaurants (cuisine);

CREATE INDEX ON menu_items USING gin (name gin_trgm_ops);

CREATE INDEX ON menu_items (restaurant_id);

CREATE INDEX ON couriers (user_id);

CREATE INDEX ON clients (user_id);

CREATE INDEX ON orders (client_id);

CREATE INDEX ON orders (restaurant_id);

CREATE INDEX ON orders (courier_id);

CREATE UNIQUE INDEX ON orders_menu_items (order_id, menu_item_id);

CREATE UNIQUE INDEX ON orders_tracking (order_id, status);


-- +goose Down
DROP TABLE IF EXISTS restaurants;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS orders_menu_items;
DROP TABLE IF EXISTS orders_tracking;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS couriers;
DROP TABLE IF EXISTS clients;

DROP TYPE IF EXISTS transport CASCADE;
DROP TYPE IF EXISTS order_status CASCADE;
DROP TYPE IF EXISTS cuisine CASCADE;

DROP EXTENSION pg_trgm;
