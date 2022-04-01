# RestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRestError

`func NewRestError() *RestError`

NewRestError instantiates a new RestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestErrorWithDefaults

`func NewRestErrorWithDefaults() *RestError`

NewRestErrorWithDefaults instantiates a new RestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RestError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *RestError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RestError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTimestamp

`func (o *RestError) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestError) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestError) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RestError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetZone

`func (o *RestError) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RestError) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RestError) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RestError) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetMessage

`func (o *RestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *RestError) GetDetails() []string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RestError) GetDetailsOk() (*[]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RestError) SetDetails(v []string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RestError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


