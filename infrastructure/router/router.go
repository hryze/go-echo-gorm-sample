package router

import (
	"github.com/labstack/echo"

	"github.com/paypay3/go-echo-gorm-sample/infrastructure/persistence"
	"github.com/paypay3/go-echo-gorm-sample/infrastructure/persistence/rdb"
	"github.com/paypay3/go-echo-gorm-sample/interface/handler"
	"github.com/paypay3/go-echo-gorm-sample/usecase"
)

func Run() error {
	taskRepository := persistence.NewTaskRepository(rdb.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()

	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

	if err := e.Start(":8090"); err != nil {
		return err
	}

	return nil
}
