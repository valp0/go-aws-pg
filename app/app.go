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
	go logExit(c)

	router := mux.NewRouter().StrictSlash(true)

	repo, err := repo.GetRepo()
	if err != nil {
		log.Fatal(err)
	}

	svc := services.NewService(repo)

	handler := handlers.NewHandler(svc)

	router.HandleFunc("/api/ping/", handler.Ping).Methods("GET")
	router.HandleFunc("/api/users/", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/", handler.PostUser).Methods("POST")
	router.HandleFunc("/api/users/{id}/", handler.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}/", handler.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/users/{id}/", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/favorites/", handler.GetFavorites).Methods("GET")
	router.HandleFunc("/api/users/{id}/favorites/", handler.PostFavorite).Methods("POST")
	router.HandleFunc("/api/users/{id}/favorites/{vidId}/", handler.DeleteFavorite).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	log.Printf("Listening on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		return err
	}

	return nil
}

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
