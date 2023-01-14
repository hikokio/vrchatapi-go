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

// APIConfigEvents struct for APIConfigEvents
type APIConfigEvents struct {
	// Unknown
	DistanceClose int32 `json:"distanceClose"`
	// Unknown
	DistanceFactor int32 `json:"distanceFactor"`
	// Unknown
	DistanceFar int32 `json:"distanceFar"`
	// Unknown
	GroupDistance int32 `json:"groupDistance"`
	// Unknown
	MaximumBunchSize int32 `json:"maximumBunchSize"`
	// Unknown
	NotVisibleFactor int32 `json:"notVisibleFactor"`
	// Unknown
	PlayerOrderBucketSize int32 `json:"playerOrderBucketSize"`
	// Unknown
	PlayerOrderFactor int32 `json:"playerOrderFactor"`
	// Unknown
	SlowUpdateFactorThreshold int32 `json:"slowUpdateFactorThreshold"`
	// Unknown
	ViewSegmentLength int32 `json:"viewSegmentLength"`
}

// NewAPIConfigEvents instantiates a new APIConfigEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfigEvents(distanceClose int32, distanceFactor int32, distanceFar int32, groupDistance int32, maximumBunchSize int32, notVisibleFactor int32, playerOrderBucketSize int32, playerOrderFactor int32, slowUpdateFactorThreshold int32, viewSegmentLength int32) *APIConfigEvents {
	this := APIConfigEvents{}
	this.DistanceClose = distanceClose
	this.DistanceFactor = distanceFactor
	this.DistanceFar = distanceFar
	this.GroupDistance = groupDistance
	this.MaximumBunchSize = maximumBunchSize
	this.NotVisibleFactor = notVisibleFactor
	this.PlayerOrderBucketSize = playerOrderBucketSize
	this.PlayerOrderFactor = playerOrderFactor
	this.SlowUpdateFactorThreshold = slowUpdateFactorThreshold
	this.ViewSegmentLength = viewSegmentLength
	return &this
}

// NewAPIConfigEventsWithDefaults instantiates a new APIConfigEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigEventsWithDefaults() *APIConfigEvents {
	this := APIConfigEvents{}
	return &this
}

// GetDistanceClose returns the DistanceClose field value
func (o *APIConfigEvents) GetDistanceClose() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceClose
}

// GetDistanceCloseOk returns a tuple with the DistanceClose field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceCloseOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DistanceClose, true
}

// SetDistanceClose sets field value
func (o *APIConfigEvents) SetDistanceClose(v int32) {
	o.DistanceClose = v
}

// GetDistanceFactor returns the DistanceFactor field value
func (o *APIConfigEvents) GetDistanceFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceFactor
}

// GetDistanceFactorOk returns a tuple with the DistanceFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceFactorOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DistanceFactor, true
}

// SetDistanceFactor sets field value
func (o *APIConfigEvents) SetDistanceFactor(v int32) {
	o.DistanceFactor = v
}

// GetDistanceFar returns the DistanceFar field value
func (o *APIConfigEvents) GetDistanceFar() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceFar
}

// GetDistanceFarOk returns a tuple with the DistanceFar field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceFarOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DistanceFar, true
}

// SetDistanceFar sets field value
func (o *APIConfigEvents) SetDistanceFar(v int32) {
	o.DistanceFar = v
}

// GetGroupDistance returns the GroupDistance field value
func (o *APIConfigEvents) GetGroupDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupDistance
}

// GetGroupDistanceOk returns a tuple with the GroupDistance field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetGroupDistanceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupDistance, true
}

// SetGroupDistance sets field value
func (o *APIConfigEvents) SetGroupDistance(v int32) {
	o.GroupDistance = v
}

// GetMaximumBunchSize returns the MaximumBunchSize field value
func (o *APIConfigEvents) GetMaximumBunchSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumBunchSize
}

// GetMaximumBunchSizeOk returns a tuple with the MaximumBunchSize field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetMaximumBunchSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaximumBunchSize, true
}

