begin;

CREATE TABLE IF NOT EXISTS bank_account(
	id serial primary key not null,
	user_id int not null, 
	bank_name varchar(15) not null,
	bank_account_name varchar(15) not null,
	bank_account_number varchar(15) not null,
	created_at timestamp not null default now(),
	updated_at timestamp,
	deleted_at timestamp
);

ALTER TABLE bank_account ADD CONSTRAINT bank_account_fk_users FOREIGN KEY (user_id) REFERENCES users(id);

commit;