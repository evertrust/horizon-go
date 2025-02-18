# PkiConnectorList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
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

### NewPkiConnectorList200ResponseInner

`func NewPkiConnectorList200ResponseInner(id string, name string, type_ string, endPoint string, template string, ca string, loginCredentials string, authenticationCredentials NullableString, region string, caArn string, offerId string, organizationId int64, profile string, issuerCADN string, issuerCACert string, signerCredentials string, productId string, apiCredentials string, caName string, certType NullableString, requesterDefaultMail string, cryptoType string, templateId int64, defaultOwner string, authenticationDomainId int64, deleteOnRevoke bool, endpointType string, domainId string, enrollmentCredentials string, caConfig string, domain string, endPointIssuingCA string, procedure string, environment string, customerId string, workflow NullableString, profilCle NullableString, customerUri string, acmeDirectoryUrl string, ) *PkiConnectorList200ResponseInner`

NewPkiConnectorList200ResponseInner instantiates a new PkiConnectorList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConnectorList200ResponseInnerWithDefaults

`func NewPkiConnectorList200ResponseInnerWithDefaults() *PkiConnectorList200ResponseInner`

NewPkiConnectorList200ResponseInnerWithDefaults instantiates a new PkiConnectorList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PkiConnectorList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkiConnectorList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkiConnectorList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PkiConnectorList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkiConnectorList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkiConnectorList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PkiConnectorList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PkiConnectorList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PkiConnectorList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *PkiConnectorList200ResponseInner) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *PkiConnectorList200ResponseInner) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *PkiConnectorList200ResponseInner) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetTemplate

`func (o *PkiConnectorList200ResponseInner) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PkiConnectorList200ResponseInner) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PkiConnectorList200ResponseInner) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetCa

`func (o *PkiConnectorList200ResponseInner) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *PkiConnectorList200ResponseInner) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *PkiConnectorList200ResponseInner) SetCa(v string)`

SetCa sets Ca field to given value.


### GetLoginCredentials

`func (o *PkiConnectorList200ResponseInner) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *PkiConnectorList200ResponseInner) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetAuthenticationCredentials

`func (o *PkiConnectorList200ResponseInner) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *PkiConnectorList200ResponseInner) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### SetAuthenticationCredentialsNil

`func (o *PkiConnectorList200ResponseInner) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *PkiConnectorList200ResponseInner) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetTimeout

`func (o *PkiConnectorList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PkiConnectorList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PkiConnectorList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PkiConnectorList200ResponseInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PkiConnectorList200ResponseInner) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PkiConnectorList200ResponseInner) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *PkiConnectorList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PkiConnectorList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PkiConnectorList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PkiConnectorList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PkiConnectorList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PkiConnectorList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *PkiConnectorList200ResponseInner) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *PkiConnectorList200ResponseInner) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *PkiConnectorList200ResponseInner) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *PkiConnectorList200ResponseInner) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *PkiConnectorList200ResponseInner) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *PkiConnectorList200ResponseInner) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *PkiConnectorList200ResponseInner) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PkiConnectorList200ResponseInner) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PkiConnectorList200ResponseInner) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PkiConnectorList200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PkiConnectorList200ResponseInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PkiConnectorList200ResponseInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRegion

`func (o *PkiConnectorList200ResponseInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PkiConnectorList200ResponseInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PkiConnectorList200ResponseInner) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCaArn

`func (o *PkiConnectorList200ResponseInner) GetCaArn() string`

GetCaArn returns the CaArn field if non-nil, zero value otherwise.

### GetCaArnOk

`func (o *PkiConnectorList200ResponseInner) GetCaArnOk() (*string, bool)`

GetCaArnOk returns a tuple with the CaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaArn

`func (o *PkiConnectorList200ResponseInner) SetCaArn(v string)`

SetCaArn sets CaArn field to given value.


### GetAccessCredentials

`func (o *PkiConnectorList200ResponseInner) GetAccessCredentials() string`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetAccessCredentialsOk() (*string, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *PkiConnectorList200ResponseInner) SetAccessCredentials(v string)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *PkiConnectorList200ResponseInner) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### SetAccessCredentialsNil

`func (o *PkiConnectorList200ResponseInner) SetAccessCredentialsNil(b bool)`

 SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil

### UnsetAccessCredentials
`func (o *PkiConnectorList200ResponseInner) UnsetAccessCredentials()`

UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
### GetTemplateArn

`func (o *PkiConnectorList200ResponseInner) GetTemplateArn() string`

GetTemplateArn returns the TemplateArn field if non-nil, zero value otherwise.

### GetTemplateArnOk

