# AutomationPolicyLifecycleGet201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeUrl** | Pointer to **string** | The ACME URL for the external ACME server. | [optional] 
**AllowedAuthorizationMethod** | Pointer to **[]string** | The allowed authorization methods for ACME external. | [optional] 
**KeyType** | Pointer to **string** | The key type used for WebRA. | [optional] 
**Module** | **string** | The module of the initialization parameters. | 
**Profile** | **string** | The profile used for WebRA. | 
**RequireEAB** | Pointer to **bool** | Indicates whether EAB is required for ACME external. | [optional] 
**Http01Port** | Pointer to **int64** | The HTTP-01 port for ACME. | [optional] 
**TlsAlpn01Port** | Pointer to **int64** | The TLS-ALPN-01 port for ACME. | [optional] 
**AuthorizationMode** | Pointer to **string** | The authorization mode for WebRA. | [optional] 
**CsrInfoIgnored** | Pointer to **bool** | Indicates whether CSR info is ignored for SCEP. | [optional] 
**EnrollmentMode** | Pointer to **string** | The enrollment mode for WebRA. | [optional] 
**PasswordPolicy** | Pointer to [**PasswordPolicy**](PasswordPolicy.md) | The password policy for WebRA. | [optional] 

## Methods

### NewAutomationPolicyLifecycleGet201Response

`func NewAutomationPolicyLifecycleGet201Response(module string, profile string, ) *AutomationPolicyLifecycleGet201Response`

NewAutomationPolicyLifecycleGet201Response instantiates a new AutomationPolicyLifecycleGet201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationPolicyLifecycleGet201ResponseWithDefaults

`func NewAutomationPolicyLifecycleGet201ResponseWithDefaults() *AutomationPolicyLifecycleGet201Response`

NewAutomationPolicyLifecycleGet201ResponseWithDefaults instantiates a new AutomationPolicyLifecycleGet201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeUrl

`func (o *AutomationPolicyLifecycleGet201Response) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *AutomationPolicyLifecycleGet201Response) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *AutomationPolicyLifecycleGet201Response) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *AutomationPolicyLifecycleGet201Response) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetAllowedAuthorizationMethod

`func (o *AutomationPolicyLifecycleGet201Response) GetAllowedAuthorizationMethod() []string`

GetAllowedAuthorizationMethod returns the AllowedAuthorizationMethod field if non-nil, zero value otherwise.

### GetAllowedAuthorizationMethodOk

`func (o *AutomationPolicyLifecycleGet201Response) GetAllowedAuthorizationMethodOk() (*[]string, bool)`

GetAllowedAuthorizationMethodOk returns a tuple with the AllowedAuthorizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthorizationMethod

`func (o *AutomationPolicyLifecycleGet201Response) SetAllowedAuthorizationMethod(v []string)`

SetAllowedAuthorizationMethod sets AllowedAuthorizationMethod field to given value.

### HasAllowedAuthorizationMethod

`func (o *AutomationPolicyLifecycleGet201Response) HasAllowedAuthorizationMethod() bool`

HasAllowedAuthorizationMethod returns a boolean if a field has been set.

### GetKeyType

`func (o *AutomationPolicyLifecycleGet201Response) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AutomationPolicyLifecycleGet201Response) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AutomationPolicyLifecycleGet201Response) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *AutomationPolicyLifecycleGet201Response) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetModule

`func (o *AutomationPolicyLifecycleGet201Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AutomationPolicyLifecycleGet201Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AutomationPolicyLifecycleGet201Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *AutomationPolicyLifecycleGet201Response) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AutomationPolicyLifecycleGet201Response) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AutomationPolicyLifecycleGet201Response) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRequireEAB

`func (o *AutomationPolicyLifecycleGet201Response) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *AutomationPolicyLifecycleGet201Response) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *AutomationPolicyLifecycleGet201Response) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.

### HasRequireEAB

`func (o *AutomationPolicyLifecycleGet201Response) HasRequireEAB() bool`

HasRequireEAB returns a boolean if a field has been set.

### GetHttp01Port

`func (o *AutomationPolicyLifecycleGet201Response) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *AutomationPolicyLifecycleGet201Response) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *AutomationPolicyLifecycleGet201Response) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *AutomationPolicyLifecycleGet201Response) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### GetTlsAlpn01Port

`func (o *AutomationPolicyLifecycleGet201Response) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *AutomationPolicyLifecycleGet201Response) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *AutomationPolicyLifecycleGet201Response) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *AutomationPolicyLifecycleGet201Response) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### GetAuthorizationMode

`func (o *AutomationPolicyLifecycleGet201Response) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *AutomationPolicyLifecycleGet201Response) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *AutomationPolicyLifecycleGet201Response) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.

### HasAuthorizationMode

`func (o *AutomationPolicyLifecycleGet201Response) HasAuthorizationMode() bool`

HasAuthorizationMode returns a boolean if a field has been set.

### GetCsrInfoIgnored

`func (o *AutomationPolicyLifecycleGet201Response) GetCsrInfoIgnored() bool`

GetCsrInfoIgnored returns the CsrInfoIgnored field if non-nil, zero value otherwise.

### GetCsrInfoIgnoredOk

`func (o *AutomationPolicyLifecycleGet201Response) GetCsrInfoIgnoredOk() (*bool, bool)`

GetCsrInfoIgnoredOk returns a tuple with the CsrInfoIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrInfoIgnored

`func (o *AutomationPolicyLifecycleGet201Response) SetCsrInfoIgnored(v bool)`

SetCsrInfoIgnored sets CsrInfoIgnored field to given value.

### HasCsrInfoIgnored

`func (o *AutomationPolicyLifecycleGet201Response) HasCsrInfoIgnored() bool`

HasCsrInfoIgnored returns a boolean if a field has been set.

### GetEnrollmentMode

`func (o *AutomationPolicyLifecycleGet201Response) GetEnrollmentMode() string`

GetEnrollmentMode returns the EnrollmentMode field if non-nil, zero value otherwise.

### GetEnrollmentModeOk

`func (o *AutomationPolicyLifecycleGet201Response) GetEnrollmentModeOk() (*string, bool)`

GetEnrollmentModeOk returns a tuple with the EnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMode

`func (o *AutomationPolicyLifecycleGet201Response) SetEnrollmentMode(v string)`

SetEnrollmentMode sets EnrollmentMode field to given value.

### HasEnrollmentMode

`func (o *AutomationPolicyLifecycleGet201Response) HasEnrollmentMode() bool`

HasEnrollmentMode returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *AutomationPolicyLifecycleGet201Response) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *AutomationPolicyLifecycleGet201Response) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *AutomationPolicyLifecycleGet201Response) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *AutomationPolicyLifecycleGet201Response) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


