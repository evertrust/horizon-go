# PKIConnectors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Ca** | **string** | Stream&#39;s technical name of the CA on which to enroll | 
**EndPoint** | **string** | Swiss base endpoint | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Template** | **string** | Stream&#39;s certificate template to use for enrollment | 
**Timeout** | **NullableString** |  | 
**Type** | **string** |  | 
**AccountEmail** | Pointer to **NullableString** | Email to associate with the account | [optional] 
**AccountKeyType** | **string** | The key type to use to generate the account key | 
**DnsChallengeProvider** | [**DnsChallengeProviders**](DnsChallengeProviders.md) | DNS Provider configuration to provision the DNS challenge | 
**DomainDictionaryProvider** | Pointer to [**NullableDomainDictionaryProviders**](DomainDictionaryProviders.md) | The dictionary provider | [optional] 
**Eab** | Pointer to **NullableString** | &#x60;password&#x60; credentials name to use for External Account Binding | [optional] 
**EabMacAlgorithm** | Pointer to **NullableString** | The MAC algorithm to use for External Account Binding. Can only be set when &#x60;eab&#x60; is defined | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**RotateAccount** | Pointer to **NullableBool** | If enabled, regenerate the account (does not need to be specified on creation) | [optional] 
**AcmeDirectoryUrl** | **string** | The directory url of the ACME endpoint | 
**CaConfig** | **string** |  | 
**Domain** | **string** |  | 
**EnrollmentCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to enroll on the PKI | 
**Profile** | **string** |  | 
**AccessCredentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used. | [optional] 
**CaArn** | **string** |  | 
**CaPolicyOid** | Pointer to **NullableString** |  | [optional] 
**CertificateUsage** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**SigningHash** | Pointer to **NullableString** |  | [optional] 
**TemplateArn** | Pointer to **NullableString** |  | [optional] 
**ValidDays** | Pointer to **NullableString** |  | [optional] 
**OfferId** | **string** |  | 
**OrganizationId** | **int64** |  | 
**RevReason** | Pointer to **NullableString** |  | [optional] 
**CnMap** | Pointer to **NullableString** |  | [optional] 
**EmailMap** | Pointer to **NullableString** |  | [optional] 
**IssuerCACert** | **string** |  | 
**IssuerCADN** | **string** |  | 
**IssuerMap** | Pointer to **NullableString** |  | [optional] 
**LegacyCMPStyle** | Pointer to **NullableBool** |  | [optional] 
**ProfileMap** | Pointer to **NullableString** |  | [optional] 
**SanDnsMap** | Pointer to **NullableString** |  | [optional] 
**SignerCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to sign on the PKI | 
**ApiCredentials** | **string** | Name of the &#x60;api-key&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**BaseUrl** | **string** | The base URL of the used digicert instance. | 
**CaCertId** | Pointer to **NullableString** |  | [optional] 
**CustomConnectorDataMapping** | Pointer to **map[string]string** | Custom mapping of connector data fields | [optional] 
**ProductId** | **string** |  | 
**SkipApproval** | Pointer to **NullableBool** |  | [optional] 
**CaName** | **string** |  | 
**EeProfile** | Pointer to **NullableString** |  | [optional] 
**AsyncParams** | Pointer to [**IntegratedCAConnectorAsyncParams**](IntegratedCAConnectorAsyncParams.md) |  | [optional] 
**CaCert** | Pointer to **NullableString** |  | [optional] 
**CaKey** | Pointer to [**NullableSecretString**](SecretString.md) |  | [optional] 
**CertType** | Pointer to **NullableString** |  | [optional] 
**CheckPop** | Pointer to **NullableBool** |  | [optional] 
**CrlLifetime** | Pointer to **NullableString** |  | [optional] 
**CrlPath** | Pointer to **NullableString** |  | [optional] 
**CrtBackDate** | Pointer to **NullableString** |  | [optional] 
**CrtLifetime** | Pointer to **NullableString** |  | [optional] 
**CryptoType** | **string** |  | 
**SignAlg** | Pointer to **NullableString** |  | [optional] 
**AuthenticationDomainId** | Pointer to **int64** |  | [optional] 
**DefaultOwner** | **string** |  | 
**DeleteOnRevoke** | **bool** |  | 
**OwnerGroups** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | **int64** |  | 
**HashAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CertificateValidity** | Pointer to **NullableInt64** |  | [optional] 
**DefaultEmail** | Pointer to **NullableString** |  | [optional] 
**DefaultPhone** | Pointer to **NullableString** |  | [optional] 
**DomainId** | **string** |  | 
**EndpointType** | **string** |  | 
**SanEmailMap** | Pointer to **NullableString** |  | [optional] 
**UidMap** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZoneLabel** | Pointer to **NullableString** | The name of the label where the zone value is stored on an enrolled certificate | [optional] 
**EndPointIssuingCA** | **string** |  | 
**FormPorteurName** | Pointer to **NullableString** |  | [optional] 
**ProfilCle** | **NullableString** |  | 
**Workflow** | **NullableString** |  | 
**CustomerId** | **string** |  | 
**Environment** | **string** | The testing environment will use https://ote-api.nameshield.net endpoint  and the production will use https://api.nameshield.net  | 
**Procedure** | **string** |  | 
**CustomerUri** | **string** |  | 
**MpkiCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI.  It should contains the mpkiId as the login and the apiKey as password.  | 
**ProductUuid** | **string** | The product Uuid that need to be retrieved from the swiss sign api&#39;s (&lt;endpoints&gt;/v2/clients) | 

