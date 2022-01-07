package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// POptions is a "p" HTML element configuration structure.
type POptions struct {
	Class string
	ID    string
	Text  string
}

// P is a controlling structure for "p" HTML element.
type P struct {
	metas.Generic
	options *POptions
}

// NewP creates new "p" HTML element controlling structure.
func NewP(opts *POptions) *P {
	p := &P{}
	p.initialize(opts)

	return p
}

// Build builds element and calls Build() for all child elements.
func (p *P) Build() *js.Object {
	p.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementP)

	if p.options == nil {
		return p.Object
	}

	if p.options.Class != "" {
		p.AddClassesFromString(p.options.Class)
	}

	if p.options.ID != "" {
		p.Object.Set(common.HTMLElementParameterID, p.options.ID)
	}

	if p.options.Text != "" {
		p.Object.Set("textContent", p.options.Text)
	}

	p.BuildChilds(p.Object)

	return p.Object
}

// Initializes controlling structure with necessary parameters.
func (p *P) initialize(opts *POptions) {
	p.options = opts

	p.InitializeGeneric()
}
