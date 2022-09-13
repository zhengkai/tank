package wiki

import (
	"io"
	"net/http"
	"os"
	"project/config"
	"project/zj"
	"strconv"
	"strings"

	jp "github.com/buger/jsonparser"
)

var s4URL = `https://skill4ltu.eu/assets/data/techtree/germany/index.json`

// STest ...
func STest() {
	ab, err := s4File()
	if err != nil {
		return
	}

	getS4Map()

	zj.J(`s4`, len(ab))
}

func getS4Map() (m map[uint32]string, err error) {
	ab, err := s4File()
	if err != nil {
		return
	}

	m = make(map[uint32]string)

	_, err = jp.ArrayEach(ab, func(ab []byte, _ jp.ValueType, _ int, err error) {

		s, _ := jp.GetString(ab, `node`, `path`)

		a := strings.SplitN(s, `/`, 4)
		if len(a) != 4 || a[1] == `` || a[2] == `` {
			return
		}

		i, _ := strconv.Atoi(a[1])
		if i > 0 {
			m[uint32(i)] = a[2]
		}
	}, `data`, `vehicles`, `edges`)

	return
}

func s4File() (ab []byte, err error) {

	file := config.OutputPath + `/s4.json`

	defer zj.Watch(&err)

	rsp, err := http.Get(s4URL)
	if err == nil {
		ab, err = io.ReadAll(rsp.Body)
		if err == nil {
			os.WriteFile(file, ab, 0666)
			return
		}
	}
	return os.ReadFile(file)
}
