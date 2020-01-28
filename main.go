package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	app := gin.Default()
	service, err := NewContainerService(NewEnvClientWrapper(), "./config.yml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	controller := NewUpdateController(service)
	app.PUT("/update", controller.HandleUpdate)
	app.Run("0.0.0.0:8080")
}
