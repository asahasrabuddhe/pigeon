package flat

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Flat is a theme
type Flat struct{}

// Name returns the name of the flat theme
func (f *Flat) Name() string {
	return "flat"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (f *Flat) HTMLTemplate() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	b, err := ioutil.ReadFile(path.Dir(filename) + "/" + f.Name() + ".tpl.html")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (f *Flat) PlainTextTemplate() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	b, err := ioutil.ReadFile(path.Dir(filename) + "/" + f.Name() + ".tpl.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
