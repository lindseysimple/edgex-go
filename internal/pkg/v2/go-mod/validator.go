//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

var (
	val   *validator.Validate
	trans ut.Translator
)

func NewValidator() *validator.Validate {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ = uni.GetTranslator("en")
	val = validator.New()
	_ = val.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} cannot be blank", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	return val
}

func Validate(a interface{}) error {
	err := val.Struct(a)
	// translate all error at once
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			return NewErrContractInvalid(e.Translate(trans))
		}
	}
	return nil
}
