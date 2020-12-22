package govatsim

import (
	"strconv"
	"strings"
	"time"
)

type ClientType string
type FacilityType int
type ClientRating int

const (
	clientSeparator = ":"
)

// TODO: get ADM value
const (
	ObserverRating    = 1
	Student1Rating    = 2
	Student2Rating    = 3
	Student3Rating    = 4
	Controller1Rating = 5
	Controller3Rating = 7
	Instructor1Rating = 8
	Instructor3Rating = 10
	SupervisorRating  = 11
)

const (
	ClientTypeATC   = "ATC"
	ClientTypePilot = "PILOT"
)

// TODO: get FSS value
const (
	FacilityTypeDelivery = 2
	FacilityTypeGround   = 3
	FacilityTypeTWR      = 4
	FacilityTypeApproach = 5
	FacilityTypeCenter   = 6
)

// Client represents an online user in the VATSIM network
type Client struct {
	Callsign                           string
	CID                                int
	RealName                           string
	ClientType                         ClientType
	Frequency                          string
	Latitude                           float64
	Longitude                          float64
	Altitude                           int
	GroundSpeed                        int
	PlannedAircraft                    string
	PlannedTASCruise                   int
	PlannedDepartureAirport            string
	PlannedAltitude                    string
	PlannedDestinationAirport          string
	Server                             string
	ProtocolRevision                   string
	Rating                             ClientRating
	Transponder                        string
	FacilityType                       FacilityType
	VisualRange                        string
	PlannedRevision                    string
	PlannedFlightType                  string
	PlannedDepartureTime               string
	PlannedActualDepartureTime         string
	PlannedHoursEnRoute                string
	PlannedMinutesEnRoute              string
	PlannedHoursFuel                   string
	PlannedMinimumFuel                 string
	PlannedAlternateAirport            string
	PlannedRemarks                     string
	PlannedRoute                       string
	PlanngeDepartureAirportLatitude    string
	PlannedDepartureAirportLongitude   string
	PlannedDestinationAirportLatitude  string
	PlannedDestinationAirportLongitude string
	ATISMessage                        string
	TimeLastATISReceived               time.Time
	TimeLogon                          time.Time
	Heading                            int
	QNHInchesOfMercury                 float64
	QNHMillibars                       int
}

func (c *Client) IsATC() bool {
	return c.ClientType == ClientTypeATC
}

func (c *Client) IsPilot() bool {
	return c.ClientType == ClientTypePilot
}

func parseClient(line string) (*Client, error) {
	fields := strings.Split(line, clientSeparator)
	cid, err := parseInt(fields[1])
	if err != nil {
		return nil, err
	}

	latitude, err := parseFloat(fields[5])
	if err != nil {
		return nil, err
	}

	longitude, err := parseFloat(fields[6])
	if err != nil {
		return nil, err
	}

	altitude, err := parseInt(fields[7])
	if err != nil {
		return nil, err
	}

	groundspeed, err := parseInt(fields[8])
	if err != nil {
		return nil, err
	}

	plannedTAS, err := parseInt(fields[10])
	if err != nil {
		return nil, err
	}

	rating, err := parseInt(fields[16])
	if err != nil {
		return nil, err
	}

	facilityType, err := parseInt(fields[18])
	if err != nil {
		return nil, err
	}

	timeLastATISReceived, err := parseTime(fields[36])
	if err != nil {
		return nil, err
	}

	timeLogon, err := parseTime(fields[37])
	if err != nil {
		return nil, err
	}

	heading, err := parseInt(fields[38])
	if err != nil {
		return nil, err
	}

	QNHin, err := parseFloat(fields[39])
	if err != nil {
		return nil, err
	}

	QNHmb, err := parseInt(fields[40])
	if err != nil {
		return nil, err
	}

	return &Client{
		Callsign:                           fields[0],
		CID:                                cid,
		RealName:                           fields[2],
		ClientType:                         ClientType(fields[3]),
		Frequency:                          fields[4],
		Latitude:                           latitude,
		Longitude:                          longitude,
		Altitude:                           altitude,
		GroundSpeed:                        groundspeed,
		PlannedAircraft:                    fields[9],
		PlannedTASCruise:                   plannedTAS,
		PlannedDepartureAirport:            fields[11],
		PlannedAltitude:                    fields[12],
		PlannedDestinationAirport:          fields[13],
		Server:                             fields[14],
		ProtocolRevision:                   fields[15],
		Rating:                             ClientRating(rating),
		Transponder:                        fields[17],
		FacilityType:                       FacilityType(facilityType),
		VisualRange:                        fields[19],
		PlannedRevision:                    fields[20],
		PlannedFlightType:                  fields[21],
		PlannedDepartureTime:               fields[22],
		PlannedActualDepartureTime:         fields[23],
		PlannedHoursEnRoute:                fields[24],
		PlannedMinutesEnRoute:              fields[25],
		PlannedHoursFuel:                   fields[26],
		PlannedMinimumFuel:                 fields[27],
		PlannedAlternateAirport:            fields[28],
		PlannedRemarks:                     fields[29],
		PlannedRoute:                       fields[30],
		PlanngeDepartureAirportLatitude:    fields[31],
		PlannedDepartureAirportLongitude:   fields[32],
		PlannedDestinationAirportLatitude:  fields[33],
		PlannedDestinationAirportLongitude: fields[34],
		ATISMessage:                        fields[35],
		TimeLastATISReceived:               timeLastATISReceived,
		TimeLogon:                          timeLogon,
		Heading:                            heading,
		QNHInchesOfMercury:                 QNHin,
		QNHMillibars:                       QNHmb,
	}, nil
}

func parseInt(field string) (int, error) {
	if field == "" {
		return 0, nil
	}

	return strconv.Atoi(field)
}

func parseFloat(field string) (float64, error) {
	if field == "" {
		return 0, nil
	}

	return strconv.ParseFloat(field, 64)
}

func parseTime(t string) (time.Time, error) {
	parsedTime, err := time.Parse("20060102150405", t)
	return parsedTime, err
}
