package project

import (
	"project/zj"
)

// Dev ...
func Dev() {

	common()
	zj.J(`dev start`)

	zj.Close()
}

// Prod ...
func Prod() {

	common()
	zj.J(`prod start`)

	zj.Close()
}
