# PkiConnectorUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** |  | 
**Template** | **string** | Stream&#39;s certificate template to use for enrollment | 
**Ca** | **string** | Stream&#39;s technical name of the CA on which to enroll | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) to use to authenticate on the PKI | 
**AuthenticationCredentials** | **NullableString** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**CaArn** | **string** |  | 
**AccessCredentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used. | [optional] 
**TemplateArn** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**ValidDays** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**SigningHash** | Pointer to **NullableString** |  | [optional] 
**CertificateUsage** | Pointer to **NullableString** |  | [optional] 
**CaPolicyOid** | Pointer to **NullableString** |  | [optional] 
**OfferId** | **string** |  | 
**OrganizationId** | **int64** |  | 
**RevReason** | Pointer to **NullableString** |  | [optional] 
**Profile** | **string** |  | 
**IssuerCADN** | **string** |  | 
**IssuerCACert** | **string** |  | 
**SignerCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to sign on the PKI | 
**EmailMap** | Pointer to **NullableString** |  | [optional] 
**SanDnsMap** | Pointer to **NullableString** |  | [optional] 
**CnMap** | Pointer to **NullableString** |  | [optional] 
**ProfileMap** | Pointer to **NullableString** |  | [optional] 
**IssuerMap** | Pointer to **NullableString** |  | [optional] 
**LegacyCMPStyle** | Pointer to **NullableBool** |  | [optional] 
**BaseUrl** | Pointer to **string** | The base URL of the used digicert instance. | [optional] 
**ProductId** | **string** |  | 
**ApiCredentials** | **string** | Name of the &#x60;api-key&#x60; [credentials](#tag/api.security.credentials) to use to authenticate on the PKI | 
**CaCertId** | Pointer to **NullableString** |  | [optional] 
**SkipApproval** | Pointer to **NullableBool** |  | [optional] 
**CustomConnectorDataMapping** | Pointer to **map[string]string** |  | [optional] 
**CaName** | **string** |  | 
**EeProfile** | Pointer to **NullableString** |  | [optional] 
**CertType** | **NullableString** |  | 
**RequesterDefaultMail** | **string** |  | 
**RequesterName** | Pointer to **NullableString** |  | [optional] 
**RequesterPhone** | Pointer to **NullableString** |  | [optional] 
**CertLifetime** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableInt64** |  | [optional] 
**CaKey** | Pointer to [**NullableSecretString**](SecretString.md) |  | [optional] 
**CaCert** | Pointer to **NullableString** |  | [optional] 
**CrlPath** | Pointer to **NullableString** |  | [optional] 
**CrlLifetime** | Pointer to **NullableString** |  | [optional] 
**SignAlg** | Pointer to **NullableString** |  | [optional] 
**CrtLifetime** | Pointer to **NullableString** |  | [optional] 
**CrtBackDate** | Pointer to **NullableString** |  | [optional] 
**CheckPop** | Pointer to **NullableBool** |  | [optional] 
**CryptoType** | **string** |  | 
**TemplateId** | **int64** |  | 
**DefaultOwner** | **string** |  | 
**AuthenticationDomainId** | **int64** |  | 
**OwnerGroups** | Pointer to **NullableString** |  | [optional] 
**DeleteOnRevoke** | **bool** |  | 
**HashAlgorithm** | Pointer to **NullableString** |  | [optional] 
**EndpointType** | **string** |  | 
**DomainId** | **string** |  | 
**CertificateValidity** | Pointer to **NullableInt64** |  | [optional] 
**DefaultEmail** | Pointer to **NullableString** |  | [optional] 
**DefaultPhone** | Pointer to **NullableString** |  | [optional] 
**SanEmailMap** | Pointer to **NullableString** |  | [optional] 
**UidMap** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZoneLabel** | Pointer to **NullableString** | The name of the label where the zone value is stored on an enrolled certificate | [optional] 
**EnrollmentCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to enroll on the PKI | 
**CaConfig** | **string** |  | 
**Domain** | **string** |  | 
**EndPointIssuingCA** | **string** | Certificate authority of the endpoint | 
**Procedure** | **string** |  | 
**Environment** | **string** | The testing environment will use https://ote-api.nameshield.net endpoint  and the production will use https://api.nameshield.net  | 
**CustomerId** | **string** |  | 
**Workflow** | **NullableString** |  | 
**ProfilCle** | **NullableString** |  | 
**FormPorteurName** | Pointer to **NullableString** |  | [optional] 
**CustomerUri** | **string** |  | 
**AcmeDirectoryUrl** | **string** | The directory url of the ACME endpoint | 

## Methods

### NewPkiConnectorUpdateRequest

