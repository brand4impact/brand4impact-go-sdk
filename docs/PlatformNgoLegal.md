# PlatformNgoLegal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Nid** | Pointer to **string** |  | [optional] 
**Brithday** | Pointer to **time.Time** |  | [optional] 
**Nationality** | Pointer to **string** |  | [optional] 
**Residence** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**PlatformNgoLegalAddress**](PlatformNgoLegalAddress.md) |  | [optional] 

## Methods

### NewPlatformNgoLegal

`func NewPlatformNgoLegal() *PlatformNgoLegal`

NewPlatformNgoLegal instantiates a new PlatformNgoLegal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformNgoLegalWithDefaults

`func NewPlatformNgoLegalWithDefaults() *PlatformNgoLegal`

NewPlatformNgoLegalWithDefaults instantiates a new PlatformNgoLegal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformNgoLegal) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformNgoLegal) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformNgoLegal) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformNgoLegal) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetFirstname

`func (o *PlatformNgoLegal) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *PlatformNgoLegal) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *PlatformNgoLegal) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *PlatformNgoLegal) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *PlatformNgoLegal) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *PlatformNgoLegal) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *PlatformNgoLegal) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *PlatformNgoLegal) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetNid

`func (o *PlatformNgoLegal) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *PlatformNgoLegal) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *PlatformNgoLegal) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *PlatformNgoLegal) HasNid() bool`

HasNid returns a boolean if a field has been set.

### GetBrithday

`func (o *PlatformNgoLegal) GetBrithday() time.Time`

GetBrithday returns the Brithday field if non-nil, zero value otherwise.

### GetBrithdayOk

`func (o *PlatformNgoLegal) GetBrithdayOk() (*time.Time, bool)`

GetBrithdayOk returns a tuple with the Brithday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrithday

`func (o *PlatformNgoLegal) SetBrithday(v time.Time)`

SetBrithday sets Brithday field to given value.

### HasBrithday

`func (o *PlatformNgoLegal) HasBrithday() bool`

HasBrithday returns a boolean if a field has been set.

### GetNationality

`func (o *PlatformNgoLegal) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PlatformNgoLegal) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PlatformNgoLegal) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PlatformNgoLegal) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetResidence

`func (o *PlatformNgoLegal) GetResidence() string`

GetResidence returns the Residence field if non-nil, zero value otherwise.

### GetResidenceOk

`func (o *PlatformNgoLegal) GetResidenceOk() (*string, bool)`

GetResidenceOk returns a tuple with the Residence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidence

`func (o *PlatformNgoLegal) SetResidence(v string)`

SetResidence sets Residence field to given value.

### HasResidence

`func (o *PlatformNgoLegal) HasResidence() bool`

HasResidence returns a boolean if a field has been set.

### GetEmail

`func (o *PlatformNgoLegal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PlatformNgoLegal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PlatformNgoLegal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PlatformNgoLegal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *PlatformNgoLegal) GetAddress() PlatformNgoLegalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlatformNgoLegal) GetAddressOk() (*PlatformNgoLegalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlatformNgoLegal) SetAddress(v PlatformNgoLegalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PlatformNgoLegal) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


