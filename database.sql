CREATE TABLE product (
  id serial PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price SMALLINT NOT NULL
);

INSERT INTO product (name, price) 
VALUES 
('Coffee', 12000),
('Milo', 9000),
('Cola', 7000),
('Sosro', 5000),
('Aqua', 2000);