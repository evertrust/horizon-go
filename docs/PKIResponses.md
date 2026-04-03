# PKIResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Ca** | **string** | Stream&#39;s technical name of the CA on which to enroll | 
**EndPoint** | **string** | Swiss base endpoint | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Template** | **string** | Stream&#39;s certificate template to use for enrollment | 
**Timeout** | **NullableString** |  | 
**Type** | **string** |  | 
**AccountEmail** | Pointer to **NullableString** | Email to associate with the account | [optional] 
**AccountKeyType** | **string** | The key type to use to generate the account key | 
**AccountUrl** | **string** | Url of the account on the ACME directory | 
**DnsChallengeProvider** | [**DnsChallengeProviders**](DnsChallengeProviders.md) | DNS Provider configuration to provision the DNS challenge. Available from &#x60;2.7.7&#x60; | 
**DomainDictionaryProvider** | Pointer to [**NullableDomainDictionaryProviders**](DomainDictionaryProviders.md) | The dictionary provider | [optional] 
**Eab** | Pointer to **NullableString** | &#x60;password&#x60; credentials name to use for External Account Binding | [optional] 
**RotateAccount** | Pointer to **NullableBool** | If enable, regenerate the account (does not need to be specified on creation) | [optional] 
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
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
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
**CustomConnectorDataMapping** | Pointer to **map[string]string** |  | [optional] 
**ProductId** | **string** |  | 
**SkipApproval** | Pointer to **NullableBool** |  | [optional] 
**CaName** | **string** |  | 
**EeProfile** | Pointer to **NullableString** |  | [optional] 
**CertLifetime** | Pointer to **NullableString** |  | [optional] 
**CertType** | **NullableString** |  | 
**ClientId** | Pointer to **NullableInt64** |  | [optional] 
**RequesterDefaultMail** | **string** |  | 
**RequesterName** | Pointer to **NullableString** |  | [optional] 
**RequesterPhone** | Pointer to **NullableString** |  | [optional] 
**CaCert** | Pointer to **NullableString** |  | [optional] 
**CaKey** | Pointer to [**NullableSecretString**](SecretString.md) |  | [optional] 
**CheckPop** | Pointer to **NullableBool** |  | [optional] 
**CrlLifetime** | Pointer to **NullableString** |  | [optional] 
**CrlPath** | Pointer to **NullableString** |  | [optional] 
**CrtBackDate** | Pointer to **NullableString** |  | [optional] 
**CrtLifetime** | Pointer to **NullableString** |  | [optional] 
**CryptoType** | **string** |  | 
**SignAlg** | Pointer to **NullableString** |  | [optional] 
**AuthenticationDomainId** | **int64** |  | 
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

### NewPKIResponses

`func NewPKIResponses(id string, authenticationCredentials string, ca string, endPoint string, loginCredentials string, name string, template string, timeout NullableString, type_ string, accountKeyType string, accountUrl string, dnsChallengeProvider DnsChallengeProviders, acmeDirectoryUrl string, caConfig string, domain string, enrollmentCredentials string, profile string, caArn string, region string, offerId string, organizationId int64, issuerCACert string, issuerCADN string, signerCredentials string, apiCredentials string, baseUrl string, productId string, caName string, certType NullableString, requesterDefaultMail string, cryptoType string, authenticationDomainId int64, defaultOwner string, deleteOnRevoke bool, templateId int64, domainId string, endpointType string, endPointIssuingCA string, profilCle NullableString, workflow NullableString, customerId string, environment string, procedure string, customerUri string, mpkiCredentials string, productUuid string, ) *PKIResponses`

NewPKIResponses instantiates a new PKIResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIResponsesWithDefaults

`func NewPKIResponsesWithDefaults() *PKIResponses`

NewPKIResponsesWithDefaults instantiates a new PKIResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PKIResponses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PKIResponses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PKIResponses) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationCredentials

`func (o *PKIResponses) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *PKIResponses) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *PKIResponses) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetCa

`func (o *PKIResponses) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *PKIResponses) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *PKIResponses) SetCa(v string)`

SetCa sets Ca field to given value.


### GetEndPoint

`func (o *PKIResponses) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *PKIResponses) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *PKIResponses) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetLoginCredentials

`func (o *PKIResponses) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *PKIResponses) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *PKIResponses) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetName

`func (o *PKIResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PKIResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PKIResponses) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *PKIResponses) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PKIResponses) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PKIResponses) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PKIResponses) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PKIResponses) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PKIResponses) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *PKIResponses) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *PKIResponses) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *PKIResponses) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *PKIResponses) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *PKIResponses) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *PKIResponses) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *PKIResponses) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PKIResponses) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PKIResponses) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PKIResponses) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PKIResponses) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PKIResponses) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTemplate

