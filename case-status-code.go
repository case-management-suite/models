package models

import (
	"reflect"
)

type CaseActionType struct {
	Create       string
	Start        string
	MoveNext     string
	MovePrevious string
	Close        string
}

var CaseActions = CaseActionType{
	Create:       "CREATE",
	Start:        "START",
	MoveNext:     "MOVE_NEXT",
	MovePrevious: "MOVE_PREVIOUS",
	Close:        "CLOSE",
}

var BaseSupportedActions = map[string]string{
	CaseActions.Create:       CaseActions.Create,
	CaseActions.Start:        CaseActions.Start,
	CaseActions.MovePrevious: CaseActions.MovePrevious,
	CaseActions.Close:        CaseActions.Close,
}

func (cst CaseActionType) GetValue(str string) Optional[string] {
	t := reflect.TypeOf(CaseActions)
	v := reflect.ValueOf(CaseActions)

	for _, f := range reflect.VisibleFields(t) {
		fVal := v.FieldByName(f.Name).String()
		if fVal == str {
			return OptionalOf(fVal)
		}
	}
	return OptionalEmpty[string]()
}

type CaseStatusType struct {
	NewCase string
	Started string
	Closed  string
	UNKOWN  string
}

func (cst CaseStatusType) GetValue(str string) Optional[string] {
	for _, f := range reflect.VisibleFields(reflect.TypeOf(CaseStatusType{})) {
		if f.Name == str {
			return OptionalOf(reflect.ValueOf(f).String())
		}
	}
	return OptionalEmpty[string]()
}

var CaseStatusList = CaseStatusType{
	NewCase: "NEW_CASE",
	Started: "STARTED",
	Closed:  "CLOSED",
	UNKOWN:  "UNKOWN",
}
