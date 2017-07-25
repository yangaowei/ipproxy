package main

import (
	"./check"
	"./crawl"
	"./db"
	//"./web"
	//"flag"
	logs "github.com/yangaowei/gologs"
	"log"
	"strconv"
	"sync"
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
			log.Printf("acquire ipproxy %s:%d %s", ipproxy.Ip, ipproxy.Port, crawl.IpType[ipproxy.Type])
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
			score := checkRule.CheckProxy(ipproxy)
			if score < 0 {
				logs.Log.Debug("this proxy is not available %s:%d", ipproxy.Ip, ipproxy.Port)
				continue
			}
			err = ipproxy.Insert(mysqlTable)
			if err != nil {
				logs.Log.Debug("insert data errr %v", err)
			} else {
				logs.Log.Debug("insert data success %v", err)
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
		ipSize := len(list)
		log.Println("check ip size ", ipSize)
		if ipSize > 0 {
			num := 50
			if num > ipSize {
				num = ipSize
			}
			ch := make(chan map[string]interface{}, num)
			go func() {
				for _, value := range list {
					ch <- value
				}
			}()
			var wg sync.WaitGroup
			for i := 0; i < num; i++ {
				go func() {
					wg.Add(1)
					if len(ch) == 0 {
						wg.Done()
						return
					}
					value := <-ch
					ip := value["ip"].(string)
					port := value["port"].(string)
					p, _ := strconv.Atoi(port)
					ipporxy := &crawl.IpProxy{Ip: ip, Port: p}
					ipporxy.SetDBHelper(mysqlTable)
					score := checkRule.CheckProxy(ipporxy)
					log.Println("check ", value["ip"], score)
					ipporxy.UpdateScore(score)
				}()
			}
			wg.Wait()
		}
		logs.Log.Debug("Check ip proxy end")
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
	logs.Log.Debug("begin crawl ip proxy")
	log.Println("rule size ", len(CrawlerList))
	//go CrawlProccess()
	// log.Println("begin check ip proxy")
	go CheckProxy()

	// go web.Run()
	<-ch
}
