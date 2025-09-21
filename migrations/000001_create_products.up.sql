CREATE TABLE products (
    id UUID PRIMARY KEY,
    sku TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    price BIGINT NOT NULL,
    currency TEXT NOT NULL
);
