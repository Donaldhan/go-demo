package middleware

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
