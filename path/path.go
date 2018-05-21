package path

import (
	"path"
	"runtime"
)

func thisFileDir() string {
	_, thisfile, _, _ := runtime.Caller(0)
	return path.Dir(thisfile)
}
