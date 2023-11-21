package system

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/tb/task-logger/common-packages/database"
	"github.com/tb/task-logger/common-packages/messaging"
	_ "github.com/tb/task-logger/common-packages/messaging"
	"gopkg.in/mgo.v2"
	"gopkg.in/redis.v3"
	"log"
)

type TbApplicationContext struct {
	MongoDBSessionMap   map[string]*mgo.Session
	RedisClusterClient  *redis.ClusterClient
	RedisClient         *redis.Client
	RabbitMQConn        *amqp.Connection
	RabbitMqMQQTClient  MQTT.Client
	RedisClusterEnabled bool
}

var TbAppContext *TbApplicationContext

func init() {
	TbAppContext = new(TbApplicationContext)
	if viper.GetBool("database.redis.clusterEnabled") {
		TbAppContext.RedisClusterEnabled = true
	} else {
		TbAppContext.RedisClusterEnabled = false
	}
}

func PrepareApplicationContext() error {

	var err error

	TbAppContext.MongoDBSessionMap, err = db.InitMongoDbSession()

	if err != nil {
		log.Println("Can't connect to mongodb database")
		return err
	} else {
		log.Println("Mongodb started successfully!")
	}
	TbAppContext.RabbitMQConn, err = messaging.GetRabbitMQConnection()

	if err != nil {
		log.Println("Can't connect to rabbit mq")
		return err
	} else {
		log.Println("Rabbit MQ started successfully!")
	}


	return nil
}

func CloseRabbitMQConnection() {
	TbAppContext.RabbitMQConn.Close()
	log.Println("Bye Rabbit MQ!")
}

func CloseMongoDBConnection() {
	mongoSchemas := viper.GetStringMapString("mongodb.schema")

	for _, value := range mongoSchemas {
		mongodbSession := TbAppContext.MongoDBSessionMap[value]
		mongodbSession.Close()
	}
	log.Println("Bye MongoDB!")
}

func CloseDatabaseConnections() {
	//closing mongodb connection
	CloseMongoDBConnection()

}
