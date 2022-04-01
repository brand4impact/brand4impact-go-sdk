# PlatformInvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Amounts** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPlatformInvoiceItem

`func NewPlatformInvoiceItem() *PlatformInvoiceItem`

NewPlatformInvoiceItem instantiates a new PlatformInvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInvoiceItemWithDefaults

`func NewPlatformInvoiceItemWithDefaults() *PlatformInvoiceItem`

NewPlatformInvoiceItemWithDefaults instantiates a new PlatformInvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformInvoiceItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformInvoiceItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformInvoiceItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformInvoiceItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformInvoiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformInvoiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformInvoiceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformInvoiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *PlatformInvoiceItem) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PlatformInvoiceItem) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PlatformInvoiceItem) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PlatformInvoiceItem) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAmounts

`func (o *PlatformInvoiceItem) GetAmounts() map[string]string`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *PlatformInvoiceItem) GetAmountsOk() (*map[string]string, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *PlatformInvoiceItem) SetAmounts(v map[string]string)`

SetAmounts sets Amounts field to given value.

### HasAmounts

`func (o *PlatformInvoiceItem) HasAmounts() bool`

HasAmounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


