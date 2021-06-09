package build

import "project/zj"

// BuildGoVersion ...
var BuildGoVersion string

// BuildTime ...
var BuildTime string

// BuildType ...
var BuildType string

// BuildHost ...
var BuildHost string

// BuildGit ...
var BuildGit string

// DumpBuildInfo ...
func DumpBuildInfo() {
	zj.Raw("\n")
	zj.Raw(BuildGoVersion, "\n")
	zj.Raw(BuildTime, "\n")
	zj.Raw(BuildType, "\n")
	zj.Raw(BuildHost, "\n")
	zj.Raw(BuildGit, "\n")
	zj.Raw("\n")
}
