# PlatformCampaign

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

## Methods

### NewPlatformCampaign

`func NewPlatformCampaign() *PlatformCampaign`

NewPlatformCampaign instantiates a new PlatformCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCampaignWithDefaults

`func NewPlatformCampaignWithDefaults() *PlatformCampaign`

NewPlatformCampaignWithDefaults instantiates a new PlatformCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformCampaign) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformCampaign) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformCampaign) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformCampaign) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PlatformCampaign) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlatformCampaign) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlatformCampaign) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlatformCampaign) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartdate

`func (o *PlatformCampaign) GetStartdate() time.Time`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *PlatformCampaign) GetStartdateOk() (*time.Time, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *PlatformCampaign) SetStartdate(v time.Time)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *PlatformCampaign) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *PlatformCampaign) GetEnddate() time.Time`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *PlatformCampaign) GetEnddateOk() (*time.Time, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *PlatformCampaign) SetEnddate(v time.Time)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *PlatformCampaign) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetPrimaryImageToken

`func (o *PlatformCampaign) GetPrimaryImageToken() string`

GetPrimaryImageToken returns the PrimaryImageToken field if non-nil, zero value otherwise.

### GetPrimaryImageTokenOk

`func (o *PlatformCampaign) GetPrimaryImageTokenOk() (*string, bool)`

GetPrimaryImageTokenOk returns a tuple with the PrimaryImageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageToken

`func (o *PlatformCampaign) SetPrimaryImageToken(v string)`

SetPrimaryImageToken sets PrimaryImageToken field to given value.

### HasPrimaryImageToken

`func (o *PlatformCampaign) HasPrimaryImageToken() bool`

HasPrimaryImageToken returns a boolean if a field has been set.

### GetSecondaryImageToken

`func (o *PlatformCampaign) GetSecondaryImageToken() string`

GetSecondaryImageToken returns the SecondaryImageToken field if non-nil, zero value otherwise.

### GetSecondaryImageTokenOk

`func (o *PlatformCampaign) GetSecondaryImageTokenOk() (*string, bool)`

GetSecondaryImageTokenOk returns a tuple with the SecondaryImageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryImageToken

`func (o *PlatformCampaign) SetSecondaryImageToken(v string)`

SetSecondaryImageToken sets SecondaryImageToken field to given value.

### HasSecondaryImageToken

`func (o *PlatformCampaign) HasSecondaryImageToken() bool`

HasSecondaryImageToken returns a boolean if a field has been set.

### GetVideoUrl

`func (o *PlatformCampaign) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *PlatformCampaign) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *PlatformCampaign) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *PlatformCampaign) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


