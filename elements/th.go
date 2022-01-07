package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// TH is a controlling structure for "th" HTML element.
type TH struct {
	metas.Generic
}

// NewTH creates new "th" HTML element controlling structure.
func NewTH() *TH {
	th := &TH{}
	th.initialize()

	return th
}

// Build builds element and calls Build() for all child elements.
func (t *TH) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *TH) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTH)
}
