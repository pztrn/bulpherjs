package widgets

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// Body is a controlling structure for "body" HTML head element.
type Body struct {
	metas.Generic
}

// NewBody creates new "body" HTML element controlling structure.
func NewBody() *Body {
	b := &Body{}
	b.initialize()

	return b
}

// Build builds child elements.
func (b *Body) Build() *js.Object {
	// This function is special - we should try to get <body> element only
	// after DOMContentLoaded event was fired, otherwise there is no
	// guarantee that <body> will ever exist.
	b.Object = js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementBody)

	b.BuildChilds(b.Object)

	return b.Object
}

// Initializes controlling structure.
func (b *Body) initialize() {
	b.InitializeGeneric()
}
