package wiki

import (
	"errors"
	"fmt"
	"net/http"
	"project/db"
	"project/zj"
	"regexp"
)

var regexpSID = regexp.MustCompile(`/tankopedia/(\d+)-(.+)/$`)

// SID ...
func SID(id uint32) (sid string, err error) {

	if id < 1 {
		return
	}

	defer zj.Watch(&err)

	url := fmt.Sprintf(
		`https://wotgame.cn/zh-cn/tankopedia/%d/`,
		id,
	)

	rsp, err := http.Get(url)
	if err != nil {
		return
	}

	re := regexpSID.FindStringSubmatch(rsp.Request.URL.Path)
	if len(re) != 3 || re[2] == `` {
		err = errors.New(`not found`)
		zj.W(rsp.Request.URL.Path, re)
		return
	}

	sid = re[2]
	db.SIDWiki(id, sid)

	zj.J(rsp.Request.URL.Path)
	zj.J(sid)
	return
}
