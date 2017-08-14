package crawl

import (
	//"../utils"
	//"../utils/surfer"
	"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"log"
	"strconv"
	// "strings"
)

type KuaiDaiLi struct {
	Index string //http://www.66ip.cn/
	Urls  []string
	Type  string
}

func (self *KuaiDaiLi) GetIndexList() (ip *KuaiDaiLi) {
	return
}

func (self *KuaiDaiLi) GetIpProxyList() (list []*IpProxy) {
	for _, url := range self.Urls {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("table>tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
			s1 := s.Find("td:nth-child(1)").Text()
			ss := s.Find("td:nth-child(2)").Text()
			sss := s.Find("td:nth-child(3)").Text()
			ssss := s.Find("td:nth-child(4)").Text()
			region := s.Find("td:nth-child(5)").Text()
			log.Println(s1, ss, sss, ssss, region)
			port, _ := strconv.Atoi(ss)
			//var proxyType int

			ipproxy := &IpProxy{Ip: s1, Port: port, Regin: region, Country: "中国", Type: 3}
			list = append(list, ipproxy)
		})
	}
	return
}

func (self *KuaiDaiLi) Name() (name string) {
	name = "http://www.kuaidaili.cn/"
	return
}
