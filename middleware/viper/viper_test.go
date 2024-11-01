package viper

import (
	"testing"
)

// cd timer
// go test -v
func TestViperDemo(t *testing.T) {
	viperDemo()
}

func TestViperConfig(t *testing.T) {
	viperConfig()
}

func TestWatchConfig(t *testing.T) {
	watchConfig()
}

func TestLoadConfigRemoteEtcd3RunInsecureMode(t *testing.T) {
	loadConfigRemoteEtcd3RunInsecureMode()
}

func TestLoadConfigRemoteEtcd2(t *testing.T) {
	loadConfigRemoteEtcd2()
}

func TestLoadConfigRemoteEtcd3(t *testing.T) {
	loadConfigRemoteEtcd3()
}
