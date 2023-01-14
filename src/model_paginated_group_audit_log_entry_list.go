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

// PaginatedGroupAuditLogEntryList struct for PaginatedGroupAuditLogEntryList
type PaginatedGroupAuditLogEntryList struct {
	//  
	Results []GroupAuditLogEntry `json:"results,omitempty"`
	// The total number of results that the query would return if there were no pagination.
	TotalCount *int32 `json:"totalCount,omitempty"`
	// Whether there are more results after this page.
	HasNext *bool `json:"hasNext,omitempty"`
}

// NewPaginatedGroupAuditLogEntryList instantiates a new PaginatedGroupAuditLogEntryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupAuditLogEntryList() *PaginatedGroupAuditLogEntryList {
	this := PaginatedGroupAuditLogEntryList{}
	return &this
}

// NewPaginatedGroupAuditLogEntryListWithDefaults instantiates a new PaginatedGroupAuditLogEntryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupAuditLogEntryListWithDefaults() *PaginatedGroupAuditLogEntryList {
	this := PaginatedGroupAuditLogEntryList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedGroupAuditLogEntryList) GetResults() []GroupAuditLogEntry {
	if o == nil || isNil(o.Results) {
		var ret []GroupAuditLogEntry
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGroupAuditLogEntryList) GetResultsOk() ([]GroupAuditLogEntry, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedGroupAuditLogEntryList) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GroupAuditLogEntry and assigns it to the Results field.
func (o *PaginatedGroupAuditLogEntryList) SetResults(v []GroupAuditLogEntry) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedGroupAuditLogEntryList) GetTotalCount() int32 {
	if o == nil || isNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGroupAuditLogEntryList) GetTotalCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedGroupAuditLogEntryList) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedGroupAuditLogEntryList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *PaginatedGroupAuditLogEntryList) GetHasNext() bool {
	if o == nil || isNil(o.HasNext) {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGroupAuditLogEntryList) GetHasNextOk() (*bool, bool) {
	if o == nil || isNil(o.HasNext) {
    return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *PaginatedGroupAuditLogEntryList) HasHasNext() bool {
	if o != nil && !isNil(o.HasNext) {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *PaginatedGroupAuditLogEntryList) SetHasNext(v bool) {
	o.HasNext = &v
}

func (o PaginatedGroupAuditLogEntryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !isNil(o.HasNext) {
		toSerialize["hasNext"] = o.HasNext
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupAuditLogEntryList struct {
	value *PaginatedGroupAuditLogEntryList
	isSet bool
}

func (v NullablePaginatedGroupAuditLogEntryList) Get() *PaginatedGroupAuditLogEntryList {
	return v.value
}

func (v *NullablePaginatedGroupAuditLogEntryList) Set(val *PaginatedGroupAuditLogEntryList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupAuditLogEntryList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupAuditLogEntryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupAuditLogEntryList(val *PaginatedGroupAuditLogEntryList) *NullablePaginatedGroupAuditLogEntryList {
	return &NullablePaginatedGroupAuditLogEntryList{value: val, isSet: true}
}

func (v NullablePaginatedGroupAuditLogEntryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupAuditLogEntryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


