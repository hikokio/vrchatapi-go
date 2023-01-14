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

// ModerateUserRequest struct for ModerateUserRequest
type ModerateUserRequest struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Moderated string `json:"moderated"`
	Type PlayerModerationType `json:"type"`
}

// NewModerateUserRequest instantiates a new ModerateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerateUserRequest(moderated string, type_ PlayerModerationType) *ModerateUserRequest {
	this := ModerateUserRequest{}
	this.Moderated = moderated
	this.Type = type_
	return &this
}

// NewModerateUserRequestWithDefaults instantiates a new ModerateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerateUserRequestWithDefaults() *ModerateUserRequest {
	this := ModerateUserRequest{}
	var type_ PlayerModerationType = PLAYERMODERATIONTYPE_UNMUTE
	this.Type = type_
	return &this
}

// GetModerated returns the Moderated field value
func (o *ModerateUserRequest) GetModerated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Moderated
}

// GetModeratedOk returns a tuple with the Moderated field value
// and a boolean to check if the value has been set.
func (o *ModerateUserRequest) GetModeratedOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Moderated, true
}

// SetModerated sets field value
func (o *ModerateUserRequest) SetModerated(v string) {
	o.Moderated = v
}

// GetType returns the Type field value
func (o *ModerateUserRequest) GetType() PlayerModerationType {
	if o == nil {
		var ret PlayerModerationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModerateUserRequest) GetTypeOk() (*PlayerModerationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModerateUserRequest) SetType(v PlayerModerationType) {
	o.Type = v
}

func (o ModerateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["moderated"] = o.Moderated
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableModerateUserRequest struct {
	value *ModerateUserRequest
	isSet bool
}

func (v NullableModerateUserRequest) Get() *ModerateUserRequest {
	return v.value
}

func (v *NullableModerateUserRequest) Set(val *ModerateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModerateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModerateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerateUserRequest(val *ModerateUserRequest) *NullableModerateUserRequest {
	return &NullableModerateUserRequest{value: val, isSet: true}
}

func (v NullableModerateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModerateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


