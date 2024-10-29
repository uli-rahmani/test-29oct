package db

import (
	"test/infra"
)

type DBRepo struct {
	User UserRepoItf
}

func NewDBRepo(db *infra.DatabaseList) DBRepo {
	return DBRepo{
		User: newUserRepo(db),
	}
}
