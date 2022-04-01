/*
Total API Reference

Este contrato se usara para generar los clientes con todos los servicios necesarios.  `!Uso exclusivo interno!`               --- 

API version: 0.0.1
Contact: support@brand4impact.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package brand4impact

import (
	"encoding/json"
	"time"
)

// PlatformProject struct for PlatformProject
type PlatformProject struct {
	Token *string `json:"token,omitempty"`
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Startdate *time.Time `json:"startdate,omitempty"`
	Enddate *time.Time `json:"enddate,omitempty"`
	PrimaryImageToken *string `json:"primaryImageToken,omitempty"`
	SecondaryImageToken *string `json:"secondaryImageToken,omitempty"`
	VideoUrl *string `json:"videoUrl,omitempty"`
	Status *string `json:"status,omitempty"`
	Budget *float64 `json:"budget,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewPlatformProject instantiates a new PlatformProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformProject() *PlatformProject {
	this := PlatformProject{}
	return &this
}

// NewPlatformProjectWithDefaults instantiates a new PlatformProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformProjectWithDefaults() *PlatformProject {
	this := PlatformProject{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PlatformProject) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PlatformProject) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PlatformProject) SetToken(v string) {
	o.Token = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformProject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformProject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlatformProject) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlatformProject) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlatformProject) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlatformProject) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlatformProject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlatformProject) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlatformProject) SetDescription(v string) {
	o.Description = &v
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *PlatformProject) GetStartdate() time.Time {
	if o == nil || o.Startdate == nil {
		var ret time.Time
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetStartdateOk() (*time.Time, bool) {
	if o == nil || o.Startdate == nil {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *PlatformProject) HasStartdate() bool {
	if o != nil && o.Startdate != nil {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given time.Time and assigns it to the Startdate field.
func (o *PlatformProject) SetStartdate(v time.Time) {
	o.Startdate = &v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *PlatformProject) GetEnddate() time.Time {
	if o == nil || o.Enddate == nil {
		var ret time.Time
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetEnddateOk() (*time.Time, bool) {
	if o == nil || o.Enddate == nil {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *PlatformProject) HasEnddate() bool {
	if o != nil && o.Enddate != nil {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given time.Time and assigns it to the Enddate field.
func (o *PlatformProject) SetEnddate(v time.Time) {
	o.Enddate = &v
}

// GetPrimaryImageToken returns the PrimaryImageToken field value if set, zero value otherwise.
func (o *PlatformProject) GetPrimaryImageToken() string {
	if o == nil || o.PrimaryImageToken == nil {
		var ret string
		return ret
	}
	return *o.PrimaryImageToken
}

// GetPrimaryImageTokenOk returns a tuple with the PrimaryImageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetPrimaryImageTokenOk() (*string, bool) {
	if o == nil || o.PrimaryImageToken == nil {
		return nil, false
	}
	return o.PrimaryImageToken, true
}

// HasPrimaryImageToken returns a boolean if a field has been set.
func (o *PlatformProject) HasPrimaryImageToken() bool {
	if o != nil && o.PrimaryImageToken != nil {
		return true
	}

	return false
}

// SetPrimaryImageToken gets a reference to the given string and assigns it to the PrimaryImageToken field.
func (o *PlatformProject) SetPrimaryImageToken(v string) {
	o.PrimaryImageToken = &v
}

// GetSecondaryImageToken returns the SecondaryImageToken field value if set, zero value otherwise.
func (o *PlatformProject) GetSecondaryImageToken() string {
	if o == nil || o.SecondaryImageToken == nil {
		var ret string
		return ret
	}
	return *o.SecondaryImageToken
}

// GetSecondaryImageTokenOk returns a tuple with the SecondaryImageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetSecondaryImageTokenOk() (*string, bool) {
	if o == nil || o.SecondaryImageToken == nil {
		return nil, false
	}
	return o.SecondaryImageToken, true
}

// HasSecondaryImageToken returns a boolean if a field has been set.
func (o *PlatformProject) HasSecondaryImageToken() bool {
	if o != nil && o.SecondaryImageToken != nil {
		return true
	}

	return false
}

// SetSecondaryImageToken gets a reference to the given string and assigns it to the SecondaryImageToken field.
func (o *PlatformProject) SetSecondaryImageToken(v string) {
	o.SecondaryImageToken = &v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *PlatformProject) GetVideoUrl() string {
	if o == nil || o.VideoUrl == nil {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetVideoUrlOk() (*string, bool) {
	if o == nil || o.VideoUrl == nil {
		return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *PlatformProject) HasVideoUrl() bool {
	if o != nil && o.VideoUrl != nil {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *PlatformProject) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlatformProject) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlatformProject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlatformProject) SetStatus(v string) {
	o.Status = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *PlatformProject) GetBudget() float64 {
	if o == nil || o.Budget == nil {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetBudgetOk() (*float64, bool) {
	if o == nil || o.Budget == nil {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *PlatformProject) HasBudget() bool {
	if o != nil && o.Budget != nil {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *PlatformProject) SetBudget(v float64) {
	o.Budget = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlatformProject) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformProject) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlatformProject) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlatformProject) SetCurrency(v string) {
	o.Currency = &v
}

func (o PlatformProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Startdate != nil {
		toSerialize["startdate"] = o.Startdate
	}
	if o.Enddate != nil {
		toSerialize["enddate"] = o.Enddate
	}
	if o.PrimaryImageToken != nil {
		toSerialize["primaryImageToken"] = o.PrimaryImageToken
	}
	if o.SecondaryImageToken != nil {
		toSerialize["secondaryImageToken"] = o.SecondaryImageToken
	}
	if o.VideoUrl != nil {
		toSerialize["videoUrl"] = o.VideoUrl
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Budget != nil {
		toSerialize["budget"] = o.Budget
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformProject struct {
	value *PlatformProject
	isSet bool
}

func (v NullablePlatformProject) Get() *PlatformProject {
	return v.value
}

func (v *NullablePlatformProject) Set(val *PlatformProject) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformProject) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformProject(val *PlatformProject) *NullablePlatformProject {
	return &NullablePlatformProject{value: val, isSet: true}
}

func (v NullablePlatformProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

