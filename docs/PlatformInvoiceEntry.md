# PlatformInvoiceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concept** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **int32** |  | [optional] 
**Base** | Pointer to [**PlatformInvoiceEntryBase**](PlatformInvoiceEntryBase.md) |  | [optional] 
**Taxes** | Pointer to [**PlatformInvoiceEntryBase**](PlatformInvoiceEntryBase.md) |  | [optional] 

## Methods

### NewPlatformInvoiceEntry

`func NewPlatformInvoiceEntry() *PlatformInvoiceEntry`

NewPlatformInvoiceEntry instantiates a new PlatformInvoiceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInvoiceEntryWithDefaults

`func NewPlatformInvoiceEntryWithDefaults() *PlatformInvoiceEntry`

NewPlatformInvoiceEntryWithDefaults instantiates a new PlatformInvoiceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcept

`func (o *PlatformInvoiceEntry) GetConcept() string`

GetConcept returns the Concept field if non-nil, zero value otherwise.

### GetConceptOk

`func (o *PlatformInvoiceEntry) GetConceptOk() (*string, bool)`

GetConceptOk returns a tuple with the Concept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcept

`func (o *PlatformInvoiceEntry) SetConcept(v string)`

SetConcept sets Concept field to given value.

### HasConcept

`func (o *PlatformInvoiceEntry) HasConcept() bool`

HasConcept returns a boolean if a field has been set.

### GetUnits

`func (o *PlatformInvoiceEntry) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PlatformInvoiceEntry) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PlatformInvoiceEntry) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *PlatformInvoiceEntry) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetBase

`func (o *PlatformInvoiceEntry) GetBase() PlatformInvoiceEntryBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PlatformInvoiceEntry) GetBaseOk() (*PlatformInvoiceEntryBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PlatformInvoiceEntry) SetBase(v PlatformInvoiceEntryBase)`

SetBase sets Base field to given value.

### HasBase

`func (o *PlatformInvoiceEntry) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetTaxes

`func (o *PlatformInvoiceEntry) GetTaxes() PlatformInvoiceEntryBase`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *PlatformInvoiceEntry) GetTaxesOk() (*PlatformInvoiceEntryBase, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *PlatformInvoiceEntry) SetTaxes(v PlatformInvoiceEntryBase)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *PlatformInvoiceEntry) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


