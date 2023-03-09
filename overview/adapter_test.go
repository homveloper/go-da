package overview_test

import (
	"log"
	"testing"
)

type IProcess interface {
	Process()
}

type Adaptee struct {
	adapterType int
}

func (a *Adaptee) Convert() {
	log.Println("Converting")
}

type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Process() {
	log.Println("Processing")
	a.adaptee.Convert()
}

func TestAdapter(t *testing.T) {
	var process IProcess = &Adapter{Adaptee{1}}
	process.Process()
}
