# ReportCSVMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The unique identifier of the report CSV. With that UUID, it is possible to download the CSV | 
**RemoveAt** | **int64** | Indicates when the report CSV will be deleted from database | 
**ReportId** | **string** | Id of the report that generated this CSV | 
**ReportName** | **string** | Name of the report that generated this CSV | 

## Methods

### NewReportCSVMetadata

`func NewReportCSVMetadata(uuid string, removeAt int64, reportId string, reportName string, ) *ReportCSVMetadata`

NewReportCSVMetadata instantiates a new ReportCSVMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCSVMetadataWithDefaults

`func NewReportCSVMetadataWithDefaults() *ReportCSVMetadata`

NewReportCSVMetadataWithDefaults instantiates a new ReportCSVMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ReportCSVMetadata) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ReportCSVMetadata) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ReportCSVMetadata) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRemoveAt

`func (o *ReportCSVMetadata) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *ReportCSVMetadata) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *ReportCSVMetadata) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetReportId

`func (o *ReportCSVMetadata) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ReportCSVMetadata) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ReportCSVMetadata) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetReportName

`func (o *ReportCSVMetadata) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *ReportCSVMetadata) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *ReportCSVMetadata) SetReportName(v string)`

SetReportName sets ReportName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


