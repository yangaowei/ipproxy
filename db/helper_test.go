package db

import (
	"log"
	"testing"
)

func TestHelper(t *testing.T) {
	log.Println(db)
	Refresh()
	log.Println(db)

	info := &MyTable{}
	info.tableName = "setting"
	info.columnNames = [][2]string{{"name", "string"}, {"field1", "string"}, {"field2", "string"}, {"field3", "string"}, {"field4", "string"}, {"field5", "string"}}

	var help DBInterface
	value := []interface{}{"1", "2", "3", "4", "5", "6"}
	help = info
	log.Println(help)
	help.AddRow(value)
	help.FlushInsert()

	result, err := info.SelectAll()
	if err == nil {
		print(result)
	}
}
