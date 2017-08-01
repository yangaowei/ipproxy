package web

import (
	//"../utils"
	"../db"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"time"
)

var (
	mysqlTable *db.MyTable
)

func init() {
	db.Refresh()
	mysqlTable = &db.MyTable{}
	mysqlTable.SetTableName("ip")
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	mysqlTable.SetColumnNames(name)
}

func test(c *gin.Context) {
	logs.Log.Informational("start:", time.Now())
	//time.Sleep(5 * time.Second)
	logs.Log.Informational("end:", time.Now())
	c.String(http.StatusOK, "Hello World!")
}

func proxy(c *gin.Context) {
	size := c.DefaultQuery("size", "1")
	dbHelp := mysqlTable
	lastCheckTime := time.Now().Unix() - 10
	sql := "select * from ip where lastCheckTime < ? and status=1 limit ?"
	args := []interface{}{}
	args = append(args, lastCheckTime, size)
	ipProxyList, _ := dbHelp.Query(sql, args)
	logs.Log.Debug("%d", len(ipProxyList))
	result := make(map[string]interface{})
	ips := []string{}
	for _, p := range ipProxyList {
		ips = append(ips, fmt.Sprintf("%s:%s", p["ip"], p["port"]))
	}
	result["msg"] = "success"
	result["ips"] = ips
	c.JSON(http.StatusOK, result)
}

func Run() {
	router := gin.Default()
	router.GET("/", test)
	router.GET("/proxy", proxy)
	router.StaticFile("/favicon.ico", "./web/resources/favicon.ico")
	logs.Log.Debug("start web server on port %d", 8001)
	router.Run(":8001")

}
