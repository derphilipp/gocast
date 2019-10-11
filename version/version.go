package version

import (
	"fmt"
	"runtime"
)

// GitCommit is the commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version is the main version number that is being run at the moment.
const Version = "0.0.1"

// BuildDate contains the date of compilation
var BuildDate = ""

// GoVersion contains the version of Go used to build it
var GoVersion = runtime.Version()

// OsArch contains the version and name of the operating system used to build the project
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
