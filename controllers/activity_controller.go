package controllers

import (
	"aswadwk/dto"
	"aswadwk/helpers"
	"aswadwk/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ActivityController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
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

	errDto := ctx.ShouldBind(&activityCreateDTO)
	if errDto != nil {
		res := helpers.ResponseError("Bad Request", "Failed to bind request", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if(activityCreateDTO.Email == "") {
		res := helpers.ResponseError("Bad Request", "Email cannot be null", gin.H{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if(activityCreateDTO.Title == "") {
		res := helpers.ResponseError("Bad Request", "Title cannot be null", gin.H{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	newTodo, err := c.activityService.Create(activityCreateDTO)

	if err != nil {
		res := helpers.ResponseError("Failed", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", newTodo)
	ctx.JSON(http.StatusCreated, res)
}

func (c *activityController) FindAll(ctx *gin.Context) {
	activities, err := c.activityService.FindAll()

	if err != nil {
		res := helpers.ResponseError("Server Error", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", activities)
	ctx.JSON(200, res)
}

func (c *activityController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	activity, err := c.activityService.FindByID(id)

	if err != nil {
		res := helpers.ResponseError("Not Found", "Activity with ID "+id+" not found", gin.H{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", activity)
	ctx.JSON(200, res)
}

func (c *activityController) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	activityUpdateDTO := dto.ActivityUpdateByIDDTO{}

	if err := ctx.ShouldBindWith(&activityUpdateDTO, binding.JSON); err != nil {
		res := helpers.ResponseError("Bad Request", err.Error(), gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedRows, err := c.activityService.UpdateByID(id, activityUpdateDTO)

	if err != nil {
		res := helpers.ResponseError("Server Error", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", updatedRows)
	ctx.JSON(200, res)
}

func (c *activityController) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.activityService.FindByID(id)

	if err != nil {
		res := helpers.ResponseError("Not Found", "Activity with ID "+id+" not found", gin.H{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	deleteActivity, err := c.activityService.DeleteByID(id)

	if err != nil {
		res := helpers.ResponseError("Server Error", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", deleteActivity)
	ctx.JSON(200, res)
}