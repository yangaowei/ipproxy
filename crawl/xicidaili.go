package crawl

import (
	"../utils"
	//"encoding/json"
	//"../utils/surfer"
	"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"log"
	"strconv"
	"strings"
)

type XiCiDaiLi struct {
	Index string //http://www.66ip.cn/
	Urls  []string
	Type  string
}

func (self *XiCiDaiLi) GetIndexList() (ip *XiCiDaiLi) {
	return
}

func (self *XiCiDaiLi) GetIpProxyList() (list []*IpProxy) {
	for _, url := range self.Urls {
		//j, e := BuildJson(url)
		html, _ := utils.GetContent(url, nil)
		htmlReader := strings.NewReader(html)
		doc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("table#ip_list").Find("tr").Each(func(i int, s *goquery.Selection) {
			s1 := s.Find("td:nth-child(2)").Text()
			ss := s.Find("td:nth-child(3)").Text()
			region := s.Find("td:nth-child(4)").Find("a").Text()
			ssss := s.Find("td:nth-child(5)").Text()
			//log.Println(s1, ss, IpType[ssss], region)
			port, _ := strconv.Atoi(ss)
			// var proxyType int
			// var ok bool
			// if proxyType, ok = IpType[sss]; !ok {
			// 	proxyType = 2
			// }
			ipproxy := &IpProxy{Ip: s1, Port: port, Regin: region, Country: "中国", Type: IpType[ssss]}
			list = append(list, ipproxy)
		})
	}
	return
}

func (self *XiCiDaiLi) Name() (name string) {
	name = "http://www.xicidaili.com/"
	return
}
