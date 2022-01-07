package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// TBody is a controlling structure for "tbody" HTML element.
type TBody struct {
	metas.Generic
}

// NewTBody creates new "tbody" HTML element controlling structure.
func NewTBody() *TBody {
	tbody := &TBody{}
	tbody.initialize()

	return tbody
}

// Build builds element and calls Build() for all child elements.
func (t *TBody) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *TBody) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTBody)
}
