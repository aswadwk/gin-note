package repositories

import (
	"aswadwk/models"

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

	return todo, nil
}
