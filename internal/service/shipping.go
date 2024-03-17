package service

import (
	"Traning/internal/confi"
	"Traning/internal/helper"
	"Traning/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type ShippingService struct {
	AppConfig *confi.AppConfig
}

func NewService(app *confi.AppConfig) *ShippingService {
	return &ShippingService{
		AppConfig: app,
	}
}

func (s *ShippingService) GetRates(shipment models.Shipment) (map[string]interface{}, error) {
	m := "Shipping Service"
	log.Println(m, "Get Rates", "Start")
	var response map[string]interface{}
	var shipReq models.Ship
	fmt.Println(shipment)
	if shipment.Fromaddr.Pincode == "" {
		return response, errors.New("From Address Pincode Cannot be nil")
	}
	shipReq.Shipment = shipment

	reqBody, _ := json.Marshal(shipReq)
	log.Println("Shipping Service", "Get Rates", "URL", s.AppConfig.EasyPost.URL)
	log.Println("Shipping Service", "Get Rates", "BODY", string(reqBody))

	res, errs := helper.Send("POST", s.AppConfig.EasyPost.URL, shipReq, s.AppConfig.EasyPost.APIKEY, "GetRates")
	if errs != nil {
		log.Println(m, "Get rates", errs)
		return response, errs
	}
	json.Unmarshal(res, &response)
	log.Println(m, "Get Rates", "Response", response)
	log.Println(m, "Get Rates", "End")
	return response, nil
}