// SetMaximumBunchSize sets field value
func (o *APIConfigEvents) SetMaximumBunchSize(v int32) {
	o.MaximumBunchSize = v
}

// GetNotVisibleFactor returns the NotVisibleFactor field value
func (o *APIConfigEvents) GetNotVisibleFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotVisibleFactor
}

// GetNotVisibleFactorOk returns a tuple with the NotVisibleFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetNotVisibleFactorOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotVisibleFactor, true
}

// SetNotVisibleFactor sets field value
func (o *APIConfigEvents) SetNotVisibleFactor(v int32) {
	o.NotVisibleFactor = v
}

// GetPlayerOrderBucketSize returns the PlayerOrderBucketSize field value
func (o *APIConfigEvents) GetPlayerOrderBucketSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerOrderBucketSize
}

// GetPlayerOrderBucketSizeOk returns a tuple with the PlayerOrderBucketSize field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetPlayerOrderBucketSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlayerOrderBucketSize, true
}

// SetPlayerOrderBucketSize sets field value
func (o *APIConfigEvents) SetPlayerOrderBucketSize(v int32) {
	o.PlayerOrderBucketSize = v
}

// GetPlayerOrderFactor returns the PlayerOrderFactor field value
func (o *APIConfigEvents) GetPlayerOrderFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerOrderFactor
}

// GetPlayerOrderFactorOk returns a tuple with the PlayerOrderFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetPlayerOrderFactorOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlayerOrderFactor, true
}

// SetPlayerOrderFactor sets field value
func (o *APIConfigEvents) SetPlayerOrderFactor(v int32) {
	o.PlayerOrderFactor = v
}

// GetSlowUpdateFactorThreshold returns the SlowUpdateFactorThreshold field value
func (o *APIConfigEvents) GetSlowUpdateFactorThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlowUpdateFactorThreshold
}

// GetSlowUpdateFactorThresholdOk returns a tuple with the SlowUpdateFactorThreshold field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetSlowUpdateFactorThresholdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SlowUpdateFactorThreshold, true
}

// SetSlowUpdateFactorThreshold sets field value
func (o *APIConfigEvents) SetSlowUpdateFactorThreshold(v int32) {
	o.SlowUpdateFactorThreshold = v
}

// GetViewSegmentLength returns the ViewSegmentLength field value
func (o *APIConfigEvents) GetViewSegmentLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ViewSegmentLength
}

// GetViewSegmentLengthOk returns a tuple with the ViewSegmentLength field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetViewSegmentLengthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ViewSegmentLength, true
}

// SetViewSegmentLength sets field value
func (o *APIConfigEvents) SetViewSegmentLength(v int32) {
	o.ViewSegmentLength = v
}

func (o APIConfigEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distanceClose"] = o.DistanceClose
	}
	if true {
		toSerialize["distanceFactor"] = o.DistanceFactor
	}
	if true {
		toSerialize["distanceFar"] = o.DistanceFar
	}
	if true {
		toSerialize["groupDistance"] = o.GroupDistance
	}
	if true {
		toSerialize["maximumBunchSize"] = o.MaximumBunchSize
	}
	if true {
		toSerialize["notVisibleFactor"] = o.NotVisibleFactor
	}
	if true {
		toSerialize["playerOrderBucketSize"] = o.PlayerOrderBucketSize
	}
	if true {
		toSerialize["playerOrderFactor"] = o.PlayerOrderFactor
	}
	if true {
		toSerialize["slowUpdateFactorThreshold"] = o.SlowUpdateFactorThreshold
	}
	if true {
		toSerialize["viewSegmentLength"] = o.ViewSegmentLength
	}
	return json.Marshal(toSerialize)
}

type NullableAPIConfigEvents struct {
	value *APIConfigEvents
	isSet bool
}

func (v NullableAPIConfigEvents) Get() *APIConfigEvents {
	return v.value
}

func (v *NullableAPIConfigEvents) Set(val *APIConfigEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfigEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfigEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfigEvents(val *APIConfigEvents) *NullableAPIConfigEvents {
	return &NullableAPIConfigEvents{value: val, isSet: true}
}

func (v NullableAPIConfigEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfigEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

