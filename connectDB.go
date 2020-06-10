package ginFastApp

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DRIVER = "mysql"
var db *gorm.DB

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