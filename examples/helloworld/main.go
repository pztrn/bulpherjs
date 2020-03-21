package main

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/elements"
)

func main() {
	document := js.Global.Get("document")
	document.Call("addEventListener", "DOMContentLoaded", func(event *js.Object) {
		js.Global.Get("console").Call("log", "Hello world!")

		constructHeader()
		constructBody()
	})
}

func constructBody() {
	mainSection := elements.NewSection(nil)

	mainDiv := elements.NewDiv(&elements.DivOptions{
		Class: "container",
		ID:    "mainContainer",
	})
	mainSection.AddChild(mainDiv)

	// Columns divs.
	columnsDiv := elements.NewDiv(&elements.DivOptions{
		Class: "columns",
		ID:    "mainColumns",
	})
	mainDiv.AddChild(columnsDiv)

	// Click-Me button should be in the middle of screen.
	centerDiv := elements.NewDiv(&elements.DivOptions{
		Class: "column is-half is-offset-two-fifths",
	})
	columnsDiv.AddChild(centerDiv)

	startTestButton := elements.NewButton(&elements.ButtonOptions{
		Class: "is-large is-rounded is-primary",
		ID:    "main_button",
		Text:  "Click Me!",
		OnClickHandler: func(e *js.Object) {
			common.Log("Button clicked! Hooray")
			js.Global.Get(common.HTMLElementDocument).Call("getElementById", "main_button").Set("textContent", "Clicked!")
		},
	})
	centerDiv.AddChild(startTestButton)

	js.Global.Get(common.HTMLElementDocument).
		Get(common.HTMLElementBody).
		Call(common.JSCallAppendChild, mainSection.Build())
}

func constructHeader() {
	link := elements.NewLink(&elements.LinkOptions{
		Href: "//cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css",
		ID:   "",
		Rel:  "stylesheet",
	})

	js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementHead).Call(common.JSCallAppendChild, link.Build())

	metaViewport := elements.NewMeta(&elements.MetaOptions{
		Content: "width=device-width, initial-scale=1",
		Name:    "viewport",
	})

	js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementHead).Call(common.JSCallAppendChild, metaViewport.Build())

	js.Global.Get(common.HTMLElementDocument).Set(common.HTMLHeadElementTitle, "BulpherJS Hello World application")
}
