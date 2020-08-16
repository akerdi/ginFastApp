package ginFastApp

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-redis/redis"
)

const DRIVER = "mysql"
var db *gorm.DB

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) GetClient() *redis.Client {
	return  r.client
}

func (app *App) ConnectDB(cb func(db *gorm.DB, err error)) {
	dbMapper := app.Config.GetDB()
	if dbMapper == nil {
		cb(nil, errors.New("Config getDB fail"))
		return
	}
	var username, password, host, name string
	var port int64
	for field, val := range dbMapper {
		if field == "username" {
			username = val.(string)
		} else if field == "password" {
			password = val.(string)
		} else if field == "host" {
			host = val.(string)
		} else if field == "name" {
			name = val.(string)
		} else if field == "port" {
			port = int64(val.(float64))
		}
	}
	fmt.Println("connect db env: ", username, password, host, port, name)
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, name)
	fmt.Println("connect db DSN", DSN)
	var err error
	db, err = gorm.Open(DRIVER, DSN)
	if err != nil {
		fmt.Println("--------connect db err::", err)
		cb(nil, err)
		return
	}
	cb(db, nil)
}

func (app *App) ConnectRedis(cb func(redisClient *RedisClient, err error))  {
	redisMapper := app.Config.GetRedis()
	if redisMapper == nil {
		cb(nil, errors.New("Config getDB fail"))
		return
	}
	var redisHost, redisPass string
	var redisPort int64
	var redisIndex int
	for field, val := range redisMapper {
		if field == "host" {
			redisHost = val.(string)
		} else if field == "pass" {
			redisPass = val.(string)
		} else if field == "port" {
			redisPort = int64(val.(float64))
		} else if field == "db_index" {
			redisIndex = int(val.(float64))
		} else {
			fmt.Printf("%s:%s not support", field, val)
		}
	}
	addr := fmt.Sprintf("%s:%d", redisHost, redisPort)
	redisOptions := redis.Options{
		Addr : addr,
		DB : redisIndex,
	}
	if redisPass != "" {
		redisOptions.Password = redisPass
	}
	client := redis.NewClient(&redisOptions)
	redisConnect := &RedisClient{
		client: client,
	}
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis client meet error : ", err)
		cb(nil, err)
		return
	}
	fmt.Println("redis client set up successfully! pong: ", pong)
	cb(redisConnect, nil)
}