package util

import (
	"io/ioutil"
	"project/zj"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/proto"
)

// WriteFile ...
func WriteFile(file string, d proto.Message) (err error) {

	defer zj.Watch(&err)

	ab, err := proto.Marshal(d)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(file, ab, 0666)
	if err != nil {
		return
	}

	err = zu.Brotli(file)
	if err != nil {
		return
	}

	err = zu.Gzip(file)
	if err != nil {
		return
	}

	return
}