`func (o *PkiConnectorList200ResponseInner) GetTemplateArnOk() (*string, bool)`

GetTemplateArnOk returns a tuple with the TemplateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArn

`func (o *PkiConnectorList200ResponseInner) SetTemplateArn(v string)`

SetTemplateArn sets TemplateArn field to given value.

### HasTemplateArn

`func (o *PkiConnectorList200ResponseInner) HasTemplateArn() bool`

HasTemplateArn returns a boolean if a field has been set.

### SetTemplateArnNil

`func (o *PkiConnectorList200ResponseInner) SetTemplateArnNil(b bool)`

 SetTemplateArnNil sets the value for TemplateArn to be an explicit nil

### UnsetTemplateArn
`func (o *PkiConnectorList200ResponseInner) UnsetTemplateArn()`

UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
### GetRoleArn

`func (o *PkiConnectorList200ResponseInner) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *PkiConnectorList200ResponseInner) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *PkiConnectorList200ResponseInner) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *PkiConnectorList200ResponseInner) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *PkiConnectorList200ResponseInner) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *PkiConnectorList200ResponseInner) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetValidDays

`func (o *PkiConnectorList200ResponseInner) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *PkiConnectorList200ResponseInner) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *PkiConnectorList200ResponseInner) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *PkiConnectorList200ResponseInner) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *PkiConnectorList200ResponseInner) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *PkiConnectorList200ResponseInner) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetRetryInterval

`func (o *PkiConnectorList200ResponseInner) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *PkiConnectorList200ResponseInner) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *PkiConnectorList200ResponseInner) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *PkiConnectorList200ResponseInner) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *PkiConnectorList200ResponseInner) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *PkiConnectorList200ResponseInner) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetSigningHash

`func (o *PkiConnectorList200ResponseInner) GetSigningHash() string`

GetSigningHash returns the SigningHash field if non-nil, zero value otherwise.

### GetSigningHashOk

`func (o *PkiConnectorList200ResponseInner) GetSigningHashOk() (*string, bool)`

GetSigningHashOk returns a tuple with the SigningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningHash

`func (o *PkiConnectorList200ResponseInner) SetSigningHash(v string)`

SetSigningHash sets SigningHash field to given value.

### HasSigningHash

`func (o *PkiConnectorList200ResponseInner) HasSigningHash() bool`

HasSigningHash returns a boolean if a field has been set.

### SetSigningHashNil

`func (o *PkiConnectorList200ResponseInner) SetSigningHashNil(b bool)`

 SetSigningHashNil sets the value for SigningHash to be an explicit nil

### UnsetSigningHash
`func (o *PkiConnectorList200ResponseInner) UnsetSigningHash()`

UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
### GetCertificateUsage

`func (o *PkiConnectorList200ResponseInner) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *PkiConnectorList200ResponseInner) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *PkiConnectorList200ResponseInner) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *PkiConnectorList200ResponseInner) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *PkiConnectorList200ResponseInner) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *PkiConnectorList200ResponseInner) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetCaPolicyOid

`func (o *PkiConnectorList200ResponseInner) GetCaPolicyOid() string`

GetCaPolicyOid returns the CaPolicyOid field if non-nil, zero value otherwise.

### GetCaPolicyOidOk

`func (o *PkiConnectorList200ResponseInner) GetCaPolicyOidOk() (*string, bool)`

GetCaPolicyOidOk returns a tuple with the CaPolicyOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPolicyOid

`func (o *PkiConnectorList200ResponseInner) SetCaPolicyOid(v string)`

SetCaPolicyOid sets CaPolicyOid field to given value.

### HasCaPolicyOid

`func (o *PkiConnectorList200ResponseInner) HasCaPolicyOid() bool`

HasCaPolicyOid returns a boolean if a field has been set.

### SetCaPolicyOidNil

`func (o *PkiConnectorList200ResponseInner) SetCaPolicyOidNil(b bool)`

 SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil

### UnsetCaPolicyOid
`func (o *PkiConnectorList200ResponseInner) UnsetCaPolicyOid()`

UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
### GetOfferId

`func (o *PkiConnectorList200ResponseInner) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PkiConnectorList200ResponseInner) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PkiConnectorList200ResponseInner) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOrganizationId

`func (o *PkiConnectorList200ResponseInner) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PkiConnectorList200ResponseInner) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PkiConnectorList200ResponseInner) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetRevReason

`func (o *PkiConnectorList200ResponseInner) GetRevReason() string`

GetRevReason returns the RevReason field if non-nil, zero value otherwise.

### GetRevReasonOk

`func (o *PkiConnectorList200ResponseInner) GetRevReasonOk() (*string, bool)`

