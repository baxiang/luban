package runtime

import (
	"os"
	"runtime"
)

func GetSysOS() string {
	return runtime.GOOS
}

// Getenv get ENV value by key name
func Getenv(name string, defaults ...string) string {
	val := os.Getenv(name)
	if val == "" && len(defaults) > 0 {
		val = defaults[0]
	}
	return val
}
