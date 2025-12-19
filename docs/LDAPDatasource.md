# LDAPDatasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 
**BaseDn** | **string** | LDAP Base DN | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for LDAP Authentication | 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**DisableHostnameValidation** | Pointer to **NullableBool** | Disable hostname validation for the LDAP connection | [optional] [default to false]
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Filter** | **string** | LDAP Filter | 
**Hostname** | **string** | Hostname of the LDAP server | 
**Name** | **string** | Name of the datasource | 
**Port** | Pointer to **NullableInt64** | Port on which to join the LDAP server | [optional] [default to 389]
**Proxy** | Pointer to **NullableString** | Name of the proxy to use to reach the LDAP server | [optional] 
**Secure** | **bool** | Use secure LDAP connection | 
**Timeout** | **string** | Timeout for the LDAP request | 
**Type** | **string** | Type of datasource | 

## Methods

### NewLDAPDatasource

`func NewLDAPDatasource(baseDn string, credentials string, filter string, hostname string, name string, secure bool, timeout string, type_ string, ) *LDAPDatasource`

NewLDAPDatasource instantiates a new LDAPDatasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDatasourceWithDefaults

`func NewLDAPDatasourceWithDefaults() *LDAPDatasource`

NewLDAPDatasourceWithDefaults instantiates a new LDAPDatasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LDAPDatasource) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LDAPDatasource) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LDAPDatasource) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LDAPDatasource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *LDAPDatasource) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *LDAPDatasource) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetBaseDn

`func (o *LDAPDatasource) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPDatasource) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPDatasource) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCredentials

`func (o *LDAPDatasource) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *LDAPDatasource) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *LDAPDatasource) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDescription

`func (o *LDAPDatasource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LDAPDatasource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LDAPDatasource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LDAPDatasource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableHostnameValidation

`func (o *LDAPDatasource) GetDisableHostnameValidation() bool`

GetDisableHostnameValidation returns the DisableHostnameValidation field if non-nil, zero value otherwise.

### GetDisableHostnameValidationOk

`func (o *LDAPDatasource) GetDisableHostnameValidationOk() (*bool, bool)`

GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameValidation

`func (o *LDAPDatasource) SetDisableHostnameValidation(v bool)`

SetDisableHostnameValidation sets DisableHostnameValidation field to given value.

### HasDisableHostnameValidation

`func (o *LDAPDatasource) HasDisableHostnameValidation() bool`

HasDisableHostnameValidation returns a boolean if a field has been set.

### SetDisableHostnameValidationNil

`func (o *LDAPDatasource) SetDisableHostnameValidationNil(b bool)`

 SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil

### UnsetDisableHostnameValidation
`func (o *LDAPDatasource) UnsetDisableHostnameValidation()`

UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
### GetDisplayName

`func (o *LDAPDatasource) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LDAPDatasource) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LDAPDatasource) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LDAPDatasource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LDAPDatasource) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LDAPDatasource) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFilter

`func (o *LDAPDatasource) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LDAPDatasource) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LDAPDatasource) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetHostname

`func (o *LDAPDatasource) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LDAPDatasource) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LDAPDatasource) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetName

`func (o *LDAPDatasource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPDatasource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPDatasource) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *LDAPDatasource) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPDatasource) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPDatasource) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPDatasource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LDAPDatasource) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LDAPDatasource) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProxy

`func (o *LDAPDatasource) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *LDAPDatasource) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *LDAPDatasource) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *LDAPDatasource) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *LDAPDatasource) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *LDAPDatasource) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSecure

`func (o *LDAPDatasource) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *LDAPDatasource) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *LDAPDatasource) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetTimeout

`func (o *LDAPDatasource) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LDAPDatasource) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LDAPDatasource) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetType

`func (o *LDAPDatasource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LDAPDatasource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LDAPDatasource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


