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
	name := [][2]string{{"ip", "string"}, {"port", "string"}, {"type", "string"}, {"country", "string"}, {"region", "string"}, {"createTime", "string"}, {"lastCheckTime", "string"}, {"status", "string"}}
	table.SetColumnNames(name)
	log.Println(table)

	ipProxy := IpProxy{"127.0.0.2", 80, "北京", "中国", 1, db.DBHelper{}}
	insertErr := ipProxy.Insert(table)
	if insertErr != nil {
		log.Println("insertErr:", insertErr)
	}

	// result, err := table.SelectAll()
	// if err == nil {
	// 	result, err := db.LoadsRows(result)
	// 	if err == nil {
	// 		log.Println("结果数量：", len(result))
	// 	}
	// }
	// args := []interface{}{}
	// args = append(args, "127.0.0.1")
	// err := ipProxy.Exec("delete from ip where ip=?", args)
	// if err == nil {
	// 	result, _ := ipProxy.Query("select * from ip where ip=?", args)
	// 	log.Println("结果数量：", len(result))
	// 	log.Println("结果：", result)
	// }

}
