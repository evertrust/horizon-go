# Chart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the chart | 
**Description** | Pointer to **NullableString** | The description of the chart | [optional] 
**Type** | **string** | The type of the chart | 
**Fields** | **[]string** | The field that will be used to group data | 
**Limit** | Pointer to **NullableInt64** | The maximum number of results to display | [optional] 
**Having** | Pointer to [**NullableHaving**](Having.md) | A condition to apply to the results of the aggregate. Only the aggregates results with more than 5 items in them can be kept for example | [optional] 
**SortOrder** | Pointer to **NullableString** | How to sort the results in the chart (if applicable) | [optional] 
**LocalQuery** | Pointer to **NullableString** | The HCQL/HRQL query to build the chart from | [optional] 
**Direction** | Pointer to **NullableString** |  | [optional] 
**Colors** | **[]string** | The colors of the chart | 
**I** | Pointer to **NullableString** | The index of the chart on the dashboard | [optional] 
**X** | Pointer to **NullableInt64** | The horizontal position of the chart on the grid | [optional] 
**Y** | Pointer to **NullableInt64** | The vertical position of the chart on the grid | [optional] 
**W** | Pointer to **NullableInt64** | The width of the chart | [optional] 
**H** | Pointer to **NullableInt64** | The height of the chart | [optional] 
**Log** | **bool** | Whether the logarithm scale is enabled or not | 

## Methods

### NewChart

`func NewChart(title string, type_ string, fields []string, colors []string, log bool, ) *Chart`

NewChart instantiates a new Chart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartWithDefaults

`func NewChartWithDefaults() *Chart`

NewChartWithDefaults instantiates a new Chart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Chart) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Chart) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Chart) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Chart) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Chart) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Chart) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Chart) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Chart) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Chart) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *Chart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Chart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Chart) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *Chart) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Chart) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Chart) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetLimit

`func (o *Chart) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Chart) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Chart) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Chart) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Chart) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Chart) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetHaving

`func (o *Chart) GetHaving() Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *Chart) GetHavingOk() (*Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *Chart) SetHaving(v Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *Chart) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### SetHavingNil

`func (o *Chart) SetHavingNil(b bool)`

 SetHavingNil sets the value for Having to be an explicit nil

### UnsetHaving
`func (o *Chart) UnsetHaving()`

UnsetHaving ensures that no value is present for Having, not even an explicit nil
### GetSortOrder

`func (o *Chart) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Chart) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Chart) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Chart) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *Chart) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Chart) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetLocalQuery

`func (o *Chart) GetLocalQuery() string`

GetLocalQuery returns the LocalQuery field if non-nil, zero value otherwise.

### GetLocalQueryOk

`func (o *Chart) GetLocalQueryOk() (*string, bool)`

GetLocalQueryOk returns a tuple with the LocalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuery

`func (o *Chart) SetLocalQuery(v string)`

SetLocalQuery sets LocalQuery field to given value.

### HasLocalQuery

`func (o *Chart) HasLocalQuery() bool`

HasLocalQuery returns a boolean if a field has been set.

### SetLocalQueryNil

`func (o *Chart) SetLocalQueryNil(b bool)`

 SetLocalQueryNil sets the value for LocalQuery to be an explicit nil

### UnsetLocalQuery
`func (o *Chart) UnsetLocalQuery()`

UnsetLocalQuery ensures that no value is present for LocalQuery, not even an explicit nil
### GetDirection

`func (o *Chart) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Chart) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Chart) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Chart) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *Chart) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *Chart) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetColors

`func (o *Chart) GetColors() []string`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *Chart) GetColorsOk() (*[]string, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *Chart) SetColors(v []string)`

SetColors sets Colors field to given value.


### GetI

`func (o *Chart) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *Chart) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *Chart) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *Chart) HasI() bool`

HasI returns a boolean if a field has been set.

### SetINil

`func (o *Chart) SetINil(b bool)`

 SetINil sets the value for I to be an explicit nil

### UnsetI
`func (o *Chart) UnsetI()`

UnsetI ensures that no value is present for I, not even an explicit nil
### GetX

`func (o *Chart) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Chart) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Chart) SetX(v int64)`

SetX sets X field to given value.

### HasX

`func (o *Chart) HasX() bool`

HasX returns a boolean if a field has been set.

### SetXNil

`func (o *Chart) SetXNil(b bool)`

 SetXNil sets the value for X to be an explicit nil

### UnsetX
`func (o *Chart) UnsetX()`

UnsetX ensures that no value is present for X, not even an explicit nil
### GetY

`func (o *Chart) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Chart) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Chart) SetY(v int64)`

SetY sets Y field to given value.

### HasY

`func (o *Chart) HasY() bool`

HasY returns a boolean if a field has been set.

### SetYNil

`func (o *Chart) SetYNil(b bool)`

 SetYNil sets the value for Y to be an explicit nil

### UnsetY
`func (o *Chart) UnsetY()`

UnsetY ensures that no value is present for Y, not even an explicit nil
### GetW

`func (o *Chart) GetW() int64`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *Chart) GetWOk() (*int64, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *Chart) SetW(v int64)`

SetW sets W field to given value.

### HasW

`func (o *Chart) HasW() bool`

HasW returns a boolean if a field has been set.

### SetWNil

`func (o *Chart) SetWNil(b bool)`

 SetWNil sets the value for W to be an explicit nil

### UnsetW
`func (o *Chart) UnsetW()`

UnsetW ensures that no value is present for W, not even an explicit nil
### GetH

`func (o *Chart) GetH() int64`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *Chart) GetHOk() (*int64, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *Chart) SetH(v int64)`

SetH sets H field to given value.

### HasH

`func (o *Chart) HasH() bool`

HasH returns a boolean if a field has been set.

### SetHNil

`func (o *Chart) SetHNil(b bool)`

 SetHNil sets the value for H to be an explicit nil

### UnsetH
`func (o *Chart) UnsetH()`

UnsetH ensures that no value is present for H, not even an explicit nil
### GetLog

`func (o *Chart) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *Chart) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *Chart) SetLog(v bool)`

SetLog sets Log field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


