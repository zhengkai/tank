package db

import "project/zj"

// SIDRow ...
type SIDRow struct {
	ID   uint32
	Wiki string
}

// SIDWiki ...
func SIDWiki(id uint32, wiki string) {
	query := `INSERT INTO sid SET tank_id = ?, wiki = ?, gg = '' ON DUPLICATE KEY UPDATE wiki = ?`
	_, err := d.Exec(query, id, wiki, wiki)
	zj.Watch(&err)
}

// SIDWikiList ...
func SIDWikiList() (list []*SIDRow, err error) {

	zj.Watch(&err)

	query := `SELECT tank_id, wiki FROM sid`
	r, err := d.Query(query)
	if err != nil {
		return
	}
	defer r.Close()

	for r.Next() {
		row := &SIDRow{}
		err = r.Scan(&row.ID, &row.Wiki)
		if err != nil {
			return
		}
		list = append(list, row)
	}

	return
}
