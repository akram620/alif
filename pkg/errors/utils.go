package errors

import (
	"fmt"
	"regexp"
	"runtime"
)

func GetCaller(skip int) (caller string) {
	_, file, no, ok := runtime.Caller(skip)
	if ok {
		re := regexp.MustCompile("(/[^/]+/[^/]+)/?$")
		filePart := re.FindString(file)
		caller = fmt.Sprintf("%s#%d", filePart, no)
	}
	return caller
}
