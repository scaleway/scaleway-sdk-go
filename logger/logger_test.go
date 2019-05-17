package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

var (
	expectedErrorf   = "ERROR: cd"
	expectedWarningf = "WARNING: ij"
	expectedInfof    = "INFO: op"
	expectedDebugf   = "DEBUG: uv"
)

func TestDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelDebug))
	testThings(t, []string{
		expectedErrorf,
		expectedWarningf,
		expectedInfof,
		expectedDebugf,
	}, buf.String())
}

func TestInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelInfo))
	testThings(t, []string{
		expectedErrorf,
		expectedWarningf,
		expectedInfof,
	}, buf.String())
}

func TestWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelWarning))
	testThings(t, []string{
		expectedErrorf,
		expectedWarningf,
	}, buf.String())
}

func TestError(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelError))
	testThings(t, []string{
		expectedErrorf,
	}, buf.String())
}

func TestEnableDebugMode(t *testing.T) {
	_defaultLogger := DefaultLogger

	DefaultLogger = newLogger(os.Stderr, LogLevelWarning)
	EnableDebugMode()
	testhelpers.Equals(t, true, DefaultLogger.ShouldLog(LogLevelDebug))

	DefaultLogger = _defaultLogger
}

func testThings(t *testing.T, expectedEvents []string, actualOutput string) {
	lines := strings.Split(actualOutput, "\n")
	for i, line := range lines[:len(lines)-1] { // last line is always empty
		tmp := strings.Split(line, " ")
		actualMessage := strings.Join(append([]string{tmp[0]}, tmp[3:]...), " ") // date and hour is not kept
		testhelpers.Equals(t, expectedEvents[i], actualMessage)
	}
}

func logThings(log Logger) {
	log.Errorf("c%s", "d")
	log.Warningf("i%s", "j")
	log.Infof("o%s", "p")
	log.Debugf("u%s", "v")
}