`func NewPkiConnectorUpdateRequest(name string, type_ string, endPoint string, template string, ca string, loginCredentials string, authenticationCredentials NullableString, region string, caArn string, offerId string, organizationId int64, profile string, issuerCADN string, issuerCACert string, signerCredentials string, productId string, apiCredentials string, caName string, certType NullableString, requesterDefaultMail string, cryptoType string, templateId int64, defaultOwner string, authenticationDomainId int64, deleteOnRevoke bool, endpointType string, domainId string, enrollmentCredentials string, caConfig string, domain string, endPointIssuingCA string, procedure string, environment string, customerId string, workflow NullableString, profilCle NullableString, customerUri string, acmeDirectoryUrl string, ) *PkiConnectorUpdateRequest`

NewPkiConnectorUpdateRequest instantiates a new PkiConnectorUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConnectorUpdateRequestWithDefaults

`func NewPkiConnectorUpdateRequestWithDefaults() *PkiConnectorUpdateRequest`

NewPkiConnectorUpdateRequestWithDefaults instantiates a new PkiConnectorUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PkiConnectorUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkiConnectorUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkiConnectorUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PkiConnectorUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PkiConnectorUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PkiConnectorUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *PkiConnectorUpdateRequest) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *PkiConnectorUpdateRequest) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *PkiConnectorUpdateRequest) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetTemplate

`func (o *PkiConnectorUpdateRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PkiConnectorUpdateRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PkiConnectorUpdateRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetCa

`func (o *PkiConnectorUpdateRequest) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *PkiConnectorUpdateRequest) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *PkiConnectorUpdateRequest) SetCa(v string)`

SetCa sets Ca field to given value.


### GetLoginCredentials

`func (o *PkiConnectorUpdateRequest) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *PkiConnectorUpdateRequest) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetAuthenticationCredentials

`func (o *PkiConnectorUpdateRequest) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *PkiConnectorUpdateRequest) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### SetAuthenticationCredentialsNil

`func (o *PkiConnectorUpdateRequest) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *PkiConnectorUpdateRequest) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetTimeout

`func (o *PkiConnectorUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PkiConnectorUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PkiConnectorUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PkiConnectorUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PkiConnectorUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PkiConnectorUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *PkiConnectorUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PkiConnectorUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PkiConnectorUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PkiConnectorUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PkiConnectorUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PkiConnectorUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *PkiConnectorUpdateRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *PkiConnectorUpdateRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *PkiConnectorUpdateRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *PkiConnectorUpdateRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *PkiConnectorUpdateRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *PkiConnectorUpdateRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRegion

`func (o *PkiConnectorUpdateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PkiConnectorUpdateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PkiConnectorUpdateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCaArn

`func (o *PkiConnectorUpdateRequest) GetCaArn() string`

GetCaArn returns the CaArn field if non-nil, zero value otherwise.

### GetCaArnOk

`func (o *PkiConnectorUpdateRequest) GetCaArnOk() (*string, bool)`

GetCaArnOk returns a tuple with the CaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaArn

`func (o *PkiConnectorUpdateRequest) SetCaArn(v string)`

SetCaArn sets CaArn field to given value.


### GetAccessCredentials

`func (o *PkiConnectorUpdateRequest) GetAccessCredentials() string`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetAccessCredentialsOk() (*string, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *PkiConnectorUpdateRequest) SetAccessCredentials(v string)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *PkiConnectorUpdateRequest) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### SetAccessCredentialsNil

`func (o *PkiConnectorUpdateRequest) SetAccessCredentialsNil(b bool)`

 SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil

### UnsetAccessCredentials
`func (o *PkiConnectorUpdateRequest) UnsetAccessCredentials()`

UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
### GetTemplateArn

`func (o *PkiConnectorUpdateRequest) GetTemplateArn() string`

GetTemplateArn returns the TemplateArn field if non-nil, zero value otherwise.

### GetTemplateArnOk

`func (o *PkiConnectorUpdateRequest) GetTemplateArnOk() (*string, bool)`

GetTemplateArnOk returns a tuple with the TemplateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArn

`func (o *PkiConnectorUpdateRequest) SetTemplateArn(v string)`

SetTemplateArn sets TemplateArn field to given value.

### HasTemplateArn

`func (o *PkiConnectorUpdateRequest) HasTemplateArn() bool`

HasTemplateArn returns a boolean if a field has been set.

### SetTemplateArnNil

`func (o *PkiConnectorUpdateRequest) SetTemplateArnNil(b bool)`

 SetTemplateArnNil sets the value for TemplateArn to be an explicit nil

### UnsetTemplateArn
`func (o *PkiConnectorUpdateRequest) UnsetTemplateArn()`

UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
### GetRoleArn

`func (o *PkiConnectorUpdateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *PkiConnectorUpdateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *PkiConnectorUpdateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *PkiConnectorUpdateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *PkiConnectorUpdateRequest) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *PkiConnectorUpdateRequest) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetValidDays

`func (o *PkiConnectorUpdateRequest) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *PkiConnectorUpdateRequest) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *PkiConnectorUpdateRequest) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *PkiConnectorUpdateRequest) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *PkiConnectorUpdateRequest) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *PkiConnectorUpdateRequest) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetRetryInterval

