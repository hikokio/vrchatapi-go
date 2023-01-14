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

// CreateGroupGalleryRequest struct for CreateGroupGalleryRequest
type CreateGroupGalleryRequest struct {
	// Name of the gallery.
	Name string `json:"name"`
	// Description of the gallery.
	Description *string `json:"description,omitempty"`
	// Whether the gallery is members only.
	MembersOnly *bool `json:"membersOnly,omitempty"`
	//  
	RoleIdsToView []string `json:"roleIdsToView,omitempty"`
	//  
	RoleIdsToSubmit []string `json:"roleIdsToSubmit,omitempty"`
	//  
	RoleIdsToAutoApprove []string `json:"roleIdsToAutoApprove,omitempty"`
	//  
	RoleIdsToManage []string `json:"roleIdsToManage,omitempty"`
}

// NewCreateGroupGalleryRequest instantiates a new CreateGroupGalleryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupGalleryRequest(name string) *CreateGroupGalleryRequest {
	this := CreateGroupGalleryRequest{}
	this.Name = name
	var membersOnly bool = false
	this.MembersOnly = &membersOnly
	return &this
}

// NewCreateGroupGalleryRequestWithDefaults instantiates a new CreateGroupGalleryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupGalleryRequestWithDefaults() *CreateGroupGalleryRequest {
	this := CreateGroupGalleryRequest{}
	var membersOnly bool = false
	this.MembersOnly = &membersOnly
	return &this
}

// GetName returns the Name field value
func (o *CreateGroupGalleryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGroupGalleryRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGroupGalleryRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGroupGalleryRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupGalleryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGroupGalleryRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMembersOnly returns the MembersOnly field value if set, zero value otherwise.
func (o *CreateGroupGalleryRequest) GetMembersOnly() bool {
	if o == nil || isNil(o.MembersOnly) {
		var ret bool
		return ret
	}
	return *o.MembersOnly
}

// GetMembersOnlyOk returns a tuple with the MembersOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupGalleryRequest) GetMembersOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.MembersOnly) {
    return nil, false
	}
	return o.MembersOnly, true
}

// HasMembersOnly returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasMembersOnly() bool {
	if o != nil && !isNil(o.MembersOnly) {
		return true
	}

	return false
}

// SetMembersOnly gets a reference to the given bool and assigns it to the MembersOnly field.
func (o *CreateGroupGalleryRequest) SetMembersOnly(v bool) {
	o.MembersOnly = &v
}

// GetRoleIdsToView returns the RoleIdsToView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupGalleryRequest) GetRoleIdsToView() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToView
}

// GetRoleIdsToViewOk returns a tuple with the RoleIdsToView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupGalleryRequest) GetRoleIdsToViewOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIdsToView) {
    return nil, false
	}
	return o.RoleIdsToView, true
}

// HasRoleIdsToView returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasRoleIdsToView() bool {
	if o != nil && isNil(o.RoleIdsToView) {
		return true
	}

	return false
}

// SetRoleIdsToView gets a reference to the given []string and assigns it to the RoleIdsToView field.
func (o *CreateGroupGalleryRequest) SetRoleIdsToView(v []string) {
	o.RoleIdsToView = v
}

// GetRoleIdsToSubmit returns the RoleIdsToSubmit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupGalleryRequest) GetRoleIdsToSubmit() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToSubmit
}

// GetRoleIdsToSubmitOk returns a tuple with the RoleIdsToSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupGalleryRequest) GetRoleIdsToSubmitOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIdsToSubmit) {
    return nil, false
	}
	return o.RoleIdsToSubmit, true
}

// HasRoleIdsToSubmit returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasRoleIdsToSubmit() bool {
	if o != nil && isNil(o.RoleIdsToSubmit) {
		return true
	}

	return false
}

// SetRoleIdsToSubmit gets a reference to the given []string and assigns it to the RoleIdsToSubmit field.
func (o *CreateGroupGalleryRequest) SetRoleIdsToSubmit(v []string) {
	o.RoleIdsToSubmit = v
}

// GetRoleIdsToAutoApprove returns the RoleIdsToAutoApprove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupGalleryRequest) GetRoleIdsToAutoApprove() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToAutoApprove
}

// GetRoleIdsToAutoApproveOk returns a tuple with the RoleIdsToAutoApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupGalleryRequest) GetRoleIdsToAutoApproveOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIdsToAutoApprove) {
    return nil, false
	}
	return o.RoleIdsToAutoApprove, true
}

// HasRoleIdsToAutoApprove returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasRoleIdsToAutoApprove() bool {
	if o != nil && isNil(o.RoleIdsToAutoApprove) {
		return true
	}

	return false
}

// SetRoleIdsToAutoApprove gets a reference to the given []string and assigns it to the RoleIdsToAutoApprove field.
func (o *CreateGroupGalleryRequest) SetRoleIdsToAutoApprove(v []string) {
	o.RoleIdsToAutoApprove = v
}

// GetRoleIdsToManage returns the RoleIdsToManage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupGalleryRequest) GetRoleIdsToManage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToManage
}

// GetRoleIdsToManageOk returns a tuple with the RoleIdsToManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupGalleryRequest) GetRoleIdsToManageOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIdsToManage) {
    return nil, false
	}
	return o.RoleIdsToManage, true
}

// HasRoleIdsToManage returns a boolean if a field has been set.
func (o *CreateGroupGalleryRequest) HasRoleIdsToManage() bool {
	if o != nil && isNil(o.RoleIdsToManage) {
		return true
	}

	return false
}

// SetRoleIdsToManage gets a reference to the given []string and assigns it to the RoleIdsToManage field.
func (o *CreateGroupGalleryRequest) SetRoleIdsToManage(v []string) {
	o.RoleIdsToManage = v
}

func (o CreateGroupGalleryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.MembersOnly) {
		toSerialize["membersOnly"] = o.MembersOnly
	}
	if o.RoleIdsToView != nil {
		toSerialize["roleIdsToView"] = o.RoleIdsToView
	}
	if o.RoleIdsToSubmit != nil {
		toSerialize["roleIdsToSubmit"] = o.RoleIdsToSubmit
	}
	if o.RoleIdsToAutoApprove != nil {
		toSerialize["roleIdsToAutoApprove"] = o.RoleIdsToAutoApprove
	}
	if o.RoleIdsToManage != nil {
		toSerialize["roleIdsToManage"] = o.RoleIdsToManage
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupGalleryRequest struct {
	value *CreateGroupGalleryRequest
	isSet bool
}

func (v NullableCreateGroupGalleryRequest) Get() *CreateGroupGalleryRequest {
	return v.value
}

func (v *NullableCreateGroupGalleryRequest) Set(val *CreateGroupGalleryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupGalleryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupGalleryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupGalleryRequest(val *CreateGroupGalleryRequest) *NullableCreateGroupGalleryRequest {
	return &NullableCreateGroupGalleryRequest{value: val, isSet: true}
}

func (v NullableCreateGroupGalleryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupGalleryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


