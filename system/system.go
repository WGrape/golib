package system

import (
	"go/build"
	"os"
)

// GetGoPath return the gopath of your go environment
func GetGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}
