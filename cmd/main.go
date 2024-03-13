package main

import (
	"Traning/internal/app"
	"Traning/internal/confi"
)

var (
	appConf *confi.AppConfig
)

func init() {
	appConf = confi.NewConfig()
}
func main() {
	app.StartApplication(appConf)
}
