package db

import (
	"context"
	"database/sql"
	"fmt"
	"test/domain"
	"test/infra"
)

type UserRepo struct {
	DBList *infra.DatabaseList
}

func newUserRepo(dbList *infra.DatabaseList) UserRepo {
	return UserRepo{
		DBList: dbList,
	}
}

const (
	uSelectUser = `
		SELECT
			id,
			"name",
			email,
			password,
			"role",
			created_at,
			updated_at,
			deleted_at
		FROM
			user
	`

	uWhere = `
		WHERE
	`

	uFilterByEmail = `
		email = ?
	`
)

type UserRepoItf interface {
	GetByEmail(ctx context.Context, email string) (*domain.UserData, error)
}

func (u UserRepo) GetByEmail(ctx context.Context, email string) (*domain.UserData, error) {
	var res domain.UserData

	q := fmt.Sprintf("%s%s%s", uSelectUser, uWhere, uFilterByEmail)
	query, args, err := u.DBList.Backend.Read.In(q, email)
	if err != nil {
		return nil, err
	}

	query = u.DBList.Backend.Read.Rebind(query)
	err = u.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}
