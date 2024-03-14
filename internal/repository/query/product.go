package query

const (
	ProductInsertQuery = `INSERT INTO product(user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at)
		VALUES (:user_id, :name, :price, :image_url, :stock, :condition, :tags, :is_purchaseable, :purchase_count, :created_at) RETURNING id`

	ProductGetByID = `SELECT id, user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at, updated_at FROM product WHERE "id" = $1 AND "deleted_at" IS NULL`
	ProductUpdate  = `UPDATE product set name=$2, price=$3, image_url=$4, condition=$5, tags=$6, is_purchaseable=$7, updated_at=$8
		WHERE id=$1`
	ProductDelete = `UPDATE product set deleted_at=$2 where id=$1`
)
