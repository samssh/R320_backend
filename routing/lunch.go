package routing

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"persian_meme_web_server/model/database"
	"time"
)

var Router *mux.Router = mux.NewRouter()

func LunchServer() error {
	err := database.LunchDB()
	if err != nil {
		return err
	}
	err = configureRouter()
	return err
}

func configureRouter() error {
	// routing
	port := 8080
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, Router),
		Addr:         "localhost:" + string(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	return err
}
