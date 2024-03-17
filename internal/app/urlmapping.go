package app

import (
	"Traning/internal/confi"
	"Traning/internal/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func RunURL(app *confi.AppConfig) {
	log.Println("Initialise all handlers")
	shippingController := controller.NewController(app)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "ping")
	})

	api := router.Group("/v1")
	api.GET("/getrates", shippingController.GetRates)
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	a := "pong"
	// 	w.Write([]byte(a))
	// })
	// http.HandleFunc("/getrates", shippingController.GetRates)

}
