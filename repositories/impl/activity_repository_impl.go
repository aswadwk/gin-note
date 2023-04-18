package repositories

import (
	"aswadwk/models"

	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func CreateActivityRepository(db *gorm.DB) *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{
		db: db,
	}
}

func (repository *ActivityRepositoryImpl) Create(activity models.Activity) (models.Activity, error) {
	err := repository.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repository *ActivityRepositoryImpl) FindAll() ([]models.Activity, error) {
	var activities []models.Activity

	err := repository.db.Find(&activities).Error

	if err != nil {
		return activities, err
	}

	return activities, nil
}

func (repository *ActivityRepositoryImpl) FindByID(id string) (models.Activity, error) {
	var activity models.Activity

	err := repository.db.Where("id = ?", id).First(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}
