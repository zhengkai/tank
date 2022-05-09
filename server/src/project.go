package project

import (
	"project/cron"
	"project/wiki"
	"project/zj"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	wiki.Test()

	// tank.Date()
	// tank.BuildHistory()
	// go cron.Run()
	// tank.Build()
	// spider.CrawlAllSimulate()

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

	select {}
}
