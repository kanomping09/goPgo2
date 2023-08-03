package main

import (
	"net/http"
	"log"
	"fmt"
	"os"


	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

)
	

func main() {
	fmt.Println("Hola")

	godotenv.Load()

	portstr := os.Getenv("PORT")
	if portstr == "" {
		log.Fatal("PORT is not found in the environment")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:		[]string{"https://*", "http://*"},
		AllowedMethods:		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:		[]string{"*"},
		ExposedHeaders:		[]string{"Link"},
		AllowCredentials:	false,
		MaxAge:				300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:	":" + portstr,
	}
	log.Printf("Server starting on port %v", portstr)
	
	err :=  srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}