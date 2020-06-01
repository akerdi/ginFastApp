package ginFastApp

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const DRIVER = "mysql"
var db *gorm.DB

type App struct {
	Config map[string]interface{}
	beforeMiddlewares []interface{}
	afterMiddlewares []interface{}
}

func New(config map[string]interface {})*App {
	return &App{Config:config}
}

func (app *App) addBeforeMiddleware(middleware interface {})  {
	app.beforeMiddlewares = append(app.beforeMiddlewares, middleware)
}
func (app *App) addAfterMiddleware(middleware interface {})  {
	app.afterMiddlewares = append(app.afterMiddlewares, middleware)
}

func (app *App) connectDB(cb func(db *gorm.DB, err error)) {
	username, ok := app.Config["DB_Username"].(string)
	if !ok {
		cb(nil, errors.New("DB_Username is not ok"))
		return
	}
	password, ok := app.Config["DB_Password"].(string)
	if !ok {
		cb(nil, errors.New("DB_Password is not ok"))
		return
	}
	host, ok := app.Config["DB_Host"].(string)
	if !ok {
		cb(nil, errors.New("DB_Host is not ok"))
		return
	}
	port, ok := app.Config["DB_Port"].(int64)
	if !ok {
		cb(nil, errors.New("DB_Port is not ok"))
		return
	}
	name, ok := app.Config["DB_Name"].(string)
	if !ok {
		cb(nil, errors.New("DB_Name is not ok"))
		return
	}
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
	port, ok := app.Config["Port"].(int64)
	if !ok {
		return nil, errors.New("Config Port is invalid. Should be int64")
	}
	portStr := fmt.Sprintf("%d", port)
	err := engine.Run(portStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("server is starting now! port : ", app.Config["Port"])
	return engine, nil
}
