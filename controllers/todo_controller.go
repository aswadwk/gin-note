package controllers

import (
	"aswadwk/dto"
	"aswadwk/helpers"
	"aswadwk/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TodoController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type todoController struct {
	// Service
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
	}
}

func (c *todoController) Create(ctx *gin.Context) {
	
	todoCreateDTO := dto.TodoCreateDTO{}

	errDto := ctx.BindJSON(&todoCreateDTO)

	if todoCreateDTO.ActivityGroupID == 0 {
		res := helpers.ResponseError("Bad Request", "activity_group_id cannot be null", gin.H{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if errDto != nil {
		res := helpers.ResponseError("Bad Request", "Failed to bind request", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if(todoCreateDTO.Title == "") {
		res := helpers.ResponseError("Bad Request", "title cannot be null", gin.H{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	newTodo, err := c.todoService.Create(todoCreateDTO)

	if err != nil {
		res := helpers.ResponseError("Failed", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", newTodo)
	ctx.JSON(http.StatusCreated, res)
}

func (c *todoController) FindAll(ctx *gin.Context) {

	activityGroupID := ctx.Query("activity_group_id")

	todoSearchDTO := dto.TodoSearchDTO{}

	if activityGroupID != "" {
		todoSearchDTO.ActivityGroupID = activityGroupID
	}

	todos, err := c.todoService.FindAll(todoSearchDTO)

	if err != nil {
		res := helpers.ResponseError("Failed", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", todos)
	ctx.JSON(http.StatusOK, res)
}

func (c *todoController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := c.todoService.FindByID(id)

	if err != nil {
		res := helpers.ResponseError("Not Found", "Todo with ID "+id+" Not Found", gin.H{})

		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", todo)
	ctx.JSON(http.StatusOK, res)
}

func (c *todoController) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	todoUpdateDTO := dto.TodoUpdateByIDDTO{}

	errDto := ctx.ShouldBind(&todoUpdateDTO)
	if errDto != nil {
		res := helpers.ResponseError("Bad Request", "Failed to bind request", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := c.todoService.FindByID(id)

	if err != nil {
		res := helpers.ResponseError("Not Found", "Todo with ID "+id+" Not Found", gin.H{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	todo, err := c.todoService.UpdateByID(id, todoUpdateDTO)

	if err != nil {
		res := helpers.ResponseError("Failed", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", todo)
	ctx.JSON(http.StatusOK, res)
}

func (c *todoController) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.todoService.FindByID(id)

	if err != nil {
		res := helpers.ResponseError("Not Found", "Todo with ID "+id+" Not Found", gin.H{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	c.todoService.DeleteByID(id)

	if err != nil {
		res := helpers.ResponseError("Failed", "Something went wrong", gin.H{})

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseSuccess("Success", "Success", gin.H{})
	ctx.JSON(http.StatusOK, res)
}