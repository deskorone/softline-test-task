CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(40) NOT NULL UNIQUE,
    email VARCHAR(40) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    phone_number VARCHAR(20) NOT NULL UNIQUE
);