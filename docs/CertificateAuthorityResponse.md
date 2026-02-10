# CertificateAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Certificate** | [**CFCertificate**](CFCertificate.md) |  | 
**CacheTimeToIdle** | Pointer to **NullableString** |  | [optional] 
**CrlUrl** | Pointer to **NullableString** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**OutdatedRevocationStatusPolicy** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Public** | **bool** |  | 
**Refresh** | Pointer to **NullableString** |  | [optional] 
**ResponderUrl** | Pointer to **NullableString** |  | [optional] 
**SubjectKeyIdentifier** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**TrustedForClientAuthentication** | **bool** |  | 
**TrustedForServerAuthentication** | **bool** |  | 

## Methods

### NewCertificateAuthorityResponse

`func NewCertificateAuthorityResponse(id string, certificate CFCertificate, name string, outdatedRevocationStatusPolicy string, public bool, trustedForClientAuthentication bool, trustedForServerAuthentication bool, ) *CertificateAuthorityResponse`

NewCertificateAuthorityResponse instantiates a new CertificateAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityResponseWithDefaults

`func NewCertificateAuthorityResponseWithDefaults() *CertificateAuthorityResponse`

NewCertificateAuthorityResponseWithDefaults instantiates a new CertificateAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateAuthorityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthorityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthorityResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCertificate

`func (o *CertificateAuthorityResponse) GetCertificate() CFCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateAuthorityResponse) GetCertificateOk() (*CFCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateAuthorityResponse) SetCertificate(v CFCertificate)`

SetCertificate sets Certificate field to given value.


### GetCacheTimeToIdle

`func (o *CertificateAuthorityResponse) GetCacheTimeToIdle() string`

GetCacheTimeToIdle returns the CacheTimeToIdle field if non-nil, zero value otherwise.

### GetCacheTimeToIdleOk

`func (o *CertificateAuthorityResponse) GetCacheTimeToIdleOk() (*string, bool)`

GetCacheTimeToIdleOk returns a tuple with the CacheTimeToIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeToIdle

`func (o *CertificateAuthorityResponse) SetCacheTimeToIdle(v string)`

SetCacheTimeToIdle sets CacheTimeToIdle field to given value.

### HasCacheTimeToIdle

`func (o *CertificateAuthorityResponse) HasCacheTimeToIdle() bool`

HasCacheTimeToIdle returns a boolean if a field has been set.

### SetCacheTimeToIdleNil

`func (o *CertificateAuthorityResponse) SetCacheTimeToIdleNil(b bool)`

 SetCacheTimeToIdleNil sets the value for CacheTimeToIdle to be an explicit nil

### UnsetCacheTimeToIdle
`func (o *CertificateAuthorityResponse) UnsetCacheTimeToIdle()`

UnsetCacheTimeToIdle ensures that no value is present for CacheTimeToIdle, not even an explicit nil
### GetCrlUrl

`func (o *CertificateAuthorityResponse) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *CertificateAuthorityResponse) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *CertificateAuthorityResponse) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *CertificateAuthorityResponse) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### SetCrlUrlNil

`func (o *CertificateAuthorityResponse) SetCrlUrlNil(b bool)`

 SetCrlUrlNil sets the value for CrlUrl to be an explicit nil

### UnsetCrlUrl
`func (o *CertificateAuthorityResponse) UnsetCrlUrl()`

UnsetCrlUrl ensures that no value is present for CrlUrl, not even an explicit nil
### GetDownloadable

`func (o *CertificateAuthorityResponse) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *CertificateAuthorityResponse) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *CertificateAuthorityResponse) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *CertificateAuthorityResponse) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthorityResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthorityResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthorityResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityResponse) GetOutdatedRevocationStatusPolicy() string`

GetOutdatedRevocationStatusPolicy returns the OutdatedRevocationStatusPolicy field if non-nil, zero value otherwise.

### GetOutdatedRevocationStatusPolicyOk

`func (o *CertificateAuthorityResponse) GetOutdatedRevocationStatusPolicyOk() (*string, bool)`

GetOutdatedRevocationStatusPolicyOk returns a tuple with the OutdatedRevocationStatusPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityResponse) SetOutdatedRevocationStatusPolicy(v string)`

