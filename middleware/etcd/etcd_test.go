package etcd

import (
	"testing"
)

// go test -v
func TestEtcdPutGet(t *testing.T) {
	etcdPutGet()
}
func TestWatchDemo(t *testing.T) {
	watchDemo()
}
func TestLeaseDemo(t *testing.T) {
	leaseDemo()
}
func TestKeepAliveDemo(t *testing.T) {
	keepAliveDemo()
}
func TestEtcdLock(t *testing.T) {
	etcdLock()
}
func TestGet(t *testing.T) {
	getValue("/godemo/config.yaml")
}
