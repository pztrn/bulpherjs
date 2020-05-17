// BulpherJS - bulma for GopherJS.
// Copyright (c) 2020 by Stanislav Nikitin <pztrn@pztrn.name>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs"
	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/elements"
)

func main() {
	app := bulpherjs.NewApplication(&bulpherjs.ApplicationOptions{
		Bulma: &bulpherjs.BulmaOptions{
			Version: "0.8.0",
		},
		Name: "BulpherJS example application a.k.a. Hello World!",
	})

	app.SetStartFunction(start)
	app.Start()
	app.SetTitle("Hello world!")
}

func start(a *bulpherjs.Application) {
	mainSection := elements.NewSection(nil)

	mainDiv := elements.NewDiv(&elements.DivOptions{
		Class: "container",
		ID:    "mainContainer",
	})
	mainSection.AddChild(mainDiv)

	// Columns divs.
	columnsDiv := elements.NewDiv(&elements.DivOptions{
		Class: "columns is-centered",
		ID:    "mainColumns",
	})
	mainDiv.AddChild(columnsDiv)

	// Click-Me button should be in the middle of screen.
	centerDiv := elements.NewDiv(&elements.DivOptions{
		Class: "column is-half",
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

	constructTable(mainDiv)

	// Build final document.
	a.HTML.Body.AddChild(mainSection)
}

func constructTable(mainDiv *elements.Div) {
	// Columns divs.
	columnsDiv := elements.NewDiv(&elements.DivOptions{
		Class: "columns is-centered",
		ID:    "tableColumns",
	})
	mainDiv.AddChild(columnsDiv)

	tableDiv := elements.NewDiv(&elements.DivOptions{
		Class: "column is-half",
	})
	columnsDiv.AddChild(tableDiv)

	testTable := elements.NewTable(&elements.TableOptions{
		IsBordered:  true,
		IsFullWidth: true,
		IsHoverable: true,
		IsNarrow:    true,
	})
	tableDiv.AddChild(testTable)

	tableHeader := elements.NewTHead()
	testTable.AddChild(tableHeader)

	tableHeaderLine := elements.NewTR()
	tableHeader.AddChild(tableHeaderLine)

	thlOne := elements.NewTH()
	thlOne.SetTextContent("Header")
	tableHeaderLine.AddChild(thlOne)

	thlTwo := elements.NewTH()
	thlTwo.SetTextContent("line")
	tableHeaderLine.AddChild(thlTwo)

	tableBody := elements.NewTBody()
	testTable.AddChild(tableBody)

	tableBodyLine := elements.NewTR()
	tableBody.AddChild(tableBodyLine)

	tblOne := elements.NewTD()
	tblOne.SetTextContent("Body")
	tableBodyLine.AddChild(tblOne)

	tblTwo := elements.NewTD()
	tblTwo.SetTextContent("line")
	tableBodyLine.AddChild(tblTwo)

	tableFooter := elements.NewTFoot()
	testTable.AddChild(tableFooter)

	tableFooterLine := elements.NewTR()
	tableFooter.AddChild(tableFooterLine)

	tflOne := elements.NewTD()
	tflOne.SetTextContent("Footer")
	tableFooterLine.AddChild(tflOne)

	tflTwo := elements.NewTD()
	tflTwo.SetTextContent("line")
	tableFooterLine.AddChild(tflTwo)
}
