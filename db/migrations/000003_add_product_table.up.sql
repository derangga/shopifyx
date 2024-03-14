begin;

CREATE TYPE "product_condition" AS ENUM (
  'NEW',
  'SECOND'
);

CREATE TABLE IF NOT EXISTS product(
  id serial PRIMARY KEY,
  user_id integer,
  name varchar,
  price int,
  image_url varchar,
  stock int,
  condition product_condition,
  tags text[],
  is_purchaseable bool,
  purchase_count int,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE INDEX ON product (user_id, name, price, tags);

ALTER TABLE product ADD FOREIGN KEY (user_id) REFERENCES users (id);

commit;