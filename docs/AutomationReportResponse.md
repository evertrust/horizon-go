# AutomationReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Renewable** | **bool** | If true the certificate can be renewed (all conditions for renewal are met - correct profile, renewal period, etc...) | 
**Runnable** | **bool** | If true the certificate can be renewed now (the execution period allows it) | 

## Methods

### NewAutomationReportResponse

`func NewAutomationReportResponse(renewable bool, runnable bool, ) *AutomationReportResponse`

NewAutomationReportResponse instantiates a new AutomationReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationReportResponseWithDefaults

`func NewAutomationReportResponseWithDefaults() *AutomationReportResponse`

NewAutomationReportResponseWithDefaults instantiates a new AutomationReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenewable

`func (o *AutomationReportResponse) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *AutomationReportResponse) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *AutomationReportResponse) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### GetRunnable

`func (o *AutomationReportResponse) GetRunnable() bool`

GetRunnable returns the Runnable field if non-nil, zero value otherwise.

### GetRunnableOk

`func (o *AutomationReportResponse) GetRunnableOk() (*bool, bool)`

GetRunnableOk returns a tuple with the Runnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnable

`func (o *AutomationReportResponse) SetRunnable(v bool)`

SetRunnable sets Runnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


