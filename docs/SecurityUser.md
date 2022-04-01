# SecurityUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtToken** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSecurityUser

`func NewSecurityUser() *SecurityUser`

NewSecurityUser instantiates a new SecurityUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityUserWithDefaults

`func NewSecurityUserWithDefaults() *SecurityUser`

NewSecurityUserWithDefaults instantiates a new SecurityUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtToken

`func (o *SecurityUser) GetJwtToken() string`

GetJwtToken returns the JwtToken field if non-nil, zero value otherwise.

### GetJwtTokenOk

`func (o *SecurityUser) GetJwtTokenOk() (*string, bool)`

GetJwtTokenOk returns a tuple with the JwtToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtToken

`func (o *SecurityUser) SetJwtToken(v string)`

SetJwtToken sets JwtToken field to given value.

### HasJwtToken

`func (o *SecurityUser) HasJwtToken() bool`

HasJwtToken returns a boolean if a field has been set.

### GetPlatform

`func (o *SecurityUser) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SecurityUser) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SecurityUser) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SecurityUser) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRoles

`func (o *SecurityUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SecurityUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SecurityUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SecurityUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


