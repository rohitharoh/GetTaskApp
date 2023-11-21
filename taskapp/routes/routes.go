package routes

import (
	"github.com/tb/task-logger/common-packages/system"
	"github.com/tb/task-logger/taskapp/controllers/api"
	"github.com/zenazn/goji"
)

func PrepareRoutes(application *system.Application) {

	//task logger


	goji.Post("/application/service/task/details", application.Route(&api.Controller{}, "TaskDetails", false, nil))

}
