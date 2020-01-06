package registry

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
	Nodes     []*Node
}

type Endpoint struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	MetaData map[string]string `json:"meta_data"`
}

type Node struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	MetaData map[string]string `json:"meta_data"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
