package check

import (
	"../crawl"
	"../db"
	"../utils"
	"../utils/surfer"
	"log"
	"strconv"
	"time"
)

type CheckProxyInterface interface {
	CheckProxy(ipproxy *crawl.IpProxy) bool
	//GetCheckIpProxy() []*crawl.IpProxy
}

func init() {
	db.Refresh()
}

type BaiduCheck struct {
}

func getCheckIpProxy(dbHelp db.DBInterface) (ipProxyList []map[string]interface{}) {
	lastCheckTime := time.Now().Unix() - 3600
	log.Println(lastCheckTime)
	sql := "select * from ip where lastCheckTime < ? and status=1"
	args := []interface{}{}
	args = append(args, lastCheckTime)
	ipProxyList, _ = dbHelp.Query(sql, args)
	return
}

func (self *BaiduCheck) CheckProxy(ipproxy *crawl.IpProxy) (result bool, score int) {
	result = true
	begin := time.Now().UnixNano()
	proxy := "http://" + ipproxy.Ip + ":" + strconv.Itoa(ipproxy.Port)
	request := &surfer.DefaultRequest{Url: "https://www.baidu.com", TryTimes: 1, EnableCookie: true, Proxy: proxy, DialTimeout: time.Second * 10}
	request.GetUrl()
	html, err := utils.GetHtml(request)
	end := time.Now().UnixNano()

	log.Printf("check proxy %s cost %ds\n", proxy, (end-begin)/1000000000.0)
	if err != nil {
		result = false
		return
	}
	//log.Println(html)
	if len(html) < 10 {
		result = false
	}
	return
}
