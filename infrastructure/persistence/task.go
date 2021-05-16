package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/paypay3/go-echo-gorm-sample/domain/model"
	"github.com/paypay3/go-echo-gorm-sample/domain/repository"
)

type taskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

func (r *taskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := r.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := r.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := r.Conn.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) Delete(task *model.Task) error {
	if err := r.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
