package db

import (
	"fmt"
	"project/pb"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

// DateStats ...
type DateStats struct {
	Date uint32
	Stat *pb.TankStats
}

// TankStats ...
func TankStats(id uint32, higher bool, date int, ts *pb.TankStats) (err error) {

	defer zj.Watch(&err)

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

	defer zj.Watch(&err)

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

// LoadTankDateStats ...
func LoadTankDateStats(id uint32, higher bool, date uint32) (list []*DateStats, err error) {

	defer zj.Watch(&err)

	colName := `hi`
	if !higher {
		colName = `def`
	}

	query := `SELECT %s, date FROM stat WHERE tank_id = ? AND date > ? ORDER BY date ASC LIMIT 500`
	query = fmt.Sprintf(query, colName)
	row, err := d.Query(query, id, date)
	if err != nil {
		return
	}
	defer row.Close()

	var ab []byte

	for row.Next() {

		d := &DateStats{}

		err = row.Scan(&ab, &d.Date)
		if err != nil {
			return
		}

		if len(ab) > 0 {
			st := &pb.TankStats{}
			err = proto.Unmarshal(ab, st)
			if err != nil {
				break
			}
			d.Stat = st
		}
		list = append(list, d)
	}

	return
}
