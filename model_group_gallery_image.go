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

// GroupGalleryImage struct for GroupGalleryImage
type GroupGalleryImage struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	GalleryId *string `json:"galleryId,omitempty"`
	FileId *string `json:"fileId,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SubmittedByUserId *string `json:"submittedByUserId,omitempty"`
	Approved *bool `json:"approved,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ApprovedByUserId *string `json:"approvedByUserId,omitempty"`
	ApprovedAt *time.Time `json:"approvedAt,omitempty"`
}

// NewGroupGalleryImage instantiates a new GroupGalleryImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGalleryImage() *GroupGalleryImage {
	this := GroupGalleryImage{}
	var approved bool = false
	this.Approved = &approved
	return &this
}

// NewGroupGalleryImageWithDefaults instantiates a new GroupGalleryImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGalleryImageWithDefaults() *GroupGalleryImage {
	this := GroupGalleryImage{}
	var approved bool = false
	this.Approved = &approved
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupGalleryImage) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupGalleryImage) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGalleryId returns the GalleryId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetGalleryId() string {
	if o == nil || isNil(o.GalleryId) {
		var ret string
		return ret
	}
	return *o.GalleryId
}

// GetGalleryIdOk returns a tuple with the GalleryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetGalleryIdOk() (*string, bool) {
	if o == nil || isNil(o.GalleryId) {
    return nil, false
	}
	return o.GalleryId, true
}

// HasGalleryId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasGalleryId() bool {
	if o != nil && !isNil(o.GalleryId) {
		return true
	}

	return false
}

// SetGalleryId gets a reference to the given string and assigns it to the GalleryId field.
func (o *GroupGalleryImage) SetGalleryId(v string) {
	o.GalleryId = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetFileId() string {
	if o == nil || isNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetFileIdOk() (*string, bool) {
	if o == nil || isNil(o.FileId) {
    return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasFileId() bool {
	if o != nil && !isNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *GroupGalleryImage) SetFileId(v string) {
	o.FileId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
    return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GroupGalleryImage) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupGalleryImage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSubmittedByUserId returns the SubmittedByUserId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetSubmittedByUserId() string {
	if o == nil || isNil(o.SubmittedByUserId) {
		var ret string
		return ret
	}
	return *o.SubmittedByUserId
}

// GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetSubmittedByUserIdOk() (*string, bool) {
	if o == nil || isNil(o.SubmittedByUserId) {
    return nil, false
	}
	return o.SubmittedByUserId, true
}

// HasSubmittedByUserId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasSubmittedByUserId() bool {
	if o != nil && !isNil(o.SubmittedByUserId) {
		return true
	}

	return false
}

// SetSubmittedByUserId gets a reference to the given string and assigns it to the SubmittedByUserId field.
func (o *GroupGalleryImage) SetSubmittedByUserId(v string) {
	o.SubmittedByUserId = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApproved() bool {
	if o == nil || isNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedOk() (*bool, bool) {
	if o == nil || isNil(o.Approved) {
    return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApproved() bool {
	if o != nil && !isNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *GroupGalleryImage) SetApproved(v bool) {
	o.Approved = &v
}

// GetApprovedByUserId returns the ApprovedByUserId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApprovedByUserId() string {
	if o == nil || isNil(o.ApprovedByUserId) {
		var ret string
		return ret
	}
	return *o.ApprovedByUserId
}

// GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedByUserIdOk() (*string, bool) {
	if o == nil || isNil(o.ApprovedByUserId) {
    return nil, false
	}
	return o.ApprovedByUserId, true
}

// HasApprovedByUserId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApprovedByUserId() bool {
	if o != nil && !isNil(o.ApprovedByUserId) {
		return true
	}

	return false
}

// SetApprovedByUserId gets a reference to the given string and assigns it to the ApprovedByUserId field.
func (o *GroupGalleryImage) SetApprovedByUserId(v string) {
	o.ApprovedByUserId = &v
}

// GetApprovedAt returns the ApprovedAt field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApprovedAt() time.Time {
	if o == nil || isNil(o.ApprovedAt) {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAt
}

// GetApprovedAtOk returns a tuple with the ApprovedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ApprovedAt) {
    return nil, false
	}
	return o.ApprovedAt, true
}

// HasApprovedAt returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApprovedAt() bool {
	if o != nil && !isNil(o.ApprovedAt) {
		return true
	}

	return false
}

// SetApprovedAt gets a reference to the given time.Time and assigns it to the ApprovedAt field.
func (o *GroupGalleryImage) SetApprovedAt(v time.Time) {
	o.ApprovedAt = &v
}

func (o GroupGalleryImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.GalleryId) {
		toSerialize["galleryId"] = o.GalleryId
	}
	if !isNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !isNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.SubmittedByUserId) {
		toSerialize["submittedByUserId"] = o.SubmittedByUserId
	}
	if !isNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !isNil(o.ApprovedByUserId) {
		toSerialize["approvedByUserId"] = o.ApprovedByUserId
	}
	if !isNil(o.ApprovedAt) {
		toSerialize["approvedAt"] = o.ApprovedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGroupGalleryImage struct {
	value *GroupGalleryImage
	isSet bool
}

func (v NullableGroupGalleryImage) Get() *GroupGalleryImage {
	return v.value
}

func (v *NullableGroupGalleryImage) Set(val *GroupGalleryImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGalleryImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGalleryImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGalleryImage(val *GroupGalleryImage) *NullableGroupGalleryImage {
	return &NullableGroupGalleryImage{value: val, isSet: true}
}

func (v NullableGroupGalleryImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGalleryImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


