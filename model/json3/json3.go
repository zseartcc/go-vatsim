package json3

const (
	TimestampLayout = "2006-01-02T15:04:05.9999999Z"
)

type JSON3 struct {
	General      General
	Pilots       []Pilot
	Controllers  []Controller
	ATIS         []Controller
	Servers      []Server
	Prefiles     []Prefile
	Facilities   []Facility
	Ratings      []Rating
	PilotRatings []PilotRating `json:"pilot_ratings"`
}
