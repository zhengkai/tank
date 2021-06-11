package project

import (
	"project/build"
	"project/cron"
	"project/db"
	"project/tank"
)

func common() {
	build.DumpBuildInfo()

	db.WaitConn(`wot:wot@/wot`)

	tank.InitMap()

	go cron.Run()
}
