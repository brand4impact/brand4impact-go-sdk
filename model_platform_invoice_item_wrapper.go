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

// PlatformInvoiceItemWrapper struct for PlatformInvoiceItemWrapper
type PlatformInvoiceItemWrapper struct {
	Items *[]PlatformInvoiceItem `json:"items,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewPlatformInvoiceItemWrapper instantiates a new PlatformInvoiceItemWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformInvoiceItemWrapper() *PlatformInvoiceItemWrapper {
	this := PlatformInvoiceItemWrapper{}
	return &this
}

// NewPlatformInvoiceItemWrapperWithDefaults instantiates a new PlatformInvoiceItemWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformInvoiceItemWrapperWithDefaults() *PlatformInvoiceItemWrapper {
	this := PlatformInvoiceItemWrapper{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PlatformInvoiceItemWrapper) GetItems() []PlatformInvoiceItem {
	if o == nil || o.Items == nil {
		var ret []PlatformInvoiceItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInvoiceItemWrapper) GetItemsOk() (*[]PlatformInvoiceItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PlatformInvoiceItemWrapper) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PlatformInvoiceItem and assigns it to the Items field.
func (o *PlatformInvoiceItemWrapper) SetItems(v []PlatformInvoiceItem) {
	o.Items = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PlatformInvoiceItemWrapper) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformInvoiceItemWrapper) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PlatformInvoiceItemWrapper) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *PlatformInvoiceItemWrapper) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o PlatformInvoiceItemWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformInvoiceItemWrapper struct {
	value *PlatformInvoiceItemWrapper
	isSet bool
}

func (v NullablePlatformInvoiceItemWrapper) Get() *PlatformInvoiceItemWrapper {
	return v.value
}

func (v *NullablePlatformInvoiceItemWrapper) Set(val *PlatformInvoiceItemWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformInvoiceItemWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformInvoiceItemWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformInvoiceItemWrapper(val *PlatformInvoiceItemWrapper) *NullablePlatformInvoiceItemWrapper {
	return &NullablePlatformInvoiceItemWrapper{value: val, isSet: true}
}

func (v NullablePlatformInvoiceItemWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformInvoiceItemWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


