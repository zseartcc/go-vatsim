package json3

type Pilot struct {
	CID                int
	Name               string
	Callsign           string
	Server             string
	PilotRating        int `json:"pilot_rating"`
	Latitude           float64
	Longitude          float64
	Altitude           int
	Groundspeed        int
	Transponder        string
	Heading            int
	QNHInchesOfMercury float64    `json:"qnh_i_hg"`
	QNHMillibars       float64    `json:"qnh_mb"`
	FlightPlan         FlightPlan `json:"flight_plan"`
	LogonTime          string     `json:"logon_time"`
	LastUpdated        string     `json:"last_updated"`
}
