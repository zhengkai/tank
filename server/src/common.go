package project

import (
	"project/build"
	"project/spider"
)

func common() {
	build.DumpBuildInfo()

	spider.CrawlAll(true)

	// db.WaitConn(`user:pass@/dbname`)
}
