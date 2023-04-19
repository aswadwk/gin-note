package services

import (
	"aswadwk/dto"
	"aswadwk/models"
	"aswadwk/repositories"
	"aswadwk/services"

	"github.com/mashingan/smapping"
)

type TodoServiceImpl struct {
	todoRepo repositories.TodoRepository
}

func CreateTodoService(todoRepo repositories.TodoRepository) services.TodoService {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (service *TodoServiceImpl) Create(todo dto.TodoCreateDTO) (models.Todo, error) {
	todoCreate := models.Todo{}

	err := smapping.FillStruct(&todoCreate, smapping.MapFields(todo))

	if err != nil {
		return todoCreate, err
	}

	return service.todoRepo.Create(todoCreate)
}

func (service *TodoServiceImpl) FindAll(params dto.TodoSearchDTO) ([]models.Todo, error) {

	todoParams := dto.TodoSearchDTO{}

	err := smapping.FillStruct(&todoParams, smapping.MapFields(params))

	if err != nil {
		return nil, err
	}

	return service.todoRepo.FindAll(todoParams)
}

func (service *TodoServiceImpl) FindByID(id string) (models.Todo, error) {
	return service.todoRepo.FindByID(id)
}

func (service *TodoServiceImpl) UpdateByID(id string, todo dto.TodoUpdateByIDDTO) (models.Todo, error) {
	todoUpdate := models.Todo{}

	err := smapping.FillStruct(&todoUpdate, smapping.MapFields(todo))

	if err != nil {
		return todoUpdate, err
	}
	
	return service.todoRepo.UpdateByID(id, todoUpdate)
}

func (service *TodoServiceImpl) DeleteByID(id string) (bool, error) {
	return service.todoRepo.DeleteByID(id)
}
