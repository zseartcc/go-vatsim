package json3

import (
	"github.com/zseartcc/go-vatsim/model/json3"
)

func PilotsToMap(pilots []json3.Pilot) map[int]*json3.Pilot {
	pilotMap := map[int]*json3.Pilot{}
	for _, pilot := range pilots {
		pilotMap[pilot.CID] = &pilot
	}

	return pilotMap
}
