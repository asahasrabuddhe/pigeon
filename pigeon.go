package pigeon

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/jaytaylor/html2text"
	"html/template"
)
import . "github.com/asahasrabuddhe/pigeon/email"

type Pigeon struct {
	Theme         Theme
	TextDirection TextDirection
	Product       Product
}

var templateFunctions = template.FuncMap{
	"url": func(s string) template.URL {
		return template.URL(s)
	},
}

// GenerateHTML generates the email body from data to an HTML Reader
// This is for modern email clients
func (p *Pigeon) GenerateHTML(email Email) (string, error) {
	err := setDefaultPigeonValues(p)
	if err != nil {
		return "", err
	}
	return p.generateTemplate(email, p.Theme.HTMLTemplate())
}

// GeneratePlainText generates the email body from data
// This is for old email clients
func (p *Pigeon) GeneratePlainText(email Email) (string, error) {
	err := setDefaultPigeonValues(p)
	if err != nil {
		return "", err
	}
	template, err := p.generateTemplate(email, p.Theme.PlainTextTemplate())
	if err != nil {
		return "", err
	}
	return html2text.FromString(template, html2text.Options{PrettyTables: true})
}

func (p *Pigeon) generateTemplate(email Email, tplt string) (string, error) {

	err := setDefaultEmailValues(&email)
	if err != nil {
		return "", err
	}

	// Generate the email from Golang template
	t, err := template.New("pigeon").Funcs(sprig.FuncMap()).Funcs(templateFunctions).Parse(tplt)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	t.Execute(&b, Template{*p, email})
	return b.String(), nil
}