`func (o *PkiConnectorUpdateRequest) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *PkiConnectorUpdateRequest) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *PkiConnectorUpdateRequest) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *PkiConnectorUpdateRequest) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *PkiConnectorUpdateRequest) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *PkiConnectorUpdateRequest) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetSigningHash

`func (o *PkiConnectorUpdateRequest) GetSigningHash() string`

GetSigningHash returns the SigningHash field if non-nil, zero value otherwise.

### GetSigningHashOk

`func (o *PkiConnectorUpdateRequest) GetSigningHashOk() (*string, bool)`

GetSigningHashOk returns a tuple with the SigningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningHash

`func (o *PkiConnectorUpdateRequest) SetSigningHash(v string)`

SetSigningHash sets SigningHash field to given value.

### HasSigningHash

`func (o *PkiConnectorUpdateRequest) HasSigningHash() bool`

HasSigningHash returns a boolean if a field has been set.

### SetSigningHashNil

`func (o *PkiConnectorUpdateRequest) SetSigningHashNil(b bool)`

 SetSigningHashNil sets the value for SigningHash to be an explicit nil

### UnsetSigningHash
`func (o *PkiConnectorUpdateRequest) UnsetSigningHash()`

UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
### GetCertificateUsage

`func (o *PkiConnectorUpdateRequest) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *PkiConnectorUpdateRequest) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *PkiConnectorUpdateRequest) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *PkiConnectorUpdateRequest) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *PkiConnectorUpdateRequest) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *PkiConnectorUpdateRequest) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetCaPolicyOid

`func (o *PkiConnectorUpdateRequest) GetCaPolicyOid() string`

GetCaPolicyOid returns the CaPolicyOid field if non-nil, zero value otherwise.

### GetCaPolicyOidOk

`func (o *PkiConnectorUpdateRequest) GetCaPolicyOidOk() (*string, bool)`

GetCaPolicyOidOk returns a tuple with the CaPolicyOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPolicyOid

`func (o *PkiConnectorUpdateRequest) SetCaPolicyOid(v string)`

SetCaPolicyOid sets CaPolicyOid field to given value.

### HasCaPolicyOid

`func (o *PkiConnectorUpdateRequest) HasCaPolicyOid() bool`

HasCaPolicyOid returns a boolean if a field has been set.

### SetCaPolicyOidNil

`func (o *PkiConnectorUpdateRequest) SetCaPolicyOidNil(b bool)`

 SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil

### UnsetCaPolicyOid
`func (o *PkiConnectorUpdateRequest) UnsetCaPolicyOid()`

UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
### GetOfferId

`func (o *PkiConnectorUpdateRequest) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PkiConnectorUpdateRequest) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PkiConnectorUpdateRequest) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOrganizationId

`func (o *PkiConnectorUpdateRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PkiConnectorUpdateRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PkiConnectorUpdateRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetRevReason

`func (o *PkiConnectorUpdateRequest) GetRevReason() string`

GetRevReason returns the RevReason field if non-nil, zero value otherwise.

### GetRevReasonOk

`func (o *PkiConnectorUpdateRequest) GetRevReasonOk() (*string, bool)`

GetRevReasonOk returns a tuple with the RevReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevReason

`func (o *PkiConnectorUpdateRequest) SetRevReason(v string)`

SetRevReason sets RevReason field to given value.

### HasRevReason

`func (o *PkiConnectorUpdateRequest) HasRevReason() bool`

HasRevReason returns a boolean if a field has been set.

### SetRevReasonNil

`func (o *PkiConnectorUpdateRequest) SetRevReasonNil(b bool)`

 SetRevReasonNil sets the value for RevReason to be an explicit nil

### UnsetRevReason
`func (o *PkiConnectorUpdateRequest) UnsetRevReason()`

UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
### GetProfile

`func (o *PkiConnectorUpdateRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PkiConnectorUpdateRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PkiConnectorUpdateRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetIssuerCADN

`func (o *PkiConnectorUpdateRequest) GetIssuerCADN() string`

GetIssuerCADN returns the IssuerCADN field if non-nil, zero value otherwise.

### GetIssuerCADNOk

`func (o *PkiConnectorUpdateRequest) GetIssuerCADNOk() (*string, bool)`

GetIssuerCADNOk returns a tuple with the IssuerCADN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCADN

`func (o *PkiConnectorUpdateRequest) SetIssuerCADN(v string)`

SetIssuerCADN sets IssuerCADN field to given value.


### GetIssuerCACert

`func (o *PkiConnectorUpdateRequest) GetIssuerCACert() string`

GetIssuerCACert returns the IssuerCACert field if non-nil, zero value otherwise.

### GetIssuerCACertOk

`func (o *PkiConnectorUpdateRequest) GetIssuerCACertOk() (*string, bool)`

GetIssuerCACertOk returns a tuple with the IssuerCACert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCACert

`func (o *PkiConnectorUpdateRequest) SetIssuerCACert(v string)`

SetIssuerCACert sets IssuerCACert field to given value.


### GetSignerCredentials

`func (o *PkiConnectorUpdateRequest) GetSignerCredentials() string`

GetSignerCredentials returns the SignerCredentials field if non-nil, zero value otherwise.

### GetSignerCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetSignerCredentialsOk() (*string, bool)`

GetSignerCredentialsOk returns a tuple with the SignerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCredentials

`func (o *PkiConnectorUpdateRequest) SetSignerCredentials(v string)`

SetSignerCredentials sets SignerCredentials field to given value.


### GetEmailMap

`func (o *PkiConnectorUpdateRequest) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *PkiConnectorUpdateRequest) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *PkiConnectorUpdateRequest) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *PkiConnectorUpdateRequest) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *PkiConnectorUpdateRequest) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *PkiConnectorUpdateRequest) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetSanDnsMap

`func (o *PkiConnectorUpdateRequest) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *PkiConnectorUpdateRequest) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *PkiConnectorUpdateRequest) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *PkiConnectorUpdateRequest) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *PkiConnectorUpdateRequest) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *PkiConnectorUpdateRequest) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetCnMap

