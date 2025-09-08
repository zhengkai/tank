package task

import (
	"project/spider"
	"project/tank"
	"project/wiki"
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
	spider.CrawlAll()
	tank.Build()
	tank.Date()
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
	tank.BuildHistory()
	wiki.Run()
}
