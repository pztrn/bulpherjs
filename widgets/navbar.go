package widgets

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/elements"
)

// NavBarOptions is a "nav" HTML body element configuration structure.
type NavBarOptions struct {
	IsDark bool
	Title  string
}

// NavBar is a controlling structure for "nav" HTML body element.
type NavBar struct {
	options *NavBarOptions

	navbar *elements.Nav
}

// NewNavBar creates new "nav" HTML element controlling structure.
func NewNavBar(opts *NavBarOptions) *NavBar {
	n := &NavBar{}
	n.initialize(opts)

	return n
}

// Build builds element and calls Build() for all child elements.
func (n *NavBar) Build() *js.Object {
	return n.navbar.Build()
}

// Initializes controlling structure.
func (n *NavBar) initialize(opts *NavBarOptions) {
	n.options = opts

	navopts := &elements.NavOptions{
		Class: "navbar",
		Data:  map[string]string{},
		Role:  "navigation",
	}

	n.navbar = elements.NewNav(navopts)

	if n.options.IsDark {
		n.navbar.AddClassesFromString("is-dark")
	}

	if n.options.Title != "" {
		brandDiv := elements.NewDiv(&elements.DivOptions{
			Class: "navbar-brand",
		})

		n.navbar.AddChild(brandDiv)

		navMainLink := elements.NewA(&elements.AOptions{
			Class: "navbar-item",
			Href:  "/",
			Text:  n.options.Title,
		})

		brandDiv.AddChild(navMainLink)
	}
}
