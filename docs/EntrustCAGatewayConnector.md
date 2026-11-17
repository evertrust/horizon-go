# EntrustCAGatewayConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CaId** | **string** | Identifier of the CA on the Entrust CA Gateway | 
**CustomConnectorDataMapping** | Pointer to **map[string]string** | Properties of the enrollment request, mapped to the name of the label whose value is sent | [optional] 
**EndPoint** | **string** | Base URL of the Entrust CA Gateway, without the &#x60;/v1&#x60; API version | 
**Name** | **string** |  | 
**ProfileId** | **string** | Identifier of the CA Gateway profile used to enroll certificates | 
**Proxy** | Pointer to **string** |  | [optional] 
**Queue** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ValidityPeriod** | Pointer to **string** | Overrides the validity of the issued certificates, as an ISO 8601 interval or duration accepted by the Entrust CA Gateway. The profile must allow the validity to be overridden. Accepted forms: - a duration only, starting when the certificate is issued: &#x60;P1Y3M10DT0H0M&#x60; - a start date and an expiry date: &#x60;2018-07-06T13:00Z/2019-07-06T13:00:00Z&#x60; - a start date and a duration: &#x60;2018-07-06T13:00Z/P1Y3M0DT0H0M&#x60; - a duration and an expiry date, the start date being computed from them: &#x60;P0Y3M0DT0H0M/2018-12-31T00:00Z&#x60; | [optional] 

## Methods

### NewEntrustCAGatewayConnector

`func NewEntrustCAGatewayConnector(authenticationCredentials string, caId string, endPoint string, name string, profileId string, type_ string, ) *EntrustCAGatewayConnector`

NewEntrustCAGatewayConnector instantiates a new EntrustCAGatewayConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrustCAGatewayConnectorWithDefaults

`func NewEntrustCAGatewayConnectorWithDefaults() *EntrustCAGatewayConnector`

NewEntrustCAGatewayConnectorWithDefaults instantiates a new EntrustCAGatewayConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *EntrustCAGatewayConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *EntrustCAGatewayConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *EntrustCAGatewayConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetCaId

`func (o *EntrustCAGatewayConnector) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *EntrustCAGatewayConnector) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *EntrustCAGatewayConnector) SetCaId(v string)`

SetCaId sets CaId field to given value.


### GetCustomConnectorDataMapping

`func (o *EntrustCAGatewayConnector) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *EntrustCAGatewayConnector) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *EntrustCAGatewayConnector) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *EntrustCAGatewayConnector) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### GetEndPoint

`func (o *EntrustCAGatewayConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *EntrustCAGatewayConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *EntrustCAGatewayConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *EntrustCAGatewayConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntrustCAGatewayConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntrustCAGatewayConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfileId

`func (o *EntrustCAGatewayConnector) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EntrustCAGatewayConnector) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *EntrustCAGatewayConnector) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetProxy

`func (o *EntrustCAGatewayConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EntrustCAGatewayConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EntrustCAGatewayConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EntrustCAGatewayConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetQueue

`func (o *EntrustCAGatewayConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EntrustCAGatewayConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EntrustCAGatewayConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EntrustCAGatewayConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetTimeout

`func (o *EntrustCAGatewayConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EntrustCAGatewayConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EntrustCAGatewayConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EntrustCAGatewayConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *EntrustCAGatewayConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrustCAGatewayConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrustCAGatewayConnector) SetType(v string)`

SetType sets Type field to given value.


### GetValidityPeriod

`func (o *EntrustCAGatewayConnector) GetValidityPeriod() string`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *EntrustCAGatewayConnector) GetValidityPeriodOk() (*string, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *EntrustCAGatewayConnector) SetValidityPeriod(v string)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *EntrustCAGatewayConnector) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


