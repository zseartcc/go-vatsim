package json3

import (
	"github.com/zseartcc/go-vatsim/model/json3"
)

func ControllersToMap(controllers []json3.Controller) map[int]*json3.Controller {
	controllerMap := map[int]*json3.Controller{}
	for _, controller := range controllers {
		controllerMap[controller.CID] = &controller
	}

	return controllerMap
}
