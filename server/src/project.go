package project

import (
	"project/cron"
	"project/tank"
	"project/wiki"
	"project/zj"
	"time"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	tank.Build()
	wiki.Run()

	zj.J(time.Now().Format(`1504`))

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
