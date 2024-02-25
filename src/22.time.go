package src

import (
	"time"
)

func Time() {
	log("time.April", time.April)
	log("time.Now().Format(time.RFC3339)", time.Now().Format(time.RFC3339))
}
