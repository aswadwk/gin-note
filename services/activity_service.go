package services

import (
	"aswadwk/dto"
	"aswadwk/models"
)

type ActivityService interface {
	Create(activity dto.ActivityCreateDTO) (models.Activity, error)
	
}