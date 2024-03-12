package query

const (
	QueryInsertBank = `INSERT INTO bank_account(bank_name, bank_account_name, bank_account_number, created_at)
		VALUES (:bank_name, :bank_account_name, :bank_account_number, now())`
)
