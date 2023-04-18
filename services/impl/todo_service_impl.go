package services

import (
	"aswadwk/dto"
	"aswadwk/models"
	"aswadwk/repositories"
	"aswadwk/services"
)

type TodoServiceImpl struct {
	todoRepo repositories.TodoRepository
}

func CreateNoteService(todoRepo repositories.TodoRepository) services.TodoService {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (service *TodoServiceImpl) Create(todo dto.TodoCreateDTO) (models.Todo, error) {
	todoCreate := models.Todo{}

	return service.todoRepo.Create(todoCreate)
}
