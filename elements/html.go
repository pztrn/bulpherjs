package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// HTML is a controlling structure for HTML page.
type HTML struct {
	metas.Generic

	Head *Head
	Body *Body
}

// NewHTML creates new HTML page.
func NewHTML() *HTML {
	h := &HTML{}
	h.initialize()

	return h
}

// Build builds HTML and all it's childs. For HTML object it'll return
// nil.
func (h *HTML) Build() *js.Object {
	h.BuildChilds(h.Object)

	js.Global.Get(common.HTMLElementDocument).Set(common.HTMLElementHTML, h.Object)

	return nil
}

// Initializes controlling structure.
func (h *HTML) initialize() {
	h.InitializeGeneric()

	h.Object = js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementDocumentElement)

	h.Head = NewHead(&HeadOptions{})
	h.Body = NewBody()

	h.AddChild(h.Head)
	h.AddChild(h.Body)

	// Set default title. Use Body.SetTitle to override in your application.
	h.Head.SetTitle("BulpherJS default application")
}
