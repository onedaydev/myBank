CREATE TABLE accounts (
        account_id VARCHAR(36) NOT NULL PRIMARY KEY,
        owner_name VARCHAR(255) NOT NULL,
        balance DECIMAL(10, 2) NOT NULL,
        currency VARCHAR(10) NOT NULL
);