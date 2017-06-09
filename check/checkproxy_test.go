package check

import (
	"../crawl"
	"../db"
	"log"
	//"reflect"
	"strconv"
	"testing"
)

func check(value map[string]interface{}) {
	log.Println(value["ip"], value["port"])
}

func TestGetCheckIpProxy(*testing.T) {
	mysqlTable := &db.MyTable{}
	mysqlTable.SetTableName("ip")
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	mysqlTable.SetColumnNames(name)
	baiduCheck := BaiduCheck{}
	// log.Println(baiduCheck)
	list := GetCheckIpProxy(mysqlTable)
	log.Println(len(list))
	// ipporxy := &crawl.IpProxy{Ip: "123.59.188.13", Port: 8118}
	// score := baiduCheck.CheckProxy(ipporxy)
	// log.Println(score)
	ch := make(chan int64)
	for _, value := range list {
		go func(value map[string]interface{}) {
			//log.Println(value["ip"], value["port"])
			ip := value["ip"].(string)
			port := value["port"].(string)
			p, _ := strconv.Atoi(port)
			ipporxy := &crawl.IpProxy{Ip: ip, Port: p}
			ipporxy.SetDBHelper(mysqlTable)
			score := baiduCheck.CheckProxy(ipporxy)
			log.Println("check ", value["ip"], score)
			ipporxy.UpdateScore(score)
			ch <- score
		}(value)
	}
	for i := 0; i < len(list); i++ {
		<-ch
	}
}
