package pigeon

import (
	. "github.com/asahasrabuddhe/pigeon/email"
	"github.com/asahasrabuddhe/pigeon/themes"
	"github.com/imdario/mergo"
)

func setDefaultEmailValues(e *Email) error {
	// Default values of an email
	defaultEmail := Email{
		Body: Body{
			Intros:     []string{},
			Dictionary: []Map{},
			Outros:     []string{},
			Signature:  "Yours truly",
			Greeting:   "Hi",
		},
	}
	// Merge the given email with default one
	// Default one overrides all zero values
	return mergo.Merge(e, defaultEmail)
}

// default values of the engine
func setDefaultPigeonValues(h *Pigeon) error {
	defaultTextDirection := LeftToRight
	defaultPigeon := Pigeon{
		Theme:         new(themes.Default),
		TextDirection: defaultTextDirection,
		Product: Product{
			Name:          "Pigeon",
			CopyrightText: "Copyright © 2017 Pigeon. All rights reserved.",
		},
	}
	// Merge the given hermes engine configuration with default one
	// Default one overrides all zero values
	err := mergo.Merge(h, defaultPigeon)
	if err != nil {
		return err
	}
	if h.TextDirection != LeftToRight && h.TextDirection != RightToLeft {
		h.TextDirection = defaultTextDirection
	}
	return nil
}
