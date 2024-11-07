package mysql

import "testing"

// go test -v
func TestMysqlInsert(t *testing.T) {
	insert()
}
func TestFind(t *testing.T) {
	query()
}

func TestUpdate(t *testing.T) {
	update()
}

func TestDelete(t *testing.T) {
	delete()
}

func TestTransaction(t *testing.T) {
	transaction()
}
