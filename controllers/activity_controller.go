package controllers

import (
	"aswadwk/dto"
	"aswadwk/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ActivityController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
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

	if err := ctx.ShouldBindWith(&activityCreateDTO, binding.JSON); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    fmt.Println(activityCreateDTO)

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

func (c *activityController) FindAll(ctx *gin.Context) {
	activities, err := c.activityService.FindAll()

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, activities)
}

func (c *activityController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	activity, err := c.activityService.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, activity)
}