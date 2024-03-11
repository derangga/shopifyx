begin;

CREATE TABLE IF NOT EXISTS users(
	id serial primary key not null,
	username varchar(20),
	"name" varchar(50),
	"password" varchar(100),
	created_at timestamp not null default now(),
	deleted_at timestamp
);

commit;