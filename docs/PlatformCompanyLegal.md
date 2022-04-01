# PlatformCompanyLegal

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
**Address** | Pointer to [**PlatformCompanyLegalAddress**](PlatformCompanyLegalAddress.md) |  | [optional] 

## Methods

### NewPlatformCompanyLegal

`func NewPlatformCompanyLegal() *PlatformCompanyLegal`

NewPlatformCompanyLegal instantiates a new PlatformCompanyLegal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCompanyLegalWithDefaults

`func NewPlatformCompanyLegalWithDefaults() *PlatformCompanyLegal`

NewPlatformCompanyLegalWithDefaults instantiates a new PlatformCompanyLegal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformCompanyLegal) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformCompanyLegal) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformCompanyLegal) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformCompanyLegal) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetFirstname

`func (o *PlatformCompanyLegal) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *PlatformCompanyLegal) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *PlatformCompanyLegal) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *PlatformCompanyLegal) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *PlatformCompanyLegal) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *PlatformCompanyLegal) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *PlatformCompanyLegal) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *PlatformCompanyLegal) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetNid

`func (o *PlatformCompanyLegal) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *PlatformCompanyLegal) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *PlatformCompanyLegal) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *PlatformCompanyLegal) HasNid() bool`

HasNid returns a boolean if a field has been set.

### GetBrithday

`func (o *PlatformCompanyLegal) GetBrithday() time.Time`

GetBrithday returns the Brithday field if non-nil, zero value otherwise.

### GetBrithdayOk

`func (o *PlatformCompanyLegal) GetBrithdayOk() (*time.Time, bool)`

GetBrithdayOk returns a tuple with the Brithday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrithday

`func (o *PlatformCompanyLegal) SetBrithday(v time.Time)`

SetBrithday sets Brithday field to given value.

### HasBrithday

`func (o *PlatformCompanyLegal) HasBrithday() bool`

HasBrithday returns a boolean if a field has been set.

### GetNationality

`func (o *PlatformCompanyLegal) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PlatformCompanyLegal) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PlatformCompanyLegal) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PlatformCompanyLegal) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetResidence

`func (o *PlatformCompanyLegal) GetResidence() string`

GetResidence returns the Residence field if non-nil, zero value otherwise.

### GetResidenceOk

`func (o *PlatformCompanyLegal) GetResidenceOk() (*string, bool)`

GetResidenceOk returns a tuple with the Residence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidence

`func (o *PlatformCompanyLegal) SetResidence(v string)`

SetResidence sets Residence field to given value.

### HasResidence

`func (o *PlatformCompanyLegal) HasResidence() bool`

HasResidence returns a boolean if a field has been set.

### GetEmail

`func (o *PlatformCompanyLegal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PlatformCompanyLegal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PlatformCompanyLegal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PlatformCompanyLegal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *PlatformCompanyLegal) GetAddress() PlatformCompanyLegalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlatformCompanyLegal) GetAddressOk() (*PlatformCompanyLegalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlatformCompanyLegal) SetAddress(v PlatformCompanyLegalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PlatformCompanyLegal) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


