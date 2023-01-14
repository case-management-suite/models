package models

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID(namespace string) string {
	return fmt.Sprintf("urn:%s:%s", namespace, uuid.NewString())
}

func NewCaseRecordUUID() string {
	id := UUID("case_record")
	// log.Debug().Stack().Str("UUID", id).Str("service", "uuid_util").Msg("New CaseRecordUUID")
	return id
}

func NewCaseActionUUID() string {
	return UUID("case_action")
}

func NewCaseNotificationUUID() string {
	return UUID("case_notification")
}
