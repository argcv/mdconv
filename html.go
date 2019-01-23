package mdconv

type HtmlElement interface {
	Node
	GetTag() string
	GetAttributes() map[string]string
}