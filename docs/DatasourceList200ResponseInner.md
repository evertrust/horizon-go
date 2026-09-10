# DatasourceList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Host** | Pointer to **NullableString** | Ip of the DNS server. If empty, Horizon Server DNS is used | [optional] 
**Lookup** | **string** | Host to lookup | 
**Name** | **string** | Name of the datasource | 
**Port** | Pointer to **NullableInt64** | Port on which to join the LDAP server | [optional] [default to 389]
**RecordTypes** | Pointer to **[]string** | Type of DNS records to fetch. All available record types are fetched if null | [optional] 
**Timeout** | **string** | Timeout for the HTTP request. | 
**Type** | **string** | Type of datasource | 
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 
**BaseDn** | **string** | LDAP Base DN | 
**Credentials** | **string** | Name of the [credentials](#tag/security.credentials) to use for authentication | 
**DisableHostnameValidation** | Pointer to **NullableBool** | Disable hostname validation for the LDAP connection | [optional] [default to false]
**Filter** | **string** | LDAP Filter | 
**Hostname** | **string** | Hostname of the LDAP server | 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while making the request | [optional] 
**Secure** | **bool** | Use secure LDAP connection | 
**AuthenticationType** | **string** | The authentication type to use while making the REST call. Is linked to &#x60;credentials&#x60;. | 
**ExpectedHttpCodes** | **[]int64** | The success HTTP codes for the request. If the return code is not in this list, the request will be considered failed. | 
**Headers** | Pointer to [**[]Header**](Header.md) | The headers of the request | [optional] 
**Method** | **string** | The HTTP method to use for the request | 
**NotFoundHttpCodes** | Pointer to **[]int64** | HTTP response codes that indicate a \&quot;not found\&quot; result. Must not overlap with &#x60;expectedHttpCodes&#x60;. When received, the datasource will return a &#x60;not_found&#x60; status with no results. Combined with the &#x60;mandatory&#x60; parameter on a datasource flow entry, this can be used to stop a flow when a required resource is not found. | [optional] 
**Payload** | Pointer to **NullableString** | The body of the request | [optional] 
**PayloadType** | Pointer to **NullableString** | For UI purposes in order to format the body correctly | [optional] 
**Url** | **string** | The URL to request | 

## Methods

### NewDatasourceList200ResponseInner

`func NewDatasourceList200ResponseInner(id string, lookup string, name string, timeout string, type_ string, baseDn string, credentials string, filter string, hostname string, secure bool, authenticationType string, expectedHttpCodes []int64, method string, url string, ) *DatasourceList200ResponseInner`

NewDatasourceList200ResponseInner instantiates a new DatasourceList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceList200ResponseInnerWithDefaults

`func NewDatasourceList200ResponseInnerWithDefaults() *DatasourceList200ResponseInner`

NewDatasourceList200ResponseInnerWithDefaults instantiates a new DatasourceList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasourceList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasourceList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasourceList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *DatasourceList200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasourceList200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasourceList200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasourceList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DatasourceList200ResponseInner) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasourceList200ResponseInner) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasourceList200ResponseInner) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasourceList200ResponseInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DatasourceList200ResponseInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DatasourceList200ResponseInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHost

`func (o *DatasourceList200ResponseInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatasourceList200ResponseInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatasourceList200ResponseInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DatasourceList200ResponseInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DatasourceList200ResponseInner) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DatasourceList200ResponseInner) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLookup

`func (o *DatasourceList200ResponseInner) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *DatasourceList200ResponseInner) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *DatasourceList200ResponseInner) SetLookup(v string)`

SetLookup sets Lookup field to given value.


### GetName

`func (o *DatasourceList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *DatasourceList200ResponseInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatasourceList200ResponseInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatasourceList200ResponseInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatasourceList200ResponseInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DatasourceList200ResponseInner) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DatasourceList200ResponseInner) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRecordTypes

`func (o *DatasourceList200ResponseInner) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *DatasourceList200ResponseInner) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *DatasourceList200ResponseInner) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *DatasourceList200ResponseInner) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.

### SetRecordTypesNil

`func (o *DatasourceList200ResponseInner) SetRecordTypesNil(b bool)`

 SetRecordTypesNil sets the value for RecordTypes to be an explicit nil

### UnsetRecordTypes
`func (o *DatasourceList200ResponseInner) UnsetRecordTypes()`

UnsetRecordTypes ensures that no value is present for RecordTypes, not even an explicit nil
### GetTimeout

`func (o *DatasourceList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DatasourceList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DatasourceList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetType

`func (o *DatasourceList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasourceList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasourceList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DatasourceList200ResponseInner) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DatasourceList200ResponseInner) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DatasourceList200ResponseInner) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DatasourceList200ResponseInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DatasourceList200ResponseInner) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DatasourceList200ResponseInner) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetBaseDn

`func (o *DatasourceList200ResponseInner) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *DatasourceList200ResponseInner) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *DatasourceList200ResponseInner) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCredentials

`func (o *DatasourceList200ResponseInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DatasourceList200ResponseInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DatasourceList200ResponseInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDisableHostnameValidation

`func (o *DatasourceList200ResponseInner) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *DatasourceList200ResponseInner) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *DatasourceList200ResponseInner) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *DatasourceList200ResponseInner) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *DatasourceList200ResponseInner) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *DatasourceList200ResponseInner) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetFilter

`func (o *DatasourceList200ResponseInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DatasourceList200ResponseInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DatasourceList200ResponseInner) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetHostname

`func (o *DatasourceList200ResponseInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DatasourceList200ResponseInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DatasourceList200ResponseInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetProxy

`func (o *DatasourceList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DatasourceList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DatasourceList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DatasourceList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DatasourceList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DatasourceList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSecure

`func (o *DatasourceList200ResponseInner) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *DatasourceList200ResponseInner) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *DatasourceList200ResponseInner) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetAuthenticationType

`func (o *DatasourceList200ResponseInner) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *DatasourceList200ResponseInner) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *DatasourceList200ResponseInner) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetExpectedHttpCodes

`func (o *DatasourceList200ResponseInner) GetExpectedHttpCodes() []int64`

GetExpectedHttpCodes returns the ExpectedHttpCodes field if non-nil, zero value otherwise.

### GetExpectedHttpCodesOk

`func (o *DatasourceList200ResponseInner) GetExpectedHttpCodesOk() (*[]int64, bool)`

GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpCodes

`func (o *DatasourceList200ResponseInner) SetExpectedHttpCodes(v []int64)`

SetExpectedHttpCodes sets ExpectedHttpCodes field to given value.


### GetHeaders

`func (o *DatasourceList200ResponseInner) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DatasourceList200ResponseInner) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DatasourceList200ResponseInner) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DatasourceList200ResponseInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *DatasourceList200ResponseInner) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *DatasourceList200ResponseInner) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *DatasourceList200ResponseInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DatasourceList200ResponseInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DatasourceList200ResponseInner) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNotFoundHttpCodes

`func (o *DatasourceList200ResponseInner) GetNotFoundHttpCodes() []int64`

GetNotFoundHttpCodes returns the NotFoundHttpCodes field if non-nil, zero value otherwise.

### GetNotFoundHttpCodesOk

`func (o *DatasourceList200ResponseInner) GetNotFoundHttpCodesOk() (*[]int64, bool)`

GetNotFoundHttpCodesOk returns a tuple with the NotFoundHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundHttpCodes

`func (o *DatasourceList200ResponseInner) SetNotFoundHttpCodes(v []int64)`

SetNotFoundHttpCodes sets NotFoundHttpCodes field to given value.

### HasNotFoundHttpCodes

`func (o *DatasourceList200ResponseInner) HasNotFoundHttpCodes() bool`

HasNotFoundHttpCodes returns a boolean if a field has been set.

### GetPayload

`func (o *DatasourceList200ResponseInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DatasourceList200ResponseInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DatasourceList200ResponseInner) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DatasourceList200ResponseInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *DatasourceList200ResponseInner) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *DatasourceList200ResponseInner) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetPayloadType

`func (o *DatasourceList200ResponseInner) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *DatasourceList200ResponseInner) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *DatasourceList200ResponseInner) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *DatasourceList200ResponseInner) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *DatasourceList200ResponseInner) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *DatasourceList200ResponseInner) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetUrl

`func (o *DatasourceList200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DatasourceList200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DatasourceList200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


