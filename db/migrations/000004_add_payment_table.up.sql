begin;

CREATE TABLE IF NOT EXISTS payment(
  id serial PRIMARY KEY,
  product_id integer,
  bank_account_id integer,
  payment_proof_image_url varchar,
  quantity integer,
  created_at timestamp,
  updated_at timestamp
);

CREATE INDEX ON "payment" (product_id, bank_account_id);

ALTER TABLE payment ADD FOREIGN KEY (product_id) REFERENCES product (id);

ALTER TABLE payment ADD FOREIGN KEY (bank_account_id) REFERENCES bank_account (id);

commit;