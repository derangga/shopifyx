package query

const (
	QueryInsertPayment = `INSERT INTO payment(product_id, bank_account_id, payment_proof_image_url, quantity, created_at)
		VALUES (:product_id, :bank_account_id, :payment_proof_image_url, :quantity, now())`
)
