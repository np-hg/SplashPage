package components

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Values for this splash
type Values struct {
	Icons      []Icon          `yaml:"icons"`
	Background Background      `yaml:"background,omitempty"`
	Auth       AuthCredentials `yaml:"auth"`
	Host       string          `yaml:"host"` // temp for now
	Port       string          `yaml:"port"`
}

// NewValues creates a service from a yaml file path
func NewValues(fpath string) Values {
	filename, _ := filepath.Abs("./values.yaml")
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
	cssString := s.Icons[0].Styles() + s.Background.Styles()
	cssString += `
	.grid {
		display: grid;
		grid-auto-columns: 200px;
		width: auto;
		overflow: hidden;
	  }
	  
	.column {
		grid-row: 1;
	  }

	.container{
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 100vh;
	}
	`
	return cssString

}

func (s *Values) prepareIcons() string {
	iconHTML := `
	<div class='container'>
		<div class='grid'>
			<div class='column'>
	`
	for idx, icon := range s.Icons {
		if idx%2 == 0 && idx > 0 {
			iconHTML += `
			</div>
			<div class='column'>
			`
		}
		iconHTML += icon.ConsumeTemplate()
	}
	iconHTML += `
			</div>
		</div>
	</div>
	`

	return iconHTML
}
