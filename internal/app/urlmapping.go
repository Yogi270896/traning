package app

import (
	"Traning/internal/confi"
	"Traning/internal/controller"
	"log"
	"net/http"
)

func RunURL(app *confi.AppConfig) {
	log.Println("Initialise all handlers")
	shippingController := controller.NewController(app)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		a := "pong"
		w.Write([]byte(a))
	})
	http.HandleFunc("/getrates", shippingController.GetRates)

}
