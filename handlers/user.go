package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"test/domain"
	"test/usecase"

	cg "test/constants/general"

	"gopkg.in/dealancer/validate.v2"
	"gopkg.in/guregu/null.v4"
)

type UserHandler struct {
	userUsecase usecase.UserUsecaseIntf
}

func NewUserHandler(usecase usecase.Usecase) UserHandler {
	return UserHandler{
		userUsecase: usecase.User,
	}
}

func (uh UserHandler) Login(res http.ResponseWriter, req *http.Request) {
	respData := &ResponseData{
		Status: cg.Fail,
	}

	ctx := req.Context()

	var param domain.LoginRequest

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		respData.Error = null.StringFrom(cg.HandlerErrorRequestDataEmpty)
		WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Error = null.StringFrom(cg.HandlerErrorRequestDataNotValid)
		WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Error = null.StringFrom(cg.HandlerErrorRequestDataFormatInvalid)
		WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	resp, message, err := uh.userUsecase.Login(ctx, param)
	if err != nil {
		respData.Error = null.StringFrom(message)
		WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &ResponseData{
		Status: cg.Success,
		Data:   resp,
	}

	WriteResponse(res, respData, http.StatusOK)
}
