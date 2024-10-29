package handlers

import (
	"test/domain/general"
	"test/usecase"
)

type Handler struct {
	User  UserHandler
	Token TokenHandler
}

func NewHandler(usecase usecase.Usecase, conf *general.SectionService) Handler {
	return Handler{
		User:  NewUserHandler(usecase),
		Token: NewTokenHandler(usecase, conf),
	}
}
