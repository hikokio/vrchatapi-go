/*
VRChat API Documentation


API version: 1.10.1
Contact: me@ariesclark.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
)

// InstancePlatforms struct for InstancePlatforms
type InstancePlatforms struct {
	Android int32 `json:"android"`
	Standalonewindows int32 `json:"standalonewindows"`
}

// NewInstancePlatforms instantiates a new InstancePlatforms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancePlatforms(android int32, standalonewindows int32) *InstancePlatforms {
	this := InstancePlatforms{}
	this.Android = android
	this.Standalonewindows = standalonewindows
	return &this
}

// NewInstancePlatformsWithDefaults instantiates a new InstancePlatforms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancePlatformsWithDefaults() *InstancePlatforms {
	this := InstancePlatforms{}
	return &this
}

// GetAndroid returns the Android field value
func (o *InstancePlatforms) GetAndroid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Android
}

// GetAndroidOk returns a tuple with the Android field value
// and a boolean to check if the value has been set.
func (o *InstancePlatforms) GetAndroidOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Android, true
}

// SetAndroid sets field value
func (o *InstancePlatforms) SetAndroid(v int32) {
	o.Android = v
}

// GetStandalonewindows returns the Standalonewindows field value
func (o *InstancePlatforms) GetStandalonewindows() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Standalonewindows
}

// GetStandalonewindowsOk returns a tuple with the Standalonewindows field value
// and a boolean to check if the value has been set.
func (o *InstancePlatforms) GetStandalonewindowsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Standalonewindows, true
}

// SetStandalonewindows sets field value
func (o *InstancePlatforms) SetStandalonewindows(v int32) {
	o.Standalonewindows = v
}

func (o InstancePlatforms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["android"] = o.Android
	}
	if true {
		toSerialize["standalonewindows"] = o.Standalonewindows
	}
	return json.Marshal(toSerialize)
}

type NullableInstancePlatforms struct {
	value *InstancePlatforms
	isSet bool
}

func (v NullableInstancePlatforms) Get() *InstancePlatforms {
	return v.value
}

func (v *NullableInstancePlatforms) Set(val *InstancePlatforms) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancePlatforms) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancePlatforms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancePlatforms(val *InstancePlatforms) *NullableInstancePlatforms {
	return &NullableInstancePlatforms{value: val, isSet: true}
}

func (v NullableInstancePlatforms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancePlatforms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


