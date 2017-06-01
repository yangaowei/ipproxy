package crawl

import (
	"../db"
	"log"
	"time"
)

type Spider interface {
	GetIpProxyList() []IpProxy
}

type IpProxy struct {
	Ip      string
	Port    int
	Regin   string
	Country string
	Type    int
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}

func (self *IpProxy) Insert(insert db.InsertInterface) error {
	value := []interface{}{self.Ip, self.Port, self.Type, self.Country, self.Regin, GetCurrentTime(), GetCurrentTime(), 1}
	dbHelper := db.DBHelper{}
	dbHelper.SetInsert(insert)
	dbHelper.Values = append(dbHelper.Values, value)
	log.Println("dbHelper:", dbHelper)
	return dbHelper.AutoInsert()
}
