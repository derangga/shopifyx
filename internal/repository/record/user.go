package record

import (
	"database/sql"
	"time"

	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/pkg/helper"
)

type User struct {
	ID        int          `db:"id"`
	Username  string       `db:"username"`
	Name      string       `db:"name"`
	Password  string       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (r *User) ToEntity() *entity.User {
	return &entity.User{
		ID:        r.ID,
		Username:  r.Username,
		Name:      r.Name,
		Password:  r.Password,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt.Time,
		DeletedAt: r.DeletedAt.Time,
	}
}

func UserEntityToRecord(req *entity.User) *User {
	return &User{
		ID:        req.ID,
		Username:  req.Username,
		Name:      req.Name,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
		UpdatedAt: helper.NullTime(req.UpdatedAt),
		DeletedAt: helper.NullTime(req.DeletedAt),
	}
}
