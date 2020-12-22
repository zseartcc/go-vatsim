package json3

type Prefile struct {
	CID         int
	Name        string
	Callsign    string
	FlightPlan  FlightPlan `json:"flight_plan"`
	LastUpdated string     `json:"last_updated"`
}
