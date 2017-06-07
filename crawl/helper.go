package crawl

import (
	"../db"
	"log"
	"time"
)

type Spider interface {
	GetIpProxyList() []*IpProxy
	Name() string
}

var (
	ListIpProxy []*IpProxy
)

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

func (self *IpProxy) SetDBHelper(dbdriver db.DBInterface) {
	self.dbHelper.SetDBDriver(dbdriver)
}

func (self *IpProxy) Insert(dbdriver db.DBInterface) error {
	self.dbHelper.SetDBDriver(dbdriver)
	exists, _ := self.Exists()
	if exists {
		log.Println("this data exists ")
		return nil
	}
	value := []interface{}{self.Ip, self.Port, self.Type, self.Country, self.Regin, GetCurrentTime(), GetCurrentTime(), 1}
	self.dbHelper.Values = append(self.dbHelper.Values, value)
	log.Println("dbHelper:", self.dbHelper)
	return self.dbHelper.AutoInsert()
}

func (self *IpProxy) Exists() (bool, error) {
	args := []interface{}{}
	args = append(args, self.Ip)
	args = append(args, self.Port)
	log.Println(args)
	result, err := self.Query("select * from ip where ip=? and port=?", args)
	if err != nil {
		return false, err
	}
	if len(result) > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (self *IpProxy) Query(sql string, args []interface{}) ([]map[string]interface{}, error) {
	return self.dbHelper.DBDriver.Query(sql, args)
}

func (self *IpProxy) Exec(sql string, args []interface{}) error {
	return self.dbHelper.DBDriver.Exec(sql, args)
}
