package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// ButtonOptions is a "button" HTML element configuration structure.
type ButtonOptions struct {
	Class          string
	ID             string
	OnClickHandler func(e *js.Object)
	Text           string
}

// Button is a controlling structure for "button" HTML element.
type Button struct {
	metas.Generic

	options *ButtonOptions
}

// NewButton creates new "button" HTML element controlling structure.
func NewButton(buttonOptions *ButtonOptions) *Button {
	b := &Button{}
	b.initialize(buttonOptions)

	return b
}

// Build builds element and calls Build() for all child elements.
func (b *Button) Build() *js.Object {
	b.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementButton)

	if b.options == nil {
		return b.Object
	}

	if b.options.Class != "" {
		b.AddClassesFromString(b.options.Class)
	}

	if b.options.ID != "" {
		b.Object.Set("id", b.options.ID)
	}

	if b.options.Text != "" {
		b.Object.Set("textContent", b.options.Text)
	}

	if b.options.OnClickHandler != nil {
		b.Object.Set("onclick", b.options.OnClickHandler)
	}

	b.BuildChilds(b.Object)

	return b.Object
}

// Initializes controlling structure with necessary parameters.
func (b *Button) initialize(buttonOptions *ButtonOptions) {
	b.options = buttonOptions

	b.InitializeGeneric()
}
