package repository

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
)

type user struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) internal.UserRepository {
	return &user{
		db: db,
	}
}

// Get implements internal.UserRepository.
func (u *user) Get(ctx context.Context, id int) (*entity.User, error) {
	panic("unimplemented")
}

// GetByUsername implements internal.UserRepository.
func (u *user) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var userRecord record.User

	err := u.db.GetContext(ctx, &userRecord, query.UserGetByUsernameQuery, username)
	if err != nil {
		return nil, err
	}

	return userRecord.ToEntity(), nil
}

// Create implements internal.UserRepository.
func (u *user) Create(ctx context.Context, data *entity.User) (*entity.User, error) {
	return handleTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (*entity.User, error) {
		userRecord := record.UserEntityToRecord(data)

		stmt, err := tx.PrepareNamedContext(ctx, query.UserInsertQuery)
		if err != nil {
			return nil, err
		}

		row := stmt.QueryRowxContext(ctx, userRecord)
		if row.Err() != nil {
			return nil, row.Err()
		}

		err = row.Scan(&data.ID)
		if err != nil {
			return nil, err
		}

		return data, nil
	})
}
