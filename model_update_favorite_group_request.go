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

// UpdateFavoriteGroupRequest struct for UpdateFavoriteGroupRequest
type UpdateFavoriteGroupRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
	Visibility *FavoriteGroupVisibility `json:"visibility,omitempty"`
	// Tags on FavoriteGroups are believed to do nothing.
	Tags []string `json:"tags,omitempty"`
}

// NewUpdateFavoriteGroupRequest instantiates a new UpdateFavoriteGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFavoriteGroupRequest() *UpdateFavoriteGroupRequest {
	this := UpdateFavoriteGroupRequest{}
	var visibility FavoriteGroupVisibility = FAVORITEGROUPVISIBILITY_PRIVATE
	this.Visibility = &visibility
	return &this
}

// NewUpdateFavoriteGroupRequestWithDefaults instantiates a new UpdateFavoriteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFavoriteGroupRequestWithDefaults() *UpdateFavoriteGroupRequest {
	this := UpdateFavoriteGroupRequest{}
	var visibility FavoriteGroupVisibility = FAVORITEGROUPVISIBILITY_PRIVATE
	this.Visibility = &visibility
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateFavoriteGroupRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFavoriteGroupRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateFavoriteGroupRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateFavoriteGroupRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateFavoriteGroupRequest) GetVisibility() FavoriteGroupVisibility {
	if o == nil || isNil(o.Visibility) {
		var ret FavoriteGroupVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFavoriteGroupRequest) GetVisibilityOk() (*FavoriteGroupVisibility, bool) {
	if o == nil || isNil(o.Visibility) {
    return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *UpdateFavoriteGroupRequest) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given FavoriteGroupVisibility and assigns it to the Visibility field.
func (o *UpdateFavoriteGroupRequest) SetVisibility(v FavoriteGroupVisibility) {
	o.Visibility = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateFavoriteGroupRequest) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFavoriteGroupRequest) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateFavoriteGroupRequest) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateFavoriteGroupRequest) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateFavoriteGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFavoriteGroupRequest struct {
	value *UpdateFavoriteGroupRequest
	isSet bool
}

func (v NullableUpdateFavoriteGroupRequest) Get() *UpdateFavoriteGroupRequest {
	return v.value
}

func (v *NullableUpdateFavoriteGroupRequest) Set(val *UpdateFavoriteGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFavoriteGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFavoriteGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFavoriteGroupRequest(val *UpdateFavoriteGroupRequest) *NullableUpdateFavoriteGroupRequest {
	return &NullableUpdateFavoriteGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateFavoriteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFavoriteGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

