package json3

type FlightPlan struct {
	FlightRules string `json:"flight_rules"`
	Aircraft    string
	Departure   string
	Arrival     string
	Alternate   string
	CruiseTAS   string `json:"cruise_tas"`
	Altitude    string
	DepTime     string
	EnrouteTime string `json:"enroute_time"`
	FuelTime    string `json:"fuel_time"`
	Remarks     string
	Route       string
}
