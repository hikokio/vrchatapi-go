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

// WorldPublishStatus struct for WorldPublishStatus
type WorldPublishStatus struct {
	CanPubilsh bool `json:"canPubilsh"`
}

// NewWorldPublishStatus instantiates a new WorldPublishStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorldPublishStatus(canPubilsh bool) *WorldPublishStatus {
	this := WorldPublishStatus{}
	this.CanPubilsh = canPubilsh
	return &this
}

// NewWorldPublishStatusWithDefaults instantiates a new WorldPublishStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorldPublishStatusWithDefaults() *WorldPublishStatus {
	this := WorldPublishStatus{}
	var canPubilsh bool = true
	this.CanPubilsh = canPubilsh
	return &this
}

// GetCanPubilsh returns the CanPubilsh field value
func (o *WorldPublishStatus) GetCanPubilsh() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanPubilsh
}

// GetCanPubilshOk returns a tuple with the CanPubilsh field value
// and a boolean to check if the value has been set.
func (o *WorldPublishStatus) GetCanPubilshOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CanPubilsh, true
}

// SetCanPubilsh sets field value
func (o *WorldPublishStatus) SetCanPubilsh(v bool) {
	o.CanPubilsh = v
}

func (o WorldPublishStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["canPubilsh"] = o.CanPubilsh
	}
	return json.Marshal(toSerialize)
}

type NullableWorldPublishStatus struct {
	value *WorldPublishStatus
	isSet bool
}

func (v NullableWorldPublishStatus) Get() *WorldPublishStatus {
	return v.value
}

func (v *NullableWorldPublishStatus) Set(val *WorldPublishStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWorldPublishStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWorldPublishStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorldPublishStatus(val *WorldPublishStatus) *NullableWorldPublishStatus {
	return &NullableWorldPublishStatus{value: val, isSet: true}
}

func (v NullableWorldPublishStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorldPublishStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

