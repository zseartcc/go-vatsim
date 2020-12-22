package govatsim

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	json3mapper "github.com/zseartcc/go-vatsim/mapper/json3"
	"github.com/zseartcc/go-vatsim/model/json3"
)

const (
	commentChar       = ';'
	keyValueSeparator = "="
	statusEndpoint    = "https://status.vatsim.net/"
)

type VATSIM struct {
	refreshRate   time.Duration
	status        *status
	url0          *url0
	json3         *json3.JSON3
	controllerMap map[int]*json3.Controller
	pilotMap      map[int]*json3.Pilot
}

func NewVATSIM() *VATSIM {
	return &VATSIM{
		refreshRate: time.Duration(30) * time.Second,
	}
}

func (v *VATSIM) Controller(controllerCID int) (*json3.Controller, error) {
	err := v.refreshJSON3()
	if err != nil {
		return nil, err
	}

	controller, ok := v.controllerMap[controllerCID]
	if !ok {
		return nil, errors.New("controller not found")
	}

	return controller, nil
}

func (v *VATSIM) Pilot(pilotCID int) (*json3.Pilot, error) {
	err := v.refreshJSON3()
	if err != nil {
		return nil, err
	}

	pilot, ok := v.pilotMap[pilotCID]
	if !ok {
		return nil, errors.New("pilot not found")
	}

	return pilot, nil
}

func (v *VATSIM) Clients() ([]Client, error) {
	err := v.initStatus()
	if err != nil {
		return nil, err
	}

	if v.url0 != nil && time.Now().Before(v.url0.general.update.Add(v.refreshRate)) {
		return v.url0.clients, nil
	}

	url0 := v.status.url0[rand.Intn(len(v.status.url0))]
	clientLines, err := getVATSIMlines(url0)
	if err != nil {
		return nil, err
	}

	v.url0, err = parseUrl0(clientLines)
	if err != nil {
		return nil, err
	}

	return v.url0.clients, nil
}

func (v *VATSIM) Client(CID int) (Client, error) {
	clients, err := v.Clients()
	if err != nil {
		return Client{}, err
	}

	// TODO: hash it
	for _, client := range clients {
		if client.CID == CID {
			return client, nil
		}
	}

	return Client{}, errors.New("client not found")
}

func (v *VATSIM) initStatus() error {
	if v.status != nil {
		return nil
	}

	statusLines, err := getVATSIMlines(statusEndpoint)
	if err != nil {
		return nil
	}

	v.status, err = parseStatus(statusLines)
	if err != nil {
		return err
	}

	return nil
}

func getJSON(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

func (v *VATSIM) refreshJSON3() error {
	err := v.initStatus()
	if err != nil {
		return err
	}

	if v.json3 != nil {
		update, err := time.Parse(json3.TimestampLayout, v.json3.General.UpdateTimestamp)
		if err != nil {
			return err
		}

		if time.Now().Before(update.Add(v.refreshRate)) {
			return nil
		}
	}

	url := v.status.json3[rand.Intn(len(v.status.json3))]
	res := json3.JSON3{}
	err = getJSON(url, &res)
	if err != nil {
		return err
	}

	v.json3 = &res
	v.controllerMap = json3mapper.ControllersToMap(res.Controllers)
	v.pilotMap = json3mapper.PilotsToMap(res.Pilots)
	return nil
}

func getVATSIMlines(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)
	splitLines := strings.Split(bodyString, "\n")
	lines := make([]string, 0, len(splitLines))
	for _, line := range splitLines {
		if len(line) == 0 || line[0] == commentChar {
			continue
		}

		lines = append(lines, strings.TrimSpace(line))
	}

	return lines, nil
}