## Methods

### NewPKIConnectors

`func NewPKIConnectors(authenticationCredentials string, ca string, endPoint string, loginCredentials string, name string, template string, timeout NullableString, type_ string, accountKeyType string, dnsChallengeProvider DnsChallengeProviders, acmeDirectoryUrl string, caConfig string, domain string, enrollmentCredentials string, profile string, caArn string, region string, offerId string, organizationId int64, issuerCACert string, issuerCADN string, signerCredentials string, apiCredentials string, baseUrl string, productId string, caName string, cryptoType string, defaultOwner string, deleteOnRevoke bool, templateId int64, domainId string, endpointType string, endPointIssuingCA string, profilCle NullableString, workflow NullableString, customerId string, environment string, procedure string, customerUri string, mpkiCredentials string, productUuid string, ) *PKIConnectors`

NewPKIConnectors instantiates a new PKIConnectors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIConnectorsWithDefaults

`func NewPKIConnectorsWithDefaults() *PKIConnectors`

NewPKIConnectorsWithDefaults instantiates a new PKIConnectors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *PKIConnectors) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *PKIConnectors) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *PKIConnectors) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetCa

`func (o *PKIConnectors) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *PKIConnectors) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *PKIConnectors) SetCa(v string)`

SetCa sets Ca field to given value.


### GetEndPoint

`func (o *PKIConnectors) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *PKIConnectors) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *PKIConnectors) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetLoginCredentials

`func (o *PKIConnectors) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *PKIConnectors) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *PKIConnectors) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetName

`func (o *PKIConnectors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PKIConnectors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PKIConnectors) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *PKIConnectors) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PKIConnectors) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PKIConnectors) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PKIConnectors) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PKIConnectors) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PKIConnectors) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *PKIConnectors) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *PKIConnectors) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *PKIConnectors) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *PKIConnectors) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *PKIConnectors) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *PKIConnectors) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTemplate

`func (o *PKIConnectors) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PKIConnectors) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PKIConnectors) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetTimeout

`func (o *PKIConnectors) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PKIConnectors) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PKIConnectors) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *PKIConnectors) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PKIConnectors) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *PKIConnectors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PKIConnectors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PKIConnectors) SetType(v string)`

SetType sets Type field to given value.


### GetAccountEmail

