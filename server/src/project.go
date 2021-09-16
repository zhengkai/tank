package project

import (
	"project/cron"
	"project/spider"
	"project/zj"
)

// Dev ...
func Dev() {

	common()

	zj.J(`dev start`)

	// tank.Build()
	spider.CrawlAll()

	select {}
}

// Prod ...
func Prod() {

	common()

	go cron.Run()

	zj.J(`prod start`)

	select {}
}
