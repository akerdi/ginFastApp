package ginFastApp

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/shaohung001/ginFastApp/middlewares"
)

func (app *App) addBeforeMiddleware(middleware gin.HandlerFunc)  {
	app.beforeMiddlewares = append(app.beforeMiddlewares, middleware)
}
func (app *App) addAfterMiddleware(middleware gin.HandlerFunc)  {
	app.afterMiddlewares = append(app.afterMiddlewares, middleware)
}

func (app *App) engineAssembleMiddlewares(engine *gin.Engine)  {
	engine.StaticFile("/favicon.icon", "./public/static/img/favicon.icon")
	engine.Static("/assets", "public/assets")
	// engine add beforeMiddleware
	for _, middleware := range app.beforeMiddlewares {
		engine.Use(middleware)
	}
	// engine add ginFast default middleware
	engine.Use(middlewares.Logger())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// engine add afterMiddleware
	for _, middleware := range app.afterMiddlewares {
		engine.Use(middleware)
	}
	
	// engine add routes
	app.applyRoutes(engine)
	
	
	engine.NoRoute(middlewares.ReturnPublic())
	engine.Use(middlewares.NotRouteResponse())
	// engine support public
	engine.Use(static.Serve("/", static.LocalFile("./public", true)))
}