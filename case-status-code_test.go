package models_test

import (
	"testing"

	"github.com/case-management-suite/models"
)

func TestParseAction(t *testing.T) {
	maybeType := models.CaseActionType{}.GetValue("START")
	if maybeType.IsEmpty() {
		t.Fatal("Type shouldn't be empty")
	}

	actual := *maybeType.Get()
	if models.CaseActions.Start != actual {
		t.Fatalf("The action should be correct: Was %v but expected %v", actual, models.CaseActions.Start)
	}
}
