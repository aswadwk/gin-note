package repositories

import (
	"aswadwk/dto"
	"aswadwk/models"
	"fmt"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func CreateTodoRepository(db *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (repository *TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	err := repository.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	fmt.Println("todo created")
	fmt.Println(todo.Title)
	return todo, nil
}

func (repository *TodoRepositoryImpl) FindAll(params dto.TodoSearchDTO) ([]models.Todo, error) {
	var todos []models.Todo

	query := repository.db.Where("1 = 1")

	if params.ActivityGroupID != "" {
		query = query.Where("activity_group_id = ?", params.ActivityGroupID)
	}

	result := query.Find(&todos)

	if result.Error != nil {
		return todos, result.Error
	}

	return todos, nil
}

func (repository *TodoRepositoryImpl) FindByID(id string) (models.Todo, error) {
	var todo models.Todo

	err := repository.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) UpdateByID(id string, todo models.Todo) (models.Todo, error) {
	err := repository.db.Where("id = ?", id).Updates(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) DeleteByID(id string) (bool, error) {
	err := repository.db.Delete(&models.Todo{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}