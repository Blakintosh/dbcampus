/** Drop the table user */

Drop table IF EXISTS users;

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120)
);
