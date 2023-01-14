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

// GroupMyMember struct for GroupMyMember
type GroupMyMember struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId *string `json:"userId,omitempty"`
	RoleIds []string `json:"roleIds,omitempty"`
	ManagerNotes *string `json:"managerNotes,omitempty"`
	MembershipStatus *string `json:"membershipStatus,omitempty"`
	IsSubscribedToAnnouncements *bool `json:"isSubscribedToAnnouncements,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	IsRepresenting *bool `json:"isRepresenting,omitempty"`
	JoinedAt *time.Time `json:"joinedAt,omitempty"`
	BannedAt NullableString `json:"bannedAt,omitempty"`
	Has2FA *bool `json:"has2FA,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// NewGroupMyMember instantiates a new GroupMyMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMyMember() *GroupMyMember {
	this := GroupMyMember{}
	var isSubscribedToAnnouncements bool = true
	this.IsSubscribedToAnnouncements = &isSubscribedToAnnouncements
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	var has2FA bool = false
	this.Has2FA = &has2FA
	return &this
}

// NewGroupMyMemberWithDefaults instantiates a new GroupMyMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMyMemberWithDefaults() *GroupMyMember {
	this := GroupMyMember{}
	var isSubscribedToAnnouncements bool = true
	this.IsSubscribedToAnnouncements = &isSubscribedToAnnouncements
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	var has2FA bool = false
	this.Has2FA = &has2FA
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupMyMember) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupMyMember) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupMyMember) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupMyMember) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupMyMember) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupMyMember) SetGroupId(v string) {
	o.GroupId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GroupMyMember) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GroupMyMember) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GroupMyMember) SetUserId(v string) {
	o.UserId = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *GroupMyMember) GetRoleIds() []string {
	if o == nil || isNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetRoleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIds) {
    return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *GroupMyMember) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *GroupMyMember) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetManagerNotes returns the ManagerNotes field value if set, zero value otherwise.
func (o *GroupMyMember) GetManagerNotes() string {
	if o == nil || isNil(o.ManagerNotes) {
		var ret string
		return ret
	}
	return *o.ManagerNotes
}

// GetManagerNotesOk returns a tuple with the ManagerNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetManagerNotesOk() (*string, bool) {
	if o == nil || isNil(o.ManagerNotes) {
    return nil, false
	}
	return o.ManagerNotes, true
}

// HasManagerNotes returns a boolean if a field has been set.
func (o *GroupMyMember) HasManagerNotes() bool {
	if o != nil && !isNil(o.ManagerNotes) {
		return true
	}

	return false
}

// SetManagerNotes gets a reference to the given string and assigns it to the ManagerNotes field.
func (o *GroupMyMember) SetManagerNotes(v string) {
	o.ManagerNotes = &v
}

// GetMembershipStatus returns the MembershipStatus field value if set, zero value otherwise.
func (o *GroupMyMember) GetMembershipStatus() string {
	if o == nil || isNil(o.MembershipStatus) {
		var ret string
		return ret
	}
	return *o.MembershipStatus
}

// GetMembershipStatusOk returns a tuple with the MembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetMembershipStatusOk() (*string, bool) {
	if o == nil || isNil(o.MembershipStatus) {
    return nil, false
	}
	return o.MembershipStatus, true
}

// HasMembershipStatus returns a boolean if a field has been set.
func (o *GroupMyMember) HasMembershipStatus() bool {
	if o != nil && !isNil(o.MembershipStatus) {
		return true
	}

	return false
}

// SetMembershipStatus gets a reference to the given string and assigns it to the MembershipStatus field.
func (o *GroupMyMember) SetMembershipStatus(v string) {
	o.MembershipStatus = &v
}

// GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field value if set, zero value otherwise.
func (o *GroupMyMember) GetIsSubscribedToAnnouncements() bool {
	if o == nil || isNil(o.IsSubscribedToAnnouncements) {
		var ret bool
		return ret
	}
	return *o.IsSubscribedToAnnouncements
}

// GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetIsSubscribedToAnnouncementsOk() (*bool, bool) {
	if o == nil || isNil(o.IsSubscribedToAnnouncements) {
    return nil, false
	}
	return o.IsSubscribedToAnnouncements, true
}

// HasIsSubscribedToAnnouncements returns a boolean if a field has been set.
func (o *GroupMyMember) HasIsSubscribedToAnnouncements() bool {
	if o != nil && !isNil(o.IsSubscribedToAnnouncements) {
		return true
	}

	return false
}

