package task

import (
	"project/spider"
	"project/tank"
	"project/wiki"
	"project/zj"
	"sync"
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
	spider.CrawlAll()
	tank.Build()
	tank.Date()
	zj.J(`task crawl done`)
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
	tank.BuildHistory()
	wiki.Run()
	zj.J(`task build done`)
	buildMux.Unlock()
}
