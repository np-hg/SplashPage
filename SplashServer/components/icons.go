package components

import (
	"bytes"
	"fmt"
	"text/template"
)

// Icon basic struct for templates
type Icon struct {
	Image string `yaml:"image"`
	Name  string `yaml:"name"`
	Link  string `yaml:"link"`
	IsB64 bool   `yaml:"isb64,omitempty"`
}

// NewIcon creates default struct, here to enforce interface
func NewIcon() HTMLComponent {
	return &Icon{}
}

// ConsumeTemplate consumes the struct into html template
func (i Icon) ConsumeTemplate() string {
	var img string
	if i.IsB64 {
		img = `<img class='img-icon' src="data:image/png;base64, {{.Image}}"/>`
	} else {
		img = `<img class='img-icon' src="{{.Image}}"/>`
	}
	templString := `
	<figure class='service-icon'>
		<a href="{{.Link}}">` +
		img +
		`</a>
	</figure>`

	tmpl, err := template.New("icon-" + i.Name).Parse(templString)
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
func (Icon) Styles() string {

	return `
	.service-icon {
		padding: 15px 25px;
		position: relative;
		display: inline-block;
		font-size: 25px;
		font-family: sans-serif;
		border-radius: 3px;	
	}
	img.img-icon {
		width: 150px;
	}
	`
}
