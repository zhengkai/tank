package project

import (
	"project/cron"
	"project/tank"
	"project/web"
	"project/wiki"
	"project/zj"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	// spider.CrawlAll()

	tank.Build()

	wiki.Run()

	zj.J(`done`)

	// tank.Date()
	// tank.BuildHistory()
	// go cron.Run()
	// tank.Build()
	// spider.CrawlAllSimulate()

	go web.Server(21024)

	select {}
}

// Prod ...
func Prod() {

	common()

	// spider.CrawlAll()
	// tank.Build()
	// time.Sleep(time.Hour)
	go cron.Run()

	zj.J(`prod start`)

	go web.Server(80)

	select {}
}
