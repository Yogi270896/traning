package service

import (
	"Traning/internal/confi"
	"Traning/internal/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
	log.Println("Shipping Service", "Get Rates", "Start")
	var res map[string]interface{}
	var shipReq models.Ship
	fmt.Println(shipment)
	if shipment.Fromaddr.Pincode == "" {
		return res, errors.New("From Address Pincode Cannot be nil")
	}
	shipReq.Shipment = shipment

	reqBody, _ := json.Marshal(shipReq)
	log.Println("Shipping Service", "Get Rates", "URL", s.AppConfig.EasyPost.URL)
	log.Println("Shipping Service", "Get Rates", "BODY", string(reqBody))

	req, _ := http.NewRequest("POST", s.AppConfig.EasyPost.URL, bytes.NewBuffer(reqBody))
	authEncoded := base64.StdEncoding.EncodeToString([]byte(s.AppConfig.EasyPost.APIKEY))

	req.Header.Set("Authorization", "Basic "+authEncoded)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rawRes, _ := client.Do(req)
	resBody, resErr := io.ReadAll(rawRes.Body)
	//fmt.Println(string(resBody))
	if resErr != nil {
		return res, resErr
	}
	defer rawRes.Body.Close()

	json.Unmarshal(resBody, &res)
	log.Println("Shipping Service", "Get Rates", "Response", res)
	log.Println("Shipping Service", "Get Rates", "End")
	return res, nil
}
