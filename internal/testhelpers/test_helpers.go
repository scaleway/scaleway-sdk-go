package testhelpers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...any) {
	tb.Helper()
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]any{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// AssertNoError fails the test if an err is not nil.
func AssertNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act any) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// HeaderContains fails if exp is not a subset of act
func HeaderContains(tb testing.TB, exp, act http.Header) {
	tb.Helper()
	for k, v := range exp {
		if !reflect.DeepEqual(act[k], v) {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf(
				"\033%s:%d:\n\n\texp: %#v\n\n\tto contain: %#v\033\n\n",
				filepath.Base(file), line, act, exp,
			)

			fmt.Printf(
				"\033%s:%d:\n\n\tfailed for key '%#v'\033\n\n",
				filepath.Base(file), line, k,
			)
			tb.FailNow()
		}
	}
}
