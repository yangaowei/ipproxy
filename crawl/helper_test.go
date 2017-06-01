package crawl

import (
	"../db"
	"log"
	"testing"
)

func TestHelper(t *testing.T) {
	db.Refresh()
	table := &db.MyTable{}
	table.SetTableName("ip")
	// result, err := table.SelectAll()
	// if err == nil {
	// 	db.Print(result)
	// }
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	table.SetColumnNames(name)
	log.Println(table)

	ipProxy := IpProxy{"127.0.0.2", 80, "北京", "中国", 1}
	insertErr := ipProxy.Insert(table)
	if insertErr != nil {
		log.Println("insertErr:", insertErr)
	}

	result, err := table.SelectAll()
	if err == nil {
		db.Print(result)
	}
}
