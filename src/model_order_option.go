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

// OrderOption the model 'OrderOption'
type OrderOption string

// List of OrderOption
const (
	ORDEROPTION_ASCENDING OrderOption = "ascending"
	ORDEROPTION_DESCENDING OrderOption = "descending"
)

// All allowed values of OrderOption enum
var AllowedOrderOptionEnumValues = []OrderOption{
	"ascending",
	"descending",
}

func (v *OrderOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderOption(value)
	for _, existing := range AllowedOrderOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderOption", value)
}

// NewOrderOptionFromValue returns a pointer to a valid OrderOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderOptionFromValue(v string) (*OrderOption, error) {
	ev := OrderOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderOption: valid values are %v", v, AllowedOrderOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderOption) IsValid() bool {
	for _, existing := range AllowedOrderOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderOption value
func (v OrderOption) Ptr() *OrderOption {
	return &v
}

type NullableOrderOption struct {
	value *OrderOption
	isSet bool
}

func (v NullableOrderOption) Get() *OrderOption {
	return v.value
}

func (v *NullableOrderOption) Set(val *OrderOption) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOption(val *OrderOption) *NullableOrderOption {
	return &NullableOrderOption{value: val, isSet: true}
}

func (v NullableOrderOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
