CREATE TABLE IF NOT EXISTS users(
    id text not null primary key,
    token text not null
);

CREATE TABLE IF NOT EXISTS wallets(
    id text not null primary key,
    owned_by   text not null,
    status    text not null,
    enabled_at TIMESTAMP null,
    Balance   float default 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at timestamp null

);

CREATE TABLE IF NOT EXISTS transactions(
    id text primary key,
    wallet_id text not null,
    type varchar(200) not null,
    amount float default 0,
    issued_by text not null,
    reference_id text not null,
    status text not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at timestamp null
);

