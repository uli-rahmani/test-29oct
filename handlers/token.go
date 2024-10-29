package handlers

import (
	"context"
	"net/http"
	"strings"

	cg "test/constants/general"
	dg "test/domain/general"

	"test/usecase"

	// uu "test/usecase/user"
	"test/utils"

	"gopkg.in/guregu/null.v4"
)

type TokenHandler struct {
	Conf *dg.SectionService
}

func NewTokenHandler(uc usecase.Usecase, conf *dg.SectionService) TokenHandler {
	utils.InitJWTConfig(conf.Authorization.JWT)
	return TokenHandler{
		Conf: conf,
	}
}

func (th TokenHandler) JWTValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		respData := ResponseData{
			Status: cg.Fail,
		}

		//List of URL that bypass this JWTValidator middleware
		if req.URL.Path == "/api/v1/renew-token" {
			next.ServeHTTP(res, req)
			return
		}

		authorizationHeader := req.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			respData.Error = null.StringFrom("Invalid Token Format")
			WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims, err := utils.CheckAccessToken(accessToken)
		if err != nil {
			respData.Error = null.StringFrom("Token expired")
			WriteResponse(res, respData, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), "session", claims["session"])
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}
