package services

import (
	"aswadwk/dto"
	"aswadwk/models"
)

type ActivityService interface {
	Create(activity dto.ActivityCreateDTO) (models.Activity, error)
	FindAll() ([]models.Activity, error)
	FindByID(id string) (models.Activity, error)
	UpdateByID(id string, activity dto.ActivityUpdateByIDDTO) (models.Activity, error)
	DeleteByID(id string) (int64, error)
}