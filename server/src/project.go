package project

import (
	"project/cron"
	"project/tank"
	"project/zj"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	tank.Build()

	select {}
}

// Prod ...
func Prod() {

	common()

	go cron.Run()

	zj.J(`prod start`)

	select {}
}
