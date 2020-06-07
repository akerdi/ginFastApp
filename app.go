package ginFastApp

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const DRIVER = "mysql"
var db *gorm.DB

type IConfig interface {
	GetPort() int
	GetHost() string
	GetDEBUG() bool
	GetDB() map[string]interface{}
	GetRedis() map[string]interface{}
}

type App struct {
	Config *IConfig
	beforeMiddlewares []interface{}
	afterMiddlewares []interface{}
}

func New(config *IConfig)*App {
	return &App{Config:config}
}

func (app *App) addBeforeMiddleware(middleware interface {})  {
	app.beforeMiddlewares = append(app.beforeMiddlewares, middleware)
}
func (app *App) addAfterMiddleware(middleware interface {})  {
	app.afterMiddlewares = append(app.afterMiddlewares, middleware)
}

func (app *App) connectDB(cb func(db *gorm.DB, err error)) {
	dbMapper := (*app.Config).GetDB()
	if dbMapper == nil {
		cb(nil, errors.New("Config getDB fail"))
		return
	}
	fmt.Println("------", dbMapper)
	username, ok := dbMapper["DB_Username"].(string)
	if !ok {
		cb(nil, errors.New("DB_Username is not ok"))
		return
	}
	password, ok := dbMapper["DB_Password"].(string)
	if !ok {
		cb(nil, errors.New("DB_Password is not ok"))
		return
	}
	host, ok := dbMapper["DB_Host"].(string)
	if !ok {
		cb(nil, errors.New("DB_Host is not ok"))
		return
	}
	port, ok := dbMapper["DB_Port"].(int64)
	if !ok {
		cb(nil, errors.New("DB_Port is not ok"))
		return
	}
	name, ok := dbMapper["DB_Name"].(string)
	if !ok {
		cb(nil, errors.New("DB_Name is not ok"))
		return
	}
	fmt.Println("=-====", username, password, host, port, name)
	return
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, name)
	var err error
	db, err = gorm.Open(DRIVER, DSN)
	if err != nil {
		cb(nil, err)
		return
	}
	cb(db, nil)
}

// Start 服务器正式启动开始
func (app *App) Start() (*gin.Engine, error)  {
	engine := gin.Default()
	port := (*app.Config).GetPort()
	portStr := fmt.Sprintf("%d", port)
	err := engine.Run(portStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("server is starting now! port : ", port)
	return engine, nil
}
