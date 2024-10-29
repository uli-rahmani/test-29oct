package routes

import (
	"log"
	"test/domain/general"
	"test/handlers"

	"github.com/gorilla/mux"
)

func GetCoreEndpoint(conf *general.SectionService, handler handlers.Handler) *mux.Router {
	parentRoute := mux.NewRouter()

	jwtRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()
	nonJWTRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()
	// publicRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()

	// Middleware for public API
	// nonJWTRoute.Use(handler.Public.AuthValidator)

	// Middleware
	if conf.Authorization.JWT.IsActive {
		log.Println("JWT token is active")
		jwtRoute.Use(handler.Token.JWTValidator)
	}

	// Get Endpoint.
	getRoutes(nonJWTRoute, jwtRoute, conf, handler)

	return parentRoute
}
