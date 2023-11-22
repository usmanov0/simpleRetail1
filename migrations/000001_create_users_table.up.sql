-- ./migrations/<timestamp>_create_users_table.up.sql

-- The `CREATE TABLE` statement to create the Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    pinfl VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(255),
    enabled BOOLEAN
);
