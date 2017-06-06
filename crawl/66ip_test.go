package crawl

import (
	//"../db"
	"log"
	"testing"
)

func TestGetIpProxyList(t *testing.T) {
	llip := LLIP{"http://www.66ip.cn/", []string{"http://www.66ip.cn/areaindex_1/index.html"}, "高匿代理"}

	log.Println(llip.GetIpProxyList())
}
