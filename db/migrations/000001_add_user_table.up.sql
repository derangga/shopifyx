begin;

CREATE TABLE IF NOT EXISTS users(
	id serial primary key not null,
	username varchar(15) not null,
	"name" varchar(50) not null,
	"password" varchar(15) not null,
	created_at timestamp not null default now(),
	deleted_at timestamp
);

commit;