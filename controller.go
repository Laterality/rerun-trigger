package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateController interface {
	HandleUpdate(c *gin.Context)
}

type updateControllerImpl struct {
	service ContainerService
}

func (ctrl *updateControllerImpl) HandleUpdate(c *gin.Context) {
	id, err := ctrl.service.Update(c.Query("tag"))
	defer (func() {
		if p := recover(); p != nil || err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": p,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"containerId": id,
			})
		}
	})()
}

func NewUpdateController(service ContainerService) UpdateController {
	controller := new(updateControllerImpl)
	controller.service = service
	return controller
}
