package repositories

import "aswadwk/models"

type ActivityRepository interface {
	Create(todo models.Activity) (models.Activity, error)
	FindAll() ([]models.Activity, error)
	FindByID(id string) (models.Activity, error)
	
	
}