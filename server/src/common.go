package project

import (
	"project/build"
	"project/db"
	"project/tank"
)

func common() {
	build.DumpBuildInfo()

	db.WaitConn(`wot:wot@/wot`)

	tank.InitMap()

	// spider.CrawlAllSimulate()

	tank.Build()
	// spider.CrawlAll()
}
