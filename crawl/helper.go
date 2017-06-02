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
	Ip       string
	Port     int
	Regin    string
	Country  string
	Type     int //1，透明  2，匿名  3，高匿
	dbHelper db.DBHelper
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}

func (self *IpProxy) SetDBHelper(dbHelp db.DBHelper) {

}

func (self *IpProxy) Insert(dbdriver db.DBInterface) error {
	value := []interface{}{self.Ip, self.Port, self.Type, self.Country, self.Regin, GetCurrentTime(), GetCurrentTime(), 1}
	self.dbHelper.SetDBDriver(dbdriver)
	self.dbHelper.Values = append(self.dbHelper.Values, value)
	log.Println("dbHelper:", self.dbHelper)
	return self.dbHelper.AutoInsert()
}

func (self *IpProxy) Query(sql string, args []interface{}) ([]map[string]interface{}, error) {
	return self.dbHelper.DBDriver.Query(sql, args)
}

func (self *IpProxy) Exec(sql string, args []interface{}) error {
	return self.dbHelper.DBDriver.Exec(sql, args)
}
