package spider

import (
	"encoding/json"
	"errors"
	"project/pb"
	"project/tank"
)

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
		err = errors.New(`empty list`)
		return
	}

	for _, v := range d.Data.Ranking {
		parseOne(v)
	}

	next = d.Data.Next == 1
	return
}

// parseOne ...
func parseOne(raw *pb.TankRaw) {
	tank.Basic(raw)
}
