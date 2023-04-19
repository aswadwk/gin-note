package services

import (
	"aswadwk/dto"
	"aswadwk/models"
)

type TodoService interface {
	Create(note dto.TodoCreateDTO) (models.Todo, error)
	FindAll(params dto.TodoSearchDTO) ([]models.Todo, error)
	FindByID(id string) (models.Todo, error)
	UpdateByID(id string, todo dto.TodoUpdateByIDDTO) (models.Todo, error)
	DeleteByID(id string) (bool, error)
}