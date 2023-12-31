package services

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/tb/task-logger/common-packages/system"
	cache "github.com/tb/task-logger/taskapp/cache"
	"github.com/tb/task-logger/taskapp/models"
	"github.com/tb/task-logger/taskapp/validations"

	"gopkg.in/mgo.v2/bson"
	"time"
)
type TaskService interface {
	TaskDetails(logger *logrus.Entry, emailId string, recordId string) ([]byte, error)
}


func TaskDetails(logger *logrus.Entry, emailId string, recordId string) ([]byte, error) {

	isValid := Validationspackage.ValidateEmail(emailId)
	if !isValid {
		return nil, system.InvalidEmailErr
	}
	if recordId == "" {
		return nil, system.NoRecordIdErr
	}
	var taskDetails *models.Task
	if taskDetails == nil {
		
		collectionName := system.TASKS_COLLECTION
		databaseName := system.GetDatabaseName(collectionName)
		sessionDb := system.TbAppContext.MongoDBSessionMap[databaseName].Clone()
		defer sessionDb.Close()
		collection := sessionDb.DB(databaseName).C(collectionName)

		err := collection.Find(bson.M{"emailId": emailId, "_id": recordId}).One(&taskDetails)
		if err != nil {
			logger.Error(err)
			if err.Error() == system.NotFoundErr.Error() {
				return nil, system.InvalidRecordId
			} else {
				return nil, err
			}
		}

		
	}
	completedOnDate := ""
	if taskDetails.Status != system.TASK_STATUS_PENDING && taskDetails.CompletedOn.Format("2006-01-02") != "0001-01-01" {
		completedOnDate = taskDetails.CompletedOn.Format("2006-01-02")
	}
	response := make(map[string]interface{})
	response["task_detail"] = map[string]string{
		"id":              taskDetails.Id,
		"title":           taskDetails.Title,
		"scheduledOn":     taskDetails.ScheduledOn,
		"description":     taskDetails.Description,
		"emailId":         taskDetails.EmailId,
		"status":          taskDetails.Status,
		"createdOnDate":   taskDetails.CreatedOn.Format("2006-01-02"),
		"completedOnDate": completedOnDate,
	}
	finalResponse, _ := json.Marshal(response)
	return finalResponse, nil
}

