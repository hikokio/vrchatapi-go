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

// Transaction 
type Transaction struct {
	Id string `json:"id"`
	Status TransactionStatus `json:"status"`
	Subscription Subscription `json:"subscription"`
	Sandbox bool `json:"sandbox"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Steam *TransactionSteamInfo `json:"steam,omitempty"`
	Agreement *TransactionAgreement `json:"agreement,omitempty"`
	Error string `json:"error"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id string, status TransactionStatus, subscription Subscription, sandbox bool, createdAt time.Time, updatedAt time.Time, error_ string) *Transaction {
	this := Transaction{}
	this.Id = id
	this.Status = status
	this.Subscription = subscription
	this.Sandbox = sandbox
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Error = error_
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	var status TransactionStatus = TRANSACTIONSTATUS_ACTIVE
	this.Status = status
	var sandbox bool = false
	this.Sandbox = sandbox
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetSubscription returns the Subscription field value
func (o *Transaction) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *Transaction) SetSubscription(v Subscription) {
	o.Subscription = v
}

// GetSandbox returns the Sandbox field value
func (o *Transaction) GetSandbox() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSandboxOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sandbox, true
}

// SetSandbox sets field value
func (o *Transaction) SetSandbox(v bool) {
	o.Sandbox = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Transaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Transaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Transaction) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Transaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSteam returns the Steam field value if set, zero value otherwise.
func (o *Transaction) GetSteam() TransactionSteamInfo {
	if o == nil || isNil(o.Steam) {
		var ret TransactionSteamInfo
		return ret
	}
	return *o.Steam
}

// GetSteamOk returns a tuple with the Steam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSteamOk() (*TransactionSteamInfo, bool) {
	if o == nil || isNil(o.Steam) {
    return nil, false
	}
	return o.Steam, true
}

// HasSteam returns a boolean if a field has been set.
func (o *Transaction) HasSteam() bool {
	if o != nil && !isNil(o.Steam) {
		return true
	}

	return false
}

// SetSteam gets a reference to the given TransactionSteamInfo and assigns it to the Steam field.
func (o *Transaction) SetSteam(v TransactionSteamInfo) {
	o.Steam = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *Transaction) GetAgreement() TransactionAgreement {
	if o == nil || isNil(o.Agreement) {
		var ret TransactionAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAgreementOk() (*TransactionAgreement, bool) {
	if o == nil || isNil(o.Agreement) {
    return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *Transaction) HasAgreement() bool {
	if o != nil && !isNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given TransactionAgreement and assigns it to the Agreement field.
func (o *Transaction) SetAgreement(v TransactionAgreement) {
	o.Agreement = &v
}

// GetError returns the Error field value
func (o *Transaction) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetErrorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Transaction) SetError(v string) {
	o.Error = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	if true {
		toSerialize["sandbox"] = o.Sandbox
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Steam) {
		toSerialize["steam"] = o.Steam
	}
	if !isNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


