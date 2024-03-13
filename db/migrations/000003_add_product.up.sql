begin;

CREATE TYPE product_condition AS ENUM (
  'NEW',
  'SECOND'
);

CREATE TABLE IF NOT EXISTS product (
  "id" serial primary key not null,
  "user_id" integer not null,
  "name" varchar not null,
  "price" int not null,
  "image_url" varchar not null,
  "stock" int not null,
  "condition" product_condition,
  "tags" text[] not null,
  "is_purchaseable" bool not null,
  "purchase_count" int,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE INDEX ON "product" ("user_id", "name", "price", "tags");
ALTER TABLE "product" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

commit;