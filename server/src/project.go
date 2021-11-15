package project

import (
	"project/cron"
	"project/spider"
	"project/tank"
	"project/zj"
	"time"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	tank.Build()
	// spider.CrawlAll()

	select {}
}

// Prod ...
func Prod() {

	common()

	spider.CrawlAll()
	tank.Build()
	time.Sleep(time.Hour)
	go cron.Run()

	zj.J(`prod start`)

	select {}
}
