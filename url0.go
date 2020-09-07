package govatsim

import (
	"strconv"
	"strings"
)

const (
	url0SectionUNKNOWN = ""
	url0SectionGeneral = "!GENERAL:"
	url0SectionClients = "!CLIENTS:"
	url0SectionServers = "!SERVERS:"
	url0SectionPrefile = "!PREFILE:"
)

type url0 struct {
	general general
	clients []Client
}

func parseUrl0(lines []string) (*url0, error) {
	url0 := url0{
		clients: []Client{},
		general: general{},
	}
	currentSection := url0SectionUNKNOWN
	for _, line := range lines {
		if strings.HasPrefix(line, url0SectionGeneral) {
			currentSection = url0SectionGeneral
			continue
		}
		if strings.HasPrefix(line, url0SectionClients) {
			currentSection = url0SectionClients
			continue
		}
		if strings.HasPrefix(line, url0SectionServers) {
			currentSection = url0SectionServers
			continue
		}
		if strings.HasPrefix(line, url0SectionPrefile) {
			currentSection = url0SectionPrefile
			continue
		}

		switch currentSection {
		case url0SectionGeneral:
			splitLine := strings.SplitN(line, keyValueSeparator, 2)
			if len(splitLine) != 2 {
				continue
			}
			key, value := strings.TrimSpace(splitLine[0]), strings.TrimSpace(splitLine[1])
			switch key {
			case "VERSION":
				url0.general.version, _ = strconv.Atoi(value)
			case "RELOAD":
				url0.general.reload, _ = strconv.Atoi(value)
			case "UPDATE":
				url0.general.update, _ = parseTime(value)
			case "CONNECTED CLIENTS":
				url0.general.connectedClients, _ = strconv.Atoi(value)
			case "UNIQUE_USERS":
				url0.general.uniqueUsers, _ = strconv.Atoi(value)
			}
		case url0SectionClients:
			client, err := parseClient(line)
			if err != nil {
				return nil, err
			}

			url0.clients = append(url0.clients, *client)
		}

	}

	return &url0, nil
}