`func (o *PkiConnectorUpdateRequest) GetCnMap() string`

GetCnMap returns the CnMap field if non-nil, zero value otherwise.

### GetCnMapOk

`func (o *PkiConnectorUpdateRequest) GetCnMapOk() (*string, bool)`

GetCnMapOk returns a tuple with the CnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnMap

`func (o *PkiConnectorUpdateRequest) SetCnMap(v string)`

SetCnMap sets CnMap field to given value.

### HasCnMap

`func (o *PkiConnectorUpdateRequest) HasCnMap() bool`

HasCnMap returns a boolean if a field has been set.

### SetCnMapNil

`func (o *PkiConnectorUpdateRequest) SetCnMapNil(b bool)`

 SetCnMapNil sets the value for CnMap to be an explicit nil

### UnsetCnMap
`func (o *PkiConnectorUpdateRequest) UnsetCnMap()`

UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
### GetProfileMap

`func (o *PkiConnectorUpdateRequest) GetProfileMap() string`

GetProfileMap returns the ProfileMap field if non-nil, zero value otherwise.

### GetProfileMapOk

`func (o *PkiConnectorUpdateRequest) GetProfileMapOk() (*string, bool)`

GetProfileMapOk returns a tuple with the ProfileMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMap

`func (o *PkiConnectorUpdateRequest) SetProfileMap(v string)`

SetProfileMap sets ProfileMap field to given value.

### HasProfileMap

`func (o *PkiConnectorUpdateRequest) HasProfileMap() bool`

HasProfileMap returns a boolean if a field has been set.

### SetProfileMapNil

`func (o *PkiConnectorUpdateRequest) SetProfileMapNil(b bool)`

 SetProfileMapNil sets the value for ProfileMap to be an explicit nil

### UnsetProfileMap
`func (o *PkiConnectorUpdateRequest) UnsetProfileMap()`

UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
### GetIssuerMap

`func (o *PkiConnectorUpdateRequest) GetIssuerMap() string`

GetIssuerMap returns the IssuerMap field if non-nil, zero value otherwise.

### GetIssuerMapOk

`func (o *PkiConnectorUpdateRequest) GetIssuerMapOk() (*string, bool)`

GetIssuerMapOk returns a tuple with the IssuerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMap

`func (o *PkiConnectorUpdateRequest) SetIssuerMap(v string)`

SetIssuerMap sets IssuerMap field to given value.

### HasIssuerMap

`func (o *PkiConnectorUpdateRequest) HasIssuerMap() bool`

HasIssuerMap returns a boolean if a field has been set.

### SetIssuerMapNil

`func (o *PkiConnectorUpdateRequest) SetIssuerMapNil(b bool)`

 SetIssuerMapNil sets the value for IssuerMap to be an explicit nil

### UnsetIssuerMap
`func (o *PkiConnectorUpdateRequest) UnsetIssuerMap()`

UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
### GetLegacyCMPStyle

`func (o *PkiConnectorUpdateRequest) GetLegacyCMPStyle() bool`

GetLegacyCMPStyle returns the LegacyCMPStyle field if non-nil, zero value otherwise.

### GetLegacyCMPStyleOk

`func (o *PkiConnectorUpdateRequest) GetLegacyCMPStyleOk() (*bool, bool)`

GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyCMPStyle

`func (o *PkiConnectorUpdateRequest) SetLegacyCMPStyle(v bool)`

SetLegacyCMPStyle sets LegacyCMPStyle field to given value.

### HasLegacyCMPStyle

`func (o *PkiConnectorUpdateRequest) HasLegacyCMPStyle() bool`

HasLegacyCMPStyle returns a boolean if a field has been set.

### SetLegacyCMPStyleNil

`func (o *PkiConnectorUpdateRequest) SetLegacyCMPStyleNil(b bool)`

 SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil

### UnsetLegacyCMPStyle
`func (o *PkiConnectorUpdateRequest) UnsetLegacyCMPStyle()`

UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
### GetBaseUrl

`func (o *PkiConnectorUpdateRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PkiConnectorUpdateRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PkiConnectorUpdateRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PkiConnectorUpdateRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetProductId

`func (o *PkiConnectorUpdateRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PkiConnectorUpdateRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PkiConnectorUpdateRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetApiCredentials

`func (o *PkiConnectorUpdateRequest) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *PkiConnectorUpdateRequest) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetCaCertId

`func (o *PkiConnectorUpdateRequest) GetCaCertId() string`

GetCaCertId returns the CaCertId field if non-nil, zero value otherwise.

### GetCaCertIdOk

`func (o *PkiConnectorUpdateRequest) GetCaCertIdOk() (*string, bool)`

GetCaCertIdOk returns a tuple with the CaCertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertId

`func (o *PkiConnectorUpdateRequest) SetCaCertId(v string)`

SetCaCertId sets CaCertId field to given value.

### HasCaCertId

`func (o *PkiConnectorUpdateRequest) HasCaCertId() bool`

HasCaCertId returns a boolean if a field has been set.

### SetCaCertIdNil

`func (o *PkiConnectorUpdateRequest) SetCaCertIdNil(b bool)`

 SetCaCertIdNil sets the value for CaCertId to be an explicit nil

### UnsetCaCertId
`func (o *PkiConnectorUpdateRequest) UnsetCaCertId()`

UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
### GetSkipApproval

`func (o *PkiConnectorUpdateRequest) GetSkipApproval() bool`

GetSkipApproval returns the SkipApproval field if non-nil, zero value otherwise.

### GetSkipApprovalOk

`func (o *PkiConnectorUpdateRequest) GetSkipApprovalOk() (*bool, bool)`

GetSkipApprovalOk returns a tuple with the SkipApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipApproval

`func (o *PkiConnectorUpdateRequest) SetSkipApproval(v bool)`

SetSkipApproval sets SkipApproval field to given value.

### HasSkipApproval

`func (o *PkiConnectorUpdateRequest) HasSkipApproval() bool`

HasSkipApproval returns a boolean if a field has been set.

### SetSkipApprovalNil

`func (o *PkiConnectorUpdateRequest) SetSkipApprovalNil(b bool)`

 SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil

### UnsetSkipApproval
`func (o *PkiConnectorUpdateRequest) UnsetSkipApproval()`

UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
### GetCustomConnectorDataMapping

`func (o *PkiConnectorUpdateRequest) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *PkiConnectorUpdateRequest) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *PkiConnectorUpdateRequest) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *PkiConnectorUpdateRequest) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *PkiConnectorUpdateRequest) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *PkiConnectorUpdateRequest) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetCaName

`func (o *PkiConnectorUpdateRequest) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *PkiConnectorUpdateRequest) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *PkiConnectorUpdateRequest) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *PkiConnectorUpdateRequest) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *PkiConnectorUpdateRequest) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *PkiConnectorUpdateRequest) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *PkiConnectorUpdateRequest) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *PkiConnectorUpdateRequest) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *PkiConnectorUpdateRequest) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetCertType

`func (o *PkiConnectorUpdateRequest) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *PkiConnectorUpdateRequest) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *PkiConnectorUpdateRequest) SetCertType(v string)`

SetCertType sets CertType field to given value.


### SetCertTypeNil

`func (o *PkiConnectorUpdateRequest) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *PkiConnectorUpdateRequest) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetRequesterDefaultMail

`func (o *PkiConnectorUpdateRequest) GetRequesterDefaultMail() string`

GetRequesterDefaultMail returns the RequesterDefaultMail field if non-nil, zero value otherwise.

### GetRequesterDefaultMailOk

`func (o *PkiConnectorUpdateRequest) GetRequesterDefaultMailOk() (*string, bool)`

GetRequesterDefaultMailOk returns a tuple with the RequesterDefaultMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterDefaultMail

`func (o *PkiConnectorUpdateRequest) SetRequesterDefaultMail(v string)`

SetRequesterDefaultMail sets RequesterDefaultMail field to given value.


### GetRequesterName

