package main

import (
	"log"
	"net/http"
	"os"

	"github.com/template/config"
	"github.com/template/internal/server"
	"github.com/template/pkg/utils"
)

func main() {

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	serverApi := server.Run(&server.Server{
		Cfg: cfg,
		
	})

	if err := http.ListenAndServe(":8080", serverApi); err != nil {
		log.Fatalln("Error listen and serve: ", err)
	}
}
