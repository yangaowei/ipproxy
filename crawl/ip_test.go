package crawl

import (
	//"../db"
	"log"
	"testing"
)

func TestGetIpProxyList(t *testing.T) {
	//llip := IP181{"http://www.66ip.cn/", []string{"http://www.66ip.cn/areaindex_1/index.html"}, "高匿代理"}
	ip181 := IP181{"http://www.ip181.com/", []string{"http://www.ip181.com/"}, "高匿代理"}
	list := ip181.GetIpProxyList()
	log.Println("------------------------------------------------------------------------")
	for _, v := range list {
		log.Println(v.Ip, v.Port, v.Type, v.Regin)
	}
}
