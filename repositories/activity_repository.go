package repositories

import "aswadwk/models"

type ActivityRepository interface {
	Create(todo models.Activity) (models.Activity, error)
	
}