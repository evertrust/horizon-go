# HorizonExportItemsDatasourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** | Type of datasource | 
**Name** | **string** | Name of the datasource | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**Host** | Pointer to **NullableString** | Ip of the DNS server. If empty, Horizon Server DNS is used | [optional] 
**Port** | Pointer to **NullableInt64** | Port on which to join the LDAP server | [optional] [default to 389]
**Timeout** | **string** | Timeout for the LDAP request | 
**RecordTypes** | Pointer to **[]string** | Type of DNS records to fetch. All available record types are fetched if null | [optional] 
**Lookup** | **string** | Host to lookup | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for LDAP Authentication | 
**Method** | **string** | The HTTP method to use for the request | 
**Url** | **string** | The URL to request | 
**AuthenticationType** | **string** | The authentication type to use while making the REST call. Is linked to &#x60;credentials&#x60;. | 
**Headers** | Pointer to [**[]Header**](Header.md) | The headers of the request | [optional] 
**PayloadType** | Pointer to **NullableString** | For UI purposes in order to format the body correctly | [optional] 
**Payload** | Pointer to **NullableString** | The body of the request | [optional] 
**ExpectedHttpCodes** | **[]int64** | The success HTTP codes for the request. If the return code is not in this list, the request will be considered failed. | 
**Proxy** | Pointer to **NullableString** | Name of the proxy to use to reach the LDAP server | [optional] 
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 
**Hostname** | **string** | Hostname of the LDAP server | 
**Secure** | **bool** | Use secure LDAP connection | 
**DisableHostnameValidation** | Pointer to **NullableBool** | Disable hostname validation for the LDAP connection | [optional] [default to false]
**BaseDn** | **string** | LDAP Base DN | 
**Filter** | **string** | LDAP Filter | 

## Methods

### NewHorizonExportItemsDatasourcesInner

`func NewHorizonExportItemsDatasourcesInner(id string, type_ string, name string, timeout string, lookup string, credentials string, method string, url string, authenticationType string, expectedHttpCodes []int64, hostname string, secure bool, baseDn string, filter string, ) *HorizonExportItemsDatasourcesInner`

NewHorizonExportItemsDatasourcesInner instantiates a new HorizonExportItemsDatasourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHorizonExportItemsDatasourcesInnerWithDefaults

`func NewHorizonExportItemsDatasourcesInnerWithDefaults() *HorizonExportItemsDatasourcesInner`

NewHorizonExportItemsDatasourcesInnerWithDefaults instantiates a new HorizonExportItemsDatasourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HorizonExportItemsDatasourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HorizonExportItemsDatasourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HorizonExportItemsDatasourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *HorizonExportItemsDatasourcesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HorizonExportItemsDatasourcesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HorizonExportItemsDatasourcesInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *HorizonExportItemsDatasourcesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HorizonExportItemsDatasourcesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HorizonExportItemsDatasourcesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *HorizonExportItemsDatasourcesInner) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HorizonExportItemsDatasourcesInner) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HorizonExportItemsDatasourcesInner) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HorizonExportItemsDatasourcesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *HorizonExportItemsDatasourcesInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *HorizonExportItemsDatasourcesInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *HorizonExportItemsDatasourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HorizonExportItemsDatasourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HorizonExportItemsDatasourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HorizonExportItemsDatasourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *HorizonExportItemsDatasourcesInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HorizonExportItemsDatasourcesInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HorizonExportItemsDatasourcesInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HorizonExportItemsDatasourcesInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HorizonExportItemsDatasourcesInner) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HorizonExportItemsDatasourcesInner) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HorizonExportItemsDatasourcesInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HorizonExportItemsDatasourcesInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HorizonExportItemsDatasourcesInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *HorizonExportItemsDatasourcesInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HorizonExportItemsDatasourcesInner) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HorizonExportItemsDatasourcesInner) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTimeout

`func (o *HorizonExportItemsDatasourcesInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HorizonExportItemsDatasourcesInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HorizonExportItemsDatasourcesInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetRecordTypes

`func (o *HorizonExportItemsDatasourcesInner) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *HorizonExportItemsDatasourcesInner) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *HorizonExportItemsDatasourcesInner) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *HorizonExportItemsDatasourcesInner) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.

### SetRecordTypesNil