`func (o *PkiConnectorUpdateRequest) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *PkiConnectorUpdateRequest) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *PkiConnectorUpdateRequest) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *PkiConnectorUpdateRequest) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *PkiConnectorUpdateRequest) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *PkiConnectorUpdateRequest) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetRequesterPhone

`func (o *PkiConnectorUpdateRequest) GetRequesterPhone() string`

GetRequesterPhone returns the RequesterPhone field if non-nil, zero value otherwise.

### GetRequesterPhoneOk

`func (o *PkiConnectorUpdateRequest) GetRequesterPhoneOk() (*string, bool)`

GetRequesterPhoneOk returns a tuple with the RequesterPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPhone

`func (o *PkiConnectorUpdateRequest) SetRequesterPhone(v string)`

SetRequesterPhone sets RequesterPhone field to given value.

### HasRequesterPhone

`func (o *PkiConnectorUpdateRequest) HasRequesterPhone() bool`

HasRequesterPhone returns a boolean if a field has been set.

### SetRequesterPhoneNil

`func (o *PkiConnectorUpdateRequest) SetRequesterPhoneNil(b bool)`

 SetRequesterPhoneNil sets the value for RequesterPhone to be an explicit nil

### UnsetRequesterPhone
`func (o *PkiConnectorUpdateRequest) UnsetRequesterPhone()`

UnsetRequesterPhone ensures that no value is present for RequesterPhone, not even an explicit nil
### GetCertLifetime

`func (o *PkiConnectorUpdateRequest) GetCertLifetime() string`

GetCertLifetime returns the CertLifetime field if non-nil, zero value otherwise.

### GetCertLifetimeOk

`func (o *PkiConnectorUpdateRequest) GetCertLifetimeOk() (*string, bool)`

GetCertLifetimeOk returns a tuple with the CertLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertLifetime

`func (o *PkiConnectorUpdateRequest) SetCertLifetime(v string)`

SetCertLifetime sets CertLifetime field to given value.

### HasCertLifetime

`func (o *PkiConnectorUpdateRequest) HasCertLifetime() bool`

HasCertLifetime returns a boolean if a field has been set.

### SetCertLifetimeNil

`func (o *PkiConnectorUpdateRequest) SetCertLifetimeNil(b bool)`

 SetCertLifetimeNil sets the value for CertLifetime to be an explicit nil

### UnsetCertLifetime
`func (o *PkiConnectorUpdateRequest) UnsetCertLifetime()`

UnsetCertLifetime ensures that no value is present for CertLifetime, not even an explicit nil
### GetClientId

`func (o *PkiConnectorUpdateRequest) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PkiConnectorUpdateRequest) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PkiConnectorUpdateRequest) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PkiConnectorUpdateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *PkiConnectorUpdateRequest) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *PkiConnectorUpdateRequest) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetCaKey

`func (o *PkiConnectorUpdateRequest) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *PkiConnectorUpdateRequest) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *PkiConnectorUpdateRequest) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *PkiConnectorUpdateRequest) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *PkiConnectorUpdateRequest) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *PkiConnectorUpdateRequest) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCaCert

`func (o *PkiConnectorUpdateRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PkiConnectorUpdateRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PkiConnectorUpdateRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PkiConnectorUpdateRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PkiConnectorUpdateRequest) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PkiConnectorUpdateRequest) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCrlPath

`func (o *PkiConnectorUpdateRequest) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *PkiConnectorUpdateRequest) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *PkiConnectorUpdateRequest) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *PkiConnectorUpdateRequest) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *PkiConnectorUpdateRequest) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *PkiConnectorUpdateRequest) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrlLifetime

`func (o *PkiConnectorUpdateRequest) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *PkiConnectorUpdateRequest) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *PkiConnectorUpdateRequest) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *PkiConnectorUpdateRequest) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *PkiConnectorUpdateRequest) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *PkiConnectorUpdateRequest) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetSignAlg

`func (o *PkiConnectorUpdateRequest) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *PkiConnectorUpdateRequest) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *PkiConnectorUpdateRequest) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *PkiConnectorUpdateRequest) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *PkiConnectorUpdateRequest) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *PkiConnectorUpdateRequest) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetCrtLifetime

`func (o *PkiConnectorUpdateRequest) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *PkiConnectorUpdateRequest) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *PkiConnectorUpdateRequest) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *PkiConnectorUpdateRequest) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *PkiConnectorUpdateRequest) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *PkiConnectorUpdateRequest) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCrtBackDate

`func (o *PkiConnectorUpdateRequest) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *PkiConnectorUpdateRequest) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *PkiConnectorUpdateRequest) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *PkiConnectorUpdateRequest) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *PkiConnectorUpdateRequest) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *PkiConnectorUpdateRequest) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCheckPop

`func (o *PkiConnectorUpdateRequest) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *PkiConnectorUpdateRequest) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *PkiConnectorUpdateRequest) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *PkiConnectorUpdateRequest) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *PkiConnectorUpdateRequest) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *PkiConnectorUpdateRequest) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCryptoType

`func (o *PkiConnectorUpdateRequest) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *PkiConnectorUpdateRequest) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *PkiConnectorUpdateRequest) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetTemplateId

`func (o *PkiConnectorUpdateRequest) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PkiConnectorUpdateRequest) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PkiConnectorUpdateRequest) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetDefaultOwner

`func (o *PkiConnectorUpdateRequest) GetDefaultOwner() string`

GetDefaultOwner returns the DefaultOwner field if non-nil, zero value otherwise.

### GetDefaultOwnerOk

`func (o *PkiConnectorUpdateRequest) GetDefaultOwnerOk() (*string, bool)`

GetDefaultOwnerOk returns a tuple with the DefaultOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOwner

`func (o *PkiConnectorUpdateRequest) SetDefaultOwner(v string)`

SetDefaultOwner sets DefaultOwner field to given value.


