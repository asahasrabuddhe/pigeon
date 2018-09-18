package default_theme

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Default is the theme by default
type Default struct{}

// Name returns the name of the default theme
func (dt *Default) Name() string {
	return "default"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Default) HTMLTemplate() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	b, err := ioutil.ReadFile(path.Dir(filename) + "/" + dt.Name() + ".tpl.html")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Default) PlainTextTemplate() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	b, err := ioutil.ReadFile(path.Dir(filename) + "/" + dt.Name() + ".tpl.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
