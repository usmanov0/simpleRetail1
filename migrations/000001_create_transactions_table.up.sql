
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP,
    description VARCHAR(255),
    types VARCHAR(255),
    status VARCHAR(255),
    amount DOUBLE PRECISION,
    available_balance DOUBLE PRECISION
);
