/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

CREATE TABLE products (
  id serial PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price SMALLINT NOT NULL,
);

INSERT INTO products (name, price) 
VALUES 
("Coffee", 12000),
("Milo", 9000),
("Cola", 7000),
("Sosro", 5000),
("Aqua", 2000);