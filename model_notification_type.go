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

// NotificationType the model 'NotificationType'
type NotificationType string

// List of NotificationType
const (
	NOTIFICATIONTYPE_FRIEND_REQUEST NotificationType = "friendRequest"
	NOTIFICATIONTYPE_INVITE NotificationType = "invite"
	NOTIFICATIONTYPE_INVITE_RESPONSE NotificationType = "inviteResponse"
	NOTIFICATIONTYPE_REQUEST_INVITE NotificationType = "requestInvite"
	NOTIFICATIONTYPE_REQUEST_INVITE_RESPONSE NotificationType = "requestInviteResponse"
	NOTIFICATIONTYPE_VOTETOKICK NotificationType = "votetokick"
)

// All allowed values of NotificationType enum
var AllowedNotificationTypeEnumValues = []NotificationType{
	"friendRequest",
	"invite",
	"inviteResponse",
	"requestInvite",
	"requestInviteResponse",
	"votetokick",
}

func (v *NotificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationType(value)
	for _, existing := range AllowedNotificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationType", value)
}

// NewNotificationTypeFromValue returns a pointer to a valid NotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationTypeFromValue(v string) (*NotificationType, error) {
	ev := NotificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationType: valid values are %v", v, AllowedNotificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationType) IsValid() bool {
	for _, existing := range AllowedNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationType value
func (v NotificationType) Ptr() *NotificationType {
	return &v
}

type NullableNotificationType struct {
	value *NotificationType
	isSet bool
}

func (v NullableNotificationType) Get() *NotificationType {
	return v.value
}

func (v *NullableNotificationType) Set(val *NotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationType(val *NotificationType) *NullableNotificationType {
	return &NullableNotificationType{value: val, isSet: true}
}

func (v NullableNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
