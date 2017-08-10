package web

import (
	//"../utils"
	"../db"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/yangaowei/gologs"
	"math/rand"
	"net/http"
	"strconv"
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
	size, _ := strconv.ParseInt(c.DefaultQuery("size", "1"), 10, 64)
	dbHelp := mysqlTable
	lastCheckTime := time.Now().Unix() - 10
	sql := "select * from ip where lastCheckTime < ? and status=1"
	args := []interface{}{}
	args = append(args, lastCheckTime)
	ipProxyList, _ := dbHelp.Query(sql, args)
	result := make(map[string]interface{})
	ips := []string{}
	for _, p := range ipProxyList {
		ips = append(ips, fmt.Sprintf("%s:%s", p["ip"], p["port"]))
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := int64(len(ips))
	//logs.Log.Debug("%s %s", r, x)
	if size < x {
		start := r.Int63n(x - size)
		ips = ips[start : start+size]
		logs.Log.Debug("start %d", start)
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
