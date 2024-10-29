package usecase

import (
	"test/domain/general"
	"test/repo/db"
)

type Usecase struct {
	User UserUsecaseIntf
}

func NewUsecase(conf *general.SectionService, dbRepo db.DBRepo) Usecase {
	return Usecase{
		User: newUserUsecase(conf, dbRepo),
	}
}
