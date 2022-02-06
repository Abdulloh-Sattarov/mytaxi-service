CREATE TABLE IF NOT EXISTS drivers(
    id uuid not null primary key,
    first_name varchar(64),
    last_name varchar(64),
    phone varchar(16),
    car_model varchar(16),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS clients(
    id uuid not null primary key,
    first_name varchar(64),
    last_name varchar(64),
    phone varchar(16),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid not null primary key,
    cost decimal(10,5),
    driver_id uuid not null references drivers(id),
    client_id uuid not null references clients(id),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);