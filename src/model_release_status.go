/*
VRChat API Documentation


API version: 1.10.1
Contact: me@ariesclark.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
	"fmt"
)

// ReleaseStatus the model 'ReleaseStatus'
type ReleaseStatus string

// List of ReleaseStatus
const (
	RELEASESTATUS_PUBLIC ReleaseStatus = "public"
	RELEASESTATUS_PRIVATE ReleaseStatus = "private"
	RELEASESTATUS_HIDDEN ReleaseStatus = "hidden"
)

// All allowed values of ReleaseStatus enum
var AllowedReleaseStatusEnumValues = []ReleaseStatus{
	"public",
	"private",
	"hidden",
}

func (v *ReleaseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReleaseStatus(value)
	for _, existing := range AllowedReleaseStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReleaseStatus", value)
}

// NewReleaseStatusFromValue returns a pointer to a valid ReleaseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReleaseStatusFromValue(v string) (*ReleaseStatus, error) {
	ev := ReleaseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReleaseStatus: valid values are %v", v, AllowedReleaseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReleaseStatus) IsValid() bool {
	for _, existing := range AllowedReleaseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReleaseStatus value
func (v ReleaseStatus) Ptr() *ReleaseStatus {
	return &v
}

type NullableReleaseStatus struct {
	value *ReleaseStatus
	isSet bool
}

func (v NullableReleaseStatus) Get() *ReleaseStatus {
	return v.value
}

func (v *NullableReleaseStatus) Set(val *ReleaseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseStatus(val *ReleaseStatus) *NullableReleaseStatus {
	return &NullableReleaseStatus{value: val, isSet: true}
}

func (v NullableReleaseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

