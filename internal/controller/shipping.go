package controller

import (
	"Traning/internal/confi"
	"Traning/internal/models"
	"Traning/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ShippingController struct {
	AppConfig *confi.AppConfig
	Service   service.ShippingService
}

func NewController(app *confi.AppConfig) *ShippingController {
	return &ShippingController{
		AppConfig: app,
		Service:   *service.NewService(app),
	}
}

func (sc *ShippingController) GetRates(w http.ResponseWriter, r *http.Request) {
	var reqModel models.Shipment
	log.Println("Shipping Contorller", "Get Rates", "Start")
	req, errs := io.ReadAll(r.Body)
	json.Unmarshal(req, &reqModel)
	if errs != nil {
		fmt.Println(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := sc.Service.GetRates(reqModel)
	if err != nil {
		resErr := err.Error()
		res, _ := json.Marshal(resErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	rs, _ := json.Marshal(res)
	w.WriteHeader(http.StatusAccepted)
	w.Write(rs)
	log.Println("Shipping Contorller", "Get Rates", "End")
}
