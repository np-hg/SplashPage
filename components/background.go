package components

import "fmt"

// Background image component
type Background struct {
	Image string `yaml:"image"`
	IsB64 bool   `yaml:"isb64,omitempty"`
}

// NewBackground creates default struct, here to enforce interface
func NewBackground() HTMLComponent {
	return &Background{}
}

// ConsumeTemplate consumes the struct into html template
func (Background) ConsumeTemplate() string {
	return ""
}

// Styles Returns the styles
func (b *Background) Styles() string {
	if b.IsB64 {
		return fmt.Sprintf(`
		body {
			background-image: url(data:image/png;base64,%s)
		}
		`, b.Image)
	}

	return fmt.Sprintf(`
		body {
			background-image: url(%s)
		}
		`, b.Image)
}
