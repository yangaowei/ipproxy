package db

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func print(rows *sql.Rows) {
	// for rows.Next() {
	// 	var name []interface{}
	// 	if err := rows.Scan(&name); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("row is ", name)
	// }
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(columns)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}

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
	value := []string{"a", "b", "c", "d", "e", "f"}
	table.addRow(value)
	table.FlushInsert()
	result, err := table.SelectAll()
	if err == nil {
		print(result)
	}

}