GetRevReasonOk returns a tuple with the RevReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevReason

`func (o *PkiConnectorList200ResponseInner) SetRevReason(v string)`

SetRevReason sets RevReason field to given value.

### HasRevReason

`func (o *PkiConnectorList200ResponseInner) HasRevReason() bool`

HasRevReason returns a boolean if a field has been set.

### SetRevReasonNil

`func (o *PkiConnectorList200ResponseInner) SetRevReasonNil(b bool)`

 SetRevReasonNil sets the value for RevReason to be an explicit nil

### UnsetRevReason
`func (o *PkiConnectorList200ResponseInner) UnsetRevReason()`

UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
### GetProfile

`func (o *PkiConnectorList200ResponseInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PkiConnectorList200ResponseInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PkiConnectorList200ResponseInner) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetIssuerCADN

`func (o *PkiConnectorList200ResponseInner) GetIssuerCADN() string`

GetIssuerCADN returns the IssuerCADN field if non-nil, zero value otherwise.

### GetIssuerCADNOk

`func (o *PkiConnectorList200ResponseInner) GetIssuerCADNOk() (*string, bool)`

GetIssuerCADNOk returns a tuple with the IssuerCADN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCADN

`func (o *PkiConnectorList200ResponseInner) SetIssuerCADN(v string)`

SetIssuerCADN sets IssuerCADN field to given value.


### GetIssuerCACert

`func (o *PkiConnectorList200ResponseInner) GetIssuerCACert() string`

GetIssuerCACert returns the IssuerCACert field if non-nil, zero value otherwise.

### GetIssuerCACertOk

`func (o *PkiConnectorList200ResponseInner) GetIssuerCACertOk() (*string, bool)`

GetIssuerCACertOk returns a tuple with the IssuerCACert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCACert

`func (o *PkiConnectorList200ResponseInner) SetIssuerCACert(v string)`

SetIssuerCACert sets IssuerCACert field to given value.


### GetSignerCredentials

`func (o *PkiConnectorList200ResponseInner) GetSignerCredentials() string`

GetSignerCredentials returns the SignerCredentials field if non-nil, zero value otherwise.

### GetSignerCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetSignerCredentialsOk() (*string, bool)`

GetSignerCredentialsOk returns a tuple with the SignerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCredentials

`func (o *PkiConnectorList200ResponseInner) SetSignerCredentials(v string)`

SetSignerCredentials sets SignerCredentials field to given value.


### GetEmailMap

`func (o *PkiConnectorList200ResponseInner) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *PkiConnectorList200ResponseInner) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *PkiConnectorList200ResponseInner) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *PkiConnectorList200ResponseInner) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *PkiConnectorList200ResponseInner) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *PkiConnectorList200ResponseInner) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetSanDnsMap

`func (o *PkiConnectorList200ResponseInner) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *PkiConnectorList200ResponseInner) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *PkiConnectorList200ResponseInner) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *PkiConnectorList200ResponseInner) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *PkiConnectorList200ResponseInner) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *PkiConnectorList200ResponseInner) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetCnMap

`func (o *PkiConnectorList200ResponseInner) GetCnMap() string`

GetCnMap returns the CnMap field if non-nil, zero value otherwise.

### GetCnMapOk

`func (o *PkiConnectorList200ResponseInner) GetCnMapOk() (*string, bool)`

GetCnMapOk returns a tuple with the CnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnMap

`func (o *PkiConnectorList200ResponseInner) SetCnMap(v string)`

SetCnMap sets CnMap field to given value.

### HasCnMap

`func (o *PkiConnectorList200ResponseInner) HasCnMap() bool`

HasCnMap returns a boolean if a field has been set.

### SetCnMapNil

`func (o *PkiConnectorList200ResponseInner) SetCnMapNil(b bool)`

 SetCnMapNil sets the value for CnMap to be an explicit nil

### UnsetCnMap
`func (o *PkiConnectorList200ResponseInner) UnsetCnMap()`

UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
### GetProfileMap

`func (o *PkiConnectorList200ResponseInner) GetProfileMap() string`

GetProfileMap returns the ProfileMap field if non-nil, zero value otherwise.

### GetProfileMapOk

`func (o *PkiConnectorList200ResponseInner) GetProfileMapOk() (*string, bool)`

GetProfileMapOk returns a tuple with the ProfileMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMap

`func (o *PkiConnectorList200ResponseInner) SetProfileMap(v string)`

SetProfileMap sets ProfileMap field to given value.

### HasProfileMap

`func (o *PkiConnectorList200ResponseInner) HasProfileMap() bool`

HasProfileMap returns a boolean if a field has been set.

