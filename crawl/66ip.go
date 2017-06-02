package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type LLIP struct {
	Index string //http://www.66ip.cn/
	Url   string
	Type  string
}

func (self *LLIP) GetIndexList() (list []*LLIP) {
	return
}
func Parse(i int, contentSelection *goquery.Selection) {
	contentSelection.Find("td").Each(func(i int, contentSelection *goquery.Selection) {
		log.Println(contentSelection.Text())
	})
	// log.Println(info)
	// //url, _ := info.Attr("href")
	// log.Println("IP", info.Text())
	log.Println("-----------------------------------")
}

func (self *LLIP) GetIpProxyList() (list []*IpProxy) {
	//url := self.Url
	html := strings.NewReader("<table><tr><td>闫</td></tr><table>")
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table").Find("tr").Each(Parse)
	return
}
