package crawl

import (
	//"../db"
	"log"
	"testing"
)

func TestDataGetIpProxyList(t *testing.T) {
	llip := DataWu{LLIP{Index: "http://www.data5u.com/"}}

	log.Println(llip.GetIpProxyList())
}
