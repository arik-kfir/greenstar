// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Operation struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Status      OperationStatus `json:"status"`
	Result      OperationResult `json:"result"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
}

type OperationChanges struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Status      OperationStatus `json:"status"`
	Result      OperationResult `json:"result"`
}

type OperationResult string

const (
	OperationResultSucceeded OperationResult = "SUCCEEDED"
	OperationResultFailed    OperationResult = "FAILED"
)

var AllOperationResult = []OperationResult{
	OperationResultSucceeded,
	OperationResultFailed,
}

func (e OperationResult) IsValid() bool {
	switch e {
	case OperationResultSucceeded, OperationResultFailed:
		return true
	}
	return false
}

func (e OperationResult) String() string {
	return string(e)
}

func (e *OperationResult) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationResult(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationResult", str)
	}
	return nil
}

func (e OperationResult) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OperationStatus string

const (
	OperationStatusPending   OperationStatus = "PENDING"
	OperationStatusAccepted  OperationStatus = "ACCEPTED"
	OperationStatusRejected  OperationStatus = "REJECTED"
	OperationStatusStarted   OperationStatus = "STARTED"
	OperationStatusCompleted OperationStatus = "COMPLETED"
)

var AllOperationStatus = []OperationStatus{
	OperationStatusPending,
	OperationStatusAccepted,
	OperationStatusRejected,
	OperationStatusStarted,
	OperationStatusCompleted,
}

func (e OperationStatus) IsValid() bool {
	switch e {
	case OperationStatusPending, OperationStatusAccepted, OperationStatusRejected, OperationStatusStarted, OperationStatusCompleted:
		return true
	}
	return false
}

func (e OperationStatus) String() string {
	return string(e)
}

func (e *OperationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationStatus", str)
	}
	return nil
}

func (e OperationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
