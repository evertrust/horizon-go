# CertificateAuthorityExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Certificate** | **string** |  | 
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

### NewCertificateAuthorityExportResponse

`func NewCertificateAuthorityExportResponse(id string, certificate string, name string, outdatedRevocationStatusPolicy string, public bool, trustedForClientAuthentication bool, trustedForServerAuthentication bool, ) *CertificateAuthorityExportResponse`

NewCertificateAuthorityExportResponse instantiates a new CertificateAuthorityExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityExportResponseWithDefaults

`func NewCertificateAuthorityExportResponseWithDefaults() *CertificateAuthorityExportResponse`

NewCertificateAuthorityExportResponseWithDefaults instantiates a new CertificateAuthorityExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateAuthorityExportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthorityExportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthorityExportResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCertificate

`func (o *CertificateAuthorityExportResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateAuthorityExportResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateAuthorityExportResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCacheTimeToIdle

`func (o *CertificateAuthorityExportResponse) GetCacheTimeToIdle() string`

GetCacheTimeToIdle returns the CacheTimeToIdle field if non-nil, zero value otherwise.

### GetCacheTimeToIdleOk

`func (o *CertificateAuthorityExportResponse) GetCacheTimeToIdleOk() (*string, bool)`

GetCacheTimeToIdleOk returns a tuple with the CacheTimeToIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeToIdle

`func (o *CertificateAuthorityExportResponse) SetCacheTimeToIdle(v string)`

SetCacheTimeToIdle sets CacheTimeToIdle field to given value.

### HasCacheTimeToIdle

`func (o *CertificateAuthorityExportResponse) HasCacheTimeToIdle() bool`

HasCacheTimeToIdle returns a boolean if a field has been set.

### SetCacheTimeToIdleNil

`func (o *CertificateAuthorityExportResponse) SetCacheTimeToIdleNil(b bool)`

 SetCacheTimeToIdleNil sets the value for CacheTimeToIdle to be an explicit nil

### UnsetCacheTimeToIdle
`func (o *CertificateAuthorityExportResponse) UnsetCacheTimeToIdle()`

UnsetCacheTimeToIdle ensures that no value is present for CacheTimeToIdle, not even an explicit nil
### GetCrlUrl

`func (o *CertificateAuthorityExportResponse) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *CertificateAuthorityExportResponse) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *CertificateAuthorityExportResponse) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *CertificateAuthorityExportResponse) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### SetCrlUrlNil

`func (o *CertificateAuthorityExportResponse) SetCrlUrlNil(b bool)`

 SetCrlUrlNil sets the value for CrlUrl to be an explicit nil

### UnsetCrlUrl
`func (o *CertificateAuthorityExportResponse) UnsetCrlUrl()`

UnsetCrlUrl ensures that no value is present for CrlUrl, not even an explicit nil
### GetDownloadable

`func (o *CertificateAuthorityExportResponse) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *CertificateAuthorityExportResponse) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *CertificateAuthorityExportResponse) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *CertificateAuthorityExportResponse) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthorityExportResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthorityExportResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthorityExportResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityExportResponse) GetOutdatedRevocationStatusPolicy() string`

GetOutdatedRevocationStatusPolicy returns the OutdatedRevocationStatusPolicy field if non-nil, zero value otherwise.

### GetOutdatedRevocationStatusPolicyOk

`func (o *CertificateAuthorityExportResponse) GetOutdatedRevocationStatusPolicyOk() (*string, bool)`

GetOutdatedRevocationStatusPolicyOk returns a tuple with the OutdatedRevocationStatusPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityExportResponse) SetOutdatedRevocationStatusPolicy(v string)`

SetOutdatedRevocationStatusPolicy sets OutdatedRevocationStatusPolicy field to given value.


### GetProxy

`func (o *CertificateAuthorityExportResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateAuthorityExportResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateAuthorityExportResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateAuthorityExportResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateAuthorityExportResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateAuthorityExportResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetPublic

`func (o *CertificateAuthorityExportResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CertificateAuthorityExportResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CertificateAuthorityExportResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetRefresh

`func (o *CertificateAuthorityExportResponse) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CertificateAuthorityExportResponse) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CertificateAuthorityExportResponse) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CertificateAuthorityExportResponse) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *CertificateAuthorityExportResponse) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *CertificateAuthorityExportResponse) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetResponderUrl

