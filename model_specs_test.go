package models_test

import (
	"reflect"
	"sort"
	"testing"

	"casemanagementsuite.com/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestGet(t *testing.T) {
	spec := models.NewCaseRecordSpec(true)
	val, err := spec.Get("Status")
	if err != nil {
		log.Logger.Fatal().Bool("Result", val)
		t.Fail()
	}
	if val != true {
		log.Logger.Fatal().Bool("Result", val).Msg("Value should be true")
		t.Fail()
	}
}

func TestInitTrue(t *testing.T) {
	spec := models.NewCaseRecordSpec(true)
	if spec.Status != true {
		log.Logger.Fatal().Bool("Result", spec.Status).Msg("Value should be false")
		t.Fail()
	}
}

func TestInitFalse(t *testing.T) {
	spec := models.NewCaseRecordSpec(false)
	if spec.Status == true {
		log.Logger.Fatal().Bool("Result", spec.Status).Msg("Value should be false")
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	spec := models.NewCaseRecordSpec(false)
	err := spec.Set("Status", true)
	if err != nil {
		log.Logger.Fatal().Bool("Result", spec.Status)
		t.Fail()
	}
	if spec.Status != true {
		log.Logger.Fatal().Bool("Result", spec.Status).Msg("Value should be true")
		t.Fail()
	}
}

func TestUnkown(t *testing.T) {
	spec := models.NewCaseRecordSpec(false)
	val, err := spec.Get("Someting")
	if err == nil {
		log.Logger.Fatal().Bool("Result", val)
		t.Fail()
	}
}

func TestIncludedFields(t *testing.T) {
	spec := models.NewCaseRecordSpec(false)
	spec.Status = true
	spec.CreatedAt = true

	fields := spec.GetIncludedFields()

	expectedFields := []string{"Status", "CreatedAt"}

	sort.Strings(fields)
	sort.Strings(expectedFields)

	actual := zerolog.Arr()
	for _, v := range fields {
		actual = actual.Str(v)
	}

	if !reflect.DeepEqual(fields, expectedFields) {
		log.Logger.Fatal().Array("actual", actual).Msg("Fields should be equal")
		t.Fail()
	}
}
