package controller

import (
	"Traning/internal/confi"
	"Traning/internal/models"
	"Traning/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (sc *ShippingController) GetRates(ctx *gin.Context) {
	log.Println("Shipping Contorller", "Get Rates", "Start")

	var reqModel models.Shipment
	if err := ctx.ShouldBindJSON(&reqModel); err != nil {
		ctx.JSON(402, err)
	}
	res, errs := sc.Service.GetRates(reqModel)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, errs.Error())
	}
	ctx.JSON(http.StatusAccepted, res)
	log.Println("Shipping Contorller", "Get Rates", "End")
}

// req, errs := io.ReadAll(r.Body)
// json.Unmarshal(req, &reqModel)
// if errs != nil {
// 	fmt.Println(http.StatusBadRequest)
// 	w.WriteHeader(http.StatusBadRequest)
// 	return
// }
// res, err := sc.Service.GetRates(reqModel)
// if err != nil {
// 	resErr := err.Error()
// 	res, _ := json.Marshal(resErr)
// 	w.WriteHeader(http.StatusBadRequest)
// 	w.Write(res)
// 	return
// }
// rs, _ := json.Marshal(res)
// w.WriteHeader(http.StatusAccepted)
// w.Write(rs)
