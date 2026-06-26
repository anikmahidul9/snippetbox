package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors    map[string]string
	NonFieldErrors []string
}

//Valid() returns true if the FieldErrors map doesn't contain any entries.


//AddFieldError() adds a validation error message to the fieldErrors map so long as no entry already exists for the given key.

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

//CheckField() adds an error message to the fielderrors map only if a validation check is not ok

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

//NotBlank() returns true if a string more thn a charecter .

func NotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

//MaxChars() returns true if a string is no more than n charecters long.

func MaxChars(s string, n int) bool {
	return utf8.RuneCountInString(s) <= n
}

//PermittedValue() returns true if a string matches one of the permitted values.

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

var EmailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func (value *Validator) Valid() bool {
	return len(value.FieldErrors) == 0 && len(value.NonFieldErrors) == 0
}

func (value *Validator) AddNonFieldError(message string) {
	value.NonFieldErrors = append(value.NonFieldErrors, message)
}
