package json3

type Server struct {
	Ident                    string
	HostnameOrIP             string `json:"hostname_or_ip"`
	Location                 string
	Name                     string
	ClientsConnectionAllowed int `json:"clients_connection_allowed"`
}
