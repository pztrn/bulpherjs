package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// SectionOptions is a "section" HTML element configuration structure.
type SectionOptions struct {
	Class string
}

// Section is a controlling structure for "section" HTML element.
type Section struct {
	metas.Generic

	options *SectionOptions
}

// NewSection creates new "section" HTML element controlling structure.
func NewSection(sectionOptions *SectionOptions) *Section {
	s := &Section{}
	s.initialize(sectionOptions)

	return s
}

// Build builds element and calls Build() for all child elements.
func (s *Section) Build() *js.Object {
	s.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementSection)

	if s.options.Class != "" {
		s.AddClassesFromString(s.options.Class)
	}

	s.BuildChilds(s.Object)

	return s.Object
}

// Initializes controlling structure with necessary parameters.
func (s *Section) initialize(sectionOptions *SectionOptions) {
	s.options = sectionOptions

	s.InitializeGeneric()
}