`func (o *PKIResponses) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PKIResponses) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PKIResponses) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetTimeout

`func (o *PKIResponses) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PKIResponses) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PKIResponses) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *PKIResponses) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PKIResponses) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *PKIResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PKIResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PKIResponses) SetType(v string)`

SetType sets Type field to given value.


### GetAccountEmail

`func (o *PKIResponses) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *PKIResponses) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *PKIResponses) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *PKIResponses) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### SetAccountEmailNil

`func (o *PKIResponses) SetAccountEmailNil(b bool)`

 SetAccountEmailNil sets the value for AccountEmail to be an explicit nil

### UnsetAccountEmail
`func (o *PKIResponses) UnsetAccountEmail()`

UnsetAccountEmail ensures that no value is present for AccountEmail, not even an explicit nil
### GetAccountKeyType

`func (o *PKIResponses) GetAccountKeyType() string`

GetAccountKeyType returns the AccountKeyType field if non-nil, zero value otherwise.

### GetAccountKeyTypeOk

`func (o *PKIResponses) GetAccountKeyTypeOk() (*string, bool)`

GetAccountKeyTypeOk returns a tuple with the AccountKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyType

`func (o *PKIResponses) SetAccountKeyType(v string)`

SetAccountKeyType sets AccountKeyType field to given value.


### GetAccountUrl

`func (o *PKIResponses) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *PKIResponses) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *PKIResponses) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.


### GetDnsChallengeProvider

`func (o *PKIResponses) GetDnsChallengeProvider() DnsChallengeProviders`

GetDnsChallengeProvider returns the DnsChallengeProvider field if non-nil, zero value otherwise.

### GetDnsChallengeProviderOk

`func (o *PKIResponses) GetDnsChallengeProviderOk() (*DnsChallengeProviders, bool)`

GetDnsChallengeProviderOk returns a tuple with the DnsChallengeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallengeProvider

`func (o *PKIResponses) SetDnsChallengeProvider(v DnsChallengeProviders)`

SetDnsChallengeProvider sets DnsChallengeProvider field to given value.


### GetDomainDictionaryProvider

`func (o *PKIResponses) GetDomainDictionaryProvider() DomainDictionaryProviders`

GetDomainDictionaryProvider returns the DomainDictionaryProvider field if non-nil, zero value otherwise.

### GetDomainDictionaryProviderOk

`func (o *PKIResponses) GetDomainDictionaryProviderOk() (*DomainDictionaryProviders, bool)`

GetDomainDictionaryProviderOk returns a tuple with the DomainDictionaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDictionaryProvider

`func (o *PKIResponses) SetDomainDictionaryProvider(v DomainDictionaryProviders)`

SetDomainDictionaryProvider sets DomainDictionaryProvider field to given value.

### HasDomainDictionaryProvider

`func (o *PKIResponses) HasDomainDictionaryProvider() bool`

HasDomainDictionaryProvider returns a boolean if a field has been set.

### SetDomainDictionaryProviderNil

`func (o *PKIResponses) SetDomainDictionaryProviderNil(b bool)`

 SetDomainDictionaryProviderNil sets the value for DomainDictionaryProvider to be an explicit nil

### UnsetDomainDictionaryProvider
`func (o *PKIResponses) UnsetDomainDictionaryProvider()`

UnsetDomainDictionaryProvider ensures that no value is present for DomainDictionaryProvider, not even an explicit nil
### GetEab

`func (o *PKIResponses) GetEab() string`

GetEab returns the Eab field if non-nil, zero value otherwise.

### GetEabOk

`func (o *PKIResponses) GetEabOk() (*string, bool)`

GetEabOk returns a tuple with the Eab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEab

`func (o *PKIResponses) SetEab(v string)`

SetEab sets Eab field to given value.

### HasEab

`func (o *PKIResponses) HasEab() bool`

HasEab returns a boolean if a field has been set.

### SetEabNil

`func (o *PKIResponses) SetEabNil(b bool)`

 SetEabNil sets the value for Eab to be an explicit nil

### UnsetEab
`func (o *PKIResponses) UnsetEab()`

UnsetEab ensures that no value is present for Eab, not even an explicit nil
### GetRotateAccount

`func (o *PKIResponses) GetRotateAccount() bool`

GetRotateAccount returns the RotateAccount field if non-nil, zero value otherwise.

### GetRotateAccountOk

`func (o *PKIResponses) GetRotateAccountOk() (*bool, bool)`

GetRotateAccountOk returns a tuple with the RotateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAccount

`func (o *PKIResponses) SetRotateAccount(v bool)`

SetRotateAccount sets RotateAccount field to given value.

### HasRotateAccount

`func (o *PKIResponses) HasRotateAccount() bool`

HasRotateAccount returns a boolean if a field has been set.

### SetRotateAccountNil

`func (o *PKIResponses) SetRotateAccountNil(b bool)`

 SetRotateAccountNil sets the value for RotateAccount to be an explicit nil

### UnsetRotateAccount
`func (o *PKIResponses) UnsetRotateAccount()`

UnsetRotateAccount ensures that no value is present for RotateAccount, not even an explicit nil
### GetAcmeDirectoryUrl

`func (o *PKIResponses) GetAcmeDirectoryUrl() string`

GetAcmeDirectoryUrl returns the AcmeDirectoryUrl field if non-nil, zero value otherwise.

### GetAcmeDirectoryUrlOk

`func (o *PKIResponses) GetAcmeDirectoryUrlOk() (*string, bool)`

GetAcmeDirectoryUrlOk returns a tuple with the AcmeDirectoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUrl

`func (o *PKIResponses) SetAcmeDirectoryUrl(v string)`

SetAcmeDirectoryUrl sets AcmeDirectoryUrl field to given value.


### GetCaConfig

`func (o *PKIResponses) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *PKIResponses) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *PKIResponses) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetDomain

