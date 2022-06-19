CREATE TABLE users (
    id SERIAL PRIMARY KEY UNIQUE,
    username VARCHAR(255),
    password VARCHAR(255),
    first_name VARCHAR(50),
    last_name VARCHAR(50) 
);

CREATE TABLE notes (
    id SERIAL PRIMARY KEY UNIQUE,
    title VARCHAR(50),
    description TEXT
); 

CREATE TABLE usersNotes (
    id SERIAL PRIMARY KEY UNIQUE,
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    note_id int REFERENCES notes(id) ON DELETE CASCADE NOT NULL
);