package ogame

import (
	"strconv"
	"time"
)

// AttackEvent all information available about an enemy attack
type AttackEvent struct {
	MissionType  MissionID
	Origin       Coordinate
	Destination  Coordinate
	ArrivalTime  time.Time
	ArriveIn     int64
	AttackerName string
	AttackerID   int64
	UnionID      int64
	Missiles     int64
	Ships        *ShipsInfos
}

func (a AttackEvent) String() string {
	return "" +
		"Mission Type: " + strconv.FormatInt(int64(a.MissionType), 10) + "\n" +
		"      Origin: " + a.Origin.String() + "\n" +
		" Destination: " + a.Destination.String() + "\n" +
		" ArrivalTime: " + a.ArrivalTime.String() + "\n" +
		"  AttackerID: " + strconv.FormatInt(a.AttackerID, 10) + "\n" +
		"     UnionID: " + strconv.FormatInt(a.UnionID, 10) + "\n" +
		"    Missiles: " + strconv.FormatInt(a.Missiles, 10)
}
