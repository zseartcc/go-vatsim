package json3

type General struct {
	Version          int
	Reload           int
	Update           string
	UpdateTimestamp  string `json:"update_timestamp"`
	ConnectedClients int    `json:"connected_clients"`
	UniqueUsers      int    `json:"unique_users"`
}
