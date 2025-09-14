package config

import (
	"os"
)

func init() {
	list := map[string]*string{
		`TANK_MYSQL`:      &MySQL,
		`TANK_STATIC_DIR`: &StaticDir,
	}
	for k, v := range list {
		s := os.Getenv(k)
		if len(s) > 1 {
			*v = s
		}
	}
}
