# DashboardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charts** | [**[]ChartResponse**](ChartResponse.md) | The dashboard&#39;s list of charts | 
**Description** | Pointer to **NullableString** | The dashboard&#39;s description | [optional] 
**Name** | **string** | The dashboard&#39;s name | 
**Type** | **string** | The type of objects the dashboard displays | 

## Methods

### NewDashboardResponse

`func NewDashboardResponse(charts []ChartResponse, name string, type_ string, ) *DashboardResponse`

NewDashboardResponse instantiates a new DashboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardResponseWithDefaults

`func NewDashboardResponseWithDefaults() *DashboardResponse`

NewDashboardResponseWithDefaults instantiates a new DashboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharts

`func (o *DashboardResponse) GetCharts() []ChartResponse`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *DashboardResponse) GetChartsOk() (*[]ChartResponse, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *DashboardResponse) SetCharts(v []ChartResponse)`

SetCharts sets Charts field to given value.


### GetDescription

`func (o *DashboardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DashboardResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DashboardResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *DashboardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DashboardResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


