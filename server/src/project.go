package project

import (
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

	zj.J(`prod start`)

	select {}
}
