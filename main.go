package main

import (
	"./check"
	"./crawl"
	"./db"
	"log"
	"strconv"
	"time"
)

var (
	CrawlerList []crawl.Spider
	FREQUENCY   = 10
	mysqlTable  *db.MyTable
	checkRule   check.CheckProxyInterface
)

func initMysql() {
	db.Refresh()
	mysqlTable = &db.MyTable{}
	mysqlTable.SetTableName("ip")
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	mysqlTable.SetColumnNames(name)
}

func Register() {

	CrawlerList = append(CrawlerList, &crawl.DataWu{})
	//CrawlerList = append(CrawlerList, &crawl.LLIP{})
}

func startSpider(spider crawl.Spider) {
	for {
		start := time.Now()
		nextRun := start.Add(time.Duration(FREQUENCY) * time.Second)
		for _, ipproxy := range spider.GetIpProxyList() {
			log.Printf("acquire ipproxy %s:%d %d", ipproxy.Ip, ipproxy.Port, ipproxy.Type)
			ipproxy.SetDBHelper(mysqlTable)
			exists, err := ipproxy.Exists()
			if err != nil {
				log.Println("check data errr ", err)
				continue
			}
			if exists {
				log.Println("this data exists ")
				continue
			}
			err = ipproxy.Insert(mysqlTable)
			if err != nil {
				log.Println("insert data errr ", err)
			}
		}
		time.Sleep(nextRun.Sub(time.Now()))
	}
}

func CrawlProccess() {
	for _, spider := range CrawlerList {
		go startSpider(spider)
	}
}

func CheckProxy() {
	for {
		start := time.Now()
		nextRun := start.Add(time.Duration(FREQUENCY) * time.Second)
		list := check.GetCheckIpProxy(mysqlTable)
		log.Println("check ip size ", len(list))
		for _, value := range list {
			go func(value map[string]interface{}) {
				//log.Println(value["ip"], value["port"])
				ip := value["ip"].(string)
				port := value["port"].(string)
				p, _ := strconv.Atoi(port)
				ipporxy := &crawl.IpProxy{Ip: ip, Port: p}
				ipporxy.SetDBHelper(mysqlTable)
				score := checkRule.CheckProxy(ipporxy)
				log.Println("check ", value["ip"], score)
				ipporxy.UpdateScore(score)
			}(value)
		}
		time.Sleep(nextRun.Sub(time.Now()))
	}
}

func init() {
	Register()
	initMysql()
	checkRule = &check.BaiduCheck{}
	log.Println("init......")
}

func main() {
	ch := make(chan int)
	log.Println("begin crawl ip proxy")
	log.Println("rule size ", len(CrawlerList))
	go CrawlProccess()
	log.Println("begin check ip proxy")
	go CheckProxy()
	<-ch
}
