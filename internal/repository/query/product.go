package query

const (
	ProductInsertQuery = `INSERT INTO product(user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at)
		VALUES (:user_id, :name, :price, :image_url, :stock, :condition, :tags, :is_purchaseable, :purchase_count, :created_at) RETURNING id`

	ProductGetByID = `SELECT id, user_id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, created_at, updated_at FROM product WHERE "id" = $1 AND "deleted_at" IS NULL`
	ProductUpdate  = `UPDATE product set name=$3, price=$4, image_url=$5, condition=$6, tags=$7, is_purchaseable=$8, updated_at=$9
		WHERE id=$1 and user_id=$2 and (deleted_at is null)`
	ProductDelete              = `UPDATE product set deleted_at=$3 WHERE id=$1 and user_id=$2`
	ProductStockUpdate         = `UPDATE product set stock=$3 WHERE id=$1 and user_id=$2 and (deleted_at is null)`
	ProductPurchaseCountUpdate = `UPDATE product set purchase_count=$3 WHERE id=$1 and user_id=$2 and (deleted_at is null)`
)
