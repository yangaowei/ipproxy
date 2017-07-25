package web

import (
	//"../utils"
	"github.com/gin-gonic/gin"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"time"
)

func test(c *gin.Context) {
	logs.Log.Informational("start:", time.Now())
	//time.Sleep(5 * time.Second)
	logs.Log.Informational("end:", time.Now())
	c.String(http.StatusOK, "Hello World!")
}

func Run() {
	router := gin.Default()
	router.GET("/", test)
	router.StaticFile("/favicon.ico", "./web/resources/favicon.ico")
	logs.Log.Debug("start web server on port %d", 8001)
	router.Run(":8001")

}
