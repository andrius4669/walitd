package render

import (
	"fmt"
	"time"
)

func date(args ...interface{}) string {
	t, ok := args[0].(time.Time)
	if ok {
		return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	} else {
		return "bad date"
	}
}
