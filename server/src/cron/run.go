package cron

import (
	"project/spider"
	"project/tank"
	"project/zj"
	"time"

	"github.com/zhengkai/zu"
)

const min = time.Second * 100
const interval = time.Hour * 2

var chDo = make(chan bool)

// Run ...
func Run() {

	if !zu.FileExists(tank.File()) {
		tank.Build()
	}

	go do()

	m := time.Hour - time.Duration(time.Now().Unix()%7200)*time.Second
	now := time.Now()
	next := now.Add(m)

	diff := next.Sub(now)
	if diff < 0 {
		diff += interval
	}

	time.Sleep(diff)
	chDo <- true

	for {
		time.Sleep(interval)

		select {
		case chDo <- true:
		default:
		}
	}
}

func do() {
	for {
		<-chDo
		zj.J(`cron`)
		spider.CrawlAll()
		tank.Build()
		tank.Date()
	}
}
