package services

import (
	"aswadwk/dto"
	"aswadwk/models"
	"aswadwk/repositories"
	"aswadwk/services"
)

type ActivityServiceImpl struct {
	activityRepo repositories.ActivityRepository
}

func CreateActivityService(activityRepo repositories.ActivityRepository) services.ActivityService {
	return &ActivityServiceImpl{
		activityRepo: activityRepo,
	}
}

func (service *ActivityServiceImpl) Create(activity dto.ActivityCreateDTO) (models.Activity, error) {
	activityCreate := models.Activity{}

	return service.activityRepo.Create(activityCreate)
}