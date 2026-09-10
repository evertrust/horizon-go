# CertificateAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTimeToIdle** | Pointer to **NullableString** |  | [optional] 
**CrlUrl** | Pointer to **NullableString** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**EmailMapping** | Pointer to **string** | A template string to apply to determine the principal email when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.san.rfc822name.1}}"]
**IdentifierMapping** | Pointer to **string** | A template string to apply to determine the principal identifier when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.dn}}"]
**Name** | Pointer to **string** |  | [optional] 
**NameMapping** | Pointer to **string** | A template string to apply to determine the principal name when a certificate from this CA is used for client authentication | [optional] [default to "{{certificate.subject.cn.1}}"]
**OutdatedRevocationStatusPolicy** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Refresh** | Pointer to **NullableString** |  | [optional] 
**ResponderUrl** | Pointer to **NullableString** |  | [optional] 
**SubjectKeyIdentifier** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**TrustedForClientAuthentication** | Pointer to **bool** |  | [optional] 
**TrustedForServerAuthentication** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateAuthority

`func NewCertificateAuthority() *CertificateAuthority`

NewCertificateAuthority instantiates a new CertificateAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityWithDefaults

`func NewCertificateAuthorityWithDefaults() *CertificateAuthority`

NewCertificateAuthorityWithDefaults instantiates a new CertificateAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheTimeToIdle

`func (o *CertificateAuthority) GetCacheTimeToIdle() string`

GetCacheTimeToIdle returns the CacheTimeToIdle field if non-nil, zero value otherwise.

### GetCacheTimeToIdleOk

`func (o *CertificateAuthority) GetCacheTimeToIdleOk() (*string, bool)`

GetCacheTimeToIdleOk returns a tuple with the CacheTimeToIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeToIdle

`func (o *CertificateAuthority) SetCacheTimeToIdle(v string)`

SetCacheTimeToIdle sets CacheTimeToIdle field to given value.

### HasCacheTimeToIdle

`func (o *CertificateAuthority) HasCacheTimeToIdle() bool`

HasCacheTimeToIdle returns a boolean if a field has been set.

### SetCacheTimeToIdleNil

`func (o *CertificateAuthority) SetCacheTimeToIdleNil(b bool)`

 SetCacheTimeToIdleNil sets the value for CacheTimeToIdle to be an explicit nil

### UnsetCacheTimeToIdle
`func (o *CertificateAuthority) UnsetCacheTimeToIdle()`

UnsetCacheTimeToIdle ensures that no value is present for CacheTimeToIdle, not even an explicit nil
### GetCrlUrl

`func (o *CertificateAuthority) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *CertificateAuthority) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *CertificateAuthority) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *CertificateAuthority) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### SetCrlUrlNil

`func (o *CertificateAuthority) SetCrlUrlNil(b bool)`

 SetCrlUrlNil sets the value for CrlUrl to be an explicit nil

### UnsetCrlUrl
`func (o *CertificateAuthority) UnsetCrlUrl()`

UnsetCrlUrl ensures that no value is present for CrlUrl, not even an explicit nil
### GetDownloadable

`func (o *CertificateAuthority) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *CertificateAuthority) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *CertificateAuthority) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *CertificateAuthority) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetEmailMapping

`func (o *CertificateAuthority) GetEmailMapping() string`

GetEmailMapping returns the EmailMapping field if non-nil, zero value otherwise.

### GetEmailMappingOk

`func (o *CertificateAuthority) GetEmailMappingOk() (*string, bool)`

GetEmailMappingOk returns a tuple with the EmailMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMapping

`func (o *CertificateAuthority) SetEmailMapping(v string)`

SetEmailMapping sets EmailMapping field to given value.

### HasEmailMapping

`func (o *CertificateAuthority) HasEmailMapping() bool`

HasEmailMapping returns a boolean if a field has been set.

### GetIdentifierMapping

`func (o *CertificateAuthority) GetIdentifierMapping() string`

GetIdentifierMapping returns the IdentifierMapping field if non-nil, zero value otherwise.

### GetIdentifierMappingOk

`func (o *CertificateAuthority) GetIdentifierMappingOk() (*string, bool)`

GetIdentifierMappingOk returns a tuple with the IdentifierMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierMapping

`func (o *CertificateAuthority) SetIdentifierMapping(v string)`

SetIdentifierMapping sets IdentifierMapping field to given value.

### HasIdentifierMapping

`func (o *CertificateAuthority) HasIdentifierMapping() bool`

HasIdentifierMapping returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthority) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateAuthority) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameMapping

`func (o *CertificateAuthority) GetNameMapping() string`

GetNameMapping returns the NameMapping field if non-nil, zero value otherwise.

### GetNameMappingOk

`func (o *CertificateAuthority) GetNameMappingOk() (*string, bool)`

GetNameMappingOk returns a tuple with the NameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMapping

`func (o *CertificateAuthority) SetNameMapping(v string)`

SetNameMapping sets NameMapping field to given value.

### HasNameMapping

`func (o *CertificateAuthority) HasNameMapping() bool`