`func (o *PKIResponses) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PKIResponses) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PKIResponses) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnrollmentCredentials

`func (o *PKIResponses) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *PKIResponses) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *PKIResponses) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetProfile

`func (o *PKIResponses) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PKIResponses) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PKIResponses) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAccessCredentials

`func (o *PKIResponses) GetAccessCredentials() string`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *PKIResponses) GetAccessCredentialsOk() (*string, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *PKIResponses) SetAccessCredentials(v string)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *PKIResponses) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### SetAccessCredentialsNil

`func (o *PKIResponses) SetAccessCredentialsNil(b bool)`

 SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil

### UnsetAccessCredentials
`func (o *PKIResponses) UnsetAccessCredentials()`

UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
### GetCaArn

`func (o *PKIResponses) GetCaArn() string`

GetCaArn returns the CaArn field if non-nil, zero value otherwise.

### GetCaArnOk

`func (o *PKIResponses) GetCaArnOk() (*string, bool)`

GetCaArnOk returns a tuple with the CaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaArn

`func (o *PKIResponses) SetCaArn(v string)`

SetCaArn sets CaArn field to given value.


### GetCaPolicyOid

`func (o *PKIResponses) GetCaPolicyOid() string`

GetCaPolicyOid returns the CaPolicyOid field if non-nil, zero value otherwise.

### GetCaPolicyOidOk

`func (o *PKIResponses) GetCaPolicyOidOk() (*string, bool)`

GetCaPolicyOidOk returns a tuple with the CaPolicyOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPolicyOid

`func (o *PKIResponses) SetCaPolicyOid(v string)`

SetCaPolicyOid sets CaPolicyOid field to given value.

### HasCaPolicyOid

`func (o *PKIResponses) HasCaPolicyOid() bool`

HasCaPolicyOid returns a boolean if a field has been set.

### SetCaPolicyOidNil

`func (o *PKIResponses) SetCaPolicyOidNil(b bool)`

 SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil

### UnsetCaPolicyOid
`func (o *PKIResponses) UnsetCaPolicyOid()`

UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
### GetCertificateUsage

`func (o *PKIResponses) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *PKIResponses) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *PKIResponses) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *PKIResponses) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *PKIResponses) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *PKIResponses) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetRegion

`func (o *PKIResponses) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PKIResponses) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PKIResponses) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRetryInterval

`func (o *PKIResponses) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *PKIResponses) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *PKIResponses) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *PKIResponses) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *PKIResponses) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *PKIResponses) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetRoleArn

`func (o *PKIResponses) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *PKIResponses) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *PKIResponses) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *PKIResponses) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *PKIResponses) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *PKIResponses) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetSigningHash

`func (o *PKIResponses) GetSigningHash() string`

GetSigningHash returns the SigningHash field if non-nil, zero value otherwise.

### GetSigningHashOk

`func (o *PKIResponses) GetSigningHashOk() (*string, bool)`

GetSigningHashOk returns a tuple with the SigningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningHash

`func (o *PKIResponses) SetSigningHash(v string)`

SetSigningHash sets SigningHash field to given value.

### HasSigningHash

`func (o *PKIResponses) HasSigningHash() bool`

HasSigningHash returns a boolean if a field has been set.

### SetSigningHashNil

`func (o *PKIResponses) SetSigningHashNil(b bool)`

 SetSigningHashNil sets the value for SigningHash to be an explicit nil

### UnsetSigningHash
`func (o *PKIResponses) UnsetSigningHash()`

UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
### GetTemplateArn

`func (o *PKIResponses) GetTemplateArn() string`

GetTemplateArn returns the TemplateArn field if non-nil, zero value otherwise.

### GetTemplateArnOk

`func (o *PKIResponses) GetTemplateArnOk() (*string, bool)`

GetTemplateArnOk returns a tuple with the TemplateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArn

`func (o *PKIResponses) SetTemplateArn(v string)`

SetTemplateArn sets TemplateArn field to given value.

### HasTemplateArn

`func (o *PKIResponses) HasTemplateArn() bool`

HasTemplateArn returns a boolean if a field has been set.

### SetTemplateArnNil

`func (o *PKIResponses) SetTemplateArnNil(b bool)`

 SetTemplateArnNil sets the value for TemplateArn to be an explicit nil

### UnsetTemplateArn
`func (o *PKIResponses) UnsetTemplateArn()`

UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
### GetValidDays

`func (o *PKIResponses) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *PKIResponses) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *PKIResponses) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *PKIResponses) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *PKIResponses) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *PKIResponses) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetOfferId

`func (o *PKIResponses) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PKIResponses) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PKIResponses) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOrganizationId

`func (o *PKIResponses) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PKIResponses) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PKIResponses) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetRevReason

`func (o *PKIResponses) GetRevReason() string`

GetRevReason returns the RevReason field if non-nil, zero value otherwise.

### GetRevReasonOk

`func (o *PKIResponses) GetRevReasonOk() (*string, bool)`

GetRevReasonOk returns a tuple with the RevReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevReason

`func (o *PKIResponses) SetRevReason(v string)`

SetRevReason sets RevReason field to given value.

### HasRevReason

`func (o *PKIResponses) HasRevReason() bool`

HasRevReason returns a boolean if a field has been set.

### SetRevReasonNil

`func (o *PKIResponses) SetRevReasonNil(b bool)`

 SetRevReasonNil sets the value for RevReason to be an explicit nil

### UnsetRevReason
`func (o *PKIResponses) UnsetRevReason()`

UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
### GetCnMap

`func (o *PKIResponses) GetCnMap() string`

GetCnMap returns the CnMap field if non-nil, zero value otherwise.

### GetCnMapOk

`func (o *PKIResponses) GetCnMapOk() (*string, bool)`

GetCnMapOk returns a tuple with the CnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnMap

`func (o *PKIResponses) SetCnMap(v string)`

SetCnMap sets CnMap field to given value.

### HasCnMap

`func (o *PKIResponses) HasCnMap() bool`

HasCnMap returns a boolean if a field has been set.

### SetCnMapNil

`func (o *PKIResponses) SetCnMapNil(b bool)`

 SetCnMapNil sets the value for CnMap to be an explicit nil

### UnsetCnMap
`func (o *PKIResponses) UnsetCnMap()`

UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
### GetEmailMap

