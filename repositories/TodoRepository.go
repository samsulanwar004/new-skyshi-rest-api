package repositories

import (
	"skyshi-rest-api/models"

	"gorm.io/gorm"
)

type (
	HandlerTodo struct {
		db *gorm.DB
	}
	TodoRepository interface {
		GetAllTodo(models.Todo) ([]models.Todo, error)
		CreateTodo(models.Todo) (models.Todo, error)
		GetTodo(models.Todo) (models.Todo, error)
		UpdateTodo(models.Todo) (models.Todo, error)
		DeleteTodo(models.Todo) (models.TodoNull, error)
	}
)

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return HandlerTodo{
		db: db,
	}
}

func (h HandlerTodo) GetAllTodo(req models.Todo) ([]models.Todo, error) {

	var err error
	var result []models.Todo
	query := h.db

	if req.ActivityGroupID != 0 {
		query = query.Where("activity_group_id = ?", req.ActivityGroupID)
	}

	if err := query.Find(&result).Error; err != nil {
		return []models.Todo{}, err
	}
	return result, err
}

func (h HandlerTodo) CreateTodo(req models.Todo) (models.Todo, error) {

	var err error
	err = h.db.Create(&req).Error
	if err != nil {
		return models.Todo{}, err
	}
	return req, nil
}

func (h HandlerTodo) GetTodo(req models.Todo) (models.Todo, error) {
	var err error
	err = h.db.Model(&req).Where("id = ?", req.ID).Take(&req).Error

	if err != nil {
		return models.Todo{}, err
	}
	return req, err
}

func (h HandlerTodo) UpdateTodo(req models.Todo) (models.Todo, error) {
	var err error
	err = h.db.Model(&req).Where("id = ?", req.ID).Updates(&req).Take(&req).Error

	if err != nil {
		return models.Todo{}, err
	}
	return req, err
}

func (h HandlerTodo) DeleteTodo(req models.Todo) (models.TodoNull, error) {

	db := h.db.Model(&req).Where("id = ?", req.ID).Take(&req).Delete(&req)

	if db.Error != nil {
		return models.TodoNull{}, db.Error
	}
	return models.TodoNull{}, nil
}
