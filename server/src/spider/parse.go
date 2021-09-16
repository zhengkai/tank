package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"project/pb"
	"project/tank"
	"regexp"
	"strconv"
	"strings"
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

	if len(d.Data.Ranking) == 0 {
		err = errEmptyList
		return
	}

	for _, v := range d.Data.Ranking {
		err = parseOne(v, d.Data.Modify)
		if err != nil {
			return
		}
	}

	next = d.Data.Next == 1
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
