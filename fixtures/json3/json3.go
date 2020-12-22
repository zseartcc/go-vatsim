package json3

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

const (
	fixtureDir = "fixtures/json3"
)

var Fixture = struct {
	General     []byte
	FlightPlan  []byte
	Pilot       []byte
	Controller  []byte
	Server      []byte
	Prefile     []byte
	Facility    []byte
	Rating      []byte
	PilotRating []byte
	JSON3       []byte
}{
	General:     unmarshalFixture("general.json"),
	FlightPlan:  unmarshalFixture("flightPlan.json"),
	Pilot:       unmarshalFixture("pilot.json"),
	Controller:  unmarshalFixture("controller.json"),
	Server:      unmarshalFixture("server.json"),
	Prefile:     unmarshalFixture("prefile.json"),
	Facility:    unmarshalFixture("facility.json"),
	Rating:      unmarshalFixture("rating.json"),
	PilotRating: unmarshalFixture("pilotRating.json"),
	JSON3:       unmarshalFixture("json3.json"),
}

func unmarshalFixture(path string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	bytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(filename), path))
	if err != nil {
		// this will fail during unmarshal, which is ok for tests
		return nil
	}

	return bytes
}
