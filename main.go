package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Port: ", port)

	router := chi.NewRouter()
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https:*", "http:*" },
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			MaxAge: 300,
		}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", checkReadiness)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}