package json3

type Controller struct {
	CID         int
	Name        string
	Callsign    string
	Frequency   string
	Facility    int
	Rating      int
	Server      string
	VisualRange int      `json:"visual_range"`
	TextATIS    []string `json:"text_atis"`
	LastUpdated string   `json:"last_updated"`
	LogonTime   string   `json:"logon_time"`
}
