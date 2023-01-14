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

// GroupMemberStatus the model 'GroupMemberStatus'
type GroupMemberStatus string

// List of GroupMemberStatus
const (
	GROUPMEMBERSTATUS_INACTIVE GroupMemberStatus = "inactive"
	GROUPMEMBERSTATUS_MEMBER GroupMemberStatus = "member"
	GROUPMEMBERSTATUS_REQUESTED GroupMemberStatus = "requested"
	GROUPMEMBERSTATUS_INVITED GroupMemberStatus = "invited"
)

// All allowed values of GroupMemberStatus enum
var AllowedGroupMemberStatusEnumValues = []GroupMemberStatus{
	"inactive",
	"member",
	"requested",
	"invited",
}

func (v *GroupMemberStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupMemberStatus(value)
	for _, existing := range AllowedGroupMemberStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupMemberStatus", value)
}

// NewGroupMemberStatusFromValue returns a pointer to a valid GroupMemberStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupMemberStatusFromValue(v string) (*GroupMemberStatus, error) {
	ev := GroupMemberStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupMemberStatus: valid values are %v", v, AllowedGroupMemberStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupMemberStatus) IsValid() bool {
	for _, existing := range AllowedGroupMemberStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupMemberStatus value
func (v GroupMemberStatus) Ptr() *GroupMemberStatus {
	return &v
}

type NullableGroupMemberStatus struct {
	value *GroupMemberStatus
	isSet bool
}

func (v NullableGroupMemberStatus) Get() *GroupMemberStatus {
	return v.value
}

func (v *NullableGroupMemberStatus) Set(val *GroupMemberStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMemberStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMemberStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMemberStatus(val *GroupMemberStatus) *NullableGroupMemberStatus {
	return &NullableGroupMemberStatus{value: val, isSet: true}
}

func (v NullableGroupMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMemberStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