`func (o *PKIConnectors) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *PKIConnectors) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *PKIConnectors) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *PKIConnectors) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### SetAccountEmailNil

`func (o *PKIConnectors) SetAccountEmailNil(b bool)`

 SetAccountEmailNil sets the value for AccountEmail to be an explicit nil

### UnsetAccountEmail
`func (o *PKIConnectors) UnsetAccountEmail()`

UnsetAccountEmail ensures that no value is present for AccountEmail, not even an explicit nil
### GetAccountKeyType

`func (o *PKIConnectors) GetAccountKeyType() string`

GetAccountKeyType returns the AccountKeyType field if non-nil, zero value otherwise.

### GetAccountKeyTypeOk

`func (o *PKIConnectors) GetAccountKeyTypeOk() (*string, bool)`

GetAccountKeyTypeOk returns a tuple with the AccountKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyType

`func (o *PKIConnectors) SetAccountKeyType(v string)`

SetAccountKeyType sets AccountKeyType field to given value.


### GetDnsChallengeProvider

`func (o *PKIConnectors) GetDnsChallengeProvider() DnsChallengeProviders`

GetDnsChallengeProvider returns the DnsChallengeProvider field if non-nil, zero value otherwise.

### GetDnsChallengeProviderOk

`func (o *PKIConnectors) GetDnsChallengeProviderOk() (*DnsChallengeProviders, bool)`

GetDnsChallengeProviderOk returns a tuple with the DnsChallengeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallengeProvider

`func (o *PKIConnectors) SetDnsChallengeProvider(v DnsChallengeProviders)`

SetDnsChallengeProvider sets DnsChallengeProvider field to given value.


### GetDomainDictionaryProvider

`func (o *PKIConnectors) GetDomainDictionaryProvider() DomainDictionaryProviders`

GetDomainDictionaryProvider returns the DomainDictionaryProvider field if non-nil, zero value otherwise.

### GetDomainDictionaryProviderOk

`func (o *PKIConnectors) GetDomainDictionaryProviderOk() (*DomainDictionaryProviders, bool)`

GetDomainDictionaryProviderOk returns a tuple with the DomainDictionaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDictionaryProvider

`func (o *PKIConnectors) SetDomainDictionaryProvider(v DomainDictionaryProviders)`

SetDomainDictionaryProvider sets DomainDictionaryProvider field to given value.

### HasDomainDictionaryProvider

`func (o *PKIConnectors) HasDomainDictionaryProvider() bool`

HasDomainDictionaryProvider returns a boolean if a field has been set.

### SetDomainDictionaryProviderNil

`func (o *PKIConnectors) SetDomainDictionaryProviderNil(b bool)`

 SetDomainDictionaryProviderNil sets the value for DomainDictionaryProvider to be an explicit nil

### UnsetDomainDictionaryProvider
`func (o *PKIConnectors) UnsetDomainDictionaryProvider()`

UnsetDomainDictionaryProvider ensures that no value is present for DomainDictionaryProvider, not even an explicit nil
### GetEab

`func (o *PKIConnectors) GetEab() string`

GetEab returns the Eab field if non-nil, zero value otherwise.

### GetEabOk

`func (o *PKIConnectors) GetEabOk() (*string, bool)`

GetEabOk returns a tuple with the Eab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEab

`func (o *PKIConnectors) SetEab(v string)`

SetEab sets Eab field to given value.

### HasEab

`func (o *PKIConnectors) HasEab() bool`

HasEab returns a boolean if a field has been set.

### SetEabNil

`func (o *PKIConnectors) SetEabNil(b bool)`

 SetEabNil sets the value for Eab to be an explicit nil

### UnsetEab
`func (o *PKIConnectors) UnsetEab()`

UnsetEab ensures that no value is present for Eab, not even an explicit nil
### GetEabMacAlgorithm

`func (o *PKIConnectors) GetEabMacAlgorithm() string`

GetEabMacAlgorithm returns the EabMacAlgorithm field if non-nil, zero value otherwise.

### GetEabMacAlgorithmOk

`func (o *PKIConnectors) GetEabMacAlgorithmOk() (*string, bool)`

GetEabMacAlgorithmOk returns a tuple with the EabMacAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabMacAlgorithm

`func (o *PKIConnectors) SetEabMacAlgorithm(v string)`

SetEabMacAlgorithm sets EabMacAlgorithm field to given value.

### HasEabMacAlgorithm

`func (o *PKIConnectors) HasEabMacAlgorithm() bool`

HasEabMacAlgorithm returns a boolean if a field has been set.

### SetEabMacAlgorithmNil

`func (o *PKIConnectors) SetEabMacAlgorithmNil(b bool)`

 SetEabMacAlgorithmNil sets the value for EabMacAlgorithm to be an explicit nil

### UnsetEabMacAlgorithm
`func (o *PKIConnectors) UnsetEabMacAlgorithm()`

UnsetEabMacAlgorithm ensures that no value is present for EabMacAlgorithm, not even an explicit nil
### GetRetryInterval

`func (o *PKIConnectors) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *PKIConnectors) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *PKIConnectors) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *PKIConnectors) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *PKIConnectors) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *PKIConnectors) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetRotateAccount

`func (o *PKIConnectors) GetRotateAccount() bool`

GetRotateAccount returns the RotateAccount field if non-nil, zero value otherwise.

### GetRotateAccountOk

`func (o *PKIConnectors) GetRotateAccountOk() (*bool, bool)`

GetRotateAccountOk returns a tuple with the RotateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAccount

`func (o *PKIConnectors) SetRotateAccount(v bool)`

SetRotateAccount sets RotateAccount field to given value.

### HasRotateAccount

`func (o *PKIConnectors) HasRotateAccount() bool`

HasRotateAccount returns a boolean if a field has been set.

### SetRotateAccountNil

`func (o *PKIConnectors) SetRotateAccountNil(b bool)`

 SetRotateAccountNil sets the value for RotateAccount to be an explicit nil

### UnsetRotateAccount
`func (o *PKIConnectors) UnsetRotateAccount()`

UnsetRotateAccount ensures that no value is present for RotateAccount, not even an explicit nil
### GetAcmeDirectoryUrl

`func (o *PKIConnectors) GetAcmeDirectoryUrl() string`

GetAcmeDirectoryUrl returns the AcmeDirectoryUrl field if non-nil, zero value otherwise.

### GetAcmeDirectoryUrlOk

`func (o *PKIConnectors) GetAcmeDirectoryUrlOk() (*string, bool)`

GetAcmeDirectoryUrlOk returns a tuple with the AcmeDirectoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUrl

`func (o *PKIConnectors) SetAcmeDirectoryUrl(v string)`

SetAcmeDirectoryUrl sets AcmeDirectoryUrl field to given value.


### GetCaConfig

`func (o *PKIConnectors) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *PKIConnectors) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *PKIConnectors) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetDomain

`func (o *PKIConnectors) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PKIConnectors) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PKIConnectors) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnrollmentCredentials

`func (o *PKIConnectors) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *PKIConnectors) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *PKIConnectors) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetProfile

`func (o *PKIConnectors) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PKIConnectors) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PKIConnectors) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAccessCredentials

`func (o *PKIConnectors) GetAccessCredentials() string`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *PKIConnectors) GetAccessCredentialsOk() (*string, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *PKIConnectors) SetAccessCredentials(v string)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *PKIConnectors) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### SetAccessCredentialsNil

`func (o *PKIConnectors) SetAccessCredentialsNil(b bool)`

 SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil

### UnsetAccessCredentials
`func (o *PKIConnectors) UnsetAccessCredentials()`

UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
### GetCaArn

`func (o *PKIConnectors) GetCaArn() string`

GetCaArn returns the CaArn field if non-nil, zero value otherwise.

### GetCaArnOk

`func (o *PKIConnectors) GetCaArnOk() (*string, bool)`

GetCaArnOk returns a tuple with the CaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaArn

`func (o *PKIConnectors) SetCaArn(v string)`

SetCaArn sets CaArn field to given value.


### GetCaPolicyOid

`func (o *PKIConnectors) GetCaPolicyOid() string`

GetCaPolicyOid returns the CaPolicyOid field if non-nil, zero value otherwise.

### GetCaPolicyOidOk

`func (o *PKIConnectors) GetCaPolicyOidOk() (*string, bool)`

GetCaPolicyOidOk returns a tuple with the CaPolicyOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPolicyOid

`func (o *PKIConnectors) SetCaPolicyOid(v string)`

SetCaPolicyOid sets CaPolicyOid field to given value.

### HasCaPolicyOid

`func (o *PKIConnectors) HasCaPolicyOid() bool`

HasCaPolicyOid returns a boolean if a field has been set.

### SetCaPolicyOidNil

`func (o *PKIConnectors) SetCaPolicyOidNil(b bool)`

 SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil

### UnsetCaPolicyOid
`func (o *PKIConnectors) UnsetCaPolicyOid()`

UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
### GetCertificateUsage

`func (o *PKIConnectors) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *PKIConnectors) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *PKIConnectors) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *PKIConnectors) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *PKIConnectors) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *PKIConnectors) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetRegion

`func (o *PKIConnectors) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PKIConnectors) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PKIConnectors) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRoleArn

`func (o *PKIConnectors) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *PKIConnectors) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *PKIConnectors) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *PKIConnectors) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *PKIConnectors) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *PKIConnectors) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetSigningHash

`func (o *PKIConnectors) GetSigningHash() string`

GetSigningHash returns the SigningHash field if non-nil, zero value otherwise.

### GetSigningHashOk

`func (o *PKIConnectors) GetSigningHashOk() (*string, bool)`

GetSigningHashOk returns a tuple with the SigningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningHash

`func (o *PKIConnectors) SetSigningHash(v string)`

SetSigningHash sets SigningHash field to given value.

### HasSigningHash

`func (o *PKIConnectors) HasSigningHash() bool`

HasSigningHash returns a boolean if a field has been set.

### SetSigningHashNil

`func (o *PKIConnectors) SetSigningHashNil(b bool)`

 SetSigningHashNil sets the value for SigningHash to be an explicit nil

### UnsetSigningHash
`func (o *PKIConnectors) UnsetSigningHash()`

UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
### GetTemplateArn

`func (o *PKIConnectors) GetTemplateArn() string`

GetTemplateArn returns the TemplateArn field if non-nil, zero value otherwise.

### GetTemplateArnOk

`func (o *PKIConnectors) GetTemplateArnOk() (*string, bool)`

GetTemplateArnOk returns a tuple with the TemplateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArn

`func (o *PKIConnectors) SetTemplateArn(v string)`

SetTemplateArn sets TemplateArn field to given value.

### HasTemplateArn

`func (o *PKIConnectors) HasTemplateArn() bool`

HasTemplateArn returns a boolean if a field has been set.

### SetTemplateArnNil

`func (o *PKIConnectors) SetTemplateArnNil(b bool)`

 SetTemplateArnNil sets the value for TemplateArn to be an explicit nil

### UnsetTemplateArn
`func (o *PKIConnectors) UnsetTemplateArn()`

UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
### GetValidDays

`func (o *PKIConnectors) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *PKIConnectors) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *PKIConnectors) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *PKIConnectors) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *PKIConnectors) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *PKIConnectors) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetOfferId

`func (o *PKIConnectors) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PKIConnectors) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PKIConnectors) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOrganizationId

`func (o *PKIConnectors) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PKIConnectors) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PKIConnectors) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetRevReason

`func (o *PKIConnectors) GetRevReason() string`

GetRevReason returns the RevReason field if non-nil, zero value otherwise.

### GetRevReasonOk

`func (o *PKIConnectors) GetRevReasonOk() (*string, bool)`

GetRevReasonOk returns a tuple with the RevReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevReason

`func (o *PKIConnectors) SetRevReason(v string)`

