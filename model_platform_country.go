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

// PlatformCountry struct for PlatformCountry
type PlatformCountry struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewPlatformCountry instantiates a new PlatformCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformCountry() *PlatformCountry {
	this := PlatformCountry{}
	return &this
}

// NewPlatformCountryWithDefaults instantiates a new PlatformCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformCountryWithDefaults() *PlatformCountry {
	this := PlatformCountry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlatformCountry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCountry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlatformCountry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlatformCountry) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PlatformCountry) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCountry) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PlatformCountry) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PlatformCountry) SetCode(v string) {
	o.Code = &v
}

func (o PlatformCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformCountry struct {
	value *PlatformCountry
	isSet bool
}

func (v NullablePlatformCountry) Get() *PlatformCountry {
	return v.value
}

func (v *NullablePlatformCountry) Set(val *PlatformCountry) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformCountry) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformCountry(val *PlatformCountry) *NullablePlatformCountry {
	return &NullablePlatformCountry{value: val, isSet: true}
}

func (v NullablePlatformCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


