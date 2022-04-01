# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Elements** | Pointer to **int64** |  | [optional] 
**Sort** | Pointer to [**ApiSort**](ApiSort.md) |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *Pagination) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Pagination) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Pagination) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *Pagination) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSize

`func (o *Pagination) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Pagination) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Pagination) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Pagination) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetElements

`func (o *Pagination) GetElements() int64`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Pagination) GetElementsOk() (*int64, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Pagination) SetElements(v int64)`

SetElements sets Elements field to given value.

### HasElements

`func (o *Pagination) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetSort

`func (o *Pagination) GetSort() ApiSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Pagination) GetSortOk() (*ApiSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Pagination) SetSort(v ApiSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Pagination) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