`func (o *PKIResponses) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *PKIResponses) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *PKIResponses) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *PKIResponses) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *PKIResponses) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *PKIResponses) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetIssuerCACert

`func (o *PKIResponses) GetIssuerCACert() string`

GetIssuerCACert returns the IssuerCACert field if non-nil, zero value otherwise.

### GetIssuerCACertOk

`func (o *PKIResponses) GetIssuerCACertOk() (*string, bool)`

GetIssuerCACertOk returns a tuple with the IssuerCACert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCACert

`func (o *PKIResponses) SetIssuerCACert(v string)`

SetIssuerCACert sets IssuerCACert field to given value.


### GetIssuerCADN

`func (o *PKIResponses) GetIssuerCADN() string`

GetIssuerCADN returns the IssuerCADN field if non-nil, zero value otherwise.

### GetIssuerCADNOk

`func (o *PKIResponses) GetIssuerCADNOk() (*string, bool)`

GetIssuerCADNOk returns a tuple with the IssuerCADN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCADN

`func (o *PKIResponses) SetIssuerCADN(v string)`

SetIssuerCADN sets IssuerCADN field to given value.


### GetIssuerMap

`func (o *PKIResponses) GetIssuerMap() string`

GetIssuerMap returns the IssuerMap field if non-nil, zero value otherwise.

### GetIssuerMapOk

`func (o *PKIResponses) GetIssuerMapOk() (*string, bool)`

GetIssuerMapOk returns a tuple with the IssuerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMap

`func (o *PKIResponses) SetIssuerMap(v string)`

SetIssuerMap sets IssuerMap field to given value.

### HasIssuerMap

`func (o *PKIResponses) HasIssuerMap() bool`

HasIssuerMap returns a boolean if a field has been set.

### SetIssuerMapNil

`func (o *PKIResponses) SetIssuerMapNil(b bool)`

 SetIssuerMapNil sets the value for IssuerMap to be an explicit nil

### UnsetIssuerMap
`func (o *PKIResponses) UnsetIssuerMap()`

UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
### GetLegacyCMPStyle

`func (o *PKIResponses) GetLegacyCMPStyle() bool`

GetLegacyCMPStyle returns the LegacyCMPStyle field if non-nil, zero value otherwise.

### GetLegacyCMPStyleOk

`func (o *PKIResponses) GetLegacyCMPStyleOk() (*bool, bool)`

GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyCMPStyle

`func (o *PKIResponses) SetLegacyCMPStyle(v bool)`

SetLegacyCMPStyle sets LegacyCMPStyle field to given value.

### HasLegacyCMPStyle

`func (o *PKIResponses) HasLegacyCMPStyle() bool`

HasLegacyCMPStyle returns a boolean if a field has been set.

### SetLegacyCMPStyleNil

`func (o *PKIResponses) SetLegacyCMPStyleNil(b bool)`

 SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil

### UnsetLegacyCMPStyle
`func (o *PKIResponses) UnsetLegacyCMPStyle()`

UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
### GetProfileMap

`func (o *PKIResponses) GetProfileMap() string`

GetProfileMap returns the ProfileMap field if non-nil, zero value otherwise.

### GetProfileMapOk

`func (o *PKIResponses) GetProfileMapOk() (*string, bool)`

GetProfileMapOk returns a tuple with the ProfileMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMap

`func (o *PKIResponses) SetProfileMap(v string)`

SetProfileMap sets ProfileMap field to given value.

### HasProfileMap

`func (o *PKIResponses) HasProfileMap() bool`

HasProfileMap returns a boolean if a field has been set.

### SetProfileMapNil

`func (o *PKIResponses) SetProfileMapNil(b bool)`

 SetProfileMapNil sets the value for ProfileMap to be an explicit nil

### UnsetProfileMap
`func (o *PKIResponses) UnsetProfileMap()`

UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
### GetSanDnsMap

`func (o *PKIResponses) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *PKIResponses) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *PKIResponses) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *PKIResponses) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *PKIResponses) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *PKIResponses) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetSignerCredentials

`func (o *PKIResponses) GetSignerCredentials() string`

GetSignerCredentials returns the SignerCredentials field if non-nil, zero value otherwise.

### GetSignerCredentialsOk

`func (o *PKIResponses) GetSignerCredentialsOk() (*string, bool)`

GetSignerCredentialsOk returns a tuple with the SignerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCredentials

`func (o *PKIResponses) SetSignerCredentials(v string)`

SetSignerCredentials sets SignerCredentials field to given value.


### GetApiCredentials

`func (o *PKIResponses) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *PKIResponses) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *PKIResponses) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetBaseUrl

`func (o *PKIResponses) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PKIResponses) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PKIResponses) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetCaCertId

`func (o *PKIResponses) GetCaCertId() string`

GetCaCertId returns the CaCertId field if non-nil, zero value otherwise.

### GetCaCertIdOk

`func (o *PKIResponses) GetCaCertIdOk() (*string, bool)`

GetCaCertIdOk returns a tuple with the CaCertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertId

`func (o *PKIResponses) SetCaCertId(v string)`

SetCaCertId sets CaCertId field to given value.

### HasCaCertId

`func (o *PKIResponses) HasCaCertId() bool`

HasCaCertId returns a boolean if a field has been set.

### SetCaCertIdNil

`func (o *PKIResponses) SetCaCertIdNil(b bool)`

 SetCaCertIdNil sets the value for CaCertId to be an explicit nil

### UnsetCaCertId
`func (o *PKIResponses) UnsetCaCertId()`

UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
### GetCustomConnectorDataMapping

`func (o *PKIResponses) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *PKIResponses) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *PKIResponses) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *PKIResponses) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *PKIResponses) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *PKIResponses) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetProductId

`func (o *PKIResponses) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PKIResponses) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PKIResponses) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetSkipApproval

`func (o *PKIResponses) GetSkipApproval() bool`

GetSkipApproval returns the SkipApproval field if non-nil, zero value otherwise.

### GetSkipApprovalOk

`func (o *PKIResponses) GetSkipApprovalOk() (*bool, bool)`

