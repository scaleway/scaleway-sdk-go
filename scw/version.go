package scw

import (
	"runtime"
	"runtime/internal/sys"
)

// TODO: versionning process
const version = "0.0.0"

const userAgent = "scaleway-sdk-go/" + version + " (" + sys.TheVersion + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"
