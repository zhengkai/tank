package page

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"project/pb"
	"project/zj"
)

// Parse ...
func Parse() {

	return

	ab, err := ioutil.ReadFile(`/www/tank/data.json`)
	if err != nil {
		return
	}

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

	zj.J(`data`, len(d.Data.Ranking))
	for _, v := range d.Data.Ranking {

		zj.J(`tank`, v.TankType, v.TankNation)
	}

	return
}
