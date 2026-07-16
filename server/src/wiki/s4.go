package wiki

import (
	"encoding/json"
	"project/pb"
	"project/util"
	"project/zj"
	"regexp"
	"strings"

	"github.com/zhengkai/zu"
)

var regexpS4Name = regexp.MustCompile(`[^a-z0-9 ]`)
var regexpMultiSpace = regexp.MustCompile(` +`)

var s4URL = `https://skill4ltu.eu/api/directus/tree-vehicles`

var s4TankID = make([]uint32, 0, 2000)

// STest ...
func STest() {
	ab, err := s4File()
	if err != nil {
		return
	}

	m, _ := getS4Map()

	zj.J(`s4`, len(ab), len(m))
}

func getS4Map() (m map[uint32]string, err error) {

	m = make(map[uint32]string)

	d := &pb.S4List{}
	ab, err := s4File()
	if err != nil {
		return
	}

	err = json.Unmarshal(ab, d)
	if err != nil {
		zj.W(err)
		return
	}

	for _, v := range d.Data {
		m[v.GetTankId()] = s4Name(v.GetName())
		zj.J(v.GetTankId(), s4Name(v.GetName()))
	}

	return
}

func s4File() (ab []byte, err error) {

	file := `data/s4.json`

	defer zj.Watch(&err)

	ab, err = zu.FetchURL(s4URL)
	if err == nil {
		util.WriteFile(file, ab)
		return
	}
	zj.W(err)
	return util.ReadFile(file)
}

func s4Name(n string) string {
	n = strings.ToLower(n)
	n = regexpS4Name.ReplaceAllString(n, ``)
	n = strings.TrimSpace(n)
	n = regexpMultiSpace.ReplaceAllString(n, `-`)
	if n == `` {
		return `a`
	}
	return n
}