### GetAuthenticationDomainId

`func (o *PkiConnectorUpdateRequest) GetAuthenticationDomainId() int64`

GetAuthenticationDomainId returns the AuthenticationDomainId field if non-nil, zero value otherwise.

### GetAuthenticationDomainIdOk

`func (o *PkiConnectorUpdateRequest) GetAuthenticationDomainIdOk() (*int64, bool)`

GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDomainId

`func (o *PkiConnectorUpdateRequest) SetAuthenticationDomainId(v int64)`

SetAuthenticationDomainId sets AuthenticationDomainId field to given value.


### GetOwnerGroups

`func (o *PkiConnectorUpdateRequest) GetOwnerGroups() string`

GetOwnerGroups returns the OwnerGroups field if non-nil, zero value otherwise.

### GetOwnerGroupsOk

`func (o *PkiConnectorUpdateRequest) GetOwnerGroupsOk() (*string, bool)`

GetOwnerGroupsOk returns a tuple with the OwnerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroups

`func (o *PkiConnectorUpdateRequest) SetOwnerGroups(v string)`

SetOwnerGroups sets OwnerGroups field to given value.

### HasOwnerGroups

`func (o *PkiConnectorUpdateRequest) HasOwnerGroups() bool`

HasOwnerGroups returns a boolean if a field has been set.

### SetOwnerGroupsNil

`func (o *PkiConnectorUpdateRequest) SetOwnerGroupsNil(b bool)`

 SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil

### UnsetOwnerGroups
`func (o *PkiConnectorUpdateRequest) UnsetOwnerGroups()`

UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
### GetDeleteOnRevoke

`func (o *PkiConnectorUpdateRequest) GetDeleteOnRevoke() bool`

GetDeleteOnRevoke returns the DeleteOnRevoke field if non-nil, zero value otherwise.

### GetDeleteOnRevokeOk

`func (o *PkiConnectorUpdateRequest) GetDeleteOnRevokeOk() (*bool, bool)`

GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnRevoke

`func (o *PkiConnectorUpdateRequest) SetDeleteOnRevoke(v bool)`

SetDeleteOnRevoke sets DeleteOnRevoke field to given value.


### GetHashAlgorithm

`func (o *PkiConnectorUpdateRequest) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *PkiConnectorUpdateRequest) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *PkiConnectorUpdateRequest) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *PkiConnectorUpdateRequest) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### SetHashAlgorithmNil

`func (o *PkiConnectorUpdateRequest) SetHashAlgorithmNil(b bool)`

 SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil

### UnsetHashAlgorithm
`func (o *PkiConnectorUpdateRequest) UnsetHashAlgorithm()`

UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
### GetEndpointType

`func (o *PkiConnectorUpdateRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *PkiConnectorUpdateRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *PkiConnectorUpdateRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetDomainId

`func (o *PkiConnectorUpdateRequest) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PkiConnectorUpdateRequest) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PkiConnectorUpdateRequest) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetCertificateValidity

`func (o *PkiConnectorUpdateRequest) GetCertificateValidity() int64`

GetCertificateValidity returns the CertificateValidity field if non-nil, zero value otherwise.

### GetCertificateValidityOk

`func (o *PkiConnectorUpdateRequest) GetCertificateValidityOk() (*int64, bool)`

GetCertificateValidityOk returns a tuple with the CertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidity

`func (o *PkiConnectorUpdateRequest) SetCertificateValidity(v int64)`

SetCertificateValidity sets CertificateValidity field to given value.

### HasCertificateValidity

`func (o *PkiConnectorUpdateRequest) HasCertificateValidity() bool`

HasCertificateValidity returns a boolean if a field has been set.

### SetCertificateValidityNil

`func (o *PkiConnectorUpdateRequest) SetCertificateValidityNil(b bool)`

 SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil

### UnsetCertificateValidity
`func (o *PkiConnectorUpdateRequest) UnsetCertificateValidity()`

UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
### GetDefaultEmail

`func (o *PkiConnectorUpdateRequest) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *PkiConnectorUpdateRequest) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *PkiConnectorUpdateRequest) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *PkiConnectorUpdateRequest) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *PkiConnectorUpdateRequest) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *PkiConnectorUpdateRequest) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *PkiConnectorUpdateRequest) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *PkiConnectorUpdateRequest) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *PkiConnectorUpdateRequest) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *PkiConnectorUpdateRequest) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *PkiConnectorUpdateRequest) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *PkiConnectorUpdateRequest) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetSanEmailMap

