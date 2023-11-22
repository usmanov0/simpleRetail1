CREATE TABLE recipients (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255),
        phone VARCHAR(255),
        account_number INT,
        description VARCHAR(255)
);
