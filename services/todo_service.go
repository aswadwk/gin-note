package services

import (
	"aswadwk/dto"
	"aswadwk/models"
)

type TodoService interface {
	Create(note dto.TodoCreateDTO) (models.Todo, error)
	
}