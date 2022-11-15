package services

import (
	"skyshi-rest-api/models"
	"skyshi-rest-api/repositories"
)

type (
	TodoService interface {
		GetAllTodo(models.Todo) ([]models.Todo, error)
		CreateTodo(models.Todo) (models.Todo, error)
		GetTodo(models.Todo) (models.Todo, error)
		UpdateTodo(models.Todo) (models.Todo, error)
		DeleteTodo(models.Todo) (models.TodoNull, error)
	}

	todoService struct {
		todoRepository repositories.TodoRepository
	}
)

func NewTodoService(_s repositories.TodoRepository) TodoService {
	return todoService{
		todoRepository: _s,
	}
}

func (_s todoService) GetAllTodo(c models.Todo) ([]models.Todo, error) {
	return _s.todoRepository.GetAllTodo(c)
}

func (_s todoService) GetTodo(c models.Todo) (models.Todo, error) {
	return _s.todoRepository.GetTodo(c)
}

func (_s todoService) CreateTodo(c models.Todo) (models.Todo, error) {
	return _s.todoRepository.CreateTodo(c)
}

func (_s todoService) UpdateTodo(c models.Todo) (models.Todo, error) {
	return _s.todoRepository.UpdateTodo(c)
}

func (_s todoService) DeleteTodo(c models.Todo) (models.TodoNull, error) {
	return _s.todoRepository.DeleteTodo(c)
}
