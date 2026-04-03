# CertificateAuthorityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** |  | 
**CacheTimeToIdle** | Pointer to **NullableString** |  | [optional] 
**CrlUrl** | Pointer to **NullableString** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**EmailMapping** | Pointer to **string** | A template string to apply to determine the principal email when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.san.rfc822name.1}}"]
**IdentifierMapping** | Pointer to **string** | A template string to apply to determine the principal identifier when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.dn}}"]
**Name** | **string** |  | 
**NameMapping** | Pointer to **string** | A template string to apply to determine the principal name when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.subject.cn.1}}"]
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

### NewCertificateAuthorityRequest

`func NewCertificateAuthorityRequest(certificate string, name string, outdatedRevocationStatusPolicy string, public bool, trustedForClientAuthentication bool, trustedForServerAuthentication bool, ) *CertificateAuthorityRequest`

NewCertificateAuthorityRequest instantiates a new CertificateAuthorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityRequestWithDefaults

`func NewCertificateAuthorityRequestWithDefaults() *CertificateAuthorityRequest`

NewCertificateAuthorityRequestWithDefaults instantiates a new CertificateAuthorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateAuthorityRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateAuthorityRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateAuthorityRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCacheTimeToIdle

`func (o *CertificateAuthorityRequest) GetCacheTimeToIdle() string`

GetCacheTimeToIdle returns the CacheTimeToIdle field if non-nil, zero value otherwise.

### GetCacheTimeToIdleOk

`func (o *CertificateAuthorityRequest) GetCacheTimeToIdleOk() (*string, bool)`

GetCacheTimeToIdleOk returns a tuple with the CacheTimeToIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeToIdle

`func (o *CertificateAuthorityRequest) SetCacheTimeToIdle(v string)`

SetCacheTimeToIdle sets CacheTimeToIdle field to given value.

### HasCacheTimeToIdle

`func (o *CertificateAuthorityRequest) HasCacheTimeToIdle() bool`

HasCacheTimeToIdle returns a boolean if a field has been set.

### SetCacheTimeToIdleNil

`func (o *CertificateAuthorityRequest) SetCacheTimeToIdleNil(b bool)`

 SetCacheTimeToIdleNil sets the value for CacheTimeToIdle to be an explicit nil

### UnsetCacheTimeToIdle
`func (o *CertificateAuthorityRequest) UnsetCacheTimeToIdle()`

UnsetCacheTimeToIdle ensures that no value is present for CacheTimeToIdle, not even an explicit nil
### GetCrlUrl

`func (o *CertificateAuthorityRequest) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *CertificateAuthorityRequest) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *CertificateAuthorityRequest) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *CertificateAuthorityRequest) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### SetCrlUrlNil

`func (o *CertificateAuthorityRequest) SetCrlUrlNil(b bool)`

 SetCrlUrlNil sets the value for CrlUrl to be an explicit nil

### UnsetCrlUrl
`func (o *CertificateAuthorityRequest) UnsetCrlUrl()`

UnsetCrlUrl ensures that no value is present for CrlUrl, not even an explicit nil
### GetDownloadable

`func (o *CertificateAuthorityRequest) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *CertificateAuthorityRequest) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *CertificateAuthorityRequest) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *CertificateAuthorityRequest) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetEmailMapping

`func (o *CertificateAuthorityRequest) GetEmailMapping() string`

GetEmailMapping returns the EmailMapping field if non-nil, zero value otherwise.

### GetEmailMappingOk

`func (o *CertificateAuthorityRequest) GetEmailMappingOk() (*string, bool)`

GetEmailMappingOk returns a tuple with the EmailMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMapping

`func (o *CertificateAuthorityRequest) SetEmailMapping(v string)`

SetEmailMapping sets EmailMapping field to given value.

### HasEmailMapping

`func (o *CertificateAuthorityRequest) HasEmailMapping() bool`

HasEmailMapping returns a boolean if a field has been set.

### GetIdentifierMapping

`func (o *CertificateAuthorityRequest) GetIdentifierMapping() string`

GetIdentifierMapping returns the IdentifierMapping field if non-nil, zero value otherwise.

### GetIdentifierMappingOk

`func (o *CertificateAuthorityRequest) GetIdentifierMappingOk() (*string, bool)`

GetIdentifierMappingOk returns a tuple with the IdentifierMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierMapping

`func (o *CertificateAuthorityRequest) SetIdentifierMapping(v string)`

SetIdentifierMapping sets IdentifierMapping field to given value.

### HasIdentifierMapping

`func (o *CertificateAuthorityRequest) HasIdentifierMapping() bool`

HasIdentifierMapping returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthorityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthorityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthorityRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNameMapping

`func (o *CertificateAuthorityRequest) GetNameMapping() string`

GetNameMapping returns the NameMapping field if non-nil, zero value otherwise.

### GetNameMappingOk

`func (o *CertificateAuthorityRequest) GetNameMappingOk() (*string, bool)`

GetNameMappingOk returns a tuple with the NameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMapping

`func (o *CertificateAuthorityRequest) SetNameMapping(v string)`

SetNameMapping sets NameMapping field to given value.

### HasNameMapping

`func (o *CertificateAuthorityRequest) HasNameMapping() bool`

HasNameMapping returns a boolean if a field has been set.

### GetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityRequest) GetOutdatedRevocationStatusPolicy() string`

GetOutdatedRevocationStatusPolicy returns the OutdatedRevocationStatusPolicy field if non-nil, zero value otherwise.

### GetOutdatedRevocationStatusPolicyOk

`func (o *CertificateAuthorityRequest) GetOutdatedRevocationStatusPolicyOk() (*string, bool)`

GetOutdatedRevocationStatusPolicyOk returns a tuple with the OutdatedRevocationStatusPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthorityRequest) SetOutdatedRevocationStatusPolicy(v string)`

SetOutdatedRevocationStatusPolicy sets OutdatedRevocationStatusPolicy field to given value.


### GetProxy

`func (o *CertificateAuthorityRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateAuthorityRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateAuthorityRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateAuthorityRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateAuthorityRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateAuthorityRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetPublic

`func (o *CertificateAuthorityRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CertificateAuthorityRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CertificateAuthorityRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetRefresh

`func (o *CertificateAuthorityRequest) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CertificateAuthorityRequest) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CertificateAuthorityRequest) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CertificateAuthorityRequest) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *CertificateAuthorityRequest) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *CertificateAuthorityRequest) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetResponderUrl

`func (o *CertificateAuthorityRequest) GetResponderUrl() string`

GetResponderUrl returns the ResponderUrl field if non-nil, zero value otherwise.

### GetResponderUrlOk

`func (o *CertificateAuthorityRequest) GetResponderUrlOk() (*string, bool)`

GetResponderUrlOk returns a tuple with the ResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderUrl

`func (o *CertificateAuthorityRequest) SetResponderUrl(v string)`

SetResponderUrl sets ResponderUrl field to given value.

### HasResponderUrl

`func (o *CertificateAuthorityRequest) HasResponderUrl() bool`

HasResponderUrl returns a boolean if a field has been set.

### SetResponderUrlNil

`func (o *CertificateAuthorityRequest) SetResponderUrlNil(b bool)`

 SetResponderUrlNil sets the value for ResponderUrl to be an explicit nil

### UnsetResponderUrl
`func (o *CertificateAuthorityRequest) UnsetResponderUrl()`

UnsetResponderUrl ensures that no value is present for ResponderUrl, not even an explicit nil
### GetSubjectKeyIdentifier

`func (o *CertificateAuthorityRequest) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CertificateAuthorityRequest) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CertificateAuthorityRequest) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.

### HasSubjectKeyIdentifier

`func (o *CertificateAuthorityRequest) HasSubjectKeyIdentifier() bool`

HasSubjectKeyIdentifier returns a boolean if a field has been set.

### SetSubjectKeyIdentifierNil

`func (o *CertificateAuthorityRequest) SetSubjectKeyIdentifierNil(b bool)`

 SetSubjectKeyIdentifierNil sets the value for SubjectKeyIdentifier to be an explicit nil

### UnsetSubjectKeyIdentifier
`func (o *CertificateAuthorityRequest) UnsetSubjectKeyIdentifier()`

UnsetSubjectKeyIdentifier ensures that no value is present for SubjectKeyIdentifier, not even an explicit nil
### GetTimeout

`func (o *CertificateAuthorityRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateAuthorityRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateAuthorityRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CertificateAuthorityRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CertificateAuthorityRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CertificateAuthorityRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustedForClientAuthentication

`func (o *CertificateAuthorityRequest) GetTrustedForClientAuthentication() bool`

GetTrustedForClientAuthentication returns the TrustedForClientAuthentication field if non-nil, zero value otherwise.

### GetTrustedForClientAuthenticationOk

`func (o *CertificateAuthorityRequest) GetTrustedForClientAuthenticationOk() (*bool, bool)`

GetTrustedForClientAuthenticationOk returns a tuple with the TrustedForClientAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForClientAuthentication

`func (o *CertificateAuthorityRequest) SetTrustedForClientAuthentication(v bool)`

SetTrustedForClientAuthentication sets TrustedForClientAuthentication field to given value.


### GetTrustedForServerAuthentication

`func (o *CertificateAuthorityRequest) GetTrustedForServerAuthentication() bool`

GetTrustedForServerAuthentication returns the TrustedForServerAuthentication field if non-nil, zero value otherwise.

### GetTrustedForServerAuthenticationOk

`func (o *CertificateAuthorityRequest) GetTrustedForServerAuthenticationOk() (*bool, bool)`

GetTrustedForServerAuthenticationOk returns a tuple with the TrustedForServerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForServerAuthentication

`func (o *CertificateAuthorityRequest) SetTrustedForServerAuthentication(v bool)`

SetTrustedForServerAuthentication sets TrustedForServerAuthentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


