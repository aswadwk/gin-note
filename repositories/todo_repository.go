package repositories

import "aswadwk/models"

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	
}