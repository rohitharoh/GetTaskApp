package Validationspackage

import (
	"github.com/sirupsen/logrus"
	"github.com/tb/task-logger/taskapp/models"
	"regexp"

	"github.com/tb/task-logger/common-packages/system"

	_ "flag"
	"strconv"
	"strings"
	"time"
)

func ValidateEmail(email string) bool {
		Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}


func ValidateName(name string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z\s?\.]+$`)
	return Re.MatchString(name)
}

func ValidateDateFormat(dateToValidate string) bool {

	if len(dateToValidate) == 16 {
		date_time_arr := strings.Split(dateToValidate, " ")
		if len(date_time_arr) == 2 {
			datePart := date_time_arr[0]
			timePart := date_time_arr[1]
			if len(datePart) != 10 || len(timePart) != 5 {
				return false
			}else {
				dateParts := strings.Split(datePart, "-")
				if len(dateParts) != 3 {
					return false
				}
				timeParts := strings.Split(timePart, ":")
				if len(timeParts) != 2 {
					return false
				}
				if len(dateParts[0]) != 4 || len(dateParts[1]) != 2 || len(dateParts[2]) != 2 || len(timeParts[0]) != 2 || len(timeParts[1]) != 2 {
					return false
				}

				providedYear, err := strconv.Atoi(dateParts[0])
				if err != nil {
					return false
				}
				if providedYear < 1970 {
					return false
				}

				providedMonth, err := strconv.Atoi(dateParts[1])
				if err != nil {
					return false
				}
				if (providedMonth < 00 || providedMonth > 12) {
					return false
				}

				providedDay, err := strconv.Atoi(dateParts[2])
				if err != nil {
					return false
				}

				if (providedDay < 00 || providedDay > 31 ) {
					return false
				}
				providedHour, err := strconv.Atoi(timeParts[0])
				if err != nil {
					return false
				}
				if (providedHour < 00 || providedHour > 23 ) {
					return false
				}

				providedMin, err := strconv.Atoi(timeParts[1])
				if err != nil {
					return false
				}
				if (providedMin < 00 || providedMin > 59 ) {
					return false
				}
				dateParseFormat := datePart+"T"+timePart+":00Z"
				_, err = time.Parse(time.RFC3339,dateParseFormat);
				if err != nil {
					return false
				}
				return true
			}


		} else {
			return false
		}
	}else {
		return false
	}






}


