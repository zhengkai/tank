package task

import (
	"project/spider"
	"project/tank"
	"project/wiki"
)

func Crawl() {
	spider.CrawlAll()
	tank.Build()
	tank.Date()
}

func Build() {
	tank.BuildHistory()
	wiki.Run()
}
