package zj

// J log
var J = defaultLog.Log

// F log printf
var F = defaultLog.Logf

// Raw ...
var Raw = defaultLog.Raw

// W warn log
var W = warnLog.Log

// IO ...
var IO = io.Log

// IOF ...
var IOF = io.Logf

// IOColor ...
var IOColor = io.ColorOnce

// N log nothing
func N(x ...interface{}) {
}
