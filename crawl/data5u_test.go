package crawl

import (
	//"../db"
	"log"
	"testing"
)

func TestDataGetIpProxyList(t *testing.T) {
	llip := DataWu{LLIP{"http://www.data5u.com/", []string{"http://www.data5u.com/"}, "高匿代理"}}

	log.Println(llip.GetIpProxyList())
}
