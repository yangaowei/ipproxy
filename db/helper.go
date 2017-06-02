package db

import (
	"database/sql"
)

type DBInterface interface {
	//AutoInsert(value []string) *MyTable
	AddRow(value []interface{}) *MyTable
	FlushInsert() error
	SelectAll() (*sql.Rows, error)
	Query(string, []interface{}) ([]map[string]interface{}, error)
	Exec(string, []interface{}) error
}

type DBHelper struct {
	Values   [][]interface{}
	DBDriver DBInterface
}

func (self *DBHelper) SetDBDriver(insert DBInterface) {
	self.DBDriver = insert
}

func (self *DBHelper) AutoInsert() error {
	for _, value := range self.Values {
		self.DBDriver.AddRow(value)
	}
	return self.DBDriver.FlushInsert()
}
