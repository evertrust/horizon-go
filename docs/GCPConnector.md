# GCPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaPool** | **string** | Identifier of the CA pool to issue from. The pool auto-selects an enabled certificate authority. | 
**CertificateLifetime** | **string** | Validity applied to every certificate issued through this connector. | 
**CertificateTemplate** | Pointer to **string** | Certificate template governing issuance policy. Accepts the template short name or its full resource path. | [optional] 
**Credentials** | Pointer to **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) holding the Google service account key (JSON). If not defined, Application Default Credentials are used (environment variable or workload identity). | [optional] 
**Endpoint** | Pointer to **string** | Overrides the default Certificate Authority Service address and port (&#x60;privateca.googleapis.com:443&#x60;). If not set, the default service URL is used. | [optional] 
**Impersonation** | Pointer to [**GCPConnectorImpersonation**](GCPConnectorImpersonation.md) |  | [optional] 
**Location** | **string** | Google Cloud location (region) of the CA pool | 
**Name** | **string** |  | 
**ProjectId** | **string** | Identifier of the Google Cloud project hosting the CA pool | 
**Proxy** | Pointer to **string** | Name of the proxy to use to connect to the GCP Api | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGCPConnector

`func NewGCPConnector(caPool string, certificateLifetime string, location string, name string, projectId string, type_ string, ) *GCPConnector`

NewGCPConnector instantiates a new GCPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPConnectorWithDefaults

`func NewGCPConnectorWithDefaults() *GCPConnector`

NewGCPConnectorWithDefaults instantiates a new GCPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaPool

`func (o *GCPConnector) GetCaPool() string`

GetCaPool returns the CaPool field if non-nil, zero value otherwise.

### GetCaPoolOk

`func (o *GCPConnector) GetCaPoolOk() (*string, bool)`

GetCaPoolOk returns a tuple with the CaPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPool

`func (o *GCPConnector) SetCaPool(v string)`

SetCaPool sets CaPool field to given value.


### GetCertificateLifetime

`func (o *GCPConnector) GetCertificateLifetime() string`

GetCertificateLifetime returns the CertificateLifetime field if non-nil, zero value otherwise.

### GetCertificateLifetimeOk

`func (o *GCPConnector) GetCertificateLifetimeOk() (*string, bool)`

GetCertificateLifetimeOk returns a tuple with the CertificateLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateLifetime

`func (o *GCPConnector) SetCertificateLifetime(v string)`

SetCertificateLifetime sets CertificateLifetime field to given value.


### GetCertificateTemplate

`func (o *GCPConnector) GetCertificateTemplate() string`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *GCPConnector) GetCertificateTemplateOk() (*string, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *GCPConnector) SetCertificateTemplate(v string)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *GCPConnector) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### GetCredentials

`func (o *GCPConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCPConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCPConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GCPConnector) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoint

`func (o *GCPConnector) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GCPConnector) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GCPConnector) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GCPConnector) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetImpersonation

`func (o *GCPConnector) GetImpersonation() GCPConnectorImpersonation`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *GCPConnector) GetImpersonationOk() (*GCPConnectorImpersonation, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *GCPConnector) SetImpersonation(v GCPConnectorImpersonation)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *GCPConnector) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetLocation

`func (o *GCPConnector) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GCPConnector) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GCPConnector) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *GCPConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GCPConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GCPConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *GCPConnector) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPConnector) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPConnector) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProxy

`func (o *GCPConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GCPConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GCPConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GCPConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetQueue

`func (o *GCPConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *GCPConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *GCPConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *GCPConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *GCPConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *GCPConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *GCPConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GCPConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GCPConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GCPConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *GCPConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


