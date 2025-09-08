package project

import (
	"project/tank"
	"project/web"
	"project/wiki"
	"project/zj"

	"github.com/zhengkai/zu"
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

	zj.J(`prod start`)

	if !zu.FileExists(tank.File()) {
		go tank.Build()
	}

	go web.Server(80)

	select {}
}
