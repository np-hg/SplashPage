package components

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

// ServiceIcons are the preset templates icons
type ServiceIcons struct {
	CustomCSS   *string `yaml:"customCSS"`
	BannerImage *string `yaml:"bannerImage"`
	IconLinks   *[]struct {
		Href          string   `yaml:"href"`
		Src           string   `yaml:"src"`
		AllowedGroups []string `yaml:"allowedGroups"`
		Show          bool
	} `yaml:"icons"`
}

// Values for this splash
type Values struct {
	Auth         AuthCredentials `yaml:"auth"`
	RawHTML      *string         `yaml:"rawHtml"`
	Host         string          `yaml:"host"` // temp for now
	Port         string          `yaml:"port"`
	ServiceIcons *ServiceIcons   `yaml:"services"`
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

// ToHTML converts the values into html string
func (v *Values) ToHTML(user User, templatedHTML string) (string, error) {
	var html string
	if v.RawHTML != nil {
		html = *v.RawHTML
		html = strings.ReplaceAll(html, "%", "%%")
		return html, nil
	}

	if v.ServiceIcons != nil {
		tmpl, err := template.New("html-template").Parse(templatedHTML)

		// if user groups is in IconLinks allowed groups
		// show flag is True
		for _, i := range *v.ServiceIcons.IconLinks {
			i.Show = isIn(user.Groups, i.AllowedGroups)
		}

		var tpl bytes.Buffer
		if err = tmpl.Execute(&tpl, v.ServiceIcons); err != nil {
			return "", err
		}

		html = tpl.String()
		html = strings.ReplaceAll(html, "%", "%%")

		return html, nil
	}
	return "", errors.New("Services or raw html must be defined")
}

func isIn(userGroups []string, valueGroups []string) bool {

	for _, ug := range userGroups {
		for _, vg := range valueGroups {
			if ug == vg {
				return true
			}
		}
	}

	return false
}
