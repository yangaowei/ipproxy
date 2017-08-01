package crawl

import (
	"../db"
	"../utils"
	"../utils/surfer"
	"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"log"
	"strconv"
	"strings"
)

type DataWu struct {
	LLIP
}

// func DataParse(i int, contentSelection *goquery.Selection) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println(err) // 这里的err其实就是panic传入的内容
// 		}
// 	}()

// 	info := contentSelection.Find("li")
// 	if info.Size() == 9 {
// 		ip := info.Nodes[0].FirstChild.Data
// 		port, _ := strconv.Atoi(info.Nodes[1].FirstChild.Data)
// 		aText := info.Find("a")
// 		region := aText.Nodes[3].FirstChild.Data
// 		proxyTypeString := aText.Nodes[0].FirstChild.Data
// 		country := aText.Nodes[2].FirstChild.Data
// 		var proxyType int
// 		if proxyTypeString == "匿名" {
// 			proxyType = 2
// 		} else if proxyTypeString == "高匿" {
// 			proxyType = 3
// 		} else {
// 			proxyType = 1
// 		}
// 		ipproxy := &IpProxy{Ip: ip, Port: port, Regin: region, Country: country, Type: proxyType, dbHelper: db.DBHelper{}}
// 		//log.Println(ipproxy)
// 		result = append(result, ipproxy)
// 	}
// }

func (self *DataWu) setUrls() {
	self.Urls = []string{"http://www.data5u.com/"}
}

func (self *DataWu) GetIpProxyList() (list []*IpProxy) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	if len(self.Urls) == 0 {
		self.setUrls()
	}
	for _, url := range self.Urls {
		request := &surfer.DefaultRequest{Url: url, TryTimes: 1, EnableCookie: true}
		request.GetUrl()
		html, _ := utils.GetHtml(request)
		htmlReader := strings.NewReader(html)
		doc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".l2").Each(func(i int, contentSelection *goquery.Selection) {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err) // 这里的err其实就是panic传入的内容
				}
			}()

			info := contentSelection.Find("li")
			if info.Size() == 9 {
				ip := info.Nodes[0].FirstChild.Data
				port, _ := strconv.Atoi(info.Nodes[1].FirstChild.Data)
				aText := info.Find("a")
				region := aText.Nodes[3].FirstChild.Data
				proxyTypeString := aText.Nodes[0].FirstChild.Data
				country := aText.Nodes[2].FirstChild.Data
				var proxyType int
				if proxyTypeString == "匿名" {
					proxyType = 2
				} else if proxyTypeString == "高匿" {
					proxyType = 3
				} else {
					proxyType = 1
				}
				ipproxy := &IpProxy{Ip: ip, Port: port, Regin: region, Country: country, Type: proxyType, dbHelper: db.DBHelper{}}
				//log.Println(ipproxy)
				list = append(list, ipproxy)
			}
		})
	}
	return
}

func (self *DataWu) Name() (name string) {
	name = "http://www.data5u.com/"
	return
}
