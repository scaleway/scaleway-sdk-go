package logger

import (
	"bytes"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

var (
	expectedError   = logEvent{"ab", LogLevelError}
	expectedErrorf  = logEvent{"cd", LogLevelError}
	expectedErrorln = logEvent{"e f", LogLevelError}

	expectedWarning   = logEvent{"gh", LogLevelWarning}
	expectedWarningf  = logEvent{"ij", LogLevelWarning}
	expectedWarningln = logEvent{"k l", LogLevelWarning}

	expectedInfo   = logEvent{"mn", LogLevelInfo}
	expectedInfof  = logEvent{"op", LogLevelInfo}
	expectedInfoln = logEvent{"q r", LogLevelInfo}

	expectedDebug   = logEvent{"st", LogLevelDebug}
	expectedDebugf  = logEvent{"uv", LogLevelDebug}
	expectedDebugln = logEvent{"w x", LogLevelDebug}
)

type logEvent struct {
	msg   string
	level LogLevel
}

func TestDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	logThings(newLogger(buf, LogLevelDebug))
	testThings(t, []logEvent{
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
	testThings(t, []logEvent{
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
	testThings(t, []logEvent{
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
	testThings(t, []logEvent{
		expectedError,
		expectedErrorf,
		expectedErrorln,
	}, buf.String())
}

func testThings(t *testing.T, expectedEvents []logEvent, actualOutput string) {
	lines := strings.Split(actualOutput, "\n")
	for i, line := range lines[:len(lines)-1] { // last line is always empty
		tmp := strings.Split(line, " ")
		actualLevel := tmp[0][:len(tmp[0])-1]       // remove the ":" at the end
		actualMessage := strings.Join(tmp[3:], " ") // everything after XXXX: XXXX/XX/XX XX:XX:XX
		testhelpers.Equals(t, expectedEvents[i].msg, actualMessage)
		testhelpers.Equals(t, severityName[expectedEvents[i].level], actualLevel)
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