SetRevReason sets RevReason field to given value.

### HasRevReason

`func (o *PKIConnectors) HasRevReason() bool`

HasRevReason returns a boolean if a field has been set.

### SetRevReasonNil

`func (o *PKIConnectors) SetRevReasonNil(b bool)`

 SetRevReasonNil sets the value for RevReason to be an explicit nil

### UnsetRevReason
`func (o *PKIConnectors) UnsetRevReason()`

UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
### GetCnMap

`func (o *PKIConnectors) GetCnMap() string`

GetCnMap returns the CnMap field if non-nil, zero value otherwise.

### GetCnMapOk

`func (o *PKIConnectors) GetCnMapOk() (*string, bool)`

GetCnMapOk returns a tuple with the CnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnMap

`func (o *PKIConnectors) SetCnMap(v string)`

SetCnMap sets CnMap field to given value.

### HasCnMap

`func (o *PKIConnectors) HasCnMap() bool`

HasCnMap returns a boolean if a field has been set.

### SetCnMapNil

`func (o *PKIConnectors) SetCnMapNil(b bool)`

 SetCnMapNil sets the value for CnMap to be an explicit nil

### UnsetCnMap
`func (o *PKIConnectors) UnsetCnMap()`

UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
### GetEmailMap

`func (o *PKIConnectors) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *PKIConnectors) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *PKIConnectors) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *PKIConnectors) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *PKIConnectors) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *PKIConnectors) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetIssuerCACert

`func (o *PKIConnectors) GetIssuerCACert() string`

GetIssuerCACert returns the IssuerCACert field if non-nil, zero value otherwise.

### GetIssuerCACertOk

`func (o *PKIConnectors) GetIssuerCACertOk() (*string, bool)`

GetIssuerCACertOk returns a tuple with the IssuerCACert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCACert

`func (o *PKIConnectors) SetIssuerCACert(v string)`

SetIssuerCACert sets IssuerCACert field to given value.


### GetIssuerCADN

`func (o *PKIConnectors) GetIssuerCADN() string`

GetIssuerCADN returns the IssuerCADN field if non-nil, zero value otherwise.

### GetIssuerCADNOk

`func (o *PKIConnectors) GetIssuerCADNOk() (*string, bool)`

GetIssuerCADNOk returns a tuple with the IssuerCADN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCADN

`func (o *PKIConnectors) SetIssuerCADN(v string)`

SetIssuerCADN sets IssuerCADN field to given value.


### GetIssuerMap

`func (o *PKIConnectors) GetIssuerMap() string`

GetIssuerMap returns the IssuerMap field if non-nil, zero value otherwise.

### GetIssuerMapOk

`func (o *PKIConnectors) GetIssuerMapOk() (*string, bool)`

GetIssuerMapOk returns a tuple with the IssuerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMap

`func (o *PKIConnectors) SetIssuerMap(v string)`

SetIssuerMap sets IssuerMap field to given value.

### HasIssuerMap

`func (o *PKIConnectors) HasIssuerMap() bool`

HasIssuerMap returns a boolean if a field has been set.

### SetIssuerMapNil

`func (o *PKIConnectors) SetIssuerMapNil(b bool)`

 SetIssuerMapNil sets the value for IssuerMap to be an explicit nil

### UnsetIssuerMap
`func (o *PKIConnectors) UnsetIssuerMap()`

UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
### GetLegacyCMPStyle

`func (o *PKIConnectors) GetLegacyCMPStyle() bool`

GetLegacyCMPStyle returns the LegacyCMPStyle field if non-nil, zero value otherwise.

### GetLegacyCMPStyleOk

`func (o *PKIConnectors) GetLegacyCMPStyleOk() (*bool, bool)`

GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyCMPStyle

`func (o *PKIConnectors) SetLegacyCMPStyle(v bool)`

SetLegacyCMPStyle sets LegacyCMPStyle field to given value.

### HasLegacyCMPStyle

`func (o *PKIConnectors) HasLegacyCMPStyle() bool`

HasLegacyCMPStyle returns a boolean if a field has been set.

### SetLegacyCMPStyleNil

`func (o *PKIConnectors) SetLegacyCMPStyleNil(b bool)`

 SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil

### UnsetLegacyCMPStyle
`func (o *PKIConnectors) UnsetLegacyCMPStyle()`

UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
### GetProfileMap

`func (o *PKIConnectors) GetProfileMap() string`

GetProfileMap returns the ProfileMap field if non-nil, zero value otherwise.

### GetProfileMapOk

`func (o *PKIConnectors) GetProfileMapOk() (*string, bool)`

GetProfileMapOk returns a tuple with the ProfileMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMap

`func (o *PKIConnectors) SetProfileMap(v string)`

SetProfileMap sets ProfileMap field to given value.

### HasProfileMap

`func (o *PKIConnectors) HasProfileMap() bool`

HasProfileMap returns a boolean if a field has been set.

### SetProfileMapNil

`func (o *PKIConnectors) SetProfileMapNil(b bool)`

 SetProfileMapNil sets the value for ProfileMap to be an explicit nil

### UnsetProfileMap
`func (o *PKIConnectors) UnsetProfileMap()`

UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
### GetSanDnsMap

`func (o *PKIConnectors) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *PKIConnectors) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *PKIConnectors) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *PKIConnectors) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *PKIConnectors) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *PKIConnectors) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetSignerCredentials

