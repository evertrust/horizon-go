# ChartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colors** | **[]string** | The colors of the chart | 
**Description** | Pointer to **NullableString** | The description of the chart | [optional] 
**Direction** | Pointer to **NullableString** | The sort direction applied to the chart values | [optional] 
**Fields** | **[]string** | The field that will be used to group data | 
**H** | Pointer to **NullableInt64** | The height of the chart | [optional] 
**Having** | Pointer to [**NullableHavingResponse**](HavingResponse.md) | A condition to apply to the results of the aggregate. Only the aggregates results with more than 5 items in them can be kept for example | [optional] 
**I** | Pointer to **NullableString** | The index of the chart on the dashboard | [optional] 
**Limit** | Pointer to **NullableInt64** | The maximum number of results to display | [optional] 
**LocalQuery** | Pointer to **NullableString** | The HCQL/HRQL query to build the chart from | [optional] 
**Log** | **bool** | Whether the logarithm scale is enabled or not | 
**SortOrder** | Pointer to **NullableString** | How to sort the results in the chart (if applicable) | [optional] 
**Title** | **string** | Title of the chart | 
**Type** | [**ChartType**](ChartType.md) |  | 
**W** | Pointer to **NullableInt64** | The width of the chart | [optional] 
**X** | Pointer to **NullableInt64** | The horizontal position of the chart on the grid | [optional] 
**Y** | Pointer to **NullableInt64** | The vertical position of the chart on the grid | [optional] 

## Methods

### NewChartResponse

`func NewChartResponse(colors []string, fields []string, log bool, title string, type_ ChartType, ) *ChartResponse`

NewChartResponse instantiates a new ChartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartResponseWithDefaults

`func NewChartResponseWithDefaults() *ChartResponse`

NewChartResponseWithDefaults instantiates a new ChartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColors

`func (o *ChartResponse) GetColors() []string`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *ChartResponse) GetColorsOk() (*[]string, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *ChartResponse) SetColors(v []string)`

SetColors sets Colors field to given value.


### GetDescription

`func (o *ChartResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChartResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChartResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChartResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ChartResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ChartResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *ChartResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ChartResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ChartResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ChartResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *ChartResponse) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *ChartResponse) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetFields

`func (o *ChartResponse) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ChartResponse) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ChartResponse) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetH

`func (o *ChartResponse) GetH() int64`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *ChartResponse) GetHOk() (*int64, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *ChartResponse) SetH(v int64)`

SetH sets H field to given value.

### HasH

`func (o *ChartResponse) HasH() bool`

HasH returns a boolean if a field has been set.

### SetHNil

`func (o *ChartResponse) SetHNil(b bool)`

 SetHNil sets the value for H to be an explicit nil

### UnsetH
`func (o *ChartResponse) UnsetH()`

UnsetH ensures that no value is present for H, not even an explicit nil
### GetHaving

`func (o *ChartResponse) GetHaving() HavingResponse`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *ChartResponse) GetHavingOk() (*HavingResponse, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *ChartResponse) SetHaving(v HavingResponse)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *ChartResponse) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### SetHavingNil

`func (o *ChartResponse) SetHavingNil(b bool)`

 SetHavingNil sets the value for Having to be an explicit nil

### UnsetHaving
`func (o *ChartResponse) UnsetHaving()`

UnsetHaving ensures that no value is present for Having, not even an explicit nil
### GetI

`func (o *ChartResponse) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *ChartResponse) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *ChartResponse) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *ChartResponse) HasI() bool`

HasI returns a boolean if a field has been set.

### SetINil

`func (o *ChartResponse) SetINil(b bool)`

 SetINil sets the value for I to be an explicit nil

### UnsetI
`func (o *ChartResponse) UnsetI()`

UnsetI ensures that no value is present for I, not even an explicit nil
### GetLimit

`func (o *ChartResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ChartResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ChartResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ChartResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ChartResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ChartResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetLocalQuery

`func (o *ChartResponse) GetLocalQuery() string`

GetLocalQuery returns the LocalQuery field if non-nil, zero value otherwise.

### GetLocalQueryOk

`func (o *ChartResponse) GetLocalQueryOk() (*string, bool)`

GetLocalQueryOk returns a tuple with the LocalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuery

`func (o *ChartResponse) SetLocalQuery(v string)`

SetLocalQuery sets LocalQuery field to given value.

### HasLocalQuery

`func (o *ChartResponse) HasLocalQuery() bool`

HasLocalQuery returns a boolean if a field has been set.

### SetLocalQueryNil

`func (o *ChartResponse) SetLocalQueryNil(b bool)`

 SetLocalQueryNil sets the value for LocalQuery to be an explicit nil

### UnsetLocalQuery
`func (o *ChartResponse) UnsetLocalQuery()`

UnsetLocalQuery ensures that no value is present for LocalQuery, not even an explicit nil
### GetLog

`func (o *ChartResponse) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ChartResponse) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ChartResponse) SetLog(v bool)`

SetLog sets Log field to given value.


### GetSortOrder

`func (o *ChartResponse) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ChartResponse) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ChartResponse) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ChartResponse) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ChartResponse) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ChartResponse) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetTitle

`func (o *ChartResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChartResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChartResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ChartResponse) GetType() ChartType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChartResponse) GetTypeOk() (*ChartType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChartResponse) SetType(v ChartType)`

SetType sets Type field to given value.


### GetW

`func (o *ChartResponse) GetW() int64`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *ChartResponse) GetWOk() (*int64, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *ChartResponse) SetW(v int64)`

SetW sets W field to given value.

### HasW

`func (o *ChartResponse) HasW() bool`

HasW returns a boolean if a field has been set.

### SetWNil

`func (o *ChartResponse) SetWNil(b bool)`

 SetWNil sets the value for W to be an explicit nil

### UnsetW
`func (o *ChartResponse) UnsetW()`

UnsetW ensures that no value is present for W, not even an explicit nil
### GetX

`func (o *ChartResponse) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ChartResponse) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ChartResponse) SetX(v int64)`

SetX sets X field to given value.

### HasX

`func (o *ChartResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### SetXNil

`func (o *ChartResponse) SetXNil(b bool)`

 SetXNil sets the value for X to be an explicit nil

### UnsetX
`func (o *ChartResponse) UnsetX()`

UnsetX ensures that no value is present for X, not even an explicit nil
### GetY

`func (o *ChartResponse) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ChartResponse) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ChartResponse) SetY(v int64)`

SetY sets Y field to given value.

### HasY

`func (o *ChartResponse) HasY() bool`

HasY returns a boolean if a field has been set.

### SetYNil

`func (o *ChartResponse) SetYNil(b bool)`

 SetYNil sets the value for Y to be an explicit nil

### UnsetY
`func (o *ChartResponse) UnsetY()`

UnsetY ensures that no value is present for Y, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


