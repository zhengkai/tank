package cron

import (
	"project/spider"
	"project/tank"
	"strconv"
	"time"

	"github.com/zhengkai/zu"
)

var chDo = make(chan bool)
var chHistory = make(chan bool)

// Run ...
func Run() {

	if !zu.FileExists(tank.File()) {
		tank.Build()
	}

	go do()
	go history()

	for {
		time.Sleep(time.Second * 10)

		select {
		case chDo <- true:
		default:
		}

		select {
		case chHistory <- true:
		default:
		}
	}
}

func do() {
	for {
		<-chDo

		// minutes
		if time.Now().Format(`04`) != `15` {
			continue
		}

		// hours
		h, _ := strconv.Atoi(time.Now().Format(`15`))
		if h%2 == 0 {
			continue
		}

		spider.CrawlAll()
		tank.Build()
		tank.Date()

		time.Sleep(time.Minute * 62)
	}
}

func history() {
	for {
		<-chHistory

		if time.Now().Format(`1504`) != `0702` {
			continue
		}

		tank.BuildHistory()

		time.Sleep(time.Minute * 10)
	}
}
