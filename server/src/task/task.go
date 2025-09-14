package task

import (
	"project/spider"
	"project/tank"
	"project/wiki"
	"project/zj"
	"sync"
	"time"
)

var crawlMux sync.Mutex
var buildMux sync.Mutex

func Crawl() bool {
	if crawlMux.TryLock() {
		go crawl()
		return true
	}
	return false
}

func crawl() {
	zj.J(`task crawl start`)
	t := time.Now()
	spider.CrawlAll()
	tank.Date()
	zj.J(`task crawl done`, time.Since(t).String())
	crawlMux.Unlock()
}

func Build() bool {
	if buildMux.TryLock() {
		go build()
		return true
	}
	return false
}

func build() {
	zj.J(`task build start`)
	t := time.Now()
	zj.J(`  task build history`)
	tank.BuildHistory()
	zj.J(`  task build list`)
	tank.Build()
	zj.J(`  task build wiki`)
	wiki.Run()
	zj.J(`task build done`, time.Since(t).String())
	buildMux.Unlock()
}