`func (o *PKIConnectors) GetSignerCredentials() string`

GetSignerCredentials returns the SignerCredentials field if non-nil, zero value otherwise.

### GetSignerCredentialsOk

`func (o *PKIConnectors) GetSignerCredentialsOk() (*string, bool)`

GetSignerCredentialsOk returns a tuple with the SignerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCredentials

`func (o *PKIConnectors) SetSignerCredentials(v string)`

SetSignerCredentials sets SignerCredentials field to given value.


### GetApiCredentials

`func (o *PKIConnectors) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *PKIConnectors) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *PKIConnectors) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetBaseUrl

`func (o *PKIConnectors) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PKIConnectors) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PKIConnectors) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetCaCertId

`func (o *PKIConnectors) GetCaCertId() string`

GetCaCertId returns the CaCertId field if non-nil, zero value otherwise.

### GetCaCertIdOk

`func (o *PKIConnectors) GetCaCertIdOk() (*string, bool)`

GetCaCertIdOk returns a tuple with the CaCertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertId

`func (o *PKIConnectors) SetCaCertId(v string)`

SetCaCertId sets CaCertId field to given value.

### HasCaCertId

`func (o *PKIConnectors) HasCaCertId() bool`

HasCaCertId returns a boolean if a field has been set.

### SetCaCertIdNil

`func (o *PKIConnectors) SetCaCertIdNil(b bool)`

 SetCaCertIdNil sets the value for CaCertId to be an explicit nil

### UnsetCaCertId
`func (o *PKIConnectors) UnsetCaCertId()`

UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
### GetCustomConnectorDataMapping

`func (o *PKIConnectors) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *PKIConnectors) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *PKIConnectors) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *PKIConnectors) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *PKIConnectors) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *PKIConnectors) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetProductId

`func (o *PKIConnectors) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PKIConnectors) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PKIConnectors) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetSkipApproval

`func (o *PKIConnectors) GetSkipApproval() bool`

GetSkipApproval returns the SkipApproval field if non-nil, zero value otherwise.

### GetSkipApprovalOk

`func (o *PKIConnectors) GetSkipApprovalOk() (*bool, bool)`

GetSkipApprovalOk returns a tuple with the SkipApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipApproval

`func (o *PKIConnectors) SetSkipApproval(v bool)`

SetSkipApproval sets SkipApproval field to given value.

### HasSkipApproval

`func (o *PKIConnectors) HasSkipApproval() bool`

HasSkipApproval returns a boolean if a field has been set.

### SetSkipApprovalNil

`func (o *PKIConnectors) SetSkipApprovalNil(b bool)`

 SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil

### UnsetSkipApproval
`func (o *PKIConnectors) UnsetSkipApproval()`

UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
### GetCaName

`func (o *PKIConnectors) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *PKIConnectors) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *PKIConnectors) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *PKIConnectors) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *PKIConnectors) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *PKIConnectors) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *PKIConnectors) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *PKIConnectors) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *PKIConnectors) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetAsyncParams

`func (o *PKIConnectors) GetAsyncParams() IntegratedCAConnectorAsyncParams`

GetAsyncParams returns the AsyncParams field if non-nil, zero value otherwise.

### GetAsyncParamsOk

`func (o *PKIConnectors) GetAsyncParamsOk() (*IntegratedCAConnectorAsyncParams, bool)`

GetAsyncParamsOk returns a tuple with the AsyncParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncParams

`func (o *PKIConnectors) SetAsyncParams(v IntegratedCAConnectorAsyncParams)`

SetAsyncParams sets AsyncParams field to given value.

### HasAsyncParams

`func (o *PKIConnectors) HasAsyncParams() bool`

HasAsyncParams returns a boolean if a field has been set.

### GetCaCert

`func (o *PKIConnectors) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PKIConnectors) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PKIConnectors) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PKIConnectors) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PKIConnectors) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PKIConnectors) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCaKey

`func (o *PKIConnectors) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *PKIConnectors) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *PKIConnectors) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *PKIConnectors) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *PKIConnectors) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *PKIConnectors) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCertType

`func (o *PKIConnectors) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *PKIConnectors) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *PKIConnectors) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *PKIConnectors) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### SetCertTypeNil

`func (o *PKIConnectors) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *PKIConnectors) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetCheckPop

`func (o *PKIConnectors) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *PKIConnectors) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *PKIConnectors) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *PKIConnectors) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *PKIConnectors) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *PKIConnectors) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCrlLifetime

`func (o *PKIConnectors) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *PKIConnectors) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *PKIConnectors) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *PKIConnectors) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *PKIConnectors) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *PKIConnectors) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetCrlPath

`func (o *PKIConnectors) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *PKIConnectors) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *PKIConnectors) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *PKIConnectors) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *PKIConnectors) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *PKIConnectors) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrtBackDate

`func (o *PKIConnectors) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *PKIConnectors) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *PKIConnectors) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *PKIConnectors) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *PKIConnectors) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *PKIConnectors) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCrtLifetime

`func (o *PKIConnectors) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *PKIConnectors) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *PKIConnectors) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *PKIConnectors) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *PKIConnectors) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *PKIConnectors) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCryptoType

`func (o *PKIConnectors) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *PKIConnectors) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *PKIConnectors) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetSignAlg

`func (o *PKIConnectors) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *PKIConnectors) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *PKIConnectors) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *PKIConnectors) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *PKIConnectors) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *PKIConnectors) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetAuthenticationDomainId

`func (o *PKIConnectors) GetAuthenticationDomainId() int64`

GetAuthenticationDomainId returns the AuthenticationDomainId field if non-nil, zero value otherwise.

### GetAuthenticationDomainIdOk

`func (o *PKIConnectors) GetAuthenticationDomainIdOk() (*int64, bool)`

GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDomainId

`func (o *PKIConnectors) SetAuthenticationDomainId(v int64)`

SetAuthenticationDomainId sets AuthenticationDomainId field to given value.

### HasAuthenticationDomainId

`func (o *PKIConnectors) HasAuthenticationDomainId() bool`

HasAuthenticationDomainId returns a boolean if a field has been set.

### GetDefaultOwner

`func (o *PKIConnectors) GetDefaultOwner() string`

GetDefaultOwner returns the DefaultOwner field if non-nil, zero value otherwise.

### GetDefaultOwnerOk

`func (o *PKIConnectors) GetDefaultOwnerOk() (*string, bool)`

GetDefaultOwnerOk returns a tuple with the DefaultOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOwner

`func (o *PKIConnectors) SetDefaultOwner(v string)`

SetDefaultOwner sets DefaultOwner field to given value.


### GetDeleteOnRevoke

`func (o *PKIConnectors) GetDeleteOnRevoke() bool`

GetDeleteOnRevoke returns the DeleteOnRevoke field if non-nil, zero value otherwise.

### GetDeleteOnRevokeOk

`func (o *PKIConnectors) GetDeleteOnRevokeOk() (*bool, bool)`

GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnRevoke

`func (o *PKIConnectors) SetDeleteOnRevoke(v bool)`

SetDeleteOnRevoke sets DeleteOnRevoke field to given value.


### GetOwnerGroups

`func (o *PKIConnectors) GetOwnerGroups() string`

GetOwnerGroups returns the OwnerGroups field if non-nil, zero value otherwise.

### GetOwnerGroupsOk

`func (o *PKIConnectors) GetOwnerGroupsOk() (*string, bool)`

GetOwnerGroupsOk returns a tuple with the OwnerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroups

`func (o *PKIConnectors) SetOwnerGroups(v string)`

SetOwnerGroups sets OwnerGroups field to given value.

### HasOwnerGroups

`func (o *PKIConnectors) HasOwnerGroups() bool`

HasOwnerGroups returns a boolean if a field has been set.

### SetOwnerGroupsNil

`func (o *PKIConnectors) SetOwnerGroupsNil(b bool)`

 SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil

### UnsetOwnerGroups
`func (o *PKIConnectors) UnsetOwnerGroups()`

UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
### GetTemplateId

`func (o *PKIConnectors) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PKIConnectors) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PKIConnectors) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetHashAlgorithm

`func (o *PKIConnectors) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *PKIConnectors) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *PKIConnectors) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *PKIConnectors) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### SetHashAlgorithmNil

`func (o *PKIConnectors) SetHashAlgorithmNil(b bool)`

 SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil

### UnsetHashAlgorithm
`func (o *PKIConnectors) UnsetHashAlgorithm()`

UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
### GetCertificateValidity

`func (o *PKIConnectors) GetCertificateValidity() int64`

GetCertificateValidity returns the CertificateValidity field if non-nil, zero value otherwise.

### GetCertificateValidityOk

`func (o *PKIConnectors) GetCertificateValidityOk() (*int64, bool)`

GetCertificateValidityOk returns a tuple with the CertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidity

`func (o *PKIConnectors) SetCertificateValidity(v int64)`

SetCertificateValidity sets CertificateValidity field to given value.

### HasCertificateValidity

`func (o *PKIConnectors) HasCertificateValidity() bool`

HasCertificateValidity returns a boolean if a field has been set.

### SetCertificateValidityNil

`func (o *PKIConnectors) SetCertificateValidityNil(b bool)`

 SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil

### UnsetCertificateValidity
`func (o *PKIConnectors) UnsetCertificateValidity()`

UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
### GetDefaultEmail

