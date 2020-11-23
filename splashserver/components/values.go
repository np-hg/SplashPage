package components

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Values for this splash
type Values struct {
	Icons      *[]Icon         `yaml:"icons,omitempty"`
	Background *Background     `yaml:"background,omitempty"`
	Header     *HeaderImage    `yaml:"titleImg,omitempty"`
	Auth       AuthCredentials `yaml:"auth"`
	RawHTML    *string         `yaml:"rawHtml"`
	Host       string          `yaml:"host"` // temp for now
	Port       string          `yaml:"port"`
}

// NewValues creates a service from a yaml file path
func NewValues(fpath string) Values {
	filename, err := filepath.Abs(fpath)

	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var values Values

	err = yaml.Unmarshal(yamlFile, &values)
	if err != nil {
		panic(err)
	}
	return values
}

// AsHTMLComponent casts to interface
func (s *Values) AsHTMLComponent() HTMLComponent {
	return s
}

// ConsumeTemplate consumes object into template
func (s *Values) ConsumeTemplate() string {
	return s.prepareIcons()
}

// Styles returns css styles
func (s *Values) Styles() string {
	// only need 1 icons style
	icons := *(s.Icons)
	cssString := icons[0].Styles() + s.Background.Styles()
	cssString += `
	.grid {
		display: grid;
		grid-template-columns: 300px 300px 300px;
		grid-column-gap: 25px;
		grid-row-gap: 25px;
		width: auto;
		overflow: hidden;
	}

	.container{
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		min-height: 100%;
		padding-top: 25px;
	}
	`
	if s.Header != nil {
		cssString += s.Header.Styles()
	}
	return cssString

}

func (s *Values) prepareIcons() string {
	icons := *(s.Icons)

	var iconHTML string
	if s.Header != nil {
		iconHTML += s.Header.ConsumeTemplate()
	}
	iconHTML += `
	<div class='container'>
		<div class='grid'>
	`
	for _, icon := range icons {
		iconHTML += icon.ConsumeTemplate()
	}
	iconHTML += `
		</div>
	</div>
	`

	return iconHTML
}
