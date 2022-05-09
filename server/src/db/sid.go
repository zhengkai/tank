package db

import (
	"project/pb"
	"project/zj"
)

// SIDRow ...
type SIDRow struct {
	ID   uint32
	Wiki string
}

// SIDWiki ...
func SIDWiki(id uint32, wiki string) {
	if id == 0 || wiki == `` {
		return
	}
	query := `INSERT INTO sid SET tank_id = ?, wiki = ?, gg = '', icon = '' ON DUPLICATE KEY UPDATE wiki = ?`
	_, err := d.Exec(query, id, wiki, wiki)
	zj.Watch(&err)
}

// SIDTg ...
func SIDTg(id uint32, tg string) {
	if id == 0 || tg == `` {
		return
	}
	query := `UPDATE sid SET gg = ? WHERE tank_id = ?`
	_, err := d.Exec(query, tg, id)
	zj.Watch(&err)
}

// SIDIcon ...
func SIDIcon(id uint32, icon string) {
	if id == 0 || icon == `` {
		return
	}
	query := `UPDATE sid SET icon = ? WHERE tank_id = ?`
	_, err := d.Exec(query, icon, id)
	zj.Watch(&err)
}

// SIDWikiList ...
func SIDWikiList() (m map[uint32]*pb.TankAlias, err error) {

	zj.Watch(&err)

	m = make(map[uint32]*pb.TankAlias)

	query := `SELECT tank_id, wiki, gg FROM sid`
	r, err := d.Query(query)
	if err != nil {
		return
	}
	defer r.Close()

	for r.Next() {
		row := &pb.TankAlias{}
		err = r.Scan(&row.ID, &row.Wiki, &row.Tanksgg)
		if err != nil {
			return
		}
		if row.ID == 0 || row.Wiki == `` {
			continue
		}
		m[row.ID] = row
	}

	return
}
