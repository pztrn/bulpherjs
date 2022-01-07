package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

const (
	bulmaVersion = "0.8.0"
)

// HeadOptions is a "head" HTML head element configuration structure.
type HeadOptions struct{}

// Head is a controlling structure for "head" HTML head element.
type Head struct {
	metas.Generic

	options *HeadOptions
}

// NewHeader creates new "head" HTML element controlling structure.
func NewHead(opts *HeadOptions) *Head {
	hw := &Head{}
	hw.initialize(opts)

	return hw
}

// Build builds header.
func (hw *Head) Build() *js.Object {
	hw.BuildChilds(hw.Object)

	return hw.Object
}

// Initializes controlling structure.
func (hw *Head) initialize(opts *HeadOptions) {
	hw.options = opts

	hw.InitializeGeneric()

	hw.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallGetElementsByTagName, common.HTMLElementHead).Index(0)

	// Add default header elements.
	hw.AddChild(NewLink(&LinkOptions{
		Href: "//cdn.jsdelivr.net/npm/bulma@" + bulmaVersion + "/css/bulma.min.css",
		ID:   "",
		Rel:  "stylesheet",
	}))

	hw.AddChild(NewMeta(&MetaOptions{
		Content: "width=device-width, initial-scale=1",
		Name:    "viewport",
	}))
}

// SetTitle sets title for whole page.
func (h *Head) SetTitle(title string) {
	js.Global.Get(common.HTMLElementDocument).Set(common.HTMLHeadElementTitle, title)
}
