package models

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (c CaseRecordSpec) GetIncludedFieldsInSnakeCase() []string {
	fields := c.GetIncludedFields()
	fieldsCopy := []string{}

	for _, v := range fields {
		if v == "Actors" {
			continue
		}
		fieldsCopy = append(fieldsCopy, ToSnakeCase(v))
	}
	return fieldsCopy
}

func (c CaseActionSpec) GetIncludedFieldsInSnakeCase() []string {
	fields := c.GetIncludedFields()
	fieldsCopy := []string{}

	for _, v := range fields {
		if v == "Actors" {
			continue
		}
		fieldsCopy = append(fieldsCopy, ToSnakeCase(v))
	}
	return fieldsCopy
}

// func (c CaseActionSpec) GetMap() map[string]bool {
// 	fields := c.GetIncludedFields()
// 	fieldsCopy := []string{}

// 	for _, v := range fields {
// 		if v == "Actors" {
// 			continue
// 		}
// 		fieldsCopy = append(fieldsCopy, ToSnakeCase(v))
// 	}
// 	return fieldsCopy
// }
