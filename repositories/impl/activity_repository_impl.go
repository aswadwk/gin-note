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
