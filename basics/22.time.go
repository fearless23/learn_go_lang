package basics

import (
	"time"
)

func Time() {
	Log("time.April", time.April)
	Log("time.Now().Format(time.RFC3339)", time.Now().Format(time.RFC3339))
}
