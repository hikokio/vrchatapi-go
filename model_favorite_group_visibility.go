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

// FavoriteGroupVisibility the model 'FavoriteGroupVisibility'
type FavoriteGroupVisibility string

// List of FavoriteGroupVisibility
const (
	FAVORITEGROUPVISIBILITY_PRIVATE FavoriteGroupVisibility = "private"
	FAVORITEGROUPVISIBILITY_FRIENDS FavoriteGroupVisibility = "friends"
	FAVORITEGROUPVISIBILITY_PUBLIC FavoriteGroupVisibility = "public"
)

// All allowed values of FavoriteGroupVisibility enum
var AllowedFavoriteGroupVisibilityEnumValues = []FavoriteGroupVisibility{
	"private",
	"friends",
	"public",
}

func (v *FavoriteGroupVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FavoriteGroupVisibility(value)
	for _, existing := range AllowedFavoriteGroupVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FavoriteGroupVisibility", value)
}

// NewFavoriteGroupVisibilityFromValue returns a pointer to a valid FavoriteGroupVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFavoriteGroupVisibilityFromValue(v string) (*FavoriteGroupVisibility, error) {
	ev := FavoriteGroupVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FavoriteGroupVisibility: valid values are %v", v, AllowedFavoriteGroupVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FavoriteGroupVisibility) IsValid() bool {
	for _, existing := range AllowedFavoriteGroupVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FavoriteGroupVisibility value
func (v FavoriteGroupVisibility) Ptr() *FavoriteGroupVisibility {
	return &v
}

type NullableFavoriteGroupVisibility struct {
	value *FavoriteGroupVisibility
	isSet bool
}

func (v NullableFavoriteGroupVisibility) Get() *FavoriteGroupVisibility {
	return v.value
}

func (v *NullableFavoriteGroupVisibility) Set(val *FavoriteGroupVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteGroupVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteGroupVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteGroupVisibility(val *FavoriteGroupVisibility) *NullableFavoriteGroupVisibility {
	return &NullableFavoriteGroupVisibility{value: val, isSet: true}
}

func (v NullableFavoriteGroupVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteGroupVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

