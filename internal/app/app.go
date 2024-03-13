package app

import (
	"Traning/internal/confi"
	"log"
	"net/http"
)

func StartApplication(app *confi.AppConfig) {

	RunURL(app)
	log.Println("It is Listening to port 8080")
	err := http.ListenAndServe(app.Server.Port, nil)
	if err != nil {
		log.Fatal("Can't able to start the server", err)

	}
}
