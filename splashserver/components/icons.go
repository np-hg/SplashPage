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
	<div class='service-icon'>
		<a href="{{.Link}}">` + img + `</a>
	</div>`

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
		background-color: white;
		position: relative;
		display: inline-block;
		font-size: 25px;
		font-family: sans-serif;
		border-radius: 3px;
		margin: 10px;
		height: 125px;
		border: calc(2px + 0.2vw) solid transparent;
		background-origin: border-box;
		background-clip: content-box, border-box;
		background-size: cover;
		box-sizing: border-box;
		box-shadow: 0 0 5px 5px rgba(0, 0, 0, 0.5);
	}

	.service-icon>a {
		position: absolute;
		top: 50%%;
		-ms-transform: translateY(-50%%);
		transform: translateY(-50%%);
		margin-left: 25%%;
	}

	.service-icon>a>img {
		position: absolute;
		text-align: center;
		top: 50%%;
		-ms-transform: translateY(-50%%);
		transform: translateY(-50%%);
	}
	`
}
