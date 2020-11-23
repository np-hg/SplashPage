package components

import (
	"bytes"
	"fmt"
	"text/template"
)

// HeaderImage is actually the image that rests above the picture
type HeaderImage struct {
	Image string `yaml:"image"`
	Name  string `yaml:"name"`
	IsB64 bool   `yaml:"isb64,omitempty"`
}

// NewHeaderImg creates default struct, here to enforce interface
func NewHeaderImg() HTMLComponent {
	return &HeaderImage{}
}

// ConsumeTemplate into html
func (i HeaderImage) ConsumeTemplate() string {
	var img string
	if i.IsB64 {
		img = `<img class='hdr-img' src="data:image/png;base64, {{.Image}}"/>`
	} else {
		img = `<img class='hdr-img' src="{{.Image}}"/>`
	}

	templString := `
	<div class='hdr-link'>` +
		img + `</div>`

	tmpl, err := template.New("hdr-" + i.Name).Parse(templString)
	if err != nil {
		fmt.Println(err)
	}

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, i); err != nil {
		fmt.Println(err)
	}

	return tpl.String()

}

// Styles Returns the styles
func (HeaderImage) Styles() string {

	return `
	.hdr-img {
		display: block;
		margin-left: auto;
		margin-right: auto;
		margin-bottom: 25px;
		width: 700px;
	}
	`
}
