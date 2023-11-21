package main

import (
	"flag"

	"github.com/spf13/viper"

	"github.com/tb/task-logger/common-packages/conf"
	"github.com/tb/task-logger/common-packages/log"
	"github.com/tb/task-logger/common-packages/system"
	"github.com/tb/task-logger/taskapp/routes"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
)

func main() {

	//Load configuration fileexport
	conf.LoadConfigFile()

	var application = &system.Application{}

	//Preparing application context, setting mongodb connections etc
	system.PrepareApplicationContext()

	//Initializing log settings
	tblog.InitLogger()

	//Apply authentication filter
	goji.Use(application.ApplyAuth)

	//Prepare routes
	routes.PrepareRoutes(application)

	//Will be called when server is shutdown
	graceful.PostHook(func() {
		//closing all database and other middleware connections
		system.CloseDatabaseConnections()

	})
	// SetV1 sets api routing ver1
//	http.ListenAndServe(":8080", nil)
	//Setting server address
	flag.Set("bind", viper.GetString("apps.taskapp.address"))

	goji.Serve()

}
