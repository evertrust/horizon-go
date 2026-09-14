# GCPConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**CaPool** | **string** | Identifier of the CA pool to issue from. The pool auto-selects an enabled certificate authority. | 
**CertificateLifetime** | **string** | Validity applied to every certificate issued through this connector. | 
**CertificateTemplate** | Pointer to **string** | Certificate template governing issuance policy. Accepts the template short name or its full resource path. | [optional] 
**Credentials** | Pointer to **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) holding the Google service account key (JSON). If not defined, Application Default Credentials are used (environment variable or workload identity). | [optional] 
**Endpoint** | Pointer to **string** | Overrides the default Certificate Authority Service address and port (&#x60;privateca.googleapis.com:443&#x60;). If not set, the default service URL is used. | [optional] 
**Impersonation** | Pointer to [**GCPConnectorImpersonation**](GCPConnectorImpersonation.md) |  | [optional] 
**Location** | **string** | Google Cloud location (region) of the CA pool. Also used by the connector to derive the Certificate Authority Service endpoint. | 
**Name** | **string** |  | 
**ProjectId** | **string** | Identifier of the Google Cloud project hosting the CA pool. May differ from the project of the service account used to authenticate. | 
**Proxy** | Pointer to **string** | Name of the proxy to use to connect to the GCP Api | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Timeout** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGCPConnectorResponse

`func NewGCPConnectorResponse(id string, caPool string, certificateLifetime string, location string, name string, projectId string, type_ string, ) *GCPConnectorResponse`

NewGCPConnectorResponse instantiates a new GCPConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPConnectorResponseWithDefaults

`func NewGCPConnectorResponseWithDefaults() *GCPConnectorResponse`

NewGCPConnectorResponseWithDefaults instantiates a new GCPConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GCPConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCPConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCPConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCaPool

`func (o *GCPConnectorResponse) GetCaPool() string`

GetCaPool returns the CaPool field if non-nil, zero value otherwise.

### GetCaPoolOk

`func (o *GCPConnectorResponse) GetCaPoolOk() (*string, bool)`

GetCaPoolOk returns a tuple with the CaPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPool

`func (o *GCPConnectorResponse) SetCaPool(v string)`

SetCaPool sets CaPool field to given value.


### GetCertificateLifetime

`func (o *GCPConnectorResponse) GetCertificateLifetime() string`

GetCertificateLifetime returns the CertificateLifetime field if non-nil, zero value otherwise.

### GetCertificateLifetimeOk

`func (o *GCPConnectorResponse) GetCertificateLifetimeOk() (*string, bool)`

GetCertificateLifetimeOk returns a tuple with the CertificateLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateLifetime

`func (o *GCPConnectorResponse) SetCertificateLifetime(v string)`

SetCertificateLifetime sets CertificateLifetime field to given value.


### GetCertificateTemplate

`func (o *GCPConnectorResponse) GetCertificateTemplate() string`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *GCPConnectorResponse) GetCertificateTemplateOk() (*string, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *GCPConnectorResponse) SetCertificateTemplate(v string)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *GCPConnectorResponse) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### GetCredentials

`func (o *GCPConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCPConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCPConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GCPConnectorResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoint

`func (o *GCPConnectorResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GCPConnectorResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GCPConnectorResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GCPConnectorResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetImpersonation

`func (o *GCPConnectorResponse) GetImpersonation() GCPConnectorImpersonation`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *GCPConnectorResponse) GetImpersonationOk() (*GCPConnectorImpersonation, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *GCPConnectorResponse) SetImpersonation(v GCPConnectorImpersonation)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *GCPConnectorResponse) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetLocation

`func (o *GCPConnectorResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GCPConnectorResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GCPConnectorResponse) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *GCPConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GCPConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GCPConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *GCPConnectorResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPConnectorResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPConnectorResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProxy

`func (o *GCPConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GCPConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GCPConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GCPConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetQueue

`func (o *GCPConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *GCPConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *GCPConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *GCPConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *GCPConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *GCPConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *GCPConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GCPConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GCPConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GCPConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GCPConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GCPConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeout

`func (o *GCPConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GCPConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GCPConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GCPConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *GCPConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


