package services

import (
	"aswadwk/dto"
	"aswadwk/models"
	"aswadwk/repositories"
	"aswadwk/services"

	"github.com/mashingan/smapping"
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

	err := smapping.FillStruct(&activityCreate, smapping.MapFields(activity))

	if err != nil {
		return activityCreate, err
	}

	return service.activityRepo.Create(activityCreate)
}

func (service *ActivityServiceImpl) FindAll() ([]models.Activity, error) {
	return service.activityRepo.FindAll()
}

func (service *ActivityServiceImpl) FindByID(id string) (models.Activity, error) {
	return service.activityRepo.FindByID(id)
}

func (service *ActivityServiceImpl) UpdateByID(id string, activity dto.ActivityUpdateByIDDTO) (models.Activity, error) {
	activityUpdate := models.Activity{}

	err := smapping.FillStruct(&activityUpdate, smapping.MapFields(activity))

	if err != nil {
		return activityUpdate, err
	}

	return service.activityRepo.UpdateByID(id, activityUpdate)
}

func (service *ActivityServiceImpl) DeleteByID(id string) (int64, error) {
	return service.activityRepo.DeleteByID(id)
}