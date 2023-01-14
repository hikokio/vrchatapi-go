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

// Avatar 
type Avatar struct {
	// Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`.
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`. **Deprecation:** `Object` has unknown usage/fields, and is always empty. Use normal `Url` field instead.
	AssetUrlObject map[string]interface{} `json:"assetUrlObject,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId string `json:"authorId"`
	AuthorName string `json:"authorName"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name string `json:"name"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus"`
	//  
	Tags []string `json:"tags"`
	ThumbnailImageUrl string `json:"thumbnailImageUrl"`
	UnityPackageUrl string `json:"unityPackageUrl"`
	// Deprecated
	UnityPackageUrlObject AvatarUnityPackageUrlObject `json:"unityPackageUrlObject"`
	UnityPackages []UnityPackage `json:"unityPackages"`
	UpdatedAt time.Time `json:"updated_at"`
	Version int32 `json:"version"`
}

// NewAvatar instantiates a new Avatar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvatar(authorId string, authorName string, createdAt time.Time, description string, featured bool, id string, imageUrl string, name string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackageUrl string, unityPackageUrlObject AvatarUnityPackageUrlObject, unityPackages []UnityPackage, updatedAt time.Time, version int32) *Avatar {
	this := Avatar{}
	this.AuthorId = authorId
	this.AuthorName = authorName
	this.CreatedAt = createdAt
	this.Description = description
	this.Featured = featured
	this.Id = id
	this.ImageUrl = imageUrl
	this.Name = name
	this.ReleaseStatus = releaseStatus
	this.Tags = tags
	this.ThumbnailImageUrl = thumbnailImageUrl
	this.UnityPackageUrl = unityPackageUrl
	this.UnityPackageUrlObject = unityPackageUrlObject
	this.UnityPackages = unityPackages
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewAvatarWithDefaults instantiates a new Avatar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvatarWithDefaults() *Avatar {
	this := Avatar{}
	var featured bool = false
	this.Featured = featured
	var releaseStatus ReleaseStatus = RELEASESTATUS_PUBLIC
	this.ReleaseStatus = releaseStatus
	var version int32 = 0
	this.Version = version
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *Avatar) GetAssetUrl() string {
	if o == nil || isNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avatar) GetAssetUrlOk() (*string, bool) {
	if o == nil || isNil(o.AssetUrl) {
    return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *Avatar) HasAssetUrl() bool {
	if o != nil && !isNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *Avatar) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetAssetUrlObject returns the AssetUrlObject field value if set, zero value otherwise.
func (o *Avatar) GetAssetUrlObject() map[string]interface{} {
	if o == nil || isNil(o.AssetUrlObject) {
		var ret map[string]interface{}
		return ret
	}
	return o.AssetUrlObject
}

// GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avatar) GetAssetUrlObjectOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AssetUrlObject) {
    return map[string]interface{}{}, false
	}
	return o.AssetUrlObject, true
}

// HasAssetUrlObject returns a boolean if a field has been set.
func (o *Avatar) HasAssetUrlObject() bool {
	if o != nil && !isNil(o.AssetUrlObject) {
		return true
	}

	return false
}

// SetAssetUrlObject gets a reference to the given map[string]interface{} and assigns it to the AssetUrlObject field.
func (o *Avatar) SetAssetUrlObject(v map[string]interface{}) {
	o.AssetUrlObject = v
}

// GetAuthorId returns the AuthorId field value
func (o *Avatar) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetAuthorIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *Avatar) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetAuthorName returns the AuthorName field value
func (o *Avatar) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetAuthorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *Avatar) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Avatar) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Avatar) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *Avatar) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Avatar) SetDescription(v string) {
	o.Description = v
}

// GetFeatured returns the Featured field value
func (o *Avatar) GetFeatured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetFeaturedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *Avatar) SetFeatured(v bool) {
	o.Featured = v
}

// GetId returns the Id field value
func (o *Avatar) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Avatar) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *Avatar) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetImageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *Avatar) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *Avatar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Avatar) SetName(v string) {
	o.Name = v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *Avatar) GetReleaseStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *Avatar) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = v
}

// GetTags returns the Tags field value
func (o *Avatar) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Avatar) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value
func (o *Avatar) GetThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ThumbnailImageUrl, true
}

// SetThumbnailImageUrl sets field value
func (o *Avatar) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value
func (o *Avatar) GetUnityPackageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnityPackageUrl
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnityPackageUrl, true
}

// SetUnityPackageUrl sets field value
func (o *Avatar) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl = v
}

// GetUnityPackageUrlObject returns the UnityPackageUrlObject field value
// Deprecated
func (o *Avatar) GetUnityPackageUrlObject() AvatarUnityPackageUrlObject {
	if o == nil {
		var ret AvatarUnityPackageUrlObject
		return ret
	}

	return o.UnityPackageUrlObject
}

// GetUnityPackageUrlObjectOk returns a tuple with the UnityPackageUrlObject field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Avatar) GetUnityPackageUrlObjectOk() (*AvatarUnityPackageUrlObject, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnityPackageUrlObject, true
}

// SetUnityPackageUrlObject sets field value
// Deprecated
func (o *Avatar) SetUnityPackageUrlObject(v AvatarUnityPackageUrlObject) {
	o.UnityPackageUrlObject = v
}

// GetUnityPackages returns the UnityPackages field value
func (o *Avatar) GetUnityPackages() []UnityPackage {
	if o == nil {
		var ret []UnityPackage
		return ret
	}

	return o.UnityPackages
}

// GetUnityPackagesOk returns a tuple with the UnityPackages field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUnityPackagesOk() ([]UnityPackage, bool) {
	if o == nil {
    return nil, false
	}
	return o.UnityPackages, true
}

// SetUnityPackages sets field value
func (o *Avatar) SetUnityPackages(v []UnityPackage) {
	o.UnityPackages = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Avatar) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Avatar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *Avatar) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetVersionOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Avatar) SetVersion(v int32) {
	o.Version = v
}

func (o Avatar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !isNil(o.AssetUrlObject) {
		toSerialize["assetUrlObject"] = o.AssetUrlObject
	}
	if true {
		toSerialize["authorId"] = o.AuthorId
	}
	if true {
		toSerialize["authorName"] = o.AuthorName
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["featured"] = o.Featured
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["releaseStatus"] = o.ReleaseStatus
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	}
	if true {
		toSerialize["unityPackageUrl"] = o.UnityPackageUrl
	}
	if true {
		toSerialize["unityPackageUrlObject"] = o.UnityPackageUrlObject
	}
	if true {
		toSerialize["unityPackages"] = o.UnityPackages
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableAvatar struct {
	value *Avatar
	isSet bool
}

func (v NullableAvatar) Get() *Avatar {
	return v.value
}

func (v *NullableAvatar) Set(val *Avatar) {
	v.value = val
	v.isSet = true
}

func (v NullableAvatar) IsSet() bool {
	return v.isSet
}

func (v *NullableAvatar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvatar(val *Avatar) *NullableAvatar {
	return &NullableAvatar{value: val, isSet: true}
}

func (v NullableAvatar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvatar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

