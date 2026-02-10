# AWSACMPCAConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Region** | **string** |  | 
**CaArn** | **string** |  | 
**AccessCredentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used. | [optional] 
**TemplateArn** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**ValidDays** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**SigningHash** | Pointer to **NullableString** |  | [optional] 
**CertificateUsage** | Pointer to **NullableString** |  | [optional] 
**CaPolicyOid** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 

## Methods

### NewAWSACMPCAConnectorResponse

`func NewAWSACMPCAConnectorResponse(id string, name string, type_ string, region string, caArn string, ) *AWSACMPCAConnectorResponse`

NewAWSACMPCAConnectorResponse instantiates a new AWSACMPCAConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSACMPCAConnectorResponseWithDefaults

`func NewAWSACMPCAConnectorResponseWithDefaults() *AWSACMPCAConnectorResponse`

NewAWSACMPCAConnectorResponseWithDefaults instantiates a new AWSACMPCAConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AWSACMPCAConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSACMPCAConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSACMPCAConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AWSACMPCAConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AWSACMPCAConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AWSACMPCAConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AWSACMPCAConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSACMPCAConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSACMPCAConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRegion

`func (o *AWSACMPCAConnectorResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSACMPCAConnectorResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSACMPCAConnectorResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCaArn

`func (o *AWSACMPCAConnectorResponse) GetCaArn() string`

GetCaArn returns the CaArn field if non-nil, zero value otherwise.

### GetCaArnOk

`func (o *AWSACMPCAConnectorResponse) GetCaArnOk() (*string, bool)`

GetCaArnOk returns a tuple with the CaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaArn

`func (o *AWSACMPCAConnectorResponse) SetCaArn(v string)`

SetCaArn sets CaArn field to given value.


### GetAccessCredentials

`func (o *AWSACMPCAConnectorResponse) GetAccessCredentials() string`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *AWSACMPCAConnectorResponse) GetAccessCredentialsOk() (*string, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *AWSACMPCAConnectorResponse) SetAccessCredentials(v string)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *AWSACMPCAConnectorResponse) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### SetAccessCredentialsNil

`func (o *AWSACMPCAConnectorResponse) SetAccessCredentialsNil(b bool)`

 SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil

### UnsetAccessCredentials
`func (o *AWSACMPCAConnectorResponse) UnsetAccessCredentials()`

UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
### GetTemplateArn

`func (o *AWSACMPCAConnectorResponse) GetTemplateArn() string`

GetTemplateArn returns the TemplateArn field if non-nil, zero value otherwise.

### GetTemplateArnOk

`func (o *AWSACMPCAConnectorResponse) GetTemplateArnOk() (*string, bool)`

GetTemplateArnOk returns a tuple with the TemplateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArn

`func (o *AWSACMPCAConnectorResponse) SetTemplateArn(v string)`

SetTemplateArn sets TemplateArn field to given value.

### HasTemplateArn

`func (o *AWSACMPCAConnectorResponse) HasTemplateArn() bool`

HasTemplateArn returns a boolean if a field has been set.

### SetTemplateArnNil

`func (o *AWSACMPCAConnectorResponse) SetTemplateArnNil(b bool)`

 SetTemplateArnNil sets the value for TemplateArn to be an explicit nil

### UnsetTemplateArn
`func (o *AWSACMPCAConnectorResponse) UnsetTemplateArn()`

UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
### GetRoleArn

`func (o *AWSACMPCAConnectorResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSACMPCAConnectorResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSACMPCAConnectorResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AWSACMPCAConnectorResponse) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *AWSACMPCAConnectorResponse) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AWSACMPCAConnectorResponse) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetValidDays

`func (o *AWSACMPCAConnectorResponse) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *AWSACMPCAConnectorResponse) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *AWSACMPCAConnectorResponse) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *AWSACMPCAConnectorResponse) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *AWSACMPCAConnectorResponse) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *AWSACMPCAConnectorResponse) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetRetryInterval

`func (o *AWSACMPCAConnectorResponse) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *AWSACMPCAConnectorResponse) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *AWSACMPCAConnectorResponse) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *AWSACMPCAConnectorResponse) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *AWSACMPCAConnectorResponse) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *AWSACMPCAConnectorResponse) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetSigningHash

`func (o *AWSACMPCAConnectorResponse) GetSigningHash() string`

GetSigningHash returns the SigningHash field if non-nil, zero value otherwise.

### GetSigningHashOk

`func (o *AWSACMPCAConnectorResponse) GetSigningHashOk() (*string, bool)`

GetSigningHashOk returns a tuple with the SigningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningHash

`func (o *AWSACMPCAConnectorResponse) SetSigningHash(v string)`

SetSigningHash sets SigningHash field to given value.

### HasSigningHash

`func (o *AWSACMPCAConnectorResponse) HasSigningHash() bool`

HasSigningHash returns a boolean if a field has been set.

### SetSigningHashNil

`func (o *AWSACMPCAConnectorResponse) SetSigningHashNil(b bool)`

 SetSigningHashNil sets the value for SigningHash to be an explicit nil

### UnsetSigningHash
`func (o *AWSACMPCAConnectorResponse) UnsetSigningHash()`

UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
### GetCertificateUsage

`func (o *AWSACMPCAConnectorResponse) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *AWSACMPCAConnectorResponse) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *AWSACMPCAConnectorResponse) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *AWSACMPCAConnectorResponse) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *AWSACMPCAConnectorResponse) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *AWSACMPCAConnectorResponse) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetCaPolicyOid

`func (o *AWSACMPCAConnectorResponse) GetCaPolicyOid() string`

GetCaPolicyOid returns the CaPolicyOid field if non-nil, zero value otherwise.

### GetCaPolicyOidOk

`func (o *AWSACMPCAConnectorResponse) GetCaPolicyOidOk() (*string, bool)`

GetCaPolicyOidOk returns a tuple with the CaPolicyOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPolicyOid

`func (o *AWSACMPCAConnectorResponse) SetCaPolicyOid(v string)`

SetCaPolicyOid sets CaPolicyOid field to given value.

### HasCaPolicyOid

`func (o *AWSACMPCAConnectorResponse) HasCaPolicyOid() bool`

HasCaPolicyOid returns a boolean if a field has been set.

### SetCaPolicyOidNil

`func (o *AWSACMPCAConnectorResponse) SetCaPolicyOidNil(b bool)`

 SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil

### UnsetCaPolicyOid
`func (o *AWSACMPCAConnectorResponse) UnsetCaPolicyOid()`

UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
### GetTimeout

`func (o *AWSACMPCAConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AWSACMPCAConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AWSACMPCAConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AWSACMPCAConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AWSACMPCAConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AWSACMPCAConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *AWSACMPCAConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AWSACMPCAConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AWSACMPCAConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AWSACMPCAConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AWSACMPCAConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AWSACMPCAConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *AWSACMPCAConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AWSACMPCAConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AWSACMPCAConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AWSACMPCAConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *AWSACMPCAConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *AWSACMPCAConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *AWSACMPCAConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AWSACMPCAConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AWSACMPCAConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AWSACMPCAConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AWSACMPCAConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AWSACMPCAConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


