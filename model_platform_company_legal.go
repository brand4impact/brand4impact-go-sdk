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
	"time"
)

// PlatformCompanyLegal struct for PlatformCompanyLegal
type PlatformCompanyLegal struct {
	Token *string `json:"token,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Nid *string `json:"nid,omitempty"`
	Brithday *time.Time `json:"brithday,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Residence *string `json:"residence,omitempty"`
	Email *string `json:"email,omitempty"`
	Address *PlatformCompanyLegalAddress `json:"address,omitempty"`
}

// NewPlatformCompanyLegal instantiates a new PlatformCompanyLegal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformCompanyLegal() *PlatformCompanyLegal {
	this := PlatformCompanyLegal{}
	return &this
}

// NewPlatformCompanyLegalWithDefaults instantiates a new PlatformCompanyLegal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformCompanyLegalWithDefaults() *PlatformCompanyLegal {
	this := PlatformCompanyLegal{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformCompanyLegal) SetToken(v string) {
	o.Token = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *PlatformCompanyLegal) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *PlatformCompanyLegal) SetLastname(v string) {
	o.Lastname = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetNid() string {
	if o == nil || o.Nid == nil {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetNidOk() (*string, bool) {
	if o == nil || o.Nid == nil {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasNid() bool {
	if o != nil && o.Nid != nil {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlatformCompanyLegal) SetNid(v string) {
	o.Nid = &v
}

// GetBrithday returns the Brithday field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetBrithday() time.Time {
	if o == nil || o.Brithday == nil {
		var ret time.Time
		return ret
	}
	return *o.Brithday
}

// GetBrithdayOk returns a tuple with the Brithday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetBrithdayOk() (*time.Time, bool) {
	if o == nil || o.Brithday == nil {
		return nil, false
	}
	return o.Brithday, true
}

// HasBrithday returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasBrithday() bool {
	if o != nil && o.Brithday != nil {
		return true
	}

	return false
}

// SetBrithday gets a reference to the given time.Time and assigns it to the Brithday field.
func (o *PlatformCompanyLegal) SetBrithday(v time.Time) {
	o.Brithday = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetNationality() string {
	if o == nil || o.Nationality == nil {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetNationalityOk() (*string, bool) {
	if o == nil || o.Nationality == nil {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasNationality() bool {
	if o != nil && o.Nationality != nil {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *PlatformCompanyLegal) SetNationality(v string) {
	o.Nationality = &v
}

// GetResidence returns the Residence field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetResidence() string {
	if o == nil || o.Residence == nil {
		var ret string
		return ret
	}
	return *o.Residence
}

// GetResidenceOk returns a tuple with the Residence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetResidenceOk() (*string, bool) {
	if o == nil || o.Residence == nil {
		return nil, false
	}
	return o.Residence, true
}

// HasResidence returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasResidence() bool {
	if o != nil && o.Residence != nil {
		return true
	}

	return false
}

// SetResidence gets a reference to the given string and assigns it to the Residence field.
func (o *PlatformCompanyLegal) SetResidence(v string) {
	o.Residence = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PlatformCompanyLegal) SetEmail(v string) {
	o.Email = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PlatformCompanyLegal) GetAddress() PlatformCompanyLegalAddress {
	if o == nil || o.Address == nil {
		var ret PlatformCompanyLegalAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformCompanyLegal) GetAddressOk() (*PlatformCompanyLegalAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PlatformCompanyLegal) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PlatformCompanyLegalAddress and assigns it to the Address field.
func (o *PlatformCompanyLegal) SetAddress(v PlatformCompanyLegalAddress) {
	o.Address = &v
}

func (o PlatformCompanyLegal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Nid != nil {
		toSerialize["nid"] = o.Nid
	}
	if o.Brithday != nil {
		toSerialize["brithday"] = o.Brithday
	}
	if o.Nationality != nil {
		toSerialize["nationality"] = o.Nationality
	}
	if o.Residence != nil {
		toSerialize["residence"] = o.Residence
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformCompanyLegal struct {
	value *PlatformCompanyLegal
	isSet bool
}

func (v NullablePlatformCompanyLegal) Get() *PlatformCompanyLegal {
	return v.value
}

func (v *NullablePlatformCompanyLegal) Set(val *PlatformCompanyLegal) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformCompanyLegal) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformCompanyLegal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformCompanyLegal(val *PlatformCompanyLegal) *NullablePlatformCompanyLegal {
	return &NullablePlatformCompanyLegal{value: val, isSet: true}
}

func (v NullablePlatformCompanyLegal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformCompanyLegal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


