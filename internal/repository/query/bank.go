package query

const (
	QueryInsertBank = `INSERT INTO bank_account(user_id, bank_name, bank_account_name, bank_account_number, created_at)
		VALUES (:user_id, :bank_name, :bank_account_name, :bank_account_number, now())`
)
