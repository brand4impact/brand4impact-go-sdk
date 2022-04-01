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

// PlatformSectorItem struct for PlatformSectorItem
type PlatformSectorItem struct {
	Token *string `json:"token,omitempty"`
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewPlatformSectorItem instantiates a new PlatformSectorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformSectorItem() *PlatformSectorItem {
	this := PlatformSectorItem{}
	return &this
}

// NewPlatformSectorItemWithDefaults instantiates a new PlatformSectorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformSectorItemWithDefaults() *PlatformSectorItem {
	this := PlatformSectorItem{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformSectorItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformSectorItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformSectorItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformSectorItem) SetToken(v string) {
	o.Token = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformSectorItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformSectorItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformSectorItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlatformSectorItem) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlatformSectorItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformSectorItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlatformSectorItem) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlatformSectorItem) SetTitle(v string) {
	o.Title = &v
}

func (o PlatformSectorItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformSectorItem struct {
	value *PlatformSectorItem
	isSet bool
}

func (v NullablePlatformSectorItem) Get() *PlatformSectorItem {
	return v.value
}

func (v *NullablePlatformSectorItem) Set(val *PlatformSectorItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformSectorItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformSectorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformSectorItem(val *PlatformSectorItem) *NullablePlatformSectorItem {
	return &NullablePlatformSectorItem{value: val, isSet: true}
}

func (v NullablePlatformSectorItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformSectorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


