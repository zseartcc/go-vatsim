package json3

import (
	"github.com/zseartcc/go-vatsim/model/json3"
)

func ControllersToMap(controllers []json3.Controller) map[int]*json3.Controller {
	controllerMap := map[int]*json3.Controller{}
	for i, controller := range controllers {
		controllerMap[controller.CID] = &controllers[i]
	}

	return controllerMap
}
