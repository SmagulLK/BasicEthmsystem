-- Create the 'accounts' table
CREATE TABLE users (
    account_id SERIAL PRIMARY KEY,
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    balance BIGINT NOT NULL,
    decimals INT NOT NULL
);

-- Create the 'hot_wallet' table
CREATE TABLE hot_wallet (
    wallet_id SERIAL PRIMARY KEY,
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    balance BIGINT NOT NULL,
    decimals INT NOT NULL
);

-- Create the 'transactions' table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    from_account_id INT REFERENCES accounts(account_id),
    to_account_id INT REFERENCES accounts(account_id),
    amount BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
