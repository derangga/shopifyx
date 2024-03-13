package query

const (
	ProductInsertQuery = `INSERT INTO product(user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at)
		VALUES (:user_id, :name, :price, :image_url, :stock, :condition, :tags, :is_purchaseable, :purchase_count, :created_at) RETURNING id`

	ProductGetByID = `SELECT id, user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at, updated_at FROM product WHERE "id" = $1 AND "deleted_at" IS NULL`
)