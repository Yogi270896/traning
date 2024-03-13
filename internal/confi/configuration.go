package confi

type AppConfig struct {
	Server   server
	EasyPost easypost
}

type server struct {
	Port string
}
type easypost struct {
	APIKEY string
	URL    string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Server: server{
			Port: ":8080",
		},
		EasyPost: easypost{
			APIKEY: "EZTK1cd698cd9f2446d18c0236b493ef6bfcggPp6PaZDxOdr4Jujsh2ZA",
			URL:    "https://api.easypost.com/beta/rates",
		},
	}
}
