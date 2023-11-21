package api

import (
	"encoding/json"

	_ "fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tb/task-logger/common-packages/system"
	
	"github.com/tb/task-logger/taskapp/models"
	"github.com/tb/task-logger/taskapp/services"
	_ "github.com/tb/task-logger/taskapp/validations"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
)

type Controller struct {
	TaskController
}



type TaskController interface {
	TaskDetails(c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error)
	
}


func (controller *Controller) TaskDetails( c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	var taskParam map[string]string
	err := decoder.Decode(&taskParam)
	if err != nil {
		logger.Error(err)
		return nil, system.InvalidPayloadError
	}

	recordId, keyExists := taskParam["recordId"]
	if !keyExists {
		return nil, system.InvalidPayloadError
	}
	response, err := services.TaskDetails(logger, c.Env["emailId"].(string), recordId)
	return response, err
}

