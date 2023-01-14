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

// SubscriptionPeriod the model 'SubscriptionPeriod'
type SubscriptionPeriod string

// List of SubscriptionPeriod
const (
	SUBSCRIPTIONPERIOD_HOUR SubscriptionPeriod = "hour"
	SUBSCRIPTIONPERIOD_DAY SubscriptionPeriod = "day"
	SUBSCRIPTIONPERIOD_WEEK SubscriptionPeriod = "week"
	SUBSCRIPTIONPERIOD_MONTH SubscriptionPeriod = "month"
	SUBSCRIPTIONPERIOD_YEAR SubscriptionPeriod = "year"
)

// All allowed values of SubscriptionPeriod enum
var AllowedSubscriptionPeriodEnumValues = []SubscriptionPeriod{
	"hour",
	"day",
	"week",
	"month",
	"year",
}

func (v *SubscriptionPeriod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionPeriod(value)
	for _, existing := range AllowedSubscriptionPeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionPeriod", value)
}

// NewSubscriptionPeriodFromValue returns a pointer to a valid SubscriptionPeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionPeriodFromValue(v string) (*SubscriptionPeriod, error) {
	ev := SubscriptionPeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionPeriod: valid values are %v", v, AllowedSubscriptionPeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionPeriod) IsValid() bool {
	for _, existing := range AllowedSubscriptionPeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionPeriod value
func (v SubscriptionPeriod) Ptr() *SubscriptionPeriod {
	return &v
}

type NullableSubscriptionPeriod struct {
	value *SubscriptionPeriod
	isSet bool
}

func (v NullableSubscriptionPeriod) Get() *SubscriptionPeriod {
	return v.value
}

func (v *NullableSubscriptionPeriod) Set(val *SubscriptionPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPeriod(val *SubscriptionPeriod) *NullableSubscriptionPeriod {
	return &NullableSubscriptionPeriod{value: val, isSet: true}
}

func (v NullableSubscriptionPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
