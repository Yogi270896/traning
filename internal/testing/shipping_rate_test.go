package testing

import (
	"Traning/internal/confi"
	"Traning/internal/models"
	"Traning/internal/service"
	"fmt"
	"testing"
)

func TestGetrate(t *testing.T) {
	app := confi.NewConfig()

	req := models.Shipment{
		Fromaddr: models.Fromaddress{
			Pincode: "90001",
			City:    "California",
			State:   "CA",
			Country: "US",
		},
		Toaddr: models.Toaddress{
			Pincode: "90005",
			City:    "California",
			State:   "CA",
			Country: "US",
		},
		Parcel: models.Parcel{
			Weight: 10,
			Length: 2,
			Width:  4,
			Height: 1,
		},
	}
	s := service.NewService(app)
	res, err := s.GetRates(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
