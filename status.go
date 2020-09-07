package govatsim

import (
	"strings"
)

type status struct {
	url0   []string
	url1   []string
	metar0 string
}

func parseStatus(lines []string) (*status, error) {
	status := status{
		url0: []string{},
		url1: []string{},
	}

	for _, line := range lines {
		splitLine := strings.SplitN(line, keyValueSeparator, 2)
		if len(splitLine) != 2 {
			continue
		}

		key, val := strings.TrimSpace(splitLine[0]), strings.TrimSpace(splitLine[1])
		switch key {
		case "url0":
			status.url0 = append(status.url0, val)
		case "url1":
			status.url1 = append(status.url1, val)
		case "metar0":
			status.metar0 = val
		}
	}

	return &status, nil
}
