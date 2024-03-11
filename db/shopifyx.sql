CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "name" varchar,
  "password" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "product" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "name" varchar,
  "price" int,
  "image_url" varchar,
  "stock" int,
  "condition" enum,
  "tags" text[],
  "is_purchaseable" bool,
  "purchase_count" int,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "bank_account" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "bank_name" varchar,
  "bank_account_name" varchar,
  "bank_account_number" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "payment" (
  "id" integer PRIMARY KEY,
  "product_id" integer,
  "bank_account_id" integer,
  "payment_proof_image_url" varchar,
  "quantity" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "product" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bank_account" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("bank_account_id") REFERENCES "bank_account" ("id");
