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

// PlatformInvoiceEntryBase struct for PlatformInvoiceEntryBase
type PlatformInvoiceEntryBase struct {
	Value *string `json:"value,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewPlatformInvoiceEntryBase instantiates a new PlatformInvoiceEntryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformInvoiceEntryBase() *PlatformInvoiceEntryBase {
	this := PlatformInvoiceEntryBase{}
	return &this
}

// NewPlatformInvoiceEntryBaseWithDefaults instantiates a new PlatformInvoiceEntryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformInvoiceEntryBaseWithDefaults() *PlatformInvoiceEntryBase {
	this := PlatformInvoiceEntryBase{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PlatformInvoiceEntryBase) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInvoiceEntryBase) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlatformInvoiceEntryBase) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PlatformInvoiceEntryBase) SetValue(v string) {
	o.Value = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PlatformInvoiceEntryBase) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInvoiceEntryBase) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PlatformInvoiceEntryBase) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PlatformInvoiceEntryBase) SetKey(v string) {
	o.Key = &v
}

func (o PlatformInvoiceEntryBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformInvoiceEntryBase struct {
	value *PlatformInvoiceEntryBase
	isSet bool
}

func (v NullablePlatformInvoiceEntryBase) Get() *PlatformInvoiceEntryBase {
	return v.value
}

func (v *NullablePlatformInvoiceEntryBase) Set(val *PlatformInvoiceEntryBase) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformInvoiceEntryBase) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformInvoiceEntryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformInvoiceEntryBase(val *PlatformInvoiceEntryBase) *NullablePlatformInvoiceEntryBase {
	return &NullablePlatformInvoiceEntryBase{value: val, isSet: true}
}

func (v NullablePlatformInvoiceEntryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformInvoiceEntryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

