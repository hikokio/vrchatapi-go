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

// Verify2FAEmailCodeResult struct for Verify2FAEmailCodeResult
type Verify2FAEmailCodeResult struct {
	Verified bool `json:"verified"`
}

// NewVerify2FAEmailCodeResult instantiates a new Verify2FAEmailCodeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerify2FAEmailCodeResult(verified bool) *Verify2FAEmailCodeResult {
	this := Verify2FAEmailCodeResult{}
	this.Verified = verified
	return &this
}

// NewVerify2FAEmailCodeResultWithDefaults instantiates a new Verify2FAEmailCodeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerify2FAEmailCodeResultWithDefaults() *Verify2FAEmailCodeResult {
	this := Verify2FAEmailCodeResult{}
	return &this
}

// GetVerified returns the Verified field value
func (o *Verify2FAEmailCodeResult) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *Verify2FAEmailCodeResult) GetVerifiedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *Verify2FAEmailCodeResult) SetVerified(v bool) {
	o.Verified = v
}

func (o Verify2FAEmailCodeResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableVerify2FAEmailCodeResult struct {
	value *Verify2FAEmailCodeResult
	isSet bool
}

func (v NullableVerify2FAEmailCodeResult) Get() *Verify2FAEmailCodeResult {
	return v.value
}

func (v *NullableVerify2FAEmailCodeResult) Set(val *Verify2FAEmailCodeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerify2FAEmailCodeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerify2FAEmailCodeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerify2FAEmailCodeResult(val *Verify2FAEmailCodeResult) *NullableVerify2FAEmailCodeResult {
	return &NullableVerify2FAEmailCodeResult{value: val, isSet: true}
}

func (v NullableVerify2FAEmailCodeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerify2FAEmailCodeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


