# DigiCertConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**BaseUrl** | **string** | The base URL of the used digicert instance. | 
**ProductId** | Pointer to **string** | One of the DigiCert product identifier an exhaustive list can be found here: https://dev.digicert.com/en/certcentral-apis/services-api/glossary.html#product-identifiers | [optional] 
**ApiCredentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI | 
**OrganizationId** | **int64** |  | 
**CaCertId** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**SkipApproval** | Pointer to **NullableBool** |  | [optional] 
**CustomConnectorDataMapping** | Pointer to **map[string]string** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDigiCertConnector

`func NewDigiCertConnector(name string, type_ string, baseUrl string, apiCredentials string, organizationId int64, ) *DigiCertConnector`

NewDigiCertConnector instantiates a new DigiCertConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertConnectorWithDefaults

`func NewDigiCertConnectorWithDefaults() *DigiCertConnector`

NewDigiCertConnectorWithDefaults instantiates a new DigiCertConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DigiCertConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigiCertConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigiCertConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DigiCertConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigiCertConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigiCertConnector) SetType(v string)`

SetType sets Type field to given value.


### GetBaseUrl

`func (o *DigiCertConnector) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DigiCertConnector) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DigiCertConnector) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetProductId

`func (o *DigiCertConnector) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DigiCertConnector) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DigiCertConnector) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *DigiCertConnector) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetApiCredentials

`func (o *DigiCertConnector) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *DigiCertConnector) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *DigiCertConnector) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetOrganizationId

`func (o *DigiCertConnector) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DigiCertConnector) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DigiCertConnector) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetCaCertId

`func (o *DigiCertConnector) GetCaCertId() string`

GetCaCertId returns the CaCertId field if non-nil, zero value otherwise.

### GetCaCertIdOk

`func (o *DigiCertConnector) GetCaCertIdOk() (*string, bool)`

GetCaCertIdOk returns a tuple with the CaCertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertId

`func (o *DigiCertConnector) SetCaCertId(v string)`

SetCaCertId sets CaCertId field to given value.

### HasCaCertId

`func (o *DigiCertConnector) HasCaCertId() bool`

HasCaCertId returns a boolean if a field has been set.

### SetCaCertIdNil

`func (o *DigiCertConnector) SetCaCertIdNil(b bool)`

 SetCaCertIdNil sets the value for CaCertId to be an explicit nil

### UnsetCaCertId
`func (o *DigiCertConnector) UnsetCaCertId()`

UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
### GetRetryInterval

`func (o *DigiCertConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *DigiCertConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *DigiCertConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *DigiCertConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *DigiCertConnector) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *DigiCertConnector) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetSkipApproval

`func (o *DigiCertConnector) GetSkipApproval() bool`

GetSkipApproval returns the SkipApproval field if non-nil, zero value otherwise.

### GetSkipApprovalOk

`func (o *DigiCertConnector) GetSkipApprovalOk() (*bool, bool)`

GetSkipApprovalOk returns a tuple with the SkipApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipApproval

`func (o *DigiCertConnector) SetSkipApproval(v bool)`

SetSkipApproval sets SkipApproval field to given value.

### HasSkipApproval

`func (o *DigiCertConnector) HasSkipApproval() bool`

HasSkipApproval returns a boolean if a field has been set.

### SetSkipApprovalNil

`func (o *DigiCertConnector) SetSkipApprovalNil(b bool)`

 SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil

### UnsetSkipApproval
`func (o *DigiCertConnector) UnsetSkipApproval()`

UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
### GetCustomConnectorDataMapping

`func (o *DigiCertConnector) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *DigiCertConnector) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *DigiCertConnector) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *DigiCertConnector) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *DigiCertConnector) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *DigiCertConnector) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetTimeout

`func (o *DigiCertConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigiCertConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigiCertConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigiCertConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DigiCertConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DigiCertConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *DigiCertConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DigiCertConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DigiCertConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DigiCertConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DigiCertConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DigiCertConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *DigiCertConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *DigiCertConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *DigiCertConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *DigiCertConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *DigiCertConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *DigiCertConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


