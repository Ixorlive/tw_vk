-- Create the User table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL
    password VARCHAR(255) NOT NULL
);
