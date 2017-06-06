package main

import (
	"./crawl"
	"log"
	"time"
)

var (
	CrawlerList []crawl.Spider
	FREQUENCY   = 10
)

func Register() {
	CrawlerList = append(CrawlerList, &crawl.DataWu{})
	//CrawlerList = append(CrawlerList, &crawl.LLIP{})
}

func startSpider(spider crawl.Spider) {
	for {
		start := time.Now()
		nextRun := start.Add(time.Duration(FREQUENCY) * time.Second)
		log.Println(spider.Name())
		time.Sleep(nextRun.Sub(time.Now()))
	}
}

func CrawlProccess() {
	for _, spider := range CrawlerList {
		go startSpider(spider)
	}
}

func main() {
	Register()
	ch := make(chan int)
	log.Println("begion crawl ip proxy")
	log.Println("rule size ", len(CrawlerList))
	CrawlProccess()
	<-ch
}
