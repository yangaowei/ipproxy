package main

import (
	"./crawl"
	"./db"
	"log"
	"time"
)

var (
	CrawlerList []crawl.Spider
	FREQUENCY   = 10
	mysqlTable  *db.MyTable
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

func init() {
	Register()
	initMysql()
	log.Println("init......")
}

func main() {
	ch := make(chan int)
	log.Println("begin crawl ip proxy")
	log.Println("rule size ", len(CrawlerList))
	CrawlProccess()
	<-ch
}
