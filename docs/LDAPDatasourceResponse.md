# LDAPDatasourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** | Type of datasource | 
**Name** | **string** | Name of the datasource | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for LDAP Authentication | 
**Hostname** | **string** | Hostname of the LDAP server | 
**Port** | Pointer to **NullableInt64** | Port on which to join the LDAP server | [optional] [default to 389]
**Proxy** | Pointer to **NullableString** | Name of the proxy to use to reach the LDAP server | [optional] 
**Timeout** | **string** | Timeout for the LDAP request | 
**Secure** | **bool** | Use secure LDAP connection | 
**DisableHostnameValidation** | Pointer to **NullableBool** | Disable hostname validation for the LDAP connection | [optional] [default to false]
**BaseDn** | **string** | LDAP Base DN | 
**Filter** | **string** | LDAP Filter | 
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 

## Methods

### NewLDAPDatasourceResponse

`func NewLDAPDatasourceResponse(id string, type_ string, name string, credentials string, hostname string, timeout string, secure bool, baseDn string, filter string, ) *LDAPDatasourceResponse`

NewLDAPDatasourceResponse instantiates a new LDAPDatasourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDatasourceResponseWithDefaults

`func NewLDAPDatasourceResponseWithDefaults() *LDAPDatasourceResponse`

NewLDAPDatasourceResponseWithDefaults instantiates a new LDAPDatasourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LDAPDatasourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LDAPDatasourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LDAPDatasourceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LDAPDatasourceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LDAPDatasourceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LDAPDatasourceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *LDAPDatasourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPDatasourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPDatasourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *LDAPDatasourceResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LDAPDatasourceResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LDAPDatasourceResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LDAPDatasourceResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LDAPDatasourceResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LDAPDatasourceResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *LDAPDatasourceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LDAPDatasourceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LDAPDatasourceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LDAPDatasourceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCredentials

`func (o *LDAPDatasourceResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *LDAPDatasourceResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *LDAPDatasourceResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *LDAPDatasourceResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LDAPDatasourceResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LDAPDatasourceResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *LDAPDatasourceResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPDatasourceResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPDatasourceResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPDatasourceResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LDAPDatasourceResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LDAPDatasourceResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProxy

`func (o *LDAPDatasourceResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *LDAPDatasourceResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *LDAPDatasourceResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *LDAPDatasourceResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *LDAPDatasourceResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *LDAPDatasourceResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *LDAPDatasourceResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LDAPDatasourceResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LDAPDatasourceResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetSecure

`func (o *LDAPDatasourceResponse) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *LDAPDatasourceResponse) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *LDAPDatasourceResponse) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetDisableHostnameValidation

`func (o *LDAPDatasourceResponse) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *LDAPDatasourceResponse) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *LDAPDatasourceResponse) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *LDAPDatasourceResponse) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *LDAPDatasourceResponse) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *LDAPDatasourceResponse) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetBaseDn

`func (o *LDAPDatasourceResponse) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPDatasourceResponse) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPDatasourceResponse) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *LDAPDatasourceResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LDAPDatasourceResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LDAPDatasourceResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributes

`func (o *LDAPDatasourceResponse) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LDAPDatasourceResponse) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LDAPDatasourceResponse) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LDAPDatasourceResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *LDAPDatasourceResponse) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *LDAPDatasourceResponse) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


