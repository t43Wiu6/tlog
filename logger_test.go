package log

import "testing"

func TestDebug(t *testing.T) {
	DEBUG = true
	Debug("debug test")
}

func TestDebugf(t *testing.T) {
	DEBUG = true
	Debugf("debug test : %s %v %d", "yes", "ok", 1)
}

func TestError(t *testing.T) {
	Error("error test")
}

func TestErrorf(t *testing.T) {
	Errorf("error test : %s %v %d", "yes", "ok", 1)
}

func TestInfo(t *testing.T) {
	Info("info test")
}

func TestInfof(t *testing.T) {
	Infof("error test : %s %v %d", "yes", "ok", 1)
}

func TestWarn(t *testing.T) {
	Warn("warn test")
}

func TestWarnf(t *testing.T) {
	Warnf("error test : %s %v %d", "yes", "ok", 1)
}

func TestFatal(t *testing.T) {
	// Fatal("fatal test")
}

func TestFatalf(t *testing.T) {
	// Fatal("fatal test")
}