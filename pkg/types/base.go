package types

// URL represents a single URL
type URL struct {
	Scheme     string      `json:"scheme"`
	Address    Address     `json:"address"`
	Path       string      `json:"path"`
	Parameters []Parameter `json:"parameters"`
}

// Address represents a single address.
// e.g. from a Slave or from a Master
type Address struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}

// Parameter represents a single key / value pair for parameters
type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Label represents a single key / value pair for labeling
type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
