package repositories

import (
	"aswadwk/dto"
	"aswadwk/models"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	FindAll(params dto.TodoSearchDTO) ([]models.Todo, error)
	FindByID(id string) (models.Todo, error)
	UpdateByID(id string, todo models.Todo) (models.Todo, error)
	DeleteByID(id string) (bool, error)
}