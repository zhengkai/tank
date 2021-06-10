package tank

import (
	"project/pb"
	"project/zj"
)

// Basic ...
func Basic(raw *pb.TankRaw) {

	t := getTankType(raw.TankType)

	var n pb.TankEnumNation

	zj.J(`base`, raw.TankName, t, n)
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
		t = pb.TankEnum_HT
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
	}

	// poland
	// italy
	// uk
	// china
	// sweden
	// japan
	// france

	return
}

func getTankShop(i uint32) (t pb.TankEnumShop) {
	switch i {
	case 1:
		t = pb.TankEnum_Gold
	case 3:
		t = pb.TankEnum_Premium
	}
	return
}
