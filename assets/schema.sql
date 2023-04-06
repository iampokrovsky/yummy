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
    'shipped',
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

CREATE TABLE IF NOT EXISTS restaurants
(
    id         bigserial PRIMARY KEY,
    name       text        NOT NULL,
    address    text        NOT NULL,
    cuisine    cuisine,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS menu_items
(
    id            bigserial PRIMARY KEY,
    restaurant_id bigint      NOT NULL,
    name          text        NOT NULL,
    price         bigint      NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz,
    deleted_at    timestamptz
);


CREATE TABLE IF NOT EXISTS users
(
    id           bigserial PRIMARY KEY,
    name         text        NOT NULL,
    email        text UNIQUE NOT NULL,
    phone_number text UNIQUE NOT NULL,
    created_at   timestamptz NOT NULL DEFAULT now(),
    updated_at   timestamptz,
    deleted_at   timestamptz
);

CREATE TABLE IF NOT EXISTS couriers
(
    id         bigserial PRIMARY KEY,
    user_id    bigint UNIQUE NOT NULL,
    transport  transport     NOT NULL,
    created_at timestamptz   NOT NULL DEFAULT now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS clients
(
    id         bigserial PRIMARY KEY,
    user_id    bigint UNIQUE NOT NULL,
    address    text          NOT NULL,
    created_at timestamptz   NOT NULL DEFAULT now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS orders
(
    id            bigserial PRIMARY KEY,
    client_id     bigint      NOT NULL,
    courier_id    bigint      NOT NULL,
    restaurant_id bigint      NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz,
    deleted_at    timestamptz
);

CREATE TABLE IF NOT EXISTS orders_menu_items
(
    order_id     bigint NOT NULL,
    menu_item_id bigint NOT NULL,
    amount       int    NOT NULL,
    PRIMARY KEY (order_id, menu_item_id)
);

CREATE TABLE IF NOT EXISTS orders_tracking
(
    order_id    bigint       NOT NULL,
    status      order_status NOT NULL,
    started_at  timestamptz  NOT NULL DEFAULT now(),
    finished_at timestamptz,
    deleted_at  timestamptz,
    PRIMARY KEY (order_id, status)
);
