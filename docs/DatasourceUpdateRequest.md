# DatasourceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of datasource | 
**Name** | **string** | Name of the datasource | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**Host** | Pointer to **NullableString** | Ip of the DNS server. If empty, Horizon Server DNS is used | [optional] 
**Port** | Pointer to **NullableInt64** | Port on which to join the LDAP server | [optional] [default to 389]
**Timeout** | **string** | Timeout for the HTTP request. | 
**RecordTypes** | Pointer to **[]string** | Type of DNS records to fetch. All available record types are fetched if null | [optional] 
**Lookup** | **string** | Host to lookup | 
**Credentials** | **string** | Name of the [credentials](#tag/api.security.credentials) to use for authentication | 
**Hostname** | **string** | Hostname of the LDAP server | 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while making the request | [optional] 
**Secure** | **bool** | Use secure LDAP connection | 
**DisableHostnameValidation** | Pointer to **NullableBool** | Disable hostname validation for the LDAP connection | [optional] [default to false]
**BaseDn** | **string** | LDAP Base DN | 
**Filter** | **string** | LDAP Filter | 
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 
**Method** | **string** | The HTTP method to use for the request | 
**Url** | **string** | The URL to request | 
**AuthenticationType** | **string** | The authentication type to use while making the REST call. Is linked to &#x60;credentials&#x60;. | 
**Headers** | Pointer to [**[]Header**](Header.md) | The headers of the request | [optional] 
**PayloadType** | Pointer to **NullableString** | For UI purposes in order to format the body correctly | [optional] 
**Payload** | Pointer to **NullableString** | The body of the request | [optional] 
**ExpectedHttpCodes** | **[]int64** | The success HTTP codes for the request. If the return code is not in this list, the request will be considered failed. | 

## Methods

### NewDatasourceUpdateRequest

`func NewDatasourceUpdateRequest(type_ string, name string, timeout string, lookup string, credentials string, hostname string, secure bool, baseDn string, filter string, method string, url string, authenticationType string, expectedHttpCodes []int64, ) *DatasourceUpdateRequest`

NewDatasourceUpdateRequest instantiates a new DatasourceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceUpdateRequestWithDefaults

`func NewDatasourceUpdateRequestWithDefaults() *DatasourceUpdateRequest`

NewDatasourceUpdateRequestWithDefaults instantiates a new DatasourceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DatasourceUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasourceUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasourceUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *DatasourceUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *DatasourceUpdateRequest) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasourceUpdateRequest) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasourceUpdateRequest) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasourceUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DatasourceUpdateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DatasourceUpdateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *DatasourceUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasourceUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasourceUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasourceUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DatasourceUpdateRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatasourceUpdateRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatasourceUpdateRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DatasourceUpdateRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DatasourceUpdateRequest) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DatasourceUpdateRequest) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DatasourceUpdateRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatasourceUpdateRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatasourceUpdateRequest) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatasourceUpdateRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DatasourceUpdateRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DatasourceUpdateRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTimeout

`func (o *DatasourceUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DatasourceUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DatasourceUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetRecordTypes

`func (o *DatasourceUpdateRequest) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *DatasourceUpdateRequest) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *DatasourceUpdateRequest) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *DatasourceUpdateRequest) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.

### SetRecordTypesNil

`func (o *DatasourceUpdateRequest) SetRecordTypesNil(b bool)`

 SetRecordTypesNil sets the value for RecordTypes to be an explicit nil

### UnsetRecordTypes
`func (o *DatasourceUpdateRequest) UnsetRecordTypes()`

UnsetRecordTypes ensures that no value is present for RecordTypes, not even an explicit nil
### GetLookup

`func (o *DatasourceUpdateRequest) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *DatasourceUpdateRequest) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *DatasourceUpdateRequest) SetLookup(v string)`

SetLookup sets Lookup field to given value.


### GetCredentials

`func (o *DatasourceUpdateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DatasourceUpdateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DatasourceUpdateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *DatasourceUpdateRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DatasourceUpdateRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DatasourceUpdateRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetProxy

`func (o *DatasourceUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DatasourceUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DatasourceUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DatasourceUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DatasourceUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DatasourceUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSecure

`func (o *DatasourceUpdateRequest) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *DatasourceUpdateRequest) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *DatasourceUpdateRequest) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetDisableHostnameValidation

`func (o *DatasourceUpdateRequest) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *DatasourceUpdateRequest) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *DatasourceUpdateRequest) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *DatasourceUpdateRequest) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *DatasourceUpdateRequest) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *DatasourceUpdateRequest) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetBaseDn

`func (o *DatasourceUpdateRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *DatasourceUpdateRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *DatasourceUpdateRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *DatasourceUpdateRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DatasourceUpdateRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DatasourceUpdateRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributes

`func (o *DatasourceUpdateRequest) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DatasourceUpdateRequest) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DatasourceUpdateRequest) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DatasourceUpdateRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DatasourceUpdateRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DatasourceUpdateRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetMethod

`func (o *DatasourceUpdateRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DatasourceUpdateRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DatasourceUpdateRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *DatasourceUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DatasourceUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DatasourceUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthenticationType

`func (o *DatasourceUpdateRequest) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *DatasourceUpdateRequest) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *DatasourceUpdateRequest) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetHeaders

`func (o *DatasourceUpdateRequest) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DatasourceUpdateRequest) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DatasourceUpdateRequest) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DatasourceUpdateRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *DatasourceUpdateRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *DatasourceUpdateRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetPayloadType

`func (o *DatasourceUpdateRequest) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *DatasourceUpdateRequest) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *DatasourceUpdateRequest) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *DatasourceUpdateRequest) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *DatasourceUpdateRequest) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *DatasourceUpdateRequest) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetPayload

`func (o *DatasourceUpdateRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DatasourceUpdateRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DatasourceUpdateRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DatasourceUpdateRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *DatasourceUpdateRequest) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *DatasourceUpdateRequest) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetExpectedHttpCodes

`func (o *DatasourceUpdateRequest) GetExpectedHttpCodes() []int64`

GetExpectedHttpCodes returns the ExpectedHttpCodes field if non-nil, zero value otherwise.

### GetExpectedHttpCodesOk

`func (o *DatasourceUpdateRequest) GetExpectedHttpCodesOk() (*[]int64, bool)`

GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpCodes

`func (o *DatasourceUpdateRequest) SetExpectedHttpCodes(v []int64)`

SetExpectedHttpCodes sets ExpectedHttpCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


