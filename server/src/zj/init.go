package zj

import (
	j "github.com/zhengkai/zj"
)

var io = j.New(&j.Config{
	Filename:   `log/io.log`,
	Prefix:     `[IO] `,
	Tunnel:     100,
	Caller:     j.CallerNone,
	TimeFormat: j.TimeMonth,
})

var defaultLog = j.New(&j.Config{
	Filename:   `log/default.log`,
	Tunnel:     100,
	Caller:     j.CallerShort,
	TimeFormat: j.TimeMonth,
})

var warnLog = j.New(&j.Config{
	Filename:   `log/warn.log`,
	Prefix:     `[WARN] `,
	Tunnel:     100,
	Caller:     j.CallerShort,
	TimeFormat: j.TimeMonth,
})

func init() {
	defaultLog.Color(`38;2;200;255;100`)
	warnLog.Color(`38;2;255;100;100`)
}

// Close ...
func Close() {
	io.Close()
	defaultLog.Close()
	warnLog.Close()
}
