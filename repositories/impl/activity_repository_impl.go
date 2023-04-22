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

func (repository *ActivityRepositoryImpl) UpdateByID(id string, activity models.Activity) (models.Activity, error) {
	var activityUpdate models.Activity

	result := repository.db.Model(&activityUpdate).Where("id = ?", id).Updates(&activity)

	if result.Error != nil {
		return activityUpdate, result.Error
	}

	if result.RowsAffected == 0 {
		return activityUpdate, gorm.ErrRecordNotFound
	}

	err := repository.db.First(&activityUpdate, id).Error
	if err != nil {
		return activityUpdate, err
	}

	return activityUpdate, nil
}

func (repository *ActivityRepositoryImpl) DeleteByID(id string) (int64, error) {
	db := repository.db.Where("id = ?", id).Delete(&models.Activity{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}