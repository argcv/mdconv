package mdconv

type Root struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type Node interface {
	// Node Type
	GetType() string

	// List Children
	GetChildren() []Node

	// Raw Data
	String() string
}
