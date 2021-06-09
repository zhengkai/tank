package project

import (
	"project/build"
)

func common() {
	build.DumpBuildInfo()

	// db.WaitConn(`user:pass@/dbname`)
}