HasNameMapping returns a boolean if a field has been set.

### GetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthority) GetOutdatedRevocationStatusPolicy() string`

GetOutdatedRevocationStatusPolicy returns the OutdatedRevocationStatusPolicy field if non-nil, zero value otherwise.

### GetOutdatedRevocationStatusPolicyOk

`func (o *CertificateAuthority) GetOutdatedRevocationStatusPolicyOk() (*string, bool)`

GetOutdatedRevocationStatusPolicyOk returns a tuple with the OutdatedRevocationStatusPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedRevocationStatusPolicy

`func (o *CertificateAuthority) SetOutdatedRevocationStatusPolicy(v string)`

SetOutdatedRevocationStatusPolicy sets OutdatedRevocationStatusPolicy field to given value.

### HasOutdatedRevocationStatusPolicy

`func (o *CertificateAuthority) HasOutdatedRevocationStatusPolicy() bool`

HasOutdatedRevocationStatusPolicy returns a boolean if a field has been set.

### GetProxy

`func (o *CertificateAuthority) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateAuthority) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateAuthority) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateAuthority) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateAuthority) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateAuthority) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetPublic

`func (o *CertificateAuthority) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CertificateAuthority) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CertificateAuthority) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CertificateAuthority) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRefresh

`func (o *CertificateAuthority) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CertificateAuthority) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CertificateAuthority) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CertificateAuthority) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *CertificateAuthority) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *CertificateAuthority) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetResponderUrl

`func (o *CertificateAuthority) GetResponderUrl() string`

GetResponderUrl returns the ResponderUrl field if non-nil, zero value otherwise.

### GetResponderUrlOk

`func (o *CertificateAuthority) GetResponderUrlOk() (*string, bool)`

GetResponderUrlOk returns a tuple with the ResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderUrl

`func (o *CertificateAuthority) SetResponderUrl(v string)`

SetResponderUrl sets ResponderUrl field to given value.

### HasResponderUrl

`func (o *CertificateAuthority) HasResponderUrl() bool`

HasResponderUrl returns a boolean if a field has been set.

### SetResponderUrlNil

`func (o *CertificateAuthority) SetResponderUrlNil(b bool)`

 SetResponderUrlNil sets the value for ResponderUrl to be an explicit nil

### UnsetResponderUrl
`func (o *CertificateAuthority) UnsetResponderUrl()`

UnsetResponderUrl ensures that no value is present for ResponderUrl, not even an explicit nil
### GetSubjectKeyIdentifier

`func (o *CertificateAuthority) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CertificateAuthority) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CertificateAuthority) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.

### HasSubjectKeyIdentifier

`func (o *CertificateAuthority) HasSubjectKeyIdentifier() bool`

HasSubjectKeyIdentifier returns a boolean if a field has been set.

### SetSubjectKeyIdentifierNil

`func (o *CertificateAuthority) SetSubjectKeyIdentifierNil(b bool)`

 SetSubjectKeyIdentifierNil sets the value for SubjectKeyIdentifier to be an explicit nil

### UnsetSubjectKeyIdentifier
`func (o *CertificateAuthority) UnsetSubjectKeyIdentifier()`

UnsetSubjectKeyIdentifier ensures that no value is present for SubjectKeyIdentifier, not even an explicit nil
### GetTimeout

`func (o *CertificateAuthority) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateAuthority) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateAuthority) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CertificateAuthority) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CertificateAuthority) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CertificateAuthority) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustedForClientAuthentication

`func (o *CertificateAuthority) GetTrustedForClientAuthentication() bool`

GetTrustedForClientAuthentication returns the TrustedForClientAuthentication field if non-nil, zero value otherwise.

### GetTrustedForClientAuthenticationOk

`func (o *CertificateAuthority) GetTrustedForClientAuthenticationOk() (*bool, bool)`

GetTrustedForClientAuthenticationOk returns a tuple with the TrustedForClientAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForClientAuthentication

`func (o *CertificateAuthority) SetTrustedForClientAuthentication(v bool)`

SetTrustedForClientAuthentication sets TrustedForClientAuthentication field to given value.

### HasTrustedForClientAuthentication

`func (o *CertificateAuthority) HasTrustedForClientAuthentication() bool`

HasTrustedForClientAuthentication returns a boolean if a field has been set.

### GetTrustedForServerAuthentication

`func (o *CertificateAuthority) GetTrustedForServerAuthentication() bool`

GetTrustedForServerAuthentication returns the TrustedForServerAuthentication field if non-nil, zero value otherwise.

### GetTrustedForServerAuthenticationOk

`func (o *CertificateAuthority) GetTrustedForServerAuthenticationOk() (*bool, bool)`

GetTrustedForServerAuthenticationOk returns a tuple with the TrustedForServerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForServerAuthentication

`func (o *CertificateAuthority) SetTrustedForServerAuthentication(v bool)`

SetTrustedForServerAuthentication sets TrustedForServerAuthentication field to given value.

### HasTrustedForServerAuthentication

`func (o *CertificateAuthority) HasTrustedForServerAuthentication() bool`

HasTrustedForServerAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


