package db

import (
	"fmt"
	"project/pb"

	"google.golang.org/protobuf/proto"
)

// TankStats ...
func TankStats(id uint32, higher bool, date int, ts *pb.TankStats) (err error) {

	ab, err := proto.Marshal(ts)
	if err != nil {
		return
	}

	query := `INSERT INTO stat SET tank_id = ?, date = ?, %s = ? ON DUPLICATE KEY UPDATE %s = ?`
	col := `def`
	if higher {
		col = `hi`
	}
	query = fmt.Sprintf(query, col, col)

	_, err = d.Exec(query, id, date, ab, ab)
	return
}

// LoadTankStats ...
func LoadTankStats(id uint32) (def *pb.TankStats, hi *pb.TankStats, date int, err error) {

	query := `SELECT def, hi, date FROM stat WHERE tank_id = ? ORDER BY date DESC LIMIT 1`
	re := d.QueryRow(query, id)

	var abd []byte
	var abh []byte

	err = re.Scan(&abd, &abh, &date)
	if err != nil {
		return
	}

	if len(abd) > 0 {
		def = &pb.TankStats{}
		err = proto.Unmarshal(abd, def)
		if err != nil {
			return
		}
	}

	if len(abh) > 0 {
		hi = &pb.TankStats{}
		err = proto.Unmarshal(abh, hi)
		if err != nil {
			return
		}
	}

	return
}
