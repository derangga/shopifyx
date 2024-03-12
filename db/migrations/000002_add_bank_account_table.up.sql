begin;

CREATE TABLE IF NOT EXISTS bank_account(
	id serial primary key not null,
	bank_name varchar(15) not null,
	bank_account_name varchar(15) not null,
	bank_account_number varchar(15) not null,
	created_at timestamp not null default now(),
	updated_at timestamp,
	deleted_at timestamp
);

commit;