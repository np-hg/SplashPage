package components

// HTMLComponent is an interface for a html component
type HTMLComponent interface {
	ConsumeTemplate() string
	Styles() string
}