`func (o *CertificateAuthorityExportResponse) GetResponderUrl() string`

GetResponderUrl returns the ResponderUrl field if non-nil, zero value otherwise.

### GetResponderUrlOk

`func (o *CertificateAuthorityExportResponse) GetResponderUrlOk() (*string, bool)`

GetResponderUrlOk returns a tuple with the ResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderUrl

`func (o *CertificateAuthorityExportResponse) SetResponderUrl(v string)`

SetResponderUrl sets ResponderUrl field to given value.

### HasResponderUrl

`func (o *CertificateAuthorityExportResponse) HasResponderUrl() bool`

HasResponderUrl returns a boolean if a field has been set.

### SetResponderUrlNil

`func (o *CertificateAuthorityExportResponse) SetResponderUrlNil(b bool)`

 SetResponderUrlNil sets the value for ResponderUrl to be an explicit nil

### UnsetResponderUrl
`func (o *CertificateAuthorityExportResponse) UnsetResponderUrl()`

UnsetResponderUrl ensures that no value is present for ResponderUrl, not even an explicit nil
### GetSubjectKeyIdentifier

`func (o *CertificateAuthorityExportResponse) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CertificateAuthorityExportResponse) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CertificateAuthorityExportResponse) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.

### HasSubjectKeyIdentifier

`func (o *CertificateAuthorityExportResponse) HasSubjectKeyIdentifier() bool`

HasSubjectKeyIdentifier returns a boolean if a field has been set.

### SetSubjectKeyIdentifierNil

`func (o *CertificateAuthorityExportResponse) SetSubjectKeyIdentifierNil(b bool)`

 SetSubjectKeyIdentifierNil sets the value for SubjectKeyIdentifier to be an explicit nil

### UnsetSubjectKeyIdentifier
`func (o *CertificateAuthorityExportResponse) UnsetSubjectKeyIdentifier()`

UnsetSubjectKeyIdentifier ensures that no value is present for SubjectKeyIdentifier, not even an explicit nil
### GetTimeout

`func (o *CertificateAuthorityExportResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateAuthorityExportResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateAuthorityExportResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CertificateAuthorityExportResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CertificateAuthorityExportResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CertificateAuthorityExportResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustedForClientAuthentication

`func (o *CertificateAuthorityExportResponse) GetTrustedForClientAuthentication() bool`

GetTrustedForClientAuthentication returns the TrustedForClientAuthentication field if non-nil, zero value otherwise.

### GetTrustedForClientAuthenticationOk

`func (o *CertificateAuthorityExportResponse) GetTrustedForClientAuthenticationOk() (*bool, bool)`

GetTrustedForClientAuthenticationOk returns a tuple with the TrustedForClientAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForClientAuthentication

`func (o *CertificateAuthorityExportResponse) SetTrustedForClientAuthentication(v bool)`

SetTrustedForClientAuthentication sets TrustedForClientAuthentication field to given value.


### GetTrustedForServerAuthentication

`func (o *CertificateAuthorityExportResponse) GetTrustedForServerAuthentication() bool`

GetTrustedForServerAuthentication returns the TrustedForServerAuthentication field if non-nil, zero value otherwise.

### GetTrustedForServerAuthenticationOk

`func (o *CertificateAuthorityExportResponse) GetTrustedForServerAuthenticationOk() (*bool, bool)`

GetTrustedForServerAuthenticationOk returns a tuple with the TrustedForServerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForServerAuthentication

`func (o *CertificateAuthorityExportResponse) SetTrustedForServerAuthentication(v bool)`

SetTrustedForServerAuthentication sets TrustedForServerAuthentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


