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

// InstanceShortNameResponse struct for InstanceShortNameResponse
type InstanceShortNameResponse struct {
	SecureName string `json:"secureName"`
	ShortName *string `json:"shortName,omitempty"`
}

// NewInstanceShortNameResponse instantiates a new InstanceShortNameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceShortNameResponse(secureName string) *InstanceShortNameResponse {
	this := InstanceShortNameResponse{}
	this.SecureName = secureName
	return &this
}

// NewInstanceShortNameResponseWithDefaults instantiates a new InstanceShortNameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceShortNameResponseWithDefaults() *InstanceShortNameResponse {
	this := InstanceShortNameResponse{}
	return &this
}

// GetSecureName returns the SecureName field value
func (o *InstanceShortNameResponse) GetSecureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureName
}

// GetSecureNameOk returns a tuple with the SecureName field value
// and a boolean to check if the value has been set.
func (o *InstanceShortNameResponse) GetSecureNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SecureName, true
}

// SetSecureName sets field value
func (o *InstanceShortNameResponse) SetSecureName(v string) {
	o.SecureName = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InstanceShortNameResponse) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceShortNameResponse) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InstanceShortNameResponse) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InstanceShortNameResponse) SetShortName(v string) {
	o.ShortName = &v
}

func (o InstanceShortNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secureName"] = o.SecureName
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceShortNameResponse struct {
	value *InstanceShortNameResponse
	isSet bool
}

func (v NullableInstanceShortNameResponse) Get() *InstanceShortNameResponse {
	return v.value
}

func (v *NullableInstanceShortNameResponse) Set(val *InstanceShortNameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceShortNameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceShortNameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceShortNameResponse(val *InstanceShortNameResponse) *NullableInstanceShortNameResponse {
	return &NullableInstanceShortNameResponse{value: val, isSet: true}
}

func (v NullableInstanceShortNameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceShortNameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


