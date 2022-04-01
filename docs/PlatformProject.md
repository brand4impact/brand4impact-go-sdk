# PlatformProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Startdate** | Pointer to **time.Time** |  | [optional] 
**Enddate** | Pointer to **time.Time** |  | [optional] 
**PrimaryImageToken** | Pointer to **string** |  | [optional] 
**SecondaryImageToken** | Pointer to **string** |  | [optional] 
**VideoUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **float64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformProject

`func NewPlatformProject() *PlatformProject`

NewPlatformProject instantiates a new PlatformProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectWithDefaults

`func NewPlatformProjectWithDefaults() *PlatformProject`

NewPlatformProjectWithDefaults instantiates a new PlatformProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformProject) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformProject) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformProject) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformProject) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PlatformProject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlatformProject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlatformProject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlatformProject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartdate

`func (o *PlatformProject) GetStartdate() time.Time`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *PlatformProject) GetStartdateOk() (*time.Time, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *PlatformProject) SetStartdate(v time.Time)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *PlatformProject) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *PlatformProject) GetEnddate() time.Time`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *PlatformProject) GetEnddateOk() (*time.Time, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *PlatformProject) SetEnddate(v time.Time)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *PlatformProject) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetPrimaryImageToken

`func (o *PlatformProject) GetPrimaryImageToken() string`

GetPrimaryImageToken returns the PrimaryImageToken field if non-nil, zero value otherwise.

### GetPrimaryImageTokenOk

`func (o *PlatformProject) GetPrimaryImageTokenOk() (*string, bool)`

GetPrimaryImageTokenOk returns a tuple with the PrimaryImageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageToken

`func (o *PlatformProject) SetPrimaryImageToken(v string)`

SetPrimaryImageToken sets PrimaryImageToken field to given value.

### HasPrimaryImageToken

`func (o *PlatformProject) HasPrimaryImageToken() bool`

HasPrimaryImageToken returns a boolean if a field has been set.

### GetSecondaryImageToken

`func (o *PlatformProject) GetSecondaryImageToken() string`

GetSecondaryImageToken returns the SecondaryImageToken field if non-nil, zero value otherwise.

### GetSecondaryImageTokenOk

`func (o *PlatformProject) GetSecondaryImageTokenOk() (*string, bool)`

GetSecondaryImageTokenOk returns a tuple with the SecondaryImageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryImageToken

`func (o *PlatformProject) SetSecondaryImageToken(v string)`

SetSecondaryImageToken sets SecondaryImageToken field to given value.

### HasSecondaryImageToken

`func (o *PlatformProject) HasSecondaryImageToken() bool`

HasSecondaryImageToken returns a boolean if a field has been set.

### GetVideoUrl

`func (o *PlatformProject) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *PlatformProject) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *PlatformProject) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *PlatformProject) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformProject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformProject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformProject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBudget

`func (o *PlatformProject) GetBudget() float64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *PlatformProject) GetBudgetOk() (*float64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *PlatformProject) SetBudget(v float64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *PlatformProject) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCurrency

`func (o *PlatformProject) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlatformProject) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlatformProject) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PlatformProject) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


