package chanselect

import "testing"

// cd chan/chanselect
// go test -v
func TestChanSelectBase(t *testing.T) {
	ChanSelectBase()
}

func TestChanSelectBaseX(t *testing.T) {
	ChanSelectBaseX()
}

func TestChanSelectTimeout(t *testing.T) {
	ChanSelectTimeout()
}

func TestChanSelectWithDefalut(t *testing.T) {
	ChanSelectWithDefalut()
}
