package query

const (
	QueryInsertBank = `INSERT INTO bank_account(user_id, bank_name, bank_account_name, bank_account_number, created_at)
		VALUES (:user_id, :bank_name, :bank_account_name, :bank_account_number, now())`

	QueryUpdateBank = `UPDATE bank_account SET bank_name = $3, bank_account_name = $4, bank_account_number = $5, updated_at = now() 
		WHERE id = $1 AND user_id = $2`
)