GetSkipApprovalOk returns a tuple with the SkipApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipApproval

`func (o *PKIResponses) SetSkipApproval(v bool)`

SetSkipApproval sets SkipApproval field to given value.

### HasSkipApproval

`func (o *PKIResponses) HasSkipApproval() bool`

HasSkipApproval returns a boolean if a field has been set.

### SetSkipApprovalNil

`func (o *PKIResponses) SetSkipApprovalNil(b bool)`

 SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil

### UnsetSkipApproval
`func (o *PKIResponses) UnsetSkipApproval()`

UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
### GetCaName

`func (o *PKIResponses) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *PKIResponses) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *PKIResponses) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *PKIResponses) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *PKIResponses) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *PKIResponses) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *PKIResponses) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *PKIResponses) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *PKIResponses) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetCertLifetime

`func (o *PKIResponses) GetCertLifetime() string`

GetCertLifetime returns the CertLifetime field if non-nil, zero value otherwise.

### GetCertLifetimeOk

`func (o *PKIResponses) GetCertLifetimeOk() (*string, bool)`

GetCertLifetimeOk returns a tuple with the CertLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertLifetime

`func (o *PKIResponses) SetCertLifetime(v string)`

SetCertLifetime sets CertLifetime field to given value.

### HasCertLifetime

`func (o *PKIResponses) HasCertLifetime() bool`

HasCertLifetime returns a boolean if a field has been set.

### SetCertLifetimeNil

`func (o *PKIResponses) SetCertLifetimeNil(b bool)`

 SetCertLifetimeNil sets the value for CertLifetime to be an explicit nil

### UnsetCertLifetime
`func (o *PKIResponses) UnsetCertLifetime()`

UnsetCertLifetime ensures that no value is present for CertLifetime, not even an explicit nil
### GetCertType

`func (o *PKIResponses) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *PKIResponses) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *PKIResponses) SetCertType(v string)`

SetCertType sets CertType field to given value.


### SetCertTypeNil

`func (o *PKIResponses) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *PKIResponses) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetClientId

`func (o *PKIResponses) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PKIResponses) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PKIResponses) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PKIResponses) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *PKIResponses) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *PKIResponses) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetRequesterDefaultMail

`func (o *PKIResponses) GetRequesterDefaultMail() string`

GetRequesterDefaultMail returns the RequesterDefaultMail field if non-nil, zero value otherwise.

### GetRequesterDefaultMailOk

`func (o *PKIResponses) GetRequesterDefaultMailOk() (*string, bool)`

GetRequesterDefaultMailOk returns a tuple with the RequesterDefaultMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterDefaultMail

`func (o *PKIResponses) SetRequesterDefaultMail(v string)`

SetRequesterDefaultMail sets RequesterDefaultMail field to given value.


### GetRequesterName

`func (o *PKIResponses) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *PKIResponses) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *PKIResponses) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *PKIResponses) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *PKIResponses) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *PKIResponses) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetRequesterPhone

`func (o *PKIResponses) GetRequesterPhone() string`

GetRequesterPhone returns the RequesterPhone field if non-nil, zero value otherwise.

### GetRequesterPhoneOk

`func (o *PKIResponses) GetRequesterPhoneOk() (*string, bool)`

GetRequesterPhoneOk returns a tuple with the RequesterPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPhone

`func (o *PKIResponses) SetRequesterPhone(v string)`

SetRequesterPhone sets RequesterPhone field to given value.

### HasRequesterPhone

`func (o *PKIResponses) HasRequesterPhone() bool`

HasRequesterPhone returns a boolean if a field has been set.

### SetRequesterPhoneNil

`func (o *PKIResponses) SetRequesterPhoneNil(b bool)`

 SetRequesterPhoneNil sets the value for RequesterPhone to be an explicit nil

### UnsetRequesterPhone
`func (o *PKIResponses) UnsetRequesterPhone()`

UnsetRequesterPhone ensures that no value is present for RequesterPhone, not even an explicit nil
### GetCaCert

`func (o *PKIResponses) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PKIResponses) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PKIResponses) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PKIResponses) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PKIResponses) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PKIResponses) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCaKey

`func (o *PKIResponses) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *PKIResponses) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *PKIResponses) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *PKIResponses) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *PKIResponses) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *PKIResponses) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCheckPop

`func (o *PKIResponses) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *PKIResponses) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *PKIResponses) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *PKIResponses) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *PKIResponses) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *PKIResponses) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCrlLifetime

`func (o *PKIResponses) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *PKIResponses) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *PKIResponses) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *PKIResponses) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *PKIResponses) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *PKIResponses) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetCrlPath

`func (o *PKIResponses) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *PKIResponses) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *PKIResponses) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *PKIResponses) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *PKIResponses) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *PKIResponses) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrtBackDate

`func (o *PKIResponses) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *PKIResponses) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *PKIResponses) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *PKIResponses) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *PKIResponses) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *PKIResponses) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCrtLifetime

`func (o *PKIResponses) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *PKIResponses) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *PKIResponses) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *PKIResponses) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *PKIResponses) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *PKIResponses) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCryptoType

`func (o *PKIResponses) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *PKIResponses) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *PKIResponses) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetSignAlg

`func (o *PKIResponses) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *PKIResponses) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *PKIResponses) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *PKIResponses) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *PKIResponses) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *PKIResponses) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetAuthenticationDomainId

`func (o *PKIResponses) GetAuthenticationDomainId() int64`

GetAuthenticationDomainId returns the AuthenticationDomainId field if non-nil, zero value otherwise.

### GetAuthenticationDomainIdOk

`func (o *PKIResponses) GetAuthenticationDomainIdOk() (*int64, bool)`

GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDomainId

`func (o *PKIResponses) SetAuthenticationDomainId(v int64)`

SetAuthenticationDomainId sets AuthenticationDomainId field to given value.


### GetDefaultOwner

`func (o *PKIResponses) GetDefaultOwner() string`

GetDefaultOwner returns the DefaultOwner field if non-nil, zero value otherwise.

### GetDefaultOwnerOk

`func (o *PKIResponses) GetDefaultOwnerOk() (*string, bool)`

GetDefaultOwnerOk returns a tuple with the DefaultOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOwner

`func (o *PKIResponses) SetDefaultOwner(v string)`

SetDefaultOwner sets DefaultOwner field to given value.


### GetDeleteOnRevoke

`func (o *PKIResponses) GetDeleteOnRevoke() bool`

GetDeleteOnRevoke returns the DeleteOnRevoke field if non-nil, zero value otherwise.

### GetDeleteOnRevokeOk

`func (o *PKIResponses) GetDeleteOnRevokeOk() (*bool, bool)`

GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnRevoke

`func (o *PKIResponses) SetDeleteOnRevoke(v bool)`

SetDeleteOnRevoke sets DeleteOnRevoke field to given value.


### GetOwnerGroups

`func (o *PKIResponses) GetOwnerGroups() string`

GetOwnerGroups returns the OwnerGroups field if non-nil, zero value otherwise.

### GetOwnerGroupsOk

`func (o *PKIResponses) GetOwnerGroupsOk() (*string, bool)`

GetOwnerGroupsOk returns a tuple with the OwnerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroups

`func (o *PKIResponses) SetOwnerGroups(v string)`

SetOwnerGroups sets OwnerGroups field to given value.

### HasOwnerGroups

`func (o *PKIResponses) HasOwnerGroups() bool`

HasOwnerGroups returns a boolean if a field has been set.

### SetOwnerGroupsNil

`func (o *PKIResponses) SetOwnerGroupsNil(b bool)`

 SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil

### UnsetOwnerGroups
`func (o *PKIResponses) UnsetOwnerGroups()`

UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
### GetTemplateId

`func (o *PKIResponses) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PKIResponses) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PKIResponses) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetHashAlgorithm

`func (o *PKIResponses) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *PKIResponses) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *PKIResponses) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *PKIResponses) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### SetHashAlgorithmNil

`func (o *PKIResponses) SetHashAlgorithmNil(b bool)`

 SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil

### UnsetHashAlgorithm
`func (o *PKIResponses) UnsetHashAlgorithm()`

UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
### GetCertificateValidity

`func (o *PKIResponses) GetCertificateValidity() int64`

GetCertificateValidity returns the CertificateValidity field if non-nil, zero value otherwise.

### GetCertificateValidityOk

`func (o *PKIResponses) GetCertificateValidityOk() (*int64, bool)`

GetCertificateValidityOk returns a tuple with the CertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidity

`func (o *PKIResponses) SetCertificateValidity(v int64)`

SetCertificateValidity sets CertificateValidity field to given value.

### HasCertificateValidity

`func (o *PKIResponses) HasCertificateValidity() bool`

HasCertificateValidity returns a boolean if a field has been set.

### SetCertificateValidityNil

`func (o *PKIResponses) SetCertificateValidityNil(b bool)`

 SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil

### UnsetCertificateValidity
`func (o *PKIResponses) UnsetCertificateValidity()`

UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
### GetDefaultEmail

`func (o *PKIResponses) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *PKIResponses) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *PKIResponses) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *PKIResponses) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *PKIResponses) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *PKIResponses) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *PKIResponses) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *PKIResponses) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *PKIResponses) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *PKIResponses) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *PKIResponses) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *PKIResponses) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetDomainId

