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

// Response struct for Response
type Response struct {
	Message *string `json:"message,omitempty"`
	StatusCode int32 `json:"status_code"`
}

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse(statusCode int32) *Response {
	this := Response{}
	this.StatusCode = statusCode
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Response) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Response) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value
func (o *Response) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *Response) SetStatusCode(v int32) {
	o.StatusCode = v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


