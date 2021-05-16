package usecase

import (
	"github.com/paypay3/go-echo-gorm-sample/domain/model"
	"github.com/paypay3/go-echo-gorm-sample/domain/repository"
)

type TaskUsecase interface {
	Create(title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepository: taskRepo}
}

func (u *taskUsecase) Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := u.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (u *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := u.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (u *taskUsecase) Update(id int, title, content string) (*model.Task, error) {
	targetTask, err := u.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err = targetTask.Set(title, content); err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepository.Update(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskUsecase) Delete(id int) error {
	task, err := u.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	if err = u.taskRepository.Delete(task); err != nil {
		return err
	}

	return nil
}
