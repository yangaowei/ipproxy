package crawl

import (
	"../utils"
	"../utils/surfer"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"log"
	"strconv"
	"strings"
)

type LLIP struct {
	Index string //http://www.66ip.cn/
	Urls  []string
	Type  string
}

func (self *LLIP) GetIndexList() (ip *LLIP) {
	return
}

// func LLParse(i int, contentSelection *goquery.Selection) {
// 	// contentSelection.Find("td").Each(func(i int, contentSelection *goquery.Selection) {
// 	// 	log.Println(contentSelection.Text())
// 	// })
// 	info := contentSelection.Find("td")
// 	if info.Size() == 5 {
// 		ip := info.Nodes[0].FirstChild.Data
// 		if ip != "ip" {
// 			port, _ := strconv.Atoi(info.Nodes[1].FirstChild.Data)
// 			proxyTypeString := info.Nodes[3].FirstChild.Data
// 			region := info.Nodes[2].FirstChild.Data
// 			var proxyType int
// 			if proxyTypeString == "高匿代理" {
// 				proxyType = 3
// 			} else {
// 				proxyType = 1
// 			}
// 			ipproxy := &IpProxy{Ip: ip, Port: port, Regin: region, Country: "中国", Type: proxyType}
// 			result = append(result, ipproxy)
// 		}
// 	}
// 	//return
// }

func (self *LLIP) GetIpProxyList() (list []*IpProxy) {
	for _, url := range self.Urls {
		request := &surfer.DefaultRequest{Url: url, TryTimes: 1, EnableCookie: true}
		request.GetUrl()
		//log.Println(request.GetHeader())
		html, _ := utils.GetHtml(request)
		html = mahonia.NewDecoder("gbk").ConvertString(html)
		htmlReader := strings.NewReader(html)
		doc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("table").Last().Find("tr").Each(func(i int, contentSelection *goquery.Selection) {
			// contentSelection.Find("td").Each(func(i int, contentSelection *goquery.Selection) {
			// 	log.Println(contentSelection.Text())
			// })
			info := contentSelection.Find("td")
			if info.Size() == 5 {
				ip := info.Nodes[0].FirstChild.Data
				if ip != "ip" {
					port, _ := strconv.Atoi(info.Nodes[1].FirstChild.Data)
					proxyTypeString := info.Nodes[3].FirstChild.Data
					region := info.Nodes[2].FirstChild.Data
					var proxyType int
					if proxyTypeString == "高匿代理" {
						proxyType = 3
					} else {
						proxyType = 1
					}
					ipproxy := &IpProxy{Ip: ip, Port: port, Regin: region, Country: "中国", Type: proxyType}
					list = append(list, ipproxy)
				}
			}
			//return
		})
	}
	return
}

func (self *LLIP) Name() (name string) {
	name = "http://www.66ip.cn/"
	return
}
