package model

import "testing"

// go test -v
func TestInsert(t *testing.T) {
	insert()
}
func TestFind(t *testing.T) {
	query()
}

func TestUpdate(t *testing.T) {
	update()
	query()
}

func TestDelete(t *testing.T) {
	delete()
}
