package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	configPath := "./config.yml"
	config, err := NewConfigFromFile(configPath)
	app := gin.Default()
	service, err := NewContainerService(NewEnvClientWrapper(), configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	controller := NewUpdateController(service)
	app.PUT("/update", controller.HandleUpdate)
	app.Run("0.0.0.0:" + strconv.Itoa(config.Server.Port))
}
