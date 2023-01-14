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

// FileData 
type FileData struct {
	Category string `json:"category"`
	FileName string `json:"fileName"`
	Md5 string `json:"md5"`
	SizeInBytes int32 `json:"sizeInBytes"`
	Status FileStatus `json:"status"`
	UploadId string `json:"uploadId"`
	Url string `json:"url"`
}

// NewFileData instantiates a new FileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileData(category string, fileName string, md5 string, sizeInBytes int32, status FileStatus, uploadId string, url string) *FileData {
	this := FileData{}
	this.Category = category
	this.FileName = fileName
	this.Md5 = md5
	this.SizeInBytes = sizeInBytes
	this.Status = status
	this.UploadId = uploadId
	this.Url = url
	return &this
}

// NewFileDataWithDefaults instantiates a new FileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDataWithDefaults() *FileData {
	this := FileData{}
	var category string = "queued"
	this.Category = category
	var status FileStatus = FILESTATUS_WAITING
	this.Status = status
	var uploadId string = ""
	this.UploadId = uploadId
	return &this
}

// GetCategory returns the Category field value
func (o *FileData) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *FileData) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *FileData) SetCategory(v string) {
	o.Category = v
}

// GetFileName returns the FileName field value
func (o *FileData) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *FileData) GetFileNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *FileData) SetFileName(v string) {
	o.FileName = v
}

// GetMd5 returns the Md5 field value
func (o *FileData) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *FileData) GetMd5Ok() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value
func (o *FileData) SetMd5(v string) {
	o.Md5 = v
}

// GetSizeInBytes returns the SizeInBytes field value
func (o *FileData) GetSizeInBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value
// and a boolean to check if the value has been set.
func (o *FileData) GetSizeInBytesOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SizeInBytes, true
}

// SetSizeInBytes sets field value
func (o *FileData) SetSizeInBytes(v int32) {
	o.SizeInBytes = v
}

// GetStatus returns the Status field value
func (o *FileData) GetStatus() FileStatus {
	if o == nil {
		var ret FileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FileData) GetStatusOk() (*FileStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FileData) SetStatus(v FileStatus) {
	o.Status = v
}

// GetUploadId returns the UploadId field value
func (o *FileData) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *FileData) GetUploadIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *FileData) SetUploadId(v string) {
	o.UploadId = v
}

// GetUrl returns the Url field value
func (o *FileData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FileData) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FileData) SetUrl(v string) {
	o.Url = v
}

func (o FileData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["fileName"] = o.FileName
	}
	if true {
		toSerialize["md5"] = o.Md5
	}
	if true {
		toSerialize["sizeInBytes"] = o.SizeInBytes
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["uploadId"] = o.UploadId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableFileData struct {
	value *FileData
	isSet bool
}

func (v NullableFileData) Get() *FileData {
	return v.value
}

func (v *NullableFileData) Set(val *FileData) {
	v.value = val
	v.isSet = true
}

func (v NullableFileData) IsSet() bool {
	return v.isSet
}

func (v *NullableFileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileData(val *FileData) *NullableFileData {
	return &NullableFileData{value: val, isSet: true}
}

func (v NullableFileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


