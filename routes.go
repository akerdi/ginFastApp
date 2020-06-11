package ginFastApp

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type IRoute interface {
	GetPath() string
	GetMethod() string
	GetMiddles() []gin.HandlerFunc
}

func (app *App) AddRoutes(route IRoute)  {
	app.Routes = append(app.Routes, route)
}

func (app *App) applyRoutes(engine *gin.Engine)  {
	for _, route := range app.Routes {
		method := strings.ToUpper(route.GetMethod())
		if method == "GET" {
			engine.GET(route.GetPath(), route.GetMiddles()...)
		} else if method == "POST" {
			engine.POST(route.GetPath(), route.GetMiddles()...)
		} else if method == "PUT" {
			engine.PUT(route.GetPath(), route.GetMiddles()...)
		} else if method == "DELETE" {
			engine.DELETE(route.GetPath(), route.GetMiddles()...)
		} else if method == "OPTIONS" {
			engine.OPTIONS(route.GetPath(), route.GetMiddles()...)
		} else {
			log.Fatalf("Current path: [%s] do method [%s] but not support yet", route.GetPath(), method)
		}
	}
}