package wiki

import (
	"project/db"
	"project/tank"
)

// Run ...
func Run() {

	wm, err := db.SIDWikiList()
	if err != nil {
		return
	}

	for _, id := range tank.LoadID() {

		_, ok := wm[id]
		if ok {
			continue
		}

		SID(id)
	}

	build()
}
