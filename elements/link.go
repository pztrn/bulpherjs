package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// LinkOptions is a "link" HTML head element configuration structure.
type LinkOptions struct {
	Href string
	ID   string
	Rel  string
}

// Link is a controlling structure for "link" HTML head element.
type Link struct {
	options *LinkOptions
}

// NewLink creates new "link" HTML head controlling structure.
func NewLink(linkOptions *LinkOptions) *Link {
	l := &Link{}
	l.initialize(linkOptions)

	return l
}

// Build builds element.
func (l *Link) Build() *js.Object {
	link := js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementLink)

	if l.options == nil {
		return link
	}

	if l.options.Href != "" {
		link.Set(common.HTMLElementParameterHref, l.options.Href)
	}

	if l.options.Rel != "" {
		link.Set(common.HTMLElementParameterRel, l.options.Rel)
	}

	return link
}

// Initializes controlling structure with necessary parameters.
func (l *Link) initialize(linkOptions *LinkOptions) {
	l.options = linkOptions
}
