CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255),
    first_name VARCHAR(50),
    last_name VARCHAR(50) 
);

CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    description TEXT
);