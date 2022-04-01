# PlatformCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vat** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**PlatformCompanyAddress**](PlatformCompanyAddress.md) |  | [optional] 
**Legal** | Pointer to [**PlatformCompanyLegal**](PlatformCompanyLegal.md) |  | [optional] 
**SectorToken** | Pointer to **string** |  | [optional] 
**TopicsToken** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPlatformCompany

`func NewPlatformCompany() *PlatformCompany`

NewPlatformCompany instantiates a new PlatformCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCompanyWithDefaults

`func NewPlatformCompanyWithDefaults() *PlatformCompany`

NewPlatformCompanyWithDefaults instantiates a new PlatformCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformCompany) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformCompany) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformCompany) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformCompany) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformCompany) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformCompany) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformCompany) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlatformCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVat

`func (o *PlatformCompany) GetVat() string`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *PlatformCompany) GetVatOk() (*string, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *PlatformCompany) SetVat(v string)`

SetVat sets Vat field to given value.

### HasVat

`func (o *PlatformCompany) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformCompany) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformCompany) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformCompany) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformCompany) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddress

`func (o *PlatformCompany) GetAddress() PlatformCompanyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlatformCompany) GetAddressOk() (*PlatformCompanyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlatformCompany) SetAddress(v PlatformCompanyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PlatformCompany) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLegal

`func (o *PlatformCompany) GetLegal() PlatformCompanyLegal`

GetLegal returns the Legal field if non-nil, zero value otherwise.

### GetLegalOk

`func (o *PlatformCompany) GetLegalOk() (*PlatformCompanyLegal, bool)`

GetLegalOk returns a tuple with the Legal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegal

`func (o *PlatformCompany) SetLegal(v PlatformCompanyLegal)`

SetLegal sets Legal field to given value.

### HasLegal

`func (o *PlatformCompany) HasLegal() bool`

HasLegal returns a boolean if a field has been set.

### GetSectorToken

`func (o *PlatformCompany) GetSectorToken() string`

GetSectorToken returns the SectorToken field if non-nil, zero value otherwise.

### GetSectorTokenOk

`func (o *PlatformCompany) GetSectorTokenOk() (*string, bool)`

GetSectorTokenOk returns a tuple with the SectorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorToken

`func (o *PlatformCompany) SetSectorToken(v string)`

SetSectorToken sets SectorToken field to given value.

### HasSectorToken

`func (o *PlatformCompany) HasSectorToken() bool`

HasSectorToken returns a boolean if a field has been set.

### GetTopicsToken

`func (o *PlatformCompany) GetTopicsToken() []string`

GetTopicsToken returns the TopicsToken field if non-nil, zero value otherwise.

### GetTopicsTokenOk

`func (o *PlatformCompany) GetTopicsTokenOk() (*[]string, bool)`

GetTopicsTokenOk returns a tuple with the TopicsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsToken

`func (o *PlatformCompany) SetTopicsToken(v []string)`

SetTopicsToken sets TopicsToken field to given value.

### HasTopicsToken

`func (o *PlatformCompany) HasTopicsToken() bool`

HasTopicsToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


