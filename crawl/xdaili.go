package crawl

import (
	"../utils"
	"encoding/json"
	//"../utils/surfer"
	//"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"log"
	"strconv"
	// "strings"
)

type XDaiLiJson struct {
	Total int64 `json:"total"`
	Rows  []struct {
		Anony      string `json:"anony"`
		CreateTime int64  `json:"createTime"`
		Ip         string `json:"ip"`
		Port       string `json:"port"`
		Regin      string `json:"position"`
	} `json:"rows"`
}

type XDaiLi struct {
	Index string //http://www.66ip.cn/
	Urls  []string
	Type  string
}

func (self *XDaiLi) GetIndexList() (ip *XDaiLi) {
	return
}

func (self *XDaiLi) GetIpProxyList() (list []*IpProxy) {
	for _, url := range self.Urls {
		//j, e := BuildJson(url)
		str, _ := utils.GetContent(url, nil)
		var s XDaiLiJson
		err := json.Unmarshal([]byte(str), &s)
		if err != nil {
			log.Println("err:", err)
		}
		for _, row := range s.Rows {
			port, _ := strconv.Atoi(row.Port)
			//var proxyType int

			ipproxy := &IpProxy{Ip: row.Ip, Port: port, Regin: row.Regin, Country: "中国", Type: 3}
			list = append(list, ipproxy)
		}
	}
	return
}

func (self *XDaiLi) Name() (name string) {
	name = "http://www.xdaili.cn/"
	return
}
