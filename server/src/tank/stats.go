package tank

import (
	"project/db"
	"project/pb"
)

// Stats ...
func Stats(raw *pb.TankRaw, date int) (tb *pb.TankStats, err error) {

	tb = &pb.TankStats{
		Frags:           raw.Frags,
		Spotted:         raw.Spotted,
		Xp:              raw.Xp,
		SurvivedBattles: raw.SurvivedBattles,
		Battles:         raw.Battles,
		Hits:            raw.Hits,
		Wins:            raw.Wins,
		Shots:           raw.Shots,
		DamageDealt:     raw.DamageDealt,
	}

	higher := raw.RankType == `higher`

	err = db.TankStats(raw.TankId, higher, date, tb)
	return
}
