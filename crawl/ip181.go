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

type IP181 struct {
	Index string //http://www.66ip.cn/
	Urls  []string
	Type  string
}

func (self *IP181) GetIndexList() (ip *IP181) {
	return
}

func (self *IP181) GetIpProxyList() (list []*IpProxy) {
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
		doc.Find("tr.warning").Each(func(i int, s *goquery.Selection) {
			s1 := s.Find("td:nth-child(1)").Text()
			ss := s.Find("td:nth-child(2)").Text()
			sss := s.Find("td:nth-child(3)").Text()
			ssss := s.Find("td:nth-child(4)").Text()
			region := s.Find("td:nth-child(6)").Text()
			log.Println(s1, ss, IpType[sss], ssss, region)
			port, _ := strconv.Atoi(ss)
			var proxyType int
			var ok bool
			if proxyType, ok = IpType[sss]; !ok {
				proxyType = 2
			}
			ipproxy := &IpProxy{Ip: s1, Port: port, Regin: region, Country: "中国", Type: proxyType}
			list = append(list, ipproxy)
		})
	}
	return
}

func (self *IP181) Name() (name string) {
	name = "http://www.ip181.com/"
	return
}
