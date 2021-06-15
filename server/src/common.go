package project

import (
	"project/build"
	"project/db"
	"project/tank"
)

func common() {
	build.DumpBuildInfo()

	db.WaitConn(`wot:wot@/wot`)

	// spider.CrawlAllSimulate()

	tank.InitMap()
}
