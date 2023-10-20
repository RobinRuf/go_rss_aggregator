package main

import (
  "fmt"
  "os"
  "log"
  "net/http"
  
  "github.com/joho/godotenv"
  "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
) 

func main() {
  godotenv.Load()

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in the environment")
  }

  router := chi.NewRouter()

	// allow everything for development purposes
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter() // create a new router
	v1Router.Get("/ready", handlerReadiness) // connecting the handlerReadiness to the path /ready (Get, so it only works with Get-Requests like we want to)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router) // mount this router to the /v1 path

  srv := &http.Server{
    Handler: router,
    Addr: ":" + portString,
  }
  
  log.Printf("Server starting on port %v", portString)
  err := srv.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("PORT:", portString)
}
