package tank

import (
	"project/db"
	"project/pb"
)

// Basic ...
func Basic(raw *pb.TankRaw) (tb *pb.TankBase, err error) {

	tb = &pb.TankBase{
		ID:     raw.TankId,
		Name:   raw.TankName,
		Tier:   raw.TankTier,
		Nation: getTankNation(raw.TankNation),
		Type:   getTankType(raw.TankType),
		Shop:   getTankShop(raw.TankSort),
	}

	db.SIDIcon(raw.TankId, raw.TankIcon)

	err = poolUpdate(tb)
	return
}

func getTankType(s string) (t pb.TankEnumType) {
	switch s {
	case `AT-SPG`:
		t = pb.TankEnum_TD
	case `mediumTank`:
		t = pb.TankEnum_MT
	case `heavyTank`:
		t = pb.TankEnum_HT
	case `lightTank`:
		t = pb.TankEnum_LT
	case `SPG`:
		t = pb.TankEnum_SPG
	}
	return
}

func getTankNation(s string) (t pb.TankEnumNation) {
	switch s {
	case `ussr`:
		t = pb.TankEnum_S
	case `usa`:
		t = pb.TankEnum_M
	case `china`:
		t = pb.TankEnum_C
	case `france`:
		t = pb.TankEnum_F
	case `uk`:
		t = pb.TankEnum_Y
	case `germany`:
		t = pb.TankEnum_D
	case `japan`:
		t = pb.TankEnum_R
	case `sweden`:
		t = pb.TankEnum_V
	case `italy`:
		t = pb.TankEnum_I
	case `poland`:
		t = pb.TankEnum_B
	case `czech`:
		t = pb.TankEnum_J
	}

	return
}

func getTankShop(i uint32) (t pb.TankEnumShop) {
	switch i {
	case 2:
		t = pb.TankEnum_Gold
	case 3:
		t = pb.TankEnum_Premium
	}
	return
}