`func (o *PKIConnectors) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *PKIConnectors) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *PKIConnectors) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *PKIConnectors) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *PKIConnectors) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *PKIConnectors) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *PKIConnectors) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *PKIConnectors) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *PKIConnectors) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *PKIConnectors) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *PKIConnectors) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *PKIConnectors) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetDomainId

`func (o *PKIConnectors) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PKIConnectors) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PKIConnectors) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetEndpointType

`func (o *PKIConnectors) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *PKIConnectors) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *PKIConnectors) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetSanEmailMap

`func (o *PKIConnectors) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *PKIConnectors) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *PKIConnectors) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *PKIConnectors) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *PKIConnectors) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *PKIConnectors) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetUidMap

`func (o *PKIConnectors) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *PKIConnectors) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *PKIConnectors) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *PKIConnectors) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *PKIConnectors) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *PKIConnectors) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *PKIConnectors) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PKIConnectors) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PKIConnectors) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PKIConnectors) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PKIConnectors) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PKIConnectors) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *PKIConnectors) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *PKIConnectors) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *PKIConnectors) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *PKIConnectors) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *PKIConnectors) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *PKIConnectors) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
### GetEndPointIssuingCA

`func (o *PKIConnectors) GetEndPointIssuingCA() string`

GetEndPointIssuingCA returns the EndPointIssuingCA field if non-nil, zero value otherwise.

### GetEndPointIssuingCAOk

`func (o *PKIConnectors) GetEndPointIssuingCAOk() (*string, bool)`

GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIssuingCA

`func (o *PKIConnectors) SetEndPointIssuingCA(v string)`

SetEndPointIssuingCA sets EndPointIssuingCA field to given value.


### GetFormPorteurName

`func (o *PKIConnectors) GetFormPorteurName() string`

GetFormPorteurName returns the FormPorteurName field if non-nil, zero value otherwise.

### GetFormPorteurNameOk

`func (o *PKIConnectors) GetFormPorteurNameOk() (*string, bool)`

GetFormPorteurNameOk returns a tuple with the FormPorteurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormPorteurName

`func (o *PKIConnectors) SetFormPorteurName(v string)`

SetFormPorteurName sets FormPorteurName field to given value.

### HasFormPorteurName

`func (o *PKIConnectors) HasFormPorteurName() bool`

HasFormPorteurName returns a boolean if a field has been set.

### SetFormPorteurNameNil

`func (o *PKIConnectors) SetFormPorteurNameNil(b bool)`

 SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil

### UnsetFormPorteurName
`func (o *PKIConnectors) UnsetFormPorteurName()`

UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
### GetProfilCle

`func (o *PKIConnectors) GetProfilCle() string`

GetProfilCle returns the ProfilCle field if non-nil, zero value otherwise.

### GetProfilCleOk

`func (o *PKIConnectors) GetProfilCleOk() (*string, bool)`

GetProfilCleOk returns a tuple with the ProfilCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilCle

`func (o *PKIConnectors) SetProfilCle(v string)`

SetProfilCle sets ProfilCle field to given value.


### SetProfilCleNil

`func (o *PKIConnectors) SetProfilCleNil(b bool)`

 SetProfilCleNil sets the value for ProfilCle to be an explicit nil

### UnsetProfilCle
`func (o *PKIConnectors) UnsetProfilCle()`

UnsetProfilCle ensures that no value is present for ProfilCle, not even an explicit nil
### GetWorkflow

`func (o *PKIConnectors) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *PKIConnectors) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *PKIConnectors) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### SetWorkflowNil

`func (o *PKIConnectors) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *PKIConnectors) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetCustomerId

`func (o *PKIConnectors) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PKIConnectors) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PKIConnectors) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEnvironment

`func (o *PKIConnectors) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PKIConnectors) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PKIConnectors) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProcedure

`func (o *PKIConnectors) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *PKIConnectors) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *PKIConnectors) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.


### GetCustomerUri

`func (o *PKIConnectors) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *PKIConnectors) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *PKIConnectors) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetMpkiCredentials

`func (o *PKIConnectors) GetMpkiCredentials() string`

GetMpkiCredentials returns the MpkiCredentials field if non-nil, zero value otherwise.

### GetMpkiCredentialsOk

`func (o *PKIConnectors) GetMpkiCredentialsOk() (*string, bool)`

GetMpkiCredentialsOk returns a tuple with the MpkiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpkiCredentials

`func (o *PKIConnectors) SetMpkiCredentials(v string)`

SetMpkiCredentials sets MpkiCredentials field to given value.


### GetProductUuid

`func (o *PKIConnectors) GetProductUuid() string`

GetProductUuid returns the ProductUuid field if non-nil, zero value otherwise.

### GetProductUuidOk

`func (o *PKIConnectors) GetProductUuidOk() (*string, bool)`

GetProductUuidOk returns a tuple with the ProductUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUuid

`func (o *PKIConnectors) SetProductUuid(v string)`

SetProductUuid sets ProductUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


