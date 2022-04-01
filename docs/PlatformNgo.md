# PlatformNgo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vat** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**PlatformNgoAddress**](PlatformNgoAddress.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Establishment** | Pointer to **time.Time** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Legal** | Pointer to [**PlatformNgoLegal**](PlatformNgoLegal.md) |  | [optional] 
**Files** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPlatformNgo

`func NewPlatformNgo() *PlatformNgo`

NewPlatformNgo instantiates a new PlatformNgo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformNgoWithDefaults

`func NewPlatformNgoWithDefaults() *PlatformNgo`

NewPlatformNgoWithDefaults instantiates a new PlatformNgo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformNgo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformNgo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformNgo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformNgo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformNgo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformNgo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformNgo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformNgo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlatformNgo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformNgo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformNgo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformNgo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVat

`func (o *PlatformNgo) GetVat() string`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *PlatformNgo) GetVatOk() (*string, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *PlatformNgo) SetVat(v string)`

SetVat sets Vat field to given value.

### HasVat

`func (o *PlatformNgo) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetAddress

`func (o *PlatformNgo) GetAddress() PlatformNgoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlatformNgo) GetAddressOk() (*PlatformNgoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlatformNgo) SetAddress(v PlatformNgoAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PlatformNgo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformNgo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformNgo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformNgo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformNgo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEstablishment

`func (o *PlatformNgo) GetEstablishment() time.Time`

GetEstablishment returns the Establishment field if non-nil, zero value otherwise.

### GetEstablishmentOk

`func (o *PlatformNgo) GetEstablishmentOk() (*time.Time, bool)`

GetEstablishmentOk returns a tuple with the Establishment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishment

`func (o *PlatformNgo) SetEstablishment(v time.Time)`

SetEstablishment sets Establishment field to given value.

### HasEstablishment

`func (o *PlatformNgo) HasEstablishment() bool`

HasEstablishment returns a boolean if a field has been set.

### GetWebsite

`func (o *PlatformNgo) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *PlatformNgo) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *PlatformNgo) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *PlatformNgo) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLegal

`func (o *PlatformNgo) GetLegal() PlatformNgoLegal`

GetLegal returns the Legal field if non-nil, zero value otherwise.

### GetLegalOk

`func (o *PlatformNgo) GetLegalOk() (*PlatformNgoLegal, bool)`

GetLegalOk returns a tuple with the Legal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegal

`func (o *PlatformNgo) SetLegal(v PlatformNgoLegal)`

SetLegal sets Legal field to given value.

### HasLegal

`func (o *PlatformNgo) HasLegal() bool`

HasLegal returns a boolean if a field has been set.

### GetFiles

`func (o *PlatformNgo) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PlatformNgo) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PlatformNgo) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PlatformNgo) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