### SetProfileMapNil

`func (o *PkiConnectorList200ResponseInner) SetProfileMapNil(b bool)`

 SetProfileMapNil sets the value for ProfileMap to be an explicit nil

### UnsetProfileMap
`func (o *PkiConnectorList200ResponseInner) UnsetProfileMap()`

UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
### GetIssuerMap

`func (o *PkiConnectorList200ResponseInner) GetIssuerMap() string`

GetIssuerMap returns the IssuerMap field if non-nil, zero value otherwise.

### GetIssuerMapOk

`func (o *PkiConnectorList200ResponseInner) GetIssuerMapOk() (*string, bool)`

GetIssuerMapOk returns a tuple with the IssuerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMap

`func (o *PkiConnectorList200ResponseInner) SetIssuerMap(v string)`

SetIssuerMap sets IssuerMap field to given value.

### HasIssuerMap

`func (o *PkiConnectorList200ResponseInner) HasIssuerMap() bool`

HasIssuerMap returns a boolean if a field has been set.

### SetIssuerMapNil

`func (o *PkiConnectorList200ResponseInner) SetIssuerMapNil(b bool)`

 SetIssuerMapNil sets the value for IssuerMap to be an explicit nil

### UnsetIssuerMap
`func (o *PkiConnectorList200ResponseInner) UnsetIssuerMap()`

UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
### GetLegacyCMPStyle

`func (o *PkiConnectorList200ResponseInner) GetLegacyCMPStyle() bool`

GetLegacyCMPStyle returns the LegacyCMPStyle field if non-nil, zero value otherwise.

### GetLegacyCMPStyleOk

`func (o *PkiConnectorList200ResponseInner) GetLegacyCMPStyleOk() (*bool, bool)`

GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyCMPStyle

`func (o *PkiConnectorList200ResponseInner) SetLegacyCMPStyle(v bool)`

SetLegacyCMPStyle sets LegacyCMPStyle field to given value.

### HasLegacyCMPStyle

`func (o *PkiConnectorList200ResponseInner) HasLegacyCMPStyle() bool`

HasLegacyCMPStyle returns a boolean if a field has been set.

### SetLegacyCMPStyleNil

`func (o *PkiConnectorList200ResponseInner) SetLegacyCMPStyleNil(b bool)`

 SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil

### UnsetLegacyCMPStyle
`func (o *PkiConnectorList200ResponseInner) UnsetLegacyCMPStyle()`

UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
### GetBaseUrl

`func (o *PkiConnectorList200ResponseInner) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PkiConnectorList200ResponseInner) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PkiConnectorList200ResponseInner) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PkiConnectorList200ResponseInner) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetProductId

`func (o *PkiConnectorList200ResponseInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PkiConnectorList200ResponseInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PkiConnectorList200ResponseInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetApiCredentials

`func (o *PkiConnectorList200ResponseInner) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *PkiConnectorList200ResponseInner) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetCaCertId

`func (o *PkiConnectorList200ResponseInner) GetCaCertId() string`

GetCaCertId returns the CaCertId field if non-nil, zero value otherwise.

### GetCaCertIdOk

`func (o *PkiConnectorList200ResponseInner) GetCaCertIdOk() (*string, bool)`

GetCaCertIdOk returns a tuple with the CaCertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertId

`func (o *PkiConnectorList200ResponseInner) SetCaCertId(v string)`

SetCaCertId sets CaCertId field to given value.

### HasCaCertId

`func (o *PkiConnectorList200ResponseInner) HasCaCertId() bool`

HasCaCertId returns a boolean if a field has been set.

### SetCaCertIdNil

`func (o *PkiConnectorList200ResponseInner) SetCaCertIdNil(b bool)`

 SetCaCertIdNil sets the value for CaCertId to be an explicit nil

### UnsetCaCertId
`func (o *PkiConnectorList200ResponseInner) UnsetCaCertId()`

UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
### GetSkipApproval

`func (o *PkiConnectorList200ResponseInner) GetSkipApproval() bool`

GetSkipApproval returns the SkipApproval field if non-nil, zero value otherwise.

### GetSkipApprovalOk

`func (o *PkiConnectorList200ResponseInner) GetSkipApprovalOk() (*bool, bool)`

GetSkipApprovalOk returns a tuple with the SkipApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipApproval

`func (o *PkiConnectorList200ResponseInner) SetSkipApproval(v bool)`

SetSkipApproval sets SkipApproval field to given value.

### HasSkipApproval

`func (o *PkiConnectorList200ResponseInner) HasSkipApproval() bool`

HasSkipApproval returns a boolean if a field has been set.

### SetSkipApprovalNil

`func (o *PkiConnectorList200ResponseInner) SetSkipApprovalNil(b bool)`

 SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil

### UnsetSkipApproval
`func (o *PkiConnectorList200ResponseInner) UnsetSkipApproval()`

UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
### GetCustomConnectorDataMapping

`func (o *PkiConnectorList200ResponseInner) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *PkiConnectorList200ResponseInner) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *PkiConnectorList200ResponseInner) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *PkiConnectorList200ResponseInner) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *PkiConnectorList200ResponseInner) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *PkiConnectorList200ResponseInner) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetCaName

`func (o *PkiConnectorList200ResponseInner) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *PkiConnectorList200ResponseInner) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *PkiConnectorList200ResponseInner) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *PkiConnectorList200ResponseInner) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *PkiConnectorList200ResponseInner) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *PkiConnectorList200ResponseInner) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *PkiConnectorList200ResponseInner) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *PkiConnectorList200ResponseInner) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *PkiConnectorList200ResponseInner) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetCertType

`func (o *PkiConnectorList200ResponseInner) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *PkiConnectorList200ResponseInner) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *PkiConnectorList200ResponseInner) SetCertType(v string)`

SetCertType sets CertType field to given value.


### SetCertTypeNil

`func (o *PkiConnectorList200ResponseInner) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *PkiConnectorList200ResponseInner) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetRequesterDefaultMail

`func (o *PkiConnectorList200ResponseInner) GetRequesterDefaultMail() string`

GetRequesterDefaultMail returns the RequesterDefaultMail field if non-nil, zero value otherwise.

### GetRequesterDefaultMailOk

`func (o *PkiConnectorList200ResponseInner) GetRequesterDefaultMailOk() (*string, bool)`

GetRequesterDefaultMailOk returns a tuple with the RequesterDefaultMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterDefaultMail

`func (o *PkiConnectorList200ResponseInner) SetRequesterDefaultMail(v string)`

SetRequesterDefaultMail sets RequesterDefaultMail field to given value.


### GetRequesterName

`func (o *PkiConnectorList200ResponseInner) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *PkiConnectorList200ResponseInner) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *PkiConnectorList200ResponseInner) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *PkiConnectorList200ResponseInner) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *PkiConnectorList200ResponseInner) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *PkiConnectorList200ResponseInner) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetRequesterPhone

`func (o *PkiConnectorList200ResponseInner) GetRequesterPhone() string`

GetRequesterPhone returns the RequesterPhone field if non-nil, zero value otherwise.

### GetRequesterPhoneOk

`func (o *PkiConnectorList200ResponseInner) GetRequesterPhoneOk() (*string, bool)`

GetRequesterPhoneOk returns a tuple with the RequesterPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPhone

`func (o *PkiConnectorList200ResponseInner) SetRequesterPhone(v string)`

SetRequesterPhone sets RequesterPhone field to given value.

### HasRequesterPhone

`func (o *PkiConnectorList200ResponseInner) HasRequesterPhone() bool`

HasRequesterPhone returns a boolean if a field has been set.

### SetRequesterPhoneNil

`func (o *PkiConnectorList200ResponseInner) SetRequesterPhoneNil(b bool)`

 SetRequesterPhoneNil sets the value for RequesterPhone to be an explicit nil

### UnsetRequesterPhone
`func (o *PkiConnectorList200ResponseInner) UnsetRequesterPhone()`

UnsetRequesterPhone ensures that no value is present for RequesterPhone, not even an explicit nil
### GetCertLifetime

`func (o *PkiConnectorList200ResponseInner) GetCertLifetime() string`

GetCertLifetime returns the CertLifetime field if non-nil, zero value otherwise.

### GetCertLifetimeOk

`func (o *PkiConnectorList200ResponseInner) GetCertLifetimeOk() (*string, bool)`

GetCertLifetimeOk returns a tuple with the CertLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertLifetime

`func (o *PkiConnectorList200ResponseInner) SetCertLifetime(v string)`

SetCertLifetime sets CertLifetime field to given value.

### HasCertLifetime

`func (o *PkiConnectorList200ResponseInner) HasCertLifetime() bool`

HasCertLifetime returns a boolean if a field has been set.

### SetCertLifetimeNil

`func (o *PkiConnectorList200ResponseInner) SetCertLifetimeNil(b bool)`

 SetCertLifetimeNil sets the value for CertLifetime to be an explicit nil

### UnsetCertLifetime
`func (o *PkiConnectorList200ResponseInner) UnsetCertLifetime()`

UnsetCertLifetime ensures that no value is present for CertLifetime, not even an explicit nil
### GetClientId

`func (o *PkiConnectorList200ResponseInner) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PkiConnectorList200ResponseInner) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PkiConnectorList200ResponseInner) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PkiConnectorList200ResponseInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *PkiConnectorList200ResponseInner) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *PkiConnectorList200ResponseInner) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetCaKey

`func (o *PkiConnectorList200ResponseInner) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *PkiConnectorList200ResponseInner) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *PkiConnectorList200ResponseInner) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *PkiConnectorList200ResponseInner) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *PkiConnectorList200ResponseInner) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *PkiConnectorList200ResponseInner) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCaCert

`func (o *PkiConnectorList200ResponseInner) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PkiConnectorList200ResponseInner) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PkiConnectorList200ResponseInner) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PkiConnectorList200ResponseInner) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PkiConnectorList200ResponseInner) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PkiConnectorList200ResponseInner) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCrlPath

`func (o *PkiConnectorList200ResponseInner) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *PkiConnectorList200ResponseInner) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *PkiConnectorList200ResponseInner) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *PkiConnectorList200ResponseInner) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *PkiConnectorList200ResponseInner) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *PkiConnectorList200ResponseInner) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrlLifetime

`func (o *PkiConnectorList200ResponseInner) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *PkiConnectorList200ResponseInner) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *PkiConnectorList200ResponseInner) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *PkiConnectorList200ResponseInner) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *PkiConnectorList200ResponseInner) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *PkiConnectorList200ResponseInner) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetSignAlg

`func (o *PkiConnectorList200ResponseInner) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *PkiConnectorList200ResponseInner) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *PkiConnectorList200ResponseInner) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *PkiConnectorList200ResponseInner) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *PkiConnectorList200ResponseInner) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *PkiConnectorList200ResponseInner) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetCrtLifetime

`func (o *PkiConnectorList200ResponseInner) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *PkiConnectorList200ResponseInner) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *PkiConnectorList200ResponseInner) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *PkiConnectorList200ResponseInner) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *PkiConnectorList200ResponseInner) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *PkiConnectorList200ResponseInner) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCrtBackDate

`func (o *PkiConnectorList200ResponseInner) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *PkiConnectorList200ResponseInner) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *PkiConnectorList200ResponseInner) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *PkiConnectorList200ResponseInner) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *PkiConnectorList200ResponseInner) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *PkiConnectorList200ResponseInner) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCheckPop

`func (o *PkiConnectorList200ResponseInner) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *PkiConnectorList200ResponseInner) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *PkiConnectorList200ResponseInner) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *PkiConnectorList200ResponseInner) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *PkiConnectorList200ResponseInner) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *PkiConnectorList200ResponseInner) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCryptoType

`func (o *PkiConnectorList200ResponseInner) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *PkiConnectorList200ResponseInner) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *PkiConnectorList200ResponseInner) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetTemplateId

`func (o *PkiConnectorList200ResponseInner) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PkiConnectorList200ResponseInner) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PkiConnectorList200ResponseInner) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetDefaultOwner

`func (o *PkiConnectorList200ResponseInner) GetDefaultOwner() string`

GetDefaultOwner returns the DefaultOwner field if non-nil, zero value otherwise.

### GetDefaultOwnerOk

`func (o *PkiConnectorList200ResponseInner) GetDefaultOwnerOk() (*string, bool)`

GetDefaultOwnerOk returns a tuple with the DefaultOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOwner

`func (o *PkiConnectorList200ResponseInner) SetDefaultOwner(v string)`

SetDefaultOwner sets DefaultOwner field to given value.


### GetAuthenticationDomainId

`func (o *PkiConnectorList200ResponseInner) GetAuthenticationDomainId() int64`

GetAuthenticationDomainId returns the AuthenticationDomainId field if non-nil, zero value otherwise.

### GetAuthenticationDomainIdOk

`func (o *PkiConnectorList200ResponseInner) GetAuthenticationDomainIdOk() (*int64, bool)`

GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDomainId

`func (o *PkiConnectorList200ResponseInner) SetAuthenticationDomainId(v int64)`

SetAuthenticationDomainId sets AuthenticationDomainId field to given value.


### GetOwnerGroups

`func (o *PkiConnectorList200ResponseInner) GetOwnerGroups() string`

GetOwnerGroups returns the OwnerGroups field if non-nil, zero value otherwise.

### GetOwnerGroupsOk

`func (o *PkiConnectorList200ResponseInner) GetOwnerGroupsOk() (*string, bool)`

GetOwnerGroupsOk returns a tuple with the OwnerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroups

`func (o *PkiConnectorList200ResponseInner) SetOwnerGroups(v string)`

