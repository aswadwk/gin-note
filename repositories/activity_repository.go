package repositories

import "aswadwk/models"

type ActivityRepository interface {
	Create(activity models.Activity) (models.Activity, error)
	FindAll() ([]models.Activity, error)
	FindByID(id string) (models.Activity, error)
	UpdateByID(id string, activity models.Activity) (models.Activity, error)
	DeleteByID(id string) (int64, error)
	
}