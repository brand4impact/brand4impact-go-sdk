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

// PlatformAsset struct for PlatformAsset
type PlatformAsset struct {
	Token *string `json:"token,omitempty"`
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	File *[]string `json:"file,omitempty"`
}

// NewPlatformAsset instantiates a new PlatformAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformAsset() *PlatformAsset {
	this := PlatformAsset{}
	return &this
}

// NewPlatformAssetWithDefaults instantiates a new PlatformAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformAssetWithDefaults() *PlatformAsset {
	this := PlatformAsset{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformAsset) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAsset) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformAsset) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformAsset) SetToken(v string) {
	o.Token = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformAsset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAsset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformAsset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlatformAsset) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlatformAsset) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAsset) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlatformAsset) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlatformAsset) SetTitle(v string) {
	o.Title = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *PlatformAsset) GetFile() []string {
	if o == nil || o.File == nil {
		var ret []string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAsset) GetFileOk() (*[]string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *PlatformAsset) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given []string and assigns it to the File field.
func (o *PlatformAsset) SetFile(v []string) {
	o.File = &v
}

func (o PlatformAsset) MarshalJSON() ([]byte, error) {
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
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformAsset struct {
	value *PlatformAsset
	isSet bool
}

func (v NullablePlatformAsset) Get() *PlatformAsset {
	return v.value
}

func (v *NullablePlatformAsset) Set(val *PlatformAsset) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformAsset) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformAsset(val *PlatformAsset) *NullablePlatformAsset {
	return &NullablePlatformAsset{value: val, isSet: true}
}

func (v NullablePlatformAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

