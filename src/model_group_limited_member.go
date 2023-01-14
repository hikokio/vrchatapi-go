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

// GroupLimitedMember struct for GroupLimitedMember
type GroupLimitedMember struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId *string `json:"userId,omitempty"`
	IsRepresenting *bool `json:"isRepresenting,omitempty"`
}

// NewGroupLimitedMember instantiates a new GroupLimitedMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLimitedMember() *GroupLimitedMember {
	this := GroupLimitedMember{}
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	return &this
}

// NewGroupLimitedMemberWithDefaults instantiates a new GroupLimitedMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLimitedMemberWithDefaults() *GroupLimitedMember {
	this := GroupLimitedMember{}
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupLimitedMember) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupLimitedMember) SetGroupId(v string) {
	o.GroupId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GroupLimitedMember) SetUserId(v string) {
	o.UserId = &v
}

// GetIsRepresenting returns the IsRepresenting field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetIsRepresenting() bool {
	if o == nil || isNil(o.IsRepresenting) {
		var ret bool
		return ret
	}
	return *o.IsRepresenting
}

// GetIsRepresentingOk returns a tuple with the IsRepresenting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIsRepresentingOk() (*bool, bool) {
	if o == nil || isNil(o.IsRepresenting) {
    return nil, false
	}
	return o.IsRepresenting, true
}

// HasIsRepresenting returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasIsRepresenting() bool {
	if o != nil && !isNil(o.IsRepresenting) {
		return true
	}

	return false
}

// SetIsRepresenting gets a reference to the given bool and assigns it to the IsRepresenting field.
func (o *GroupLimitedMember) SetIsRepresenting(v bool) {
	o.IsRepresenting = &v
}

func (o GroupLimitedMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.IsRepresenting) {
		toSerialize["isRepresenting"] = o.IsRepresenting
	}
	return json.Marshal(toSerialize)
}

type NullableGroupLimitedMember struct {
	value *GroupLimitedMember
	isSet bool
}

func (v NullableGroupLimitedMember) Get() *GroupLimitedMember {
	return v.value
}

func (v *NullableGroupLimitedMember) Set(val *GroupLimitedMember) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLimitedMember) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLimitedMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLimitedMember(val *GroupLimitedMember) *NullableGroupLimitedMember {
	return &NullableGroupLimitedMember{value: val, isSet: true}
}

func (v NullableGroupLimitedMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLimitedMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


