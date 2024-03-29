package query

const (
	UserInsertQuery = `INSERT INTO users(username, name, password, created_at)
		VALUES (:username, :name, :password, :created_at) RETURNING id`

	UserGetByUsernameQuery = `SELECT id, username, name, password FROM users WHERE "username" = $1 AND "deleted_at" IS NULL`
)
