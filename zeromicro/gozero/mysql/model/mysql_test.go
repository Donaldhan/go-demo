package model

import "testing"

// go test -v
func TestMysqlInsert(t *testing.T) {
	mysqlInsert()
}
func TestFind(t *testing.T) {
	find()
}

func TestUpdate(t *testing.T) {
	update()
}

func TestDelete(t *testing.T) {
	delete()
}
