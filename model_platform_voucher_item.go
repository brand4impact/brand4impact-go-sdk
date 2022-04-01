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

// PlatformVoucherItem struct for PlatformVoucherItem
type PlatformVoucherItem struct {
	Token *string `json:"token,omitempty"`
	Id *string `json:"id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewPlatformVoucherItem instantiates a new PlatformVoucherItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformVoucherItem() *PlatformVoucherItem {
	this := PlatformVoucherItem{}
	return &this
}

// NewPlatformVoucherItemWithDefaults instantiates a new PlatformVoucherItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformVoucherItemWithDefaults() *PlatformVoucherItem {
	this := PlatformVoucherItem{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformVoucherItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucherItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformVoucherItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformVoucherItem) SetToken(v string) {
	o.Token = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformVoucherItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucherItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformVoucherItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlatformVoucherItem) SetId(v string) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PlatformVoucherItem) GetAmount() float64 {
	if o == nil || o.Amount == nil {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucherItem) GetAmountOk() (*float64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PlatformVoucherItem) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PlatformVoucherItem) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlatformVoucherItem) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformVoucherItem) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlatformVoucherItem) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlatformVoucherItem) SetCurrency(v string) {
	o.Currency = &v
}

func (o PlatformVoucherItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformVoucherItem struct {
	value *PlatformVoucherItem
	isSet bool
}

func (v NullablePlatformVoucherItem) Get() *PlatformVoucherItem {
	return v.value
}

func (v *NullablePlatformVoucherItem) Set(val *PlatformVoucherItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformVoucherItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformVoucherItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformVoucherItem(val *PlatformVoucherItem) *NullablePlatformVoucherItem {
	return &NullablePlatformVoucherItem{value: val, isSet: true}
}

func (v NullablePlatformVoucherItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformVoucherItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

