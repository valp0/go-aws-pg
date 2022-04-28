package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/valp0/go-aws-pg/handlers"
	"github.com/valp0/go-aws-pg/repo"
	"github.com/valp0/go-aws-pg/services"
)

const (
	port = 8080
)

// RunServer will setup a server and defines the routes it will respond to.
func RunServer() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	router := mux.NewRouter().StrictSlash(true)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	repo, err := repo.GetRepo()
	if err != nil {
		log.Fatal(err)
	}
	go logExit(c)

	svc := services.NewService(repo)
	handler := handlers.NewHandler(svc)

	router.HandleFunc("/api/ping/", handler.Ping).Methods("GET")
	router.HandleFunc("/auth/get-token/", handler.GetToken).Methods("POST")
	router.HandleFunc("/api/users/", checkToken(handler.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users/", checkToken(handler.PostUser)).Methods("POST")
	router.HandleFunc("/api/users/{id}/", checkToken(handler.GetUser)).Methods("GET")
	router.HandleFunc("/api/users/{id}/", checkToken(handler.UpdateUser)).Methods("PATCH")
	router.HandleFunc("/api/users/{id}/", checkToken(handler.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/favorites/", checkToken(handler.GetFavorites)).Methods("GET")
	router.HandleFunc("/api/users/{id}/favorites/", checkToken(handler.PostFavorite)).Methods("POST")
	router.HandleFunc("/api/users/{id}/favorites/{vidId}/", checkToken(handler.DeleteFavorite)).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	log.Printf("Listening on port %d\n", port)

	lambda.Start(httpadapter.NewV2(router).ProxyWithContext)

	// if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
	// 	return err
	// }

	return nil
}

// Graceful shutdown.
func logExit(c chan os.Signal) {
	repo, _ := repo.GetRepo()
	for range c {
		if err := repo.CloseDB(); err != nil {
			log.Println(err)
		}
		fmt.Print("\r")
		log.Fatal("Process terminated")
	}
}

// Wrapper function to check JWT.
func checkToken(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return handlers.CheckToken()(http.HandlerFunc(f)).ServeHTTP
}
