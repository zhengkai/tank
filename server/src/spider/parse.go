package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"project/db"
	"project/pb"
	"project/tank"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var patternDate = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2}) `)

var errEmptyList = errors.New(`empty list`)

// Parse ...
func Parse(ab []byte) (next bool, err error) {

	d := &pb.Rank{}
	err = json.Unmarshal(ab, d)
	if err != nil {
		return
	}

	if d.Errmsg != `success` {
		err = errors.New(`json errmsg not "success"`)
		return
	}

	date := d.GetData().GetModify()
	if date == `` {
		err = errors.New(`no date`)
		return
	}
	r := d.GetData().GetRanking()
	if len(r) == 0 {
		err = errEmptyList
		return
	}

	for _, v := range r {
		err = parseOne(v, date)
		if err != nil {
			return
		}
	}

	next = d.GetData().GetNext() == 1
	return
}

// parseOne ...
func parseOne(raw *pb.TankRaw, date string) (err error) {

	s := patternDate.FindAllStringSubmatch(date, 1)
	if len(s) != 1 {
		err = fmt.Errorf(`unknown date: %s`, date)
		return
	}

	di, err := strconv.Atoi(strings.Join(s[0][1:], ``))
	if err != nil {
		return
	}

	_, err = tank.Basic(raw)
	if err != nil {
		return
	}

	_, err = tank.Stats(raw, di)
	if err != nil {
		return
	}

	return
}

// ParsePercent ...
func ParsePercent(ab []byte, percent int) (next bool, err error) {

	d := &pb.RankPercent{}
	err = json.Unmarshal(ab, d)
	if err != nil {
		return
	}

	date := time.Now().Format(`20060203`)
	di, _ := strconv.Atoi(date)

	for _, v := range d.GetData().GetRanking() {
		db.UpdatePercent(v.TankId, di, percent, int(v.Mastery))
		tank.UpdatePercent(v.TankId, uint32(v.Mastery), percent)
	}

	next = d.GetData().GetNext() == 1
	return
}
