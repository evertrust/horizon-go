# GCPConnectorImpersonation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lifetime** | **string** |  | 
**Target** | **string** | Email of the service account to impersonate. | 

## Methods

### NewGCPConnectorImpersonation

`func NewGCPConnectorImpersonation(lifetime string, target string, ) *GCPConnectorImpersonation`

NewGCPConnectorImpersonation instantiates a new GCPConnectorImpersonation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPConnectorImpersonationWithDefaults

`func NewGCPConnectorImpersonationWithDefaults() *GCPConnectorImpersonation`

NewGCPConnectorImpersonationWithDefaults instantiates a new GCPConnectorImpersonation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifetime

`func (o *GCPConnectorImpersonation) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *GCPConnectorImpersonation) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *GCPConnectorImpersonation) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetTarget

`func (o *GCPConnectorImpersonation) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GCPConnectorImpersonation) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GCPConnectorImpersonation) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


