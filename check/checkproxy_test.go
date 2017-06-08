package check

import (
	"../crawl"
	"../db"
	"log"
	//"reflect"
	"strconv"
	"testing"
)

func TestGetCheckIpProxy(*testing.T) {
	mysqlTable := &db.MyTable{}
	mysqlTable.SetTableName("ip")
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	mysqlTable.SetColumnNames(name)
	baiduCheck := BaiduCheck{}
	// log.Println(baiduCheck)
	list := getCheckIpProxy(mysqlTable)
	log.Println(len(list))
	ipporxy := &crawl.IpProxy{Ip: "123.59.188.13", Port: 8118}
	result, score := baiduCheck.CheckProxy(ipporxy)
	log.Println(result, score)
	for _, value := range list {
		ip := value["ip"].(string)
		port := value["port"].(string)
		p, _ := strconv.Atoi(port)
		ipporxy = &crawl.IpProxy{Ip: ip, Port: p}
		result, score = baiduCheck.CheckProxy(ipporxy)
		log.Println("check ", value["ip"], " result ", result, score)
		break
	}
}
