# DataSourceTestRequestDs

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
**Credentials** | **string** | Name of the [credentials](#tag/security.credentials) to use for authentication | 
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

### NewDataSourceTestRequestDs

`func NewDataSourceTestRequestDs(type_ string, name string, timeout string, lookup string, credentials string, hostname string, secure bool, baseDn string, filter string, method string, url string, authenticationType string, expectedHttpCodes []int64, ) *DataSourceTestRequestDs`

NewDataSourceTestRequestDs instantiates a new DataSourceTestRequestDs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceTestRequestDsWithDefaults

`func NewDataSourceTestRequestDsWithDefaults() *DataSourceTestRequestDs`

NewDataSourceTestRequestDsWithDefaults instantiates a new DataSourceTestRequestDs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataSourceTestRequestDs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataSourceTestRequestDs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataSourceTestRequestDs) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *DataSourceTestRequestDs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceTestRequestDs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceTestRequestDs) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *DataSourceTestRequestDs) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DataSourceTestRequestDs) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DataSourceTestRequestDs) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DataSourceTestRequestDs) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DataSourceTestRequestDs) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DataSourceTestRequestDs) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *DataSourceTestRequestDs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataSourceTestRequestDs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataSourceTestRequestDs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataSourceTestRequestDs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DataSourceTestRequestDs) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DataSourceTestRequestDs) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DataSourceTestRequestDs) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DataSourceTestRequestDs) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DataSourceTestRequestDs) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DataSourceTestRequestDs) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DataSourceTestRequestDs) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DataSourceTestRequestDs) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DataSourceTestRequestDs) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DataSourceTestRequestDs) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DataSourceTestRequestDs) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DataSourceTestRequestDs) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTimeout

`func (o *DataSourceTestRequestDs) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DataSourceTestRequestDs) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DataSourceTestRequestDs) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetRecordTypes

`func (o *DataSourceTestRequestDs) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *DataSourceTestRequestDs) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *DataSourceTestRequestDs) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *DataSourceTestRequestDs) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.

### SetRecordTypesNil

`func (o *DataSourceTestRequestDs) SetRecordTypesNil(b bool)`

 SetRecordTypesNil sets the value for RecordTypes to be an explicit nil

### UnsetRecordTypes
`func (o *DataSourceTestRequestDs) UnsetRecordTypes()`

UnsetRecordTypes ensures that no value is present for RecordTypes, not even an explicit nil
### GetLookup

`func (o *DataSourceTestRequestDs) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *DataSourceTestRequestDs) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *DataSourceTestRequestDs) SetLookup(v string)`

SetLookup sets Lookup field to given value.


### GetCredentials

`func (o *DataSourceTestRequestDs) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DataSourceTestRequestDs) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DataSourceTestRequestDs) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *DataSourceTestRequestDs) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DataSourceTestRequestDs) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DataSourceTestRequestDs) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetProxy

`func (o *DataSourceTestRequestDs) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DataSourceTestRequestDs) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DataSourceTestRequestDs) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DataSourceTestRequestDs) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DataSourceTestRequestDs) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DataSourceTestRequestDs) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSecure

`func (o *DataSourceTestRequestDs) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *DataSourceTestRequestDs) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *DataSourceTestRequestDs) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetDisableHostnameValidation

`func (o *DataSourceTestRequestDs) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *DataSourceTestRequestDs) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *DataSourceTestRequestDs) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *DataSourceTestRequestDs) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *DataSourceTestRequestDs) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *DataSourceTestRequestDs) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetBaseDn

`func (o *DataSourceTestRequestDs) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *DataSourceTestRequestDs) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *DataSourceTestRequestDs) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *DataSourceTestRequestDs) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DataSourceTestRequestDs) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DataSourceTestRequestDs) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributes

`func (o *DataSourceTestRequestDs) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DataSourceTestRequestDs) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DataSourceTestRequestDs) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DataSourceTestRequestDs) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DataSourceTestRequestDs) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DataSourceTestRequestDs) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetMethod

`func (o *DataSourceTestRequestDs) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DataSourceTestRequestDs) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DataSourceTestRequestDs) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *DataSourceTestRequestDs) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DataSourceTestRequestDs) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DataSourceTestRequestDs) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthenticationType

`func (o *DataSourceTestRequestDs) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *DataSourceTestRequestDs) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *DataSourceTestRequestDs) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetHeaders

`func (o *DataSourceTestRequestDs) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DataSourceTestRequestDs) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DataSourceTestRequestDs) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DataSourceTestRequestDs) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *DataSourceTestRequestDs) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *DataSourceTestRequestDs) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetPayloadType

`func (o *DataSourceTestRequestDs) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *DataSourceTestRequestDs) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *DataSourceTestRequestDs) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *DataSourceTestRequestDs) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *DataSourceTestRequestDs) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *DataSourceTestRequestDs) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetPayload

`func (o *DataSourceTestRequestDs) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DataSourceTestRequestDs) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DataSourceTestRequestDs) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DataSourceTestRequestDs) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *DataSourceTestRequestDs) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *DataSourceTestRequestDs) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetExpectedHttpCodes

`func (o *DataSourceTestRequestDs) GetExpectedHttpCodes() []int64`

GetExpectedHttpCodes returns the ExpectedHttpCodes field if non-nil, zero value otherwise.

### GetExpectedHttpCodesOk

`func (o *DataSourceTestRequestDs) GetExpectedHttpCodesOk() (*[]int64, bool)`

GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpCodes

`func (o *DataSourceTestRequestDs) SetExpectedHttpCodes(v []int64)`

SetExpectedHttpCodes sets ExpectedHttpCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