`func (o *HorizonExportItemsDatasourcesInner) SetRecordTypesNil(b bool)`

 SetRecordTypesNil sets the value for RecordTypes to be an explicit nil

### UnsetRecordTypes
`func (o *HorizonExportItemsDatasourcesInner) UnsetRecordTypes()`

UnsetRecordTypes ensures that no value is present for RecordTypes, not even an explicit nil
### GetLookup

`func (o *HorizonExportItemsDatasourcesInner) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *HorizonExportItemsDatasourcesInner) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *HorizonExportItemsDatasourcesInner) SetLookup(v string)`

SetLookup sets Lookup field to given value.


### GetCredentials

`func (o *HorizonExportItemsDatasourcesInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HorizonExportItemsDatasourcesInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HorizonExportItemsDatasourcesInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetMethod

`func (o *HorizonExportItemsDatasourcesInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HorizonExportItemsDatasourcesInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HorizonExportItemsDatasourcesInner) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *HorizonExportItemsDatasourcesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HorizonExportItemsDatasourcesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HorizonExportItemsDatasourcesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthenticationType

`func (o *HorizonExportItemsDatasourcesInner) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *HorizonExportItemsDatasourcesInner) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *HorizonExportItemsDatasourcesInner) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetHeaders

`func (o *HorizonExportItemsDatasourcesInner) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HorizonExportItemsDatasourcesInner) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HorizonExportItemsDatasourcesInner) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HorizonExportItemsDatasourcesInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *HorizonExportItemsDatasourcesInner) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *HorizonExportItemsDatasourcesInner) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetPayloadType

`func (o *HorizonExportItemsDatasourcesInner) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *HorizonExportItemsDatasourcesInner) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *HorizonExportItemsDatasourcesInner) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *HorizonExportItemsDatasourcesInner) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *HorizonExportItemsDatasourcesInner) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *HorizonExportItemsDatasourcesInner) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetPayload

`func (o *HorizonExportItemsDatasourcesInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *HorizonExportItemsDatasourcesInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *HorizonExportItemsDatasourcesInner) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *HorizonExportItemsDatasourcesInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *HorizonExportItemsDatasourcesInner) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *HorizonExportItemsDatasourcesInner) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetExpectedHttpCodes

`func (o *HorizonExportItemsDatasourcesInner) GetExpectedHttpCodes() []int64`

GetExpectedHttpCodes returns the ExpectedHttpCodes field if non-nil, zero value otherwise.

### GetExpectedHttpCodesOk

`func (o *HorizonExportItemsDatasourcesInner) GetExpectedHttpCodesOk() (*[]int64, bool)`

GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpCodes

`func (o *HorizonExportItemsDatasourcesInner) SetExpectedHttpCodes(v []int64)`

SetExpectedHttpCodes sets ExpectedHttpCodes field to given value.


### GetProxy

`func (o *HorizonExportItemsDatasourcesInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *HorizonExportItemsDatasourcesInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *HorizonExportItemsDatasourcesInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *HorizonExportItemsDatasourcesInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *HorizonExportItemsDatasourcesInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *HorizonExportItemsDatasourcesInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetAttributes

`func (o *HorizonExportItemsDatasourcesInner) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HorizonExportItemsDatasourcesInner) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HorizonExportItemsDatasourcesInner) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HorizonExportItemsDatasourcesInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *HorizonExportItemsDatasourcesInner) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *HorizonExportItemsDatasourcesInner) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetHostname

`func (o *HorizonExportItemsDatasourcesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HorizonExportItemsDatasourcesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HorizonExportItemsDatasourcesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSecure

`func (o *HorizonExportItemsDatasourcesInner) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *HorizonExportItemsDatasourcesInner) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *HorizonExportItemsDatasourcesInner) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetDisableHostnameValidation

`func (o *HorizonExportItemsDatasourcesInner) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *HorizonExportItemsDatasourcesInner) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *HorizonExportItemsDatasourcesInner) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *HorizonExportItemsDatasourcesInner) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *HorizonExportItemsDatasourcesInner) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *HorizonExportItemsDatasourcesInner) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetBaseDn

`func (o *HorizonExportItemsDatasourcesInner) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *HorizonExportItemsDatasourcesInner) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *HorizonExportItemsDatasourcesInner) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *HorizonExportItemsDatasourcesInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *HorizonExportItemsDatasourcesInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *HorizonExportItemsDatasourcesInner) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