`func (o *PkiConnectorUpdateRequest) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *PkiConnectorUpdateRequest) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *PkiConnectorUpdateRequest) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *PkiConnectorUpdateRequest) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *PkiConnectorUpdateRequest) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *PkiConnectorUpdateRequest) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetUidMap

`func (o *PkiConnectorUpdateRequest) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *PkiConnectorUpdateRequest) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *PkiConnectorUpdateRequest) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *PkiConnectorUpdateRequest) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *PkiConnectorUpdateRequest) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *PkiConnectorUpdateRequest) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *PkiConnectorUpdateRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PkiConnectorUpdateRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PkiConnectorUpdateRequest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PkiConnectorUpdateRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PkiConnectorUpdateRequest) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PkiConnectorUpdateRequest) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *PkiConnectorUpdateRequest) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *PkiConnectorUpdateRequest) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *PkiConnectorUpdateRequest) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *PkiConnectorUpdateRequest) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *PkiConnectorUpdateRequest) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *PkiConnectorUpdateRequest) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
### GetEnrollmentCredentials

`func (o *PkiConnectorUpdateRequest) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *PkiConnectorUpdateRequest) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *PkiConnectorUpdateRequest) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetCaConfig

`func (o *PkiConnectorUpdateRequest) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *PkiConnectorUpdateRequest) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *PkiConnectorUpdateRequest) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetDomain

`func (o *PkiConnectorUpdateRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PkiConnectorUpdateRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PkiConnectorUpdateRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEndPointIssuingCA

`func (o *PkiConnectorUpdateRequest) GetEndPointIssuingCA() string`

GetEndPointIssuingCA returns the EndPointIssuingCA field if non-nil, zero value otherwise.

### GetEndPointIssuingCAOk

`func (o *PkiConnectorUpdateRequest) GetEndPointIssuingCAOk() (*string, bool)`

GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIssuingCA

`func (o *PkiConnectorUpdateRequest) SetEndPointIssuingCA(v string)`

SetEndPointIssuingCA sets EndPointIssuingCA field to given value.


### GetProcedure

`func (o *PkiConnectorUpdateRequest) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *PkiConnectorUpdateRequest) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *PkiConnectorUpdateRequest) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.


### GetEnvironment

`func (o *PkiConnectorUpdateRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PkiConnectorUpdateRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PkiConnectorUpdateRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetCustomerId

`func (o *PkiConnectorUpdateRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PkiConnectorUpdateRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PkiConnectorUpdateRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetWorkflow

`func (o *PkiConnectorUpdateRequest) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *PkiConnectorUpdateRequest) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *PkiConnectorUpdateRequest) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### SetWorkflowNil

`func (o *PkiConnectorUpdateRequest) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *PkiConnectorUpdateRequest) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetProfilCle

`func (o *PkiConnectorUpdateRequest) GetProfilCle() string`

GetProfilCle returns the ProfilCle field if non-nil, zero value otherwise.

### GetProfilCleOk

`func (o *PkiConnectorUpdateRequest) GetProfilCleOk() (*string, bool)`

GetProfilCleOk returns a tuple with the ProfilCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilCle

`func (o *PkiConnectorUpdateRequest) SetProfilCle(v string)`

SetProfilCle sets ProfilCle field to given value.


### SetProfilCleNil

`func (o *PkiConnectorUpdateRequest) SetProfilCleNil(b bool)`

 SetProfilCleNil sets the value for ProfilCle to be an explicit nil

### UnsetProfilCle
`func (o *PkiConnectorUpdateRequest) UnsetProfilCle()`

UnsetProfilCle ensures that no value is present for ProfilCle, not even an explicit nil
### GetFormPorteurName

`func (o *PkiConnectorUpdateRequest) GetFormPorteurName() string`

GetFormPorteurName returns the FormPorteurName field if non-nil, zero value otherwise.

### GetFormPorteurNameOk

`func (o *PkiConnectorUpdateRequest) GetFormPorteurNameOk() (*string, bool)`

GetFormPorteurNameOk returns a tuple with the FormPorteurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormPorteurName

`func (o *PkiConnectorUpdateRequest) SetFormPorteurName(v string)`

SetFormPorteurName sets FormPorteurName field to given value.

### HasFormPorteurName

`func (o *PkiConnectorUpdateRequest) HasFormPorteurName() bool`

HasFormPorteurName returns a boolean if a field has been set.

### SetFormPorteurNameNil

`func (o *PkiConnectorUpdateRequest) SetFormPorteurNameNil(b bool)`

 SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil

### UnsetFormPorteurName
`func (o *PkiConnectorUpdateRequest) UnsetFormPorteurName()`

UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
### GetCustomerUri

`func (o *PkiConnectorUpdateRequest) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *PkiConnectorUpdateRequest) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *PkiConnectorUpdateRequest) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetAcmeDirectoryUrl

`func (o *PkiConnectorUpdateRequest) GetAcmeDirectoryUrl() string`

GetAcmeDirectoryUrl returns the AcmeDirectoryUrl field if non-nil, zero value otherwise.

### GetAcmeDirectoryUrlOk

`func (o *PkiConnectorUpdateRequest) GetAcmeDirectoryUrlOk() (*string, bool)`

GetAcmeDirectoryUrlOk returns a tuple with the AcmeDirectoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUrl

`func (o *PkiConnectorUpdateRequest) SetAcmeDirectoryUrl(v string)`

SetAcmeDirectoryUrl sets AcmeDirectoryUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


