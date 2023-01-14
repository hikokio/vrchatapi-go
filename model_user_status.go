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

// UserStatus Defines the User's current status, for example \"ask me\", \"join me\" or \"offline. This status is a combined indicator of their online activity and privacy preference.
type UserStatus string

// List of UserStatus
const (
	USERSTATUS_ACTIVE UserStatus = "active"
	USERSTATUS_JOIN_ME UserStatus = "join me"
	USERSTATUS_ASK_ME UserStatus = "ask me"
	USERSTATUS_BUSY UserStatus = "busy"
	USERSTATUS_OFFLINE UserStatus = "offline"
)

// All allowed values of UserStatus enum
var AllowedUserStatusEnumValues = []UserStatus{
	"active",
	"join me",
	"ask me",
	"busy",
	"offline",
}

func (v *UserStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserStatus(value)
	for _, existing := range AllowedUserStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserStatus", value)
}

// NewUserStatusFromValue returns a pointer to a valid UserStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserStatusFromValue(v string) (*UserStatus, error) {
	ev := UserStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserStatus: valid values are %v", v, AllowedUserStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserStatus) IsValid() bool {
	for _, existing := range AllowedUserStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserStatus value
func (v UserStatus) Ptr() *UserStatus {
	return &v
}

type NullableUserStatus struct {
	value *UserStatus
	isSet bool
}

func (v NullableUserStatus) Get() *UserStatus {
	return v.value
}

func (v *NullableUserStatus) Set(val *UserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatus(val *UserStatus) *NullableUserStatus {
	return &NullableUserStatus{value: val, isSet: true}
}

func (v NullableUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
