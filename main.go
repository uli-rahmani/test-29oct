package main

import (
	"fmt"
	"log"
	"net/http"
	"test/config"
	"test/routes"

	"github.com/gorilla/handlers"
)

func main() {
	conf, err := config.GetCoreConfig()
	if err != nil {
		panic(err)
	}

	handler, err := config.NewRepoContext(conf)
	if err != nil {
		panic(err)
	}
	origins := handlers.AllowedOrigins(
		[]string{},
	)
	credentials := handlers.AllowCredentials()

	router := routes.GetCoreEndpoint(conf, handler)

	port := fmt.Sprintf(":%s", conf.App.Port)
	log.Println("server listen to port ", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(nil, nil, origins, credentials)(router)))
}
