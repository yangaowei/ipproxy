package db

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestRefresh(t *testing.T) {
	log.Println(db)
	Refresh()
	log.Println(db)

	info := &MyTable{}
	info.tableName = "setting"
	result, err := info.SelectAll()
	if err == nil {
		print(result)
	}
}

func TestAutoInsert(t *testing.T) {
	table := &MyTable{}
	table.tableName = "setting"
	table.columnNames = [][2]string{{"name", "string"}, {"field1", "string"}, {"field2", "string"}, {"field3", "string"}, {"field4", "string"}, {"field5", "string"}}
	value := []interface{}{"a", "b", "c", "d", "e", "f"}
	table.AddRow(value)
	table.FlushInsert()
	result, err := table.SelectAll()
	if err == nil {
		print(result)
	}

}
