# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charts** | [**[]Chart**](Chart.md) | The dashboard&#39;s list of charts | 
**Description** | Pointer to **NullableString** | The dashboard&#39;s description | [optional] 
**Name** | **string** | The dashboard&#39;s name | 
**Type** | **string** | The type of objects the dashboard displays | 

## Methods

### NewDashboard

`func NewDashboard(charts []Chart, name string, type_ string, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharts

`func (o *Dashboard) GetCharts() []Chart`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *Dashboard) GetChartsOk() (*[]Chart, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *Dashboard) SetCharts(v []Chart)`

SetCharts sets Charts field to given value.


### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Dashboard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Dashboard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Dashboard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Dashboard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Dashboard) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


