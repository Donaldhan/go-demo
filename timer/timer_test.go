package timer

import "testing"

// cd timer
// go test -v
func TestTimerWaitChannelDemoWithoutWrite(t *testing.T) {
	TimerWaitChannelDemoWithoutWrite()
}

func TestTimerWaitChannelDemoAfterSelectWrite(t *testing.T) {
	TimerWaitChannelDemoAfterSelectWrite()
}
func TestTimerWaitChannelDemoBeforeSelectWrite(t *testing.T) {
	TimerWaitChannelDemoBeforeSelectWrite()
}
func TestTimerWaitChannelDemoGoSelectThenWrite(t *testing.T) {
	TimerWaitChannelDemoGoSelectThenWrite()
}

func TestAsynChannelSelectForModeWithTimeoutDemo(t *testing.T) {
	AsynChannelSelectForModeWithTimeoutDemo()
}

func TestSyncChannelSelectForModeWithTimeoutDemo(t *testing.T) {
	SyncChannelSelectForModeWithTimeoutDemo()
}

func TestAsyncChannelSelectForDefaultCaseModeDemo(t *testing.T) {
	AsyncChannelSelectForDefaultCaseModeDemo()
}
func TestSyncChannelSelectForDefaultCaseModeDemo(t *testing.T) {
	SyncChannelSelectForDefaultCaseModeDemo()
}
func TestTimerDelayFunction(t *testing.T) {
	TimerDelayFunction()
}

func TestTimerAfterDemo(t *testing.T) {
	TimerAfterDemo()
}

func TestTimerAfterFuncDemo(t *testing.T) {
	TimerAfterFuncDemo()
}

func TestTickerDemo(t *testing.T) {
	TickerDemo()
}

func TestTickerTaskDemo(t *testing.T) {
	TickerTaskDemo()
}
