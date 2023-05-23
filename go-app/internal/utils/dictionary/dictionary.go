// Package dictionary - справочники.
package dictionary

import "github.com/go-playground/validator/v10"

// LayoutToString - шаблон для изменения формата time.Time в string.
const LayoutToString = "2006-01-02 15:04:05"

// HourDivider - Количество часов для деления.
const HourDivider = 24

// Validate - движок валидации.
var Validate = validator.New()
