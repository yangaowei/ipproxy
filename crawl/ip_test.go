package crawl

import (
	//"../db"
	"log"
	"testing"
)

func TestGetIpProxyList(t *testing.T) {
	//llip := IP181{"http://www.66ip.cn/", []string{"http://www.66ip.cn/areaindex_1/index.html"}, "高匿代理"}
	// ip181 := IP181{"http://www.ip181.com/", []string{"http://www.ip181.com/"}, "高匿代理"}
	kdl := KuaiDaiLi{"http://www.kuaidaili.com/", []string{"http://www.kuaidaili.com/free/inha/1/"}, "高匿代理"}
	list := kdl.GetIpProxyList()
	log.Println("------------------------------------------------------------------------")
	for _, v := range list {
		log.Println(v.Ip, v.Port, v.Type, v.Regin)
	}
}