SetOwnerGroups sets OwnerGroups field to given value.

### HasOwnerGroups

`func (o *PkiConnectorList200ResponseInner) HasOwnerGroups() bool`

HasOwnerGroups returns a boolean if a field has been set.

### SetOwnerGroupsNil

`func (o *PkiConnectorList200ResponseInner) SetOwnerGroupsNil(b bool)`

 SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil

### UnsetOwnerGroups
`func (o *PkiConnectorList200ResponseInner) UnsetOwnerGroups()`

UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
### GetDeleteOnRevoke

`func (o *PkiConnectorList200ResponseInner) GetDeleteOnRevoke() bool`

GetDeleteOnRevoke returns the DeleteOnRevoke field if non-nil, zero value otherwise.

### GetDeleteOnRevokeOk

`func (o *PkiConnectorList200ResponseInner) GetDeleteOnRevokeOk() (*bool, bool)`

GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnRevoke

`func (o *PkiConnectorList200ResponseInner) SetDeleteOnRevoke(v bool)`

SetDeleteOnRevoke sets DeleteOnRevoke field to given value.


### GetHashAlgorithm

`func (o *PkiConnectorList200ResponseInner) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *PkiConnectorList200ResponseInner) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *PkiConnectorList200ResponseInner) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *PkiConnectorList200ResponseInner) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### SetHashAlgorithmNil

`func (o *PkiConnectorList200ResponseInner) SetHashAlgorithmNil(b bool)`

 SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil

### UnsetHashAlgorithm
`func (o *PkiConnectorList200ResponseInner) UnsetHashAlgorithm()`

UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
### GetEndpointType

`func (o *PkiConnectorList200ResponseInner) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *PkiConnectorList200ResponseInner) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *PkiConnectorList200ResponseInner) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetDomainId

`func (o *PkiConnectorList200ResponseInner) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PkiConnectorList200ResponseInner) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PkiConnectorList200ResponseInner) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetCertificateValidity

`func (o *PkiConnectorList200ResponseInner) GetCertificateValidity() int64`

GetCertificateValidity returns the CertificateValidity field if non-nil, zero value otherwise.

### GetCertificateValidityOk

`func (o *PkiConnectorList200ResponseInner) GetCertificateValidityOk() (*int64, bool)`

GetCertificateValidityOk returns a tuple with the CertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidity

`func (o *PkiConnectorList200ResponseInner) SetCertificateValidity(v int64)`

SetCertificateValidity sets CertificateValidity field to given value.

### HasCertificateValidity

`func (o *PkiConnectorList200ResponseInner) HasCertificateValidity() bool`

HasCertificateValidity returns a boolean if a field has been set.

### SetCertificateValidityNil

`func (o *PkiConnectorList200ResponseInner) SetCertificateValidityNil(b bool)`

 SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil

### UnsetCertificateValidity
`func (o *PkiConnectorList200ResponseInner) UnsetCertificateValidity()`

UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
### GetDefaultEmail

`func (o *PkiConnectorList200ResponseInner) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *PkiConnectorList200ResponseInner) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *PkiConnectorList200ResponseInner) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *PkiConnectorList200ResponseInner) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *PkiConnectorList200ResponseInner) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *PkiConnectorList200ResponseInner) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *PkiConnectorList200ResponseInner) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *PkiConnectorList200ResponseInner) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *PkiConnectorList200ResponseInner) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *PkiConnectorList200ResponseInner) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *PkiConnectorList200ResponseInner) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *PkiConnectorList200ResponseInner) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetSanEmailMap

`func (o *PkiConnectorList200ResponseInner) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *PkiConnectorList200ResponseInner) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *PkiConnectorList200ResponseInner) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *PkiConnectorList200ResponseInner) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *PkiConnectorList200ResponseInner) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *PkiConnectorList200ResponseInner) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetUidMap

`func (o *PkiConnectorList200ResponseInner) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *PkiConnectorList200ResponseInner) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *PkiConnectorList200ResponseInner) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *PkiConnectorList200ResponseInner) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *PkiConnectorList200ResponseInner) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *PkiConnectorList200ResponseInner) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *PkiConnectorList200ResponseInner) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PkiConnectorList200ResponseInner) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PkiConnectorList200ResponseInner) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PkiConnectorList200ResponseInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PkiConnectorList200ResponseInner) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PkiConnectorList200ResponseInner) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *PkiConnectorList200ResponseInner) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *PkiConnectorList200ResponseInner) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *PkiConnectorList200ResponseInner) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *PkiConnectorList200ResponseInner) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *PkiConnectorList200ResponseInner) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *PkiConnectorList200ResponseInner) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
### GetEnrollmentCredentials

