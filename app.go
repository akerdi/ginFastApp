package ginFastApp

import "github.com/gin-gonic/gin"

func InitApp() int64 {
	return 89
}

func AhuangName() string {
	return  "ahuang"
}

func Init() *gin.Engine {
	g := gin.New()
	return  g
}

func Start(g *gin.Engine, listenAddress string) error {
	return g.Run(listenAddress)
}