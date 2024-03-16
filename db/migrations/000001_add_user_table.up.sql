begin;

CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username varchar(15) unique not null,
  "name" varchar(50),
  "password" varchar(100),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE INDEX ON "users" (id, username);

commit;