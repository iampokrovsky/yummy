-- +goose Up
CREATE TABLE orders
(
    id            bigserial PRIMARY KEY,
    client_id     bigint NOT NULL,
    courier_id    bigint NOT NULL,
    restaurant_id bigint NOT NULL,
    total_price   bigint NOT NULL
);

-- +goose Down
DROP TABLE orders;
