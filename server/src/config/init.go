package config

import (
	"os"
)

func init() {
	list := map[string]*string{
		`TANK_MYSQL`:  &MySQL,
		`OUTPUT_PATH`: &OutputPath,
		`TMP_PATH`:    &TmpPath,
	}
	for k, v := range list {
		s := os.Getenv(k)
		if len(s) > 1 {
			*v = s
		}
	}
}
