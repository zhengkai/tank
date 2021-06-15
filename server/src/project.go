package project

import (
	"project/cron"
	"project/zj"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	select {}
}

// Prod ...
func Prod() {

	common()

	go cron.Run()

	zj.J(`prod start`)

	select {}
}
