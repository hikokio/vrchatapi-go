/*
VRChat API Documentation


API version: 1.10.1
Contact: me@ariesclark.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
	"time"
)

// GroupRole struct for GroupRole
type GroupRole struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsSelfAssignable *bool `json:"isSelfAssignable,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsManagementRole *bool `json:"isManagementRole,omitempty"`
	RequiresTwoFactor *bool `json:"requiresTwoFactor,omitempty"`
	RequiresPurchase *bool `json:"requiresPurchase,omitempty"`
	Order *int32 `json:"order,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGroupRole instantiates a new GroupRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRole() *GroupRole {
	this := GroupRole{}
	var isSelfAssignable bool = false
	this.IsSelfAssignable = &isSelfAssignable
	var isManagementRole bool = false
	this.IsManagementRole = &isManagementRole
	var requiresTwoFactor bool = false
	this.RequiresTwoFactor = &requiresTwoFactor
	var requiresPurchase bool = false
	this.RequiresPurchase = &requiresPurchase
	return &this
}

// NewGroupRoleWithDefaults instantiates a new GroupRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRoleWithDefaults() *GroupRole {
	this := GroupRole{}
	var isSelfAssignable bool = false
	this.IsSelfAssignable = &isSelfAssignable
	var isManagementRole bool = false
	this.IsManagementRole = &isManagementRole
	var requiresTwoFactor bool = false
	this.RequiresTwoFactor = &requiresTwoFactor
	var requiresPurchase bool = false
	this.RequiresPurchase = &requiresPurchase
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupRole) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupRole) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupRole) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupRole) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupRole) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupRole) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupRole) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupRole) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupRole) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupRole) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupRole) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupRole) SetDescription(v string) {
	o.Description = &v
}

// GetIsSelfAssignable returns the IsSelfAssignable field value if set, zero value otherwise.
func (o *GroupRole) GetIsSelfAssignable() bool {
	if o == nil || isNil(o.IsSelfAssignable) {
		var ret bool
		return ret
	}
	return *o.IsSelfAssignable
}

// GetIsSelfAssignableOk returns a tuple with the IsSelfAssignable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIsSelfAssignableOk() (*bool, bool) {
	if o == nil || isNil(o.IsSelfAssignable) {
    return nil, false
	}
	return o.IsSelfAssignable, true
}

// HasIsSelfAssignable returns a boolean if a field has been set.
func (o *GroupRole) HasIsSelfAssignable() bool {
	if o != nil && !isNil(o.IsSelfAssignable) {
		return true
	}

	return false
}

// SetIsSelfAssignable gets a reference to the given bool and assigns it to the IsSelfAssignable field.
func (o *GroupRole) SetIsSelfAssignable(v bool) {
	o.IsSelfAssignable = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupRole) GetPermissions() []string {
	if o == nil || isNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupRole) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GroupRole) SetPermissions(v []string) {
	o.Permissions = v
}

// GetIsManagementRole returns the IsManagementRole field value if set, zero value otherwise.
func (o *GroupRole) GetIsManagementRole() bool {
	if o == nil || isNil(o.IsManagementRole) {
		var ret bool
		return ret
	}
	return *o.IsManagementRole
}

// GetIsManagementRoleOk returns a tuple with the IsManagementRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIsManagementRoleOk() (*bool, bool) {
	if o == nil || isNil(o.IsManagementRole) {
    return nil, false
	}
	return o.IsManagementRole, true
}

// HasIsManagementRole returns a boolean if a field has been set.
func (o *GroupRole) HasIsManagementRole() bool {
	if o != nil && !isNil(o.IsManagementRole) {
		return true
	}

	return false
}

// SetIsManagementRole gets a reference to the given bool and assigns it to the IsManagementRole field.
func (o *GroupRole) SetIsManagementRole(v bool) {
	o.IsManagementRole = &v
}

// GetRequiresTwoFactor returns the RequiresTwoFactor field value if set, zero value otherwise.
func (o *GroupRole) GetRequiresTwoFactor() bool {
	if o == nil || isNil(o.RequiresTwoFactor) {
		var ret bool
		return ret
	}
	return *o.RequiresTwoFactor
}

// GetRequiresTwoFactorOk returns a tuple with the RequiresTwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetRequiresTwoFactorOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresTwoFactor) {
    return nil, false
	}
	return o.RequiresTwoFactor, true
}

// HasRequiresTwoFactor returns a boolean if a field has been set.
func (o *GroupRole) HasRequiresTwoFactor() bool {
	if o != nil && !isNil(o.RequiresTwoFactor) {
		return true
	}

	return false
}

// SetRequiresTwoFactor gets a reference to the given bool and assigns it to the RequiresTwoFactor field.
func (o *GroupRole) SetRequiresTwoFactor(v bool) {
	o.RequiresTwoFactor = &v
}

// GetRequiresPurchase returns the RequiresPurchase field value if set, zero value otherwise.
func (o *GroupRole) GetRequiresPurchase() bool {
	if o == nil || isNil(o.RequiresPurchase) {
		var ret bool
		return ret
	}
	return *o.RequiresPurchase
}

// GetRequiresPurchaseOk returns a tuple with the RequiresPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetRequiresPurchaseOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPurchase) {
    return nil, false
	}
	return o.RequiresPurchase, true
}

// HasRequiresPurchase returns a boolean if a field has been set.
func (o *GroupRole) HasRequiresPurchase() bool {
	if o != nil && !isNil(o.RequiresPurchase) {
		return true
	}

	return false
}

// SetRequiresPurchase gets a reference to the given bool and assigns it to the RequiresPurchase field.
func (o *GroupRole) SetRequiresPurchase(v bool) {
	o.RequiresPurchase = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GroupRole) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GroupRole) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *GroupRole) SetOrder(v int32) {
	o.Order = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupRole) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupRole) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupRole) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupRole) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupRole) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GroupRole) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GroupRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsSelfAssignable) {
		toSerialize["isSelfAssignable"] = o.IsSelfAssignable
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !isNil(o.IsManagementRole) {
		toSerialize["isManagementRole"] = o.IsManagementRole
	}
	if !isNil(o.RequiresTwoFactor) {
		toSerialize["requiresTwoFactor"] = o.RequiresTwoFactor
	}
	if !isNil(o.RequiresPurchase) {
		toSerialize["requiresPurchase"] = o.RequiresPurchase
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGroupRole struct {
	value *GroupRole
	isSet bool
}

func (v NullableGroupRole) Get() *GroupRole {
	return v.value
}

func (v *NullableGroupRole) Set(val *GroupRole) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRole) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRole(val *GroupRole) *NullableGroupRole {
	return &NullableGroupRole{value: val, isSet: true}
}

func (v NullableGroupRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


