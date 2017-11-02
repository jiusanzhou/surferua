package surferua

import (
	"math/rand"
	"time"
	"strings"
)

func init() {
	Seed(time.Now().UnixNano())
}

func Seed(seed int64) {
	rand.Seed(seed)
}

const Mozilla = "Mozilla"
const MozillaWithVersion = "Mozilla/5.0"

type UserAgent struct {
	// Common version: 5.0
	Version string

	// Browser information
	Browser *Browser

	// Platform information
	Platform *Platform
}

func (ua *UserAgent) String() (us string) {

	// firefox has different tpl
	// 1. use `.` to connect semver instead of `_`
	// 2. ends with `rv:`
	if strings.Contains(ua.Browser.Name, "irefox") {
		return ua.Version + " (" + strings.Replace(ua.Platform.String(), "_", ".", -1) + "; rv:" + ua.Browser.Semver.String() + ") " + ua.Browser.String()
	}
	return ua.Version + " (" + ua.Platform.String() +  ") " + ua.Browser.String()
}

func New() (ua *UserAgent) {

	ua = &UserAgent{
		Version: MozillaWithVersion,
	}

	// if we need specific platform or browser
	// we just need to set the value again with the specific function.
	// this is a easy way to attach our global

	// random platform with version

	// random browser with version
	if ua.Browser == nil {
		ua.Browser = browserDB[rand.Intn(browserDBSize)].Random()
	}

	if ua.Platform == nil {
		x := rand.Intn(platformTypeSize)
		ua.Platform = platformDB[x][rand.Intn(platformSize[x])].Random()
	}

	return
}