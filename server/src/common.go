package project

import (
	"project/build"
	"project/config"
	"project/db"
	"project/tank"
)

func common() {
	build.DumpBuildInfo()

	db.WaitConn(config.MySQL)

	// spider.CrawlAllSimulate()

	tank.InitMap()
}
