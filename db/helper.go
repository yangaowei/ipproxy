package db

type InsertInterface interface {
	//AutoInsert(value []string) *MyTable
	AddRow(value []interface{}) *MyTable
	FlushInsert() error
}

type DBHelper struct {
	Values [][]interface{}
	Insert InsertInterface
}

func (self *DBHelper) SetInsert(insert InsertInterface) {
	self.Insert = insert
}

func (self *DBHelper) AutoInsert() error {
	for _, value := range self.Values {
		self.Insert.AddRow(value)
	}
	return self.Insert.FlushInsert()
}
