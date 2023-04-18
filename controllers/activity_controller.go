package controllers

import (
	"aswadwk/dto"
	"aswadwk/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ActivityController interface {
	Create(ctx *gin.Context)
}

type activityController struct {
	// Service
	activityService services.ActivityService
}

func NewActivityController(activityService services.ActivityService) ActivityController {
	return &activityController{
		activityService: activityService,
	}
}

func (c *activityController) Create(ctx *gin.Context) {
	activityCreateDTO := dto.ActivityCreateDTO{}
	ctx.ShouldBind(&activityCreateDTO)

	fmt.Println(activityCreateDTO)

	todo, err := c.activityService.Create(activityCreateDTO)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, todo)
}