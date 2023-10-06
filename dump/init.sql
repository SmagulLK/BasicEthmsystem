-- Create the 'accounts' table
CREATE TABLE users (
    account_id SERIAL PRIMARY KEY,
    public_key TEXT NOT NULL ,
    private_key TEXT NOT NULL,
    balance BIGINT NOT NULL,
    addres TEXT NOT NULL
);

-- Create the 'hot_wallet' table
CREATE TABLE hot_wallet (
    wallet_id SERIAL PRIMARY KEY,
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    balance BIGINT NOT NULL
);

-- Create the 'transactions' table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    from_account_pubk INT REFERENCES users(account_id),
    to_account_pubk INT REFERENCES users(account_id),
    -- amount BIGINT NOT NULL,
    -- numeric_amount numeric(30, 0)
    amount_str VARCHAR(64),
    -- status TEXT NOT NULL,
    private_key TEXT NOT NULL,
    adress_to TEXT NOT NULL,
    hex TEXT NOT NULL
);
