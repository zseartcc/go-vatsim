package govatsim

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	commentChar       = ';'
	keyValueSeparator = "="
	statusEndpoint    = "http://status.vatsim.net/"
)

type VATSIM struct {
	refreshRate time.Duration
	status      *status
	url0        *url0
}

func NewVATSIM() *VATSIM {
	return &VATSIM{
		refreshRate: time.Duration(90) * time.Second,
	}
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
