package services

import (
	"skyshi-rest-api/models"
	"skyshi-rest-api/repositories"
)

type (
	ActivityService interface {
		GetAllActivity(models.Activity) ([]models.Activity, error)
		CreateActivity(models.Activity) (models.Activity, error)
		GetActivity(models.Activity) (models.Activity, error)
		UpdateActivity(models.Activity) (models.Activity, error)
		DeleteActivity(models.Activity) (models.ActivityNull, error)
	}

	activityService struct {
		activityRepository repositories.ActivityRepository
	}
)

func NewActivityService(_s repositories.ActivityRepository) ActivityService {
	return activityService{
		activityRepository: _s,
	}
}

func (_s activityService) GetAllActivity(c models.Activity) ([]models.Activity, error) {
	return _s.activityRepository.GetAllActivity(c)
}

func (_s activityService) GetActivity(c models.Activity) (models.Activity, error) {
	return _s.activityRepository.GetActivity(c)
}

func (_s activityService) CreateActivity(c models.Activity) (models.Activity, error) {
	return _s.activityRepository.CreateActivity(c)
}

func (_s activityService) UpdateActivity(c models.Activity) (models.Activity, error) {
	return _s.activityRepository.UpdateActivity(c)
}

func (_s activityService) DeleteActivity(c models.Activity) (models.ActivityNull, error) {
	return _s.activityRepository.DeleteActivity(c)
}
