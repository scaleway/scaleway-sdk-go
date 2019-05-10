package logger

import (
	"bytes"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

var (
	expectedError   = "ERROR: ab"
	expectedErrorf  = "ERROR: cd"
	expectedErrorln = "ERROR: e f"

	expectedWarning   = "WARNING: gh"
	expectedWarningf  = "WARNING: ij"
	expectedWarningln = "WARNING: k l"

	expectedInfo   = "INFO: mn"
	expectedInfof  = "INFO: op"
	expectedInfoln = "INFO: q r"

	expectedDebug   = "DEBUG: st"
	expectedDebugf  = "DEBUG: uv"
	expectedDebugln = "DEBUG: w x"
)

func TestDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelDebug))
	testThings(t, []string{
		expectedError,
		expectedErrorf,
		expectedErrorln,
		expectedWarning,
		expectedWarningf,
		expectedWarningln,
		expectedInfo,
		expectedInfof,
		expectedInfoln,
		expectedDebug,
		expectedDebugf,
		expectedDebugln,
	}, buf.String())
}

func TestInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelInfo))
	testThings(t, []string{
		expectedError,
		expectedErrorf,
		expectedErrorln,
		expectedWarning,
		expectedWarningf,
		expectedWarningln,
		expectedInfo,
		expectedInfof,
		expectedInfoln,
	}, buf.String())
}

func TestWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelWarning))
	testThings(t, []string{
		expectedError,
		expectedErrorf,
		expectedErrorln,
		expectedWarning,
		expectedWarningf,
		expectedWarningln,
	}, buf.String())
}

func TestError(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelError))
	testThings(t, []string{
		expectedError,
		expectedErrorf,
		expectedErrorln,
	}, buf.String())
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
	log.Error("a", "b")
	log.Errorf("c%s", "d")
	log.Errorln("e", "f")

	log.Warning("g", "h")
	log.Warningf("i%s", "j")
	log.Warningln("k", "l")

	log.Info("m", "n")
	log.Infof("o%s", "p")
	log.Infoln("q", "r")

	log.Debug("s", "t")
	log.Debugf("u%s", "v")
	log.Debugln("w", "x")
}
