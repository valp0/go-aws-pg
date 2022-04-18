package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/handlers"
)

const (
	port = 8080
)

// RunServer will setup a server and defines the routes it will respond to.
func RunServer() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go logExit(c)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/ping/", handlers.Ping).Methods("GET")
	router.HandleFunc("/api/users/", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/", handlers.PostUser).Methods("POST")
	router.HandleFunc("/api/users/{id}/", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}/", handlers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/users/{id}/", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/favorites/", handlers.GetFavorites).Methods("GET")
	router.HandleFunc("/api/users/{id}/favorites/", handlers.PostFavorite).Methods("POST")
	router.HandleFunc("/api/users/{id}/favorites/{vidId}/", handlers.DeleteFavorite).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	log.Printf("Listening on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		return err
	}

	return nil
}

func logExit(c chan os.Signal) {
	for range c {
		fmt.Print("\r")
		log.Fatal("Process terminated")
	}
}