// SetIsSubscribedToAnnouncements gets a reference to the given bool and assigns it to the IsSubscribedToAnnouncements field.
func (o *GroupMyMember) SetIsSubscribedToAnnouncements(v bool) {
	o.IsSubscribedToAnnouncements = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GroupMyMember) GetVisibility() string {
	if o == nil || isNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetVisibilityOk() (*string, bool) {
	if o == nil || isNil(o.Visibility) {
    return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *GroupMyMember) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GroupMyMember) SetVisibility(v string) {
	o.Visibility = &v
}

// GetIsRepresenting returns the IsRepresenting field value if set, zero value otherwise.
func (o *GroupMyMember) GetIsRepresenting() bool {
	if o == nil || isNil(o.IsRepresenting) {
		var ret bool
		return ret
	}
	return *o.IsRepresenting
}

// GetIsRepresentingOk returns a tuple with the IsRepresenting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetIsRepresentingOk() (*bool, bool) {
	if o == nil || isNil(o.IsRepresenting) {
    return nil, false
	}
	return o.IsRepresenting, true
}

// HasIsRepresenting returns a boolean if a field has been set.
func (o *GroupMyMember) HasIsRepresenting() bool {
	if o != nil && !isNil(o.IsRepresenting) {
		return true
	}

	return false
}

// SetIsRepresenting gets a reference to the given bool and assigns it to the IsRepresenting field.
func (o *GroupMyMember) SetIsRepresenting(v bool) {
	o.IsRepresenting = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *GroupMyMember) GetJoinedAt() time.Time {
	if o == nil || isNil(o.JoinedAt) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.JoinedAt) {
    return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *GroupMyMember) HasJoinedAt() bool {
	if o != nil && !isNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given time.Time and assigns it to the JoinedAt field.
func (o *GroupMyMember) SetJoinedAt(v time.Time) {
	o.JoinedAt = &v
}

// GetBannedAt returns the BannedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupMyMember) GetBannedAt() string {
	if o == nil || isNil(o.BannedAt.Get()) {
		var ret string
		return ret
	}
	return *o.BannedAt.Get()
}

// GetBannedAtOk returns a tuple with the BannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMyMember) GetBannedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BannedAt.Get(), o.BannedAt.IsSet()
}

// HasBannedAt returns a boolean if a field has been set.
func (o *GroupMyMember) HasBannedAt() bool {
	if o != nil && o.BannedAt.IsSet() {
		return true
	}

	return false
}

// SetBannedAt gets a reference to the given NullableString and assigns it to the BannedAt field.
func (o *GroupMyMember) SetBannedAt(v string) {
	o.BannedAt.Set(&v)
}
// SetBannedAtNil sets the value for BannedAt to be an explicit nil
func (o *GroupMyMember) SetBannedAtNil() {
	o.BannedAt.Set(nil)
}

// UnsetBannedAt ensures that no value is present for BannedAt, not even an explicit nil
func (o *GroupMyMember) UnsetBannedAt() {
	o.BannedAt.Unset()
}

// GetHas2FA returns the Has2FA field value if set, zero value otherwise.
func (o *GroupMyMember) GetHas2FA() bool {
	if o == nil || isNil(o.Has2FA) {
		var ret bool
		return ret
	}
	return *o.Has2FA
}

// GetHas2FAOk returns a tuple with the Has2FA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetHas2FAOk() (*bool, bool) {
	if o == nil || isNil(o.Has2FA) {
    return nil, false
	}
	return o.Has2FA, true
}

// HasHas2FA returns a boolean if a field has been set.
func (o *GroupMyMember) HasHas2FA() bool {
	if o != nil && !isNil(o.Has2FA) {
		return true
	}

	return false
}

// SetHas2FA gets a reference to the given bool and assigns it to the Has2FA field.
func (o *GroupMyMember) SetHas2FA(v bool) {
	o.Has2FA = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupMyMember) GetPermissions() []string {
	if o == nil || isNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMyMember) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupMyMember) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GroupMyMember) SetPermissions(v []string) {
	o.Permissions = v
}

func (o GroupMyMember) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !isNil(o.ManagerNotes) {
		toSerialize["managerNotes"] = o.ManagerNotes
	}
	if !isNil(o.MembershipStatus) {
		toSerialize["membershipStatus"] = o.MembershipStatus
	}
	if !isNil(o.IsSubscribedToAnnouncements) {
		toSerialize["isSubscribedToAnnouncements"] = o.IsSubscribedToAnnouncements
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !isNil(o.IsRepresenting) {
		toSerialize["isRepresenting"] = o.IsRepresenting
	}
	if !isNil(o.JoinedAt) {
		toSerialize["joinedAt"] = o.JoinedAt
	}
	if o.BannedAt.IsSet() {
		toSerialize["bannedAt"] = o.BannedAt.Get()
	}
	if !isNil(o.Has2FA) {
		toSerialize["has2FA"] = o.Has2FA
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableGroupMyMember struct {
	value *GroupMyMember
	isSet bool
}

func (v NullableGroupMyMember) Get() *GroupMyMember {
	return v.value
}

func (v *NullableGroupMyMember) Set(val *GroupMyMember) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMyMember) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMyMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMyMember(val *GroupMyMember) *NullableGroupMyMember {
	return &NullableGroupMyMember{value: val, isSet: true}
}

func (v NullableGroupMyMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMyMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