`func (o *PkiConnectorList200ResponseInner) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *PkiConnectorList200ResponseInner) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *PkiConnectorList200ResponseInner) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetCaConfig

`func (o *PkiConnectorList200ResponseInner) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *PkiConnectorList200ResponseInner) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *PkiConnectorList200ResponseInner) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetDomain

`func (o *PkiConnectorList200ResponseInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PkiConnectorList200ResponseInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PkiConnectorList200ResponseInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEndPointIssuingCA

`func (o *PkiConnectorList200ResponseInner) GetEndPointIssuingCA() string`

GetEndPointIssuingCA returns the EndPointIssuingCA field if non-nil, zero value otherwise.

### GetEndPointIssuingCAOk

`func (o *PkiConnectorList200ResponseInner) GetEndPointIssuingCAOk() (*string, bool)`

GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIssuingCA

`func (o *PkiConnectorList200ResponseInner) SetEndPointIssuingCA(v string)`

SetEndPointIssuingCA sets EndPointIssuingCA field to given value.


### GetProcedure

`func (o *PkiConnectorList200ResponseInner) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *PkiConnectorList200ResponseInner) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *PkiConnectorList200ResponseInner) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.


### GetEnvironment

`func (o *PkiConnectorList200ResponseInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PkiConnectorList200ResponseInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PkiConnectorList200ResponseInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetCustomerId

`func (o *PkiConnectorList200ResponseInner) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PkiConnectorList200ResponseInner) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PkiConnectorList200ResponseInner) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetWorkflow

`func (o *PkiConnectorList200ResponseInner) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *PkiConnectorList200ResponseInner) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *PkiConnectorList200ResponseInner) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### SetWorkflowNil

`func (o *PkiConnectorList200ResponseInner) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *PkiConnectorList200ResponseInner) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetProfilCle

`func (o *PkiConnectorList200ResponseInner) GetProfilCle() string`

GetProfilCle returns the ProfilCle field if non-nil, zero value otherwise.

### GetProfilCleOk

`func (o *PkiConnectorList200ResponseInner) GetProfilCleOk() (*string, bool)`

GetProfilCleOk returns a tuple with the ProfilCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilCle

`func (o *PkiConnectorList200ResponseInner) SetProfilCle(v string)`

SetProfilCle sets ProfilCle field to given value.


### SetProfilCleNil

`func (o *PkiConnectorList200ResponseInner) SetProfilCleNil(b bool)`

 SetProfilCleNil sets the value for ProfilCle to be an explicit nil

### UnsetProfilCle
`func (o *PkiConnectorList200ResponseInner) UnsetProfilCle()`

UnsetProfilCle ensures that no value is present for ProfilCle, not even an explicit nil
### GetFormPorteurName

`func (o *PkiConnectorList200ResponseInner) GetFormPorteurName() string`

GetFormPorteurName returns the FormPorteurName field if non-nil, zero value otherwise.

### GetFormPorteurNameOk

`func (o *PkiConnectorList200ResponseInner) GetFormPorteurNameOk() (*string, bool)`

GetFormPorteurNameOk returns a tuple with the FormPorteurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormPorteurName

`func (o *PkiConnectorList200ResponseInner) SetFormPorteurName(v string)`

SetFormPorteurName sets FormPorteurName field to given value.

### HasFormPorteurName

`func (o *PkiConnectorList200ResponseInner) HasFormPorteurName() bool`

HasFormPorteurName returns a boolean if a field has been set.

### SetFormPorteurNameNil

`func (o *PkiConnectorList200ResponseInner) SetFormPorteurNameNil(b bool)`

 SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil

### UnsetFormPorteurName
`func (o *PkiConnectorList200ResponseInner) UnsetFormPorteurName()`

UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
### GetCustomerUri

`func (o *PkiConnectorList200ResponseInner) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *PkiConnectorList200ResponseInner) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *PkiConnectorList200ResponseInner) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetAcmeDirectoryUrl

`func (o *PkiConnectorList200ResponseInner) GetAcmeDirectoryUrl() string`

GetAcmeDirectoryUrl returns the AcmeDirectoryUrl field if non-nil, zero value otherwise.

### GetAcmeDirectoryUrlOk

`func (o *PkiConnectorList200ResponseInner) GetAcmeDirectoryUrlOk() (*string, bool)`

GetAcmeDirectoryUrlOk returns a tuple with the AcmeDirectoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUrl

`func (o *PkiConnectorList200ResponseInner) SetAcmeDirectoryUrl(v string)`

SetAcmeDirectoryUrl sets AcmeDirectoryUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