`func (o *PKIResponses) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PKIResponses) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PKIResponses) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetEndpointType

`func (o *PKIResponses) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *PKIResponses) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *PKIResponses) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetSanEmailMap

`func (o *PKIResponses) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *PKIResponses) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *PKIResponses) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *PKIResponses) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *PKIResponses) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *PKIResponses) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetUidMap

`func (o *PKIResponses) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *PKIResponses) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *PKIResponses) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *PKIResponses) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *PKIResponses) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *PKIResponses) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *PKIResponses) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PKIResponses) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PKIResponses) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PKIResponses) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PKIResponses) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PKIResponses) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *PKIResponses) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *PKIResponses) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *PKIResponses) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *PKIResponses) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *PKIResponses) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *PKIResponses) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
### GetEndPointIssuingCA

`func (o *PKIResponses) GetEndPointIssuingCA() string`

GetEndPointIssuingCA returns the EndPointIssuingCA field if non-nil, zero value otherwise.

### GetEndPointIssuingCAOk

`func (o *PKIResponses) GetEndPointIssuingCAOk() (*string, bool)`

GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIssuingCA

`func (o *PKIResponses) SetEndPointIssuingCA(v string)`

SetEndPointIssuingCA sets EndPointIssuingCA field to given value.


### GetFormPorteurName

`func (o *PKIResponses) GetFormPorteurName() string`

GetFormPorteurName returns the FormPorteurName field if non-nil, zero value otherwise.

### GetFormPorteurNameOk

`func (o *PKIResponses) GetFormPorteurNameOk() (*string, bool)`

GetFormPorteurNameOk returns a tuple with the FormPorteurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormPorteurName

`func (o *PKIResponses) SetFormPorteurName(v string)`

SetFormPorteurName sets FormPorteurName field to given value.

### HasFormPorteurName

`func (o *PKIResponses) HasFormPorteurName() bool`

HasFormPorteurName returns a boolean if a field has been set.

### SetFormPorteurNameNil

`func (o *PKIResponses) SetFormPorteurNameNil(b bool)`

 SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil

### UnsetFormPorteurName
`func (o *PKIResponses) UnsetFormPorteurName()`

UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
### GetProfilCle

`func (o *PKIResponses) GetProfilCle() string`

GetProfilCle returns the ProfilCle field if non-nil, zero value otherwise.

### GetProfilCleOk

`func (o *PKIResponses) GetProfilCleOk() (*string, bool)`

GetProfilCleOk returns a tuple with the ProfilCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilCle

`func (o *PKIResponses) SetProfilCle(v string)`

SetProfilCle sets ProfilCle field to given value.


### SetProfilCleNil

`func (o *PKIResponses) SetProfilCleNil(b bool)`

 SetProfilCleNil sets the value for ProfilCle to be an explicit nil

### UnsetProfilCle
`func (o *PKIResponses) UnsetProfilCle()`

UnsetProfilCle ensures that no value is present for ProfilCle, not even an explicit nil
### GetWorkflow

`func (o *PKIResponses) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *PKIResponses) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *PKIResponses) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### SetWorkflowNil

`func (o *PKIResponses) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *PKIResponses) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetCustomerId

`func (o *PKIResponses) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PKIResponses) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PKIResponses) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEnvironment

`func (o *PKIResponses) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PKIResponses) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PKIResponses) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProcedure

`func (o *PKIResponses) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *PKIResponses) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *PKIResponses) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.


### GetCustomerUri

`func (o *PKIResponses) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *PKIResponses) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *PKIResponses) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetMpkiCredentials

`func (o *PKIResponses) GetMpkiCredentials() string`

GetMpkiCredentials returns the MpkiCredentials field if non-nil, zero value otherwise.

### GetMpkiCredentialsOk

`func (o *PKIResponses) GetMpkiCredentialsOk() (*string, bool)`

GetMpkiCredentialsOk returns a tuple with the MpkiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpkiCredentials

`func (o *PKIResponses) SetMpkiCredentials(v string)`

SetMpkiCredentials sets MpkiCredentials field to given value.


### GetProductUuid

`func (o *PKIResponses) GetProductUuid() string`

GetProductUuid returns the ProductUuid field if non-nil, zero value otherwise.

### GetProductUuidOk

`func (o *PKIResponses) GetProductUuidOk() (*string, bool)`

GetProductUuidOk returns a tuple with the ProductUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUuid

`func (o *PKIResponses) SetProductUuid(v string)`

SetProductUuid sets ProductUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


