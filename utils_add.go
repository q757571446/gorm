package gorm

import (
	"fmt"
	"regexp"
	"runtime"
)

var internalDbRegexp = regexp.MustCompile(`internal/dao(.*)/db.go`)

func fileWithLineNumSql() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!(goSrcRegexp.MatchString(file) || internalDbRegexp.MatchString(file)) || goTestRegexp.MatchString(file)) {
			return fmt.Sprintf("%v:%v", file, line)
		}
	}
	return ""
}
