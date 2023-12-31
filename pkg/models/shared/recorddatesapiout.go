// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/utils"
	"time"
)

// RecordDatesAPIOut - Date information associated with the record
type RecordDatesAPIOut struct {
	// The moment this record was completed if the record is currently in an end step measured in milliseconds since the Unix epoch
	Completed *time.Time `json:"completed,omitempty"`
	// The moment this Risk Cloud resource was created measured in milliseconds since the Unix epoch.
	Created *time.Time `json:"created,omitempty"`
	// The number of days until this record is due
	DaysUntilDue *int64 `json:"daysUntilDue,omitempty"`
	// The due date of this record and step SLA measured in milliseconds since the Unix epoch
	DueDate *time.Time `json:"dueDate,omitempty"`
	// The moment this record was last completed regardless of if the record has transitioned from an end step measured in milliseconds since the Unix epoch
	LastCompleted *time.Time `json:"lastCompleted,omitempty"`
	// The due date of this record measured in milliseconds since the Unix epoch
	RecordDueDate *time.Time `json:"recordDueDate,omitempty"`
	// The moment this Risk Cloud resource was last updated measured in milliseconds since the Unix epoch.
	Updated *time.Time `json:"updated,omitempty"`
}

func (r RecordDatesAPIOut) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RecordDatesAPIOut) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RecordDatesAPIOut) GetCompleted() *time.Time {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *RecordDatesAPIOut) GetCreated() *time.Time {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *RecordDatesAPIOut) GetDaysUntilDue() *int64 {
	if o == nil {
		return nil
	}
	return o.DaysUntilDue
}

func (o *RecordDatesAPIOut) GetDueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *RecordDatesAPIOut) GetLastCompleted() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastCompleted
}

func (o *RecordDatesAPIOut) GetRecordDueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.RecordDueDate
}

func (o *RecordDatesAPIOut) GetUpdated() *time.Time {
	if o == nil {
		return nil
	}
	return o.Updated
}
