/*
Total API Reference

Este contrato se usara para generar los clientes con todos los servicios necesarios.  `!Uso exclusivo interno!`               --- 

API version: 0.0.1
Contact: support@brand4impact.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package brand4impact

import (
	"encoding/json"
)

// PlatformVoucher struct for PlatformVoucher
type PlatformVoucher struct {
	Token *string `json:"token,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewPlatformVoucher instantiates a new PlatformVoucher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformVoucher() *PlatformVoucher {
	this := PlatformVoucher{}
	return &this
}

// NewPlatformVoucherWithDefaults instantiates a new PlatformVoucher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformVoucherWithDefaults() *PlatformVoucher {
	this := PlatformVoucher{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformVoucher) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucher) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformVoucher) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformVoucher) SetToken(v string) {
	o.Token = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformVoucher) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucher) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformVoucher) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlatformVoucher) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlatformVoucher) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucher) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlatformVoucher) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlatformVoucher) SetStatus(v string) {
	o.Status = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PlatformVoucher) GetAmount() float64 {
	if o == nil || o.Amount == nil {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucher) GetAmountOk() (*float64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PlatformVoucher) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PlatformVoucher) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlatformVoucher) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucher) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlatformVoucher) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlatformVoucher) SetCurrency(v string) {
	o.Currency = &v
}

func (o PlatformVoucher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformVoucher struct {
	value *PlatformVoucher
	isSet bool
}

func (v NullablePlatformVoucher) Get() *PlatformVoucher {
	return v.value
}

func (v *NullablePlatformVoucher) Set(val *PlatformVoucher) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformVoucher) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformVoucher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformVoucher(val *PlatformVoucher) *NullablePlatformVoucher {
	return &NullablePlatformVoucher{value: val, isSet: true}
}

func (v NullablePlatformVoucher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformVoucher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


