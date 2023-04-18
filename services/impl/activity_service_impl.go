package services

import (
	"aswadwk/dto"
	"aswadwk/models"
	"aswadwk/repositories"
	"aswadwk/services"
	"fmt"

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

	err := smapping.FillStruct(&activityCreate, smapping.MapFields(&activity))

	if err != nil {
		return activityCreate, err
	}

	fmt.Println("Activity Create")
	fmt.Println(activityCreate)

	return service.activityRepo.Create(activityCreate)
}