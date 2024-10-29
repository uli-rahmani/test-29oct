package routes

import (
	"net/http"
	"test/domain/general"
	"test/handlers"

	"github.com/gorilla/mux"
)

func getRoutes(router, routerJWT *mux.Router, conf *general.SectionService, handler handlers.Handler) {
	router.HandleFunc("/login", handler.User.Login).Methods(http.MethodPost)

}
