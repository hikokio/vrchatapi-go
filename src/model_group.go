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

// Group struct for Group
type Group struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortCode *string `json:"shortCode,omitempty"`
	Discriminator *string `json:"discriminator,omitempty"`
	Description *string `json:"description,omitempty"`
	IconUrl NullableString `json:"iconUrl,omitempty"`
	BannerUrl NullableString `json:"bannerUrl,omitempty"`
	Privacy *GroupPrivacy `json:"privacy,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId *string `json:"ownerId,omitempty"`
	Rules NullableString `json:"rules,omitempty"`
	Links []string `json:"links,omitempty"`
	Languages []string `json:"languages,omitempty"`
	IconId NullableString `json:"iconId,omitempty"`
	BannerId NullableString `json:"bannerId,omitempty"`
	MemberCount *int32 `json:"memberCount,omitempty"`
	MemberCountSyncedAt *time.Time `json:"memberCountSyncedAt,omitempty"`
	IsVerified *bool `json:"isVerified,omitempty"`
	JoinState *GroupJoinState `json:"joinState,omitempty"`
	//  
	Tags []string `json:"tags,omitempty"`
	//  
	Galleries []GroupGallery `json:"galleries,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	OnlineMemberCount *int32 `json:"onlineMemberCount,omitempty"`
	MembershipStatus *GroupMemberStatus `json:"membershipStatus,omitempty"`
	MyMember *GroupMyMember `json:"myMember,omitempty"`
	// Only returned if ?includeRoles=true is specified.
	Roles []GroupRole `json:"roles,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	var privacy GroupPrivacy = GROUPPRIVACY_DEFAULT
	this.Privacy = &privacy
	var isVerified bool = false
	this.IsVerified = &isVerified
	var joinState GroupJoinState = GROUPJOINSTATE_OPEN
	this.JoinState = &joinState
	var membershipStatus GroupMemberStatus = GROUPMEMBERSTATUS_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	var privacy GroupPrivacy = GROUPPRIVACY_DEFAULT
	this.Privacy = &privacy
	var isVerified bool = false
	this.IsVerified = &isVerified
	var joinState GroupJoinState = GROUPJOINSTATE_OPEN
	this.JoinState = &joinState
	var membershipStatus GroupMemberStatus = GROUPMEMBERSTATUS_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Group) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name = &v
}

// GetShortCode returns the ShortCode field value if set, zero value otherwise.
func (o *Group) GetShortCode() string {
	if o == nil || isNil(o.ShortCode) {
		var ret string
		return ret
	}
	return *o.ShortCode
}

// GetShortCodeOk returns a tuple with the ShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetShortCodeOk() (*string, bool) {
	if o == nil || isNil(o.ShortCode) {
    return nil, false
	}
	return o.ShortCode, true
}

// HasShortCode returns a boolean if a field has been set.
func (o *Group) HasShortCode() bool {
	if o != nil && !isNil(o.ShortCode) {
		return true
	}

	return false
}

// SetShortCode gets a reference to the given string and assigns it to the ShortCode field.
func (o *Group) SetShortCode(v string) {
	o.ShortCode = &v
}

// GetDiscriminator returns the Discriminator field value if set, zero value otherwise.
func (o *Group) GetDiscriminator() string {
	if o == nil || isNil(o.Discriminator) {
		var ret string
		return ret
	}
	return *o.Discriminator
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDiscriminatorOk() (*string, bool) {
	if o == nil || isNil(o.Discriminator) {
    return nil, false
	}
	return o.Discriminator, true
}

// HasDiscriminator returns a boolean if a field has been set.
func (o *Group) HasDiscriminator() bool {
	if o != nil && !isNil(o.Discriminator) {
		return true
	}

	return false
}

// SetDiscriminator gets a reference to the given string and assigns it to the Discriminator field.
func (o *Group) SetDiscriminator(v string) {
	o.Discriminator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Group) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Group) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Group) SetDescription(v string) {
	o.Description = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIconUrl() string {
	if o == nil || isNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIconUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *Group) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *Group) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *Group) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *Group) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetBannerUrl() string {
	if o == nil || isNil(o.BannerUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BannerUrl.Get()
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetBannerUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BannerUrl.Get(), o.BannerUrl.IsSet()
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *Group) HasBannerUrl() bool {
	if o != nil && o.BannerUrl.IsSet() {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given NullableString and assigns it to the BannerUrl field.
func (o *Group) SetBannerUrl(v string) {
	o.BannerUrl.Set(&v)
}
// SetBannerUrlNil sets the value for BannerUrl to be an explicit nil
func (o *Group) SetBannerUrlNil() {
	o.BannerUrl.Set(nil)
}

// UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
func (o *Group) UnsetBannerUrl() {
	o.BannerUrl.Unset()
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *Group) GetPrivacy() GroupPrivacy {
	if o == nil || isNil(o.Privacy) {
		var ret GroupPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetPrivacyOk() (*GroupPrivacy, bool) {
	if o == nil || isNil(o.Privacy) {
    return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *Group) HasPrivacy() bool {
	if o != nil && !isNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given GroupPrivacy and assigns it to the Privacy field.
func (o *Group) SetPrivacy(v GroupPrivacy) {
	o.Privacy = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Group) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
    return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Group) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Group) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetRules() string {
	if o == nil || isNil(o.Rules.Get()) {
		var ret string
		return ret
	}
	return *o.Rules.Get()
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetRulesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules.Get(), o.Rules.IsSet()
}

// HasRules returns a boolean if a field has been set.
func (o *Group) HasRules() bool {
	if o != nil && o.Rules.IsSet() {
		return true
	}

	return false
}

// SetRules gets a reference to the given NullableString and assigns it to the Rules field.
func (o *Group) SetRules(v string) {
	o.Rules.Set(&v)
}
// SetRulesNil sets the value for Rules to be an explicit nil
func (o *Group) SetRulesNil() {
	o.Rules.Set(nil)
}

// UnsetRules ensures that no value is present for Rules, not even an explicit nil
func (o *Group) UnsetRules() {
	o.Rules.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() []string {
	if o == nil || isNil(o.Links) {
		var ret []string
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() ([]string, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []string and assigns it to the Links field.
func (o *Group) SetLinks(v []string) {
	o.Links = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *Group) GetLanguages() []string {
	if o == nil || isNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLanguagesOk() ([]string, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *Group) HasLanguages() bool {
	if o != nil && !isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *Group) SetLanguages(v []string) {
	o.Languages = v
}

// GetIconId returns the IconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIconId() string {
	if o == nil || isNil(o.IconId.Get()) {
		var ret string
		return ret
	}
	return *o.IconId.Get()
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIconIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.IconId.Get(), o.IconId.IsSet()
}

// HasIconId returns a boolean if a field has been set.
func (o *Group) HasIconId() bool {
	if o != nil && o.IconId.IsSet() {
		return true
	}

	return false
}

// SetIconId gets a reference to the given NullableString and assigns it to the IconId field.
func (o *Group) SetIconId(v string) {
	o.IconId.Set(&v)
}
// SetIconIdNil sets the value for IconId to be an explicit nil
func (o *Group) SetIconIdNil() {
	o.IconId.Set(nil)
}

// UnsetIconId ensures that no value is present for IconId, not even an explicit nil
func (o *Group) UnsetIconId() {
	o.IconId.Unset()
}

// GetBannerId returns the BannerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetBannerId() string {
	if o == nil || isNil(o.BannerId.Get()) {
		var ret string
		return ret
	}
	return *o.BannerId.Get()
}

// GetBannerIdOk returns a tuple with the BannerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetBannerIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BannerId.Get(), o.BannerId.IsSet()
}

// HasBannerId returns a boolean if a field has been set.
func (o *Group) HasBannerId() bool {
	if o != nil && o.BannerId.IsSet() {
		return true
	}

	return false
}

// SetBannerId gets a reference to the given NullableString and assigns it to the BannerId field.
func (o *Group) SetBannerId(v string) {
	o.BannerId.Set(&v)
}
// SetBannerIdNil sets the value for BannerId to be an explicit nil
func (o *Group) SetBannerIdNil() {
	o.BannerId.Set(nil)
}

// UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
func (o *Group) UnsetBannerId() {
	o.BannerId.Unset()
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *Group) GetMemberCount() int32 {
	if o == nil || isNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMemberCountOk() (*int32, bool) {
	if o == nil || isNil(o.MemberCount) {
    return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *Group) HasMemberCount() bool {
	if o != nil && !isNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *Group) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetMemberCountSyncedAt returns the MemberCountSyncedAt field value if set, zero value otherwise.
func (o *Group) GetMemberCountSyncedAt() time.Time {
	if o == nil || isNil(o.MemberCountSyncedAt) {
		var ret time.Time
		return ret
	}
	return *o.MemberCountSyncedAt
}

// GetMemberCountSyncedAtOk returns a tuple with the MemberCountSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMemberCountSyncedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.MemberCountSyncedAt) {
    return nil, false
	}
	return o.MemberCountSyncedAt, true
}

// HasMemberCountSyncedAt returns a boolean if a field has been set.
func (o *Group) HasMemberCountSyncedAt() bool {
	if o != nil && !isNil(o.MemberCountSyncedAt) {
		return true
	}

	return false
}

// SetMemberCountSyncedAt gets a reference to the given time.Time and assigns it to the MemberCountSyncedAt field.
func (o *Group) SetMemberCountSyncedAt(v time.Time) {
	o.MemberCountSyncedAt = &v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *Group) GetIsVerified() bool {
	if o == nil || isNil(o.IsVerified) {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || isNil(o.IsVerified) {
    return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *Group) HasIsVerified() bool {
	if o != nil && !isNil(o.IsVerified) {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *Group) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetJoinState returns the JoinState field value if set, zero value otherwise.
func (o *Group) GetJoinState() GroupJoinState {
	if o == nil || isNil(o.JoinState) {
		var ret GroupJoinState
		return ret
	}
	return *o.JoinState
}

// GetJoinStateOk returns a tuple with the JoinState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetJoinStateOk() (*GroupJoinState, bool) {
	if o == nil || isNil(o.JoinState) {
    return nil, false
	}
	return o.JoinState, true
}

// HasJoinState returns a boolean if a field has been set.
func (o *Group) HasJoinState() bool {
	if o != nil && !isNil(o.JoinState) {
		return true
	}

	return false
}

// SetJoinState gets a reference to the given GroupJoinState and assigns it to the JoinState field.
func (o *Group) SetJoinState(v GroupJoinState) {
	o.JoinState = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Group) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Group) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Group) SetTags(v []string) {
	o.Tags = v
}

// GetGalleries returns the Galleries field value if set, zero value otherwise.
func (o *Group) GetGalleries() []GroupGallery {
	if o == nil || isNil(o.Galleries) {
		var ret []GroupGallery
		return ret
	}
	return o.Galleries
}

// GetGalleriesOk returns a tuple with the Galleries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetGalleriesOk() ([]GroupGallery, bool) {
	if o == nil || isNil(o.Galleries) {
    return nil, false
	}
	return o.Galleries, true
}

// HasGalleries returns a boolean if a field has been set.
func (o *Group) HasGalleries() bool {
	if o != nil && !isNil(o.Galleries) {
		return true
	}

	return false
}

// SetGalleries gets a reference to the given []GroupGallery and assigns it to the Galleries field.
func (o *Group) SetGalleries(v []GroupGallery) {
	o.Galleries = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Group) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Group) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Group) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetOnlineMemberCount returns the OnlineMemberCount field value if set, zero value otherwise.
func (o *Group) GetOnlineMemberCount() int32 {
	if o == nil || isNil(o.OnlineMemberCount) {
		var ret int32
		return ret
	}
	return *o.OnlineMemberCount
}

// GetOnlineMemberCountOk returns a tuple with the OnlineMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOnlineMemberCountOk() (*int32, bool) {
	if o == nil || isNil(o.OnlineMemberCount) {
    return nil, false
	}
	return o.OnlineMemberCount, true
}

// HasOnlineMemberCount returns a boolean if a field has been set.
func (o *Group) HasOnlineMemberCount() bool {
	if o != nil && !isNil(o.OnlineMemberCount) {
		return true
	}

	return false
}

// SetOnlineMemberCount gets a reference to the given int32 and assigns it to the OnlineMemberCount field.
func (o *Group) SetOnlineMemberCount(v int32) {
	o.OnlineMemberCount = &v
}

// GetMembershipStatus returns the MembershipStatus field value if set, zero value otherwise.
func (o *Group) GetMembershipStatus() GroupMemberStatus {
	if o == nil || isNil(o.MembershipStatus) {
		var ret GroupMemberStatus
		return ret
	}
	return *o.MembershipStatus
}

// GetMembershipStatusOk returns a tuple with the MembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMembershipStatusOk() (*GroupMemberStatus, bool) {
	if o == nil || isNil(o.MembershipStatus) {
    return nil, false
	}
	return o.MembershipStatus, true
}

// HasMembershipStatus returns a boolean if a field has been set.
func (o *Group) HasMembershipStatus() bool {
	if o != nil && !isNil(o.MembershipStatus) {
		return true
	}

	return false
}

// SetMembershipStatus gets a reference to the given GroupMemberStatus and assigns it to the MembershipStatus field.
func (o *Group) SetMembershipStatus(v GroupMemberStatus) {
	o.MembershipStatus = &v
}

// GetMyMember returns the MyMember field value if set, zero value otherwise.
func (o *Group) GetMyMember() GroupMyMember {
	if o == nil || isNil(o.MyMember) {
		var ret GroupMyMember
		return ret
	}
	return *o.MyMember
}

// GetMyMemberOk returns a tuple with the MyMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMyMemberOk() (*GroupMyMember, bool) {
	if o == nil || isNil(o.MyMember) {
    return nil, false
	}
	return o.MyMember, true
}

// HasMyMember returns a boolean if a field has been set.
func (o *Group) HasMyMember() bool {
	if o != nil && !isNil(o.MyMember) {
		return true
	}

	return false
}

// SetMyMember gets a reference to the given GroupMyMember and assigns it to the MyMember field.
func (o *Group) SetMyMember(v GroupMyMember) {
	o.MyMember = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetRoles() []GroupRole {
	if o == nil {
		var ret []GroupRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetRolesOk() ([]GroupRole, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Group) HasRoles() bool {
	if o != nil && isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []GroupRole and assigns it to the Roles field.
func (o *Group) SetRoles(v []GroupRole) {
	o.Roles = v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ShortCode) {
		toSerialize["shortCode"] = o.ShortCode
	}
	if !isNil(o.Discriminator) {
		toSerialize["discriminator"] = o.Discriminator
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.IconUrl.IsSet() {
		toSerialize["iconUrl"] = o.IconUrl.Get()
	}
	if o.BannerUrl.IsSet() {
		toSerialize["bannerUrl"] = o.BannerUrl.Get()
	}
	if !isNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.Rules.IsSet() {
		toSerialize["rules"] = o.Rules.Get()
	}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if o.IconId.IsSet() {
		toSerialize["iconId"] = o.IconId.Get()
	}
	if o.BannerId.IsSet() {
		toSerialize["bannerId"] = o.BannerId.Get()
	}
	if !isNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !isNil(o.MemberCountSyncedAt) {
		toSerialize["memberCountSyncedAt"] = o.MemberCountSyncedAt
	}
	if !isNil(o.IsVerified) {
		toSerialize["isVerified"] = o.IsVerified
	}
	if !isNil(o.JoinState) {
		toSerialize["joinState"] = o.JoinState
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Galleries) {
		toSerialize["galleries"] = o.Galleries
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.OnlineMemberCount) {
		toSerialize["onlineMemberCount"] = o.OnlineMemberCount
	}
	if !isNil(o.MembershipStatus) {
		toSerialize["membershipStatus"] = o.MembershipStatus
	}
	if !isNil(o.MyMember) {
		toSerialize["myMember"] = o.MyMember
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