SetOutdatedRevocationStatusPolicy sets OutdatedRevocationStatusPolicy field to given value.


### GetProxy

`func (o *CertificateAuthorityResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateAuthorityResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateAuthorityResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateAuthorityResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateAuthorityResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateAuthorityResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetPublic

`func (o *CertificateAuthorityResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CertificateAuthorityResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CertificateAuthorityResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetRefresh

`func (o *CertificateAuthorityResponse) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CertificateAuthorityResponse) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CertificateAuthorityResponse) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CertificateAuthorityResponse) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *CertificateAuthorityResponse) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *CertificateAuthorityResponse) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetResponderUrl

`func (o *CertificateAuthorityResponse) GetResponderUrl() string`

GetResponderUrl returns the ResponderUrl field if non-nil, zero value otherwise.

### GetResponderUrlOk

`func (o *CertificateAuthorityResponse) GetResponderUrlOk() (*string, bool)`

GetResponderUrlOk returns a tuple with the ResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderUrl

`func (o *CertificateAuthorityResponse) SetResponderUrl(v string)`

SetResponderUrl sets ResponderUrl field to given value.

### HasResponderUrl

`func (o *CertificateAuthorityResponse) HasResponderUrl() bool`

HasResponderUrl returns a boolean if a field has been set.

### SetResponderUrlNil

`func (o *CertificateAuthorityResponse) SetResponderUrlNil(b bool)`

 SetResponderUrlNil sets the value for ResponderUrl to be an explicit nil

### UnsetResponderUrl
`func (o *CertificateAuthorityResponse) UnsetResponderUrl()`

UnsetResponderUrl ensures that no value is present for ResponderUrl, not even an explicit nil
### GetSubjectKeyIdentifier

`func (o *CertificateAuthorityResponse) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CertificateAuthorityResponse) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CertificateAuthorityResponse) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.

### HasSubjectKeyIdentifier

`func (o *CertificateAuthorityResponse) HasSubjectKeyIdentifier() bool`

HasSubjectKeyIdentifier returns a boolean if a field has been set.

### SetSubjectKeyIdentifierNil

`func (o *CertificateAuthorityResponse) SetSubjectKeyIdentifierNil(b bool)`

 SetSubjectKeyIdentifierNil sets the value for SubjectKeyIdentifier to be an explicit nil

### UnsetSubjectKeyIdentifier
`func (o *CertificateAuthorityResponse) UnsetSubjectKeyIdentifier()`

UnsetSubjectKeyIdentifier ensures that no value is present for SubjectKeyIdentifier, not even an explicit nil
### GetTimeout

`func (o *CertificateAuthorityResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateAuthorityResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateAuthorityResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CertificateAuthorityResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CertificateAuthorityResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CertificateAuthorityResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustedForClientAuthentication

`func (o *CertificateAuthorityResponse) GetTrustedForClientAuthentication() bool`

GetTrustedForClientAuthentication returns the TrustedForClientAuthentication field if non-nil, zero value otherwise.

### GetTrustedForClientAuthenticationOk

`func (o *CertificateAuthorityResponse) GetTrustedForClientAuthenticationOk() (*bool, bool)`

GetTrustedForClientAuthenticationOk returns a tuple with the TrustedForClientAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForClientAuthentication

`func (o *CertificateAuthorityResponse) SetTrustedForClientAuthentication(v bool)`

SetTrustedForClientAuthentication sets TrustedForClientAuthentication field to given value.


### GetTrustedForServerAuthentication

`func (o *CertificateAuthorityResponse) GetTrustedForServerAuthentication() bool`

GetTrustedForServerAuthentication returns the TrustedForServerAuthentication field if non-nil, zero value otherwise.

### GetTrustedForServerAuthenticationOk

`func (o *CertificateAuthorityResponse) GetTrustedForServerAuthenticationOk() (*bool, bool)`

GetTrustedForServerAuthenticationOk returns a tuple with the TrustedForServerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForServerAuthentication

`func (o *CertificateAuthorityResponse) SetTrustedForServerAuthentication(v bool)`

SetTrustedForServerAuthentication sets TrustedForServerAuthentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


