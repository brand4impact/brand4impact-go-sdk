# PlatformInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Startdate** | Pointer to **time.Time** |  | [optional] 
**Enddate** | Pointer to **time.Time** |  | [optional] 
**Issuedate** | Pointer to **time.Time** |  | [optional] 
**ClientName** | Pointer to **string** |  | [optional] 
**ClientVat** | Pointer to **string** |  | [optional] 
**ClientAddress** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**map[string][]PlatformInvoiceEntry**](array.md) |  | [optional] 

## Methods

### NewPlatformInvoice

`func NewPlatformInvoice() *PlatformInvoice`

NewPlatformInvoice instantiates a new PlatformInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInvoiceWithDefaults

`func NewPlatformInvoiceWithDefaults() *PlatformInvoice`

NewPlatformInvoiceWithDefaults instantiates a new PlatformInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PlatformInvoice) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformInvoice) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformInvoice) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformInvoice) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *PlatformInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *PlatformInvoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PlatformInvoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PlatformInvoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PlatformInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStartdate

`func (o *PlatformInvoice) GetStartdate() time.Time`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *PlatformInvoice) GetStartdateOk() (*time.Time, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *PlatformInvoice) SetStartdate(v time.Time)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *PlatformInvoice) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *PlatformInvoice) GetEnddate() time.Time`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *PlatformInvoice) GetEnddateOk() (*time.Time, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *PlatformInvoice) SetEnddate(v time.Time)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *PlatformInvoice) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetIssuedate

`func (o *PlatformInvoice) GetIssuedate() time.Time`

GetIssuedate returns the Issuedate field if non-nil, zero value otherwise.

### GetIssuedateOk

`func (o *PlatformInvoice) GetIssuedateOk() (*time.Time, bool)`

GetIssuedateOk returns a tuple with the Issuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedate

`func (o *PlatformInvoice) SetIssuedate(v time.Time)`

SetIssuedate sets Issuedate field to given value.

### HasIssuedate

`func (o *PlatformInvoice) HasIssuedate() bool`

HasIssuedate returns a boolean if a field has been set.

### GetClientName

`func (o *PlatformInvoice) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *PlatformInvoice) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *PlatformInvoice) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *PlatformInvoice) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientVat

`func (o *PlatformInvoice) GetClientVat() string`

GetClientVat returns the ClientVat field if non-nil, zero value otherwise.

### GetClientVatOk

`func (o *PlatformInvoice) GetClientVatOk() (*string, bool)`

GetClientVatOk returns a tuple with the ClientVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVat

`func (o *PlatformInvoice) SetClientVat(v string)`

SetClientVat sets ClientVat field to given value.

### HasClientVat

`func (o *PlatformInvoice) HasClientVat() bool`

HasClientVat returns a boolean if a field has been set.

### GetClientAddress

`func (o *PlatformInvoice) GetClientAddress() string`

GetClientAddress returns the ClientAddress field if non-nil, zero value otherwise.

### GetClientAddressOk

`func (o *PlatformInvoice) GetClientAddressOk() (*string, bool)`

GetClientAddressOk returns a tuple with the ClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAddress

`func (o *PlatformInvoice) SetClientAddress(v string)`

SetClientAddress sets ClientAddress field to given value.

### HasClientAddress

`func (o *PlatformInvoice) HasClientAddress() bool`

HasClientAddress returns a boolean if a field has been set.

### GetEntries

`func (o *PlatformInvoice) GetEntries() map[string][]PlatformInvoiceEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *PlatformInvoice) GetEntriesOk() (*map[string][]PlatformInvoiceEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *PlatformInvoice) SetEntries(v map[string][]PlatformInvoiceEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *PlatformInvoice) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


