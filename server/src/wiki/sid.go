package wiki

import (
	"errors"
	"fmt"
	"project/db"
	"project/util"
	"project/zj"
	"regexp"
	"slices"
)

var regexpSID = regexp.MustCompile(`/tankopedia/(\d+)-(.+)/$`)

func defaultSID() {
	db.SIDWiki(15185, `GB84_Chieftain_Mk6`)
	db.SIDWiki(21537, `A143_M_V_Y`)
	db.SIDWiki(21281, `A144_M_VI_Y`)
	db.SIDWiki(21025, `A139_M_III_Y`)
	db.SIDWiki(20769, `A147_M_II_Y`)
	db.SIDWiki(20513, `A142_Pawlack_Tank`)
	db.SIDWiki(57633, `A112_T71E2`)
	db.SIDWiki(62017, `F83_AMX_M4_Mle1949_Bis_FL`)
	db.SIDWiki(32001, `R178_Object_780`)
	db.SIDWiki(16913, `G98_Waffentrager_E100_P`)
	db.SIDWiki(33281, `R185_T_34_L_11_1941`)
	db.SIDWiki(34593, `A148_Convertible_Medium_Tank_T3`)
	db.SIDWiki(25617, `G158_VK2801_105_SPXXI`)
	db.SIDWiki(55377, `GB108_A46`)
	db.SIDWiki(59169, `A117_T26E5_Patriot`)
}

var ignoreSID = []uint32{
	50721,
	64769,
	48145,
	62001,
	59217,
	65601,
	49681,
}

// SID ...
func SID(id uint32) (sid string, err error) {

	if id < 1 {
		err = errors.New(`empty id`)
		return
	}

	if slices.Contains(ignoreSID, id) {
		return
	}

	defer zj.Watch(&err)

	url := fmt.Sprintf(
		`https://wotgame.cn/zh-cn/tankopedia/%d/`,
		id,
	)

	rsp, err := util.HTTPNoProxyGet(url)
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
