package main

import (
	"log"
	"net/http"
	"tahanasir/service-catalog/internal/config"
	"tahanasir/service-catalog/internal/database"
	"tahanasir/service-catalog/internal/transport"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.ParseEnv()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewDBConn(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	storage := database.NewStorage(db)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/v1/services", transport.GetAllServices(storage))
	router.Get("/v1/services/{id}", transport.GetService(storage))

	log.Fatal(http.ListenAndServe(":8080", router))
}
