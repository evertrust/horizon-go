# WebRARenewRequestOnApproveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**NullableCertificate**](Certificate.md) | The certificate that was generated for this request. This is only available after the request has been approved | [optional] 
**Error** | Pointer to **string** | Information about the error that occurred when executing this request | [optional] 
**ExternalId** | Pointer to **map[string]interface{}** | Information about the request on the underlying PKI | [optional] 
**Module** | **string** |  | 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password to decrypt the PKCS12 file. | [optional] 
**Pkcs12** | Pointer to [**NullableSecretString**](SecretString.md) | The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode | [optional] 
**RetryAt** | Pointer to **int64** | Timestamp after which the &#x60;in_progress&#x60; request will be polled again | [optional] 
**RetryCount** | Pointer to **int64** | Number of retries on the polling of the certificate emission | [optional] 
**Template** | Pointer to [**WebRARenewRequestTemplateResponse**](WebRARenewRequestTemplateResponse.md) | The user-data that will be used to generate the certificate | [optional] 
**Workflow** | **string** |  | 
**Id** | **string** | Object internal ID | 
**Approver** | Pointer to **NullableString** | The approver&#39;s principal identifier | [optional] 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**Contact** | Pointer to **NullableString** | The request&#39;s contact email | [optional] 
**Dn** | Pointer to **string** | Certificate&#39;s Distinguished Name | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**ExpirationDate** | Pointer to **int64** | The date the request will expire. This is set by the system | [optional] 
**GlobalHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the Horizon database | [optional] 
**HolderId** | **string** | The computed holderID for this request. This is set by the system based on DN and SANs | 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The labels set in this request | [optional] 
**LastModificationDate** | **int64** | The date the request was last modified. This is set by the system | 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) | The metadata set in this request | [optional] 
**Profile** | **string** | The associated profile name | 
**ProfileHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile | [optional] 
**RegistrationDate** | **int64** | The date the request was created. This is set by the system | 
**RemoveAt** | **int64** | The date the requested will be deleted. This is set by the system | 
**Requester** | Pointer to **NullableString** | The requester&#39;s principal identifier | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Team** | Pointer to **NullableString** | The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them | [optional] 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) | The result of the execution of triggers on this request | [optional] 

## Methods

### NewWebRARenewRequestOnApproveResponse

`func NewWebRARenewRequestOnApproveResponse(module string, workflow string, id string, holderId string, lastModificationDate int64, profile string, registrationDate int64, removeAt int64, status RequestStatus, ) *WebRARenewRequestOnApproveResponse`

NewWebRARenewRequestOnApproveResponse instantiates a new WebRARenewRequestOnApproveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestOnApproveResponseWithDefaults

`func NewWebRARenewRequestOnApproveResponseWithDefaults() *WebRARenewRequestOnApproveResponse`

NewWebRARenewRequestOnApproveResponseWithDefaults instantiates a new WebRARenewRequestOnApproveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *WebRARenewRequestOnApproveResponse) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WebRARenewRequestOnApproveResponse) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WebRARenewRequestOnApproveResponse) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WebRARenewRequestOnApproveResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WebRARenewRequestOnApproveResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WebRARenewRequestOnApproveResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetError

`func (o *WebRARenewRequestOnApproveResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebRARenewRequestOnApproveResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebRARenewRequestOnApproveResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebRARenewRequestOnApproveResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExternalId

`func (o *WebRARenewRequestOnApproveResponse) GetExternalId() map[string]interface{}`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *WebRARenewRequestOnApproveResponse) GetExternalIdOk() (*map[string]interface{}, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *WebRARenewRequestOnApproveResponse) SetExternalId(v map[string]interface{})`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *WebRARenewRequestOnApproveResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetModule

`func (o *WebRARenewRequestOnApproveResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARenewRequestOnApproveResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARenewRequestOnApproveResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetPassword

`func (o *WebRARenewRequestOnApproveResponse) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebRARenewRequestOnApproveResponse) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebRARenewRequestOnApproveResponse) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebRARenewRequestOnApproveResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebRARenewRequestOnApproveResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebRARenewRequestOnApproveResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPkcs12

`func (o *WebRARenewRequestOnApproveResponse) GetPkcs12() SecretString`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *WebRARenewRequestOnApproveResponse) GetPkcs12Ok() (*SecretString, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *WebRARenewRequestOnApproveResponse) SetPkcs12(v SecretString)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *WebRARenewRequestOnApproveResponse) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### SetPkcs12Nil

`func (o *WebRARenewRequestOnApproveResponse) SetPkcs12Nil(b bool)`

 SetPkcs12Nil sets the value for Pkcs12 to be an explicit nil

### UnsetPkcs12
`func (o *WebRARenewRequestOnApproveResponse) UnsetPkcs12()`

UnsetPkcs12 ensures that no value is present for Pkcs12, not even an explicit nil
### GetRetryAt

`func (o *WebRARenewRequestOnApproveResponse) GetRetryAt() int64`

GetRetryAt returns the RetryAt field if non-nil, zero value otherwise.

### GetRetryAtOk

`func (o *WebRARenewRequestOnApproveResponse) GetRetryAtOk() (*int64, bool)`

GetRetryAtOk returns a tuple with the RetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAt

`func (o *WebRARenewRequestOnApproveResponse) SetRetryAt(v int64)`

SetRetryAt sets RetryAt field to given value.

### HasRetryAt

`func (o *WebRARenewRequestOnApproveResponse) HasRetryAt() bool`

HasRetryAt returns a boolean if a field has been set.

### GetRetryCount

`func (o *WebRARenewRequestOnApproveResponse) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WebRARenewRequestOnApproveResponse) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WebRARenewRequestOnApproveResponse) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WebRARenewRequestOnApproveResponse) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetTemplate

`func (o *WebRARenewRequestOnApproveResponse) GetTemplate() WebRARenewRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARenewRequestOnApproveResponse) GetTemplateOk() (*WebRARenewRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARenewRequestOnApproveResponse) SetTemplate(v WebRARenewRequestTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRARenewRequestOnApproveResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRARenewRequestOnApproveResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARenewRequestOnApproveResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARenewRequestOnApproveResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetId

`func (o *WebRARenewRequestOnApproveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRARenewRequestOnApproveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRARenewRequestOnApproveResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApprover

`func (o *WebRARenewRequestOnApproveResponse) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *WebRARenewRequestOnApproveResponse) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *WebRARenewRequestOnApproveResponse) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *WebRARenewRequestOnApproveResponse) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *WebRARenewRequestOnApproveResponse) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *WebRARenewRequestOnApproveResponse) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetApproverComment

`func (o *WebRARenewRequestOnApproveResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRARenewRequestOnApproveResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRARenewRequestOnApproveResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRARenewRequestOnApproveResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRARenewRequestOnApproveResponse) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRARenewRequestOnApproveResponse) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetContact

`func (o *WebRARenewRequestOnApproveResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *WebRARenewRequestOnApproveResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *WebRARenewRequestOnApproveResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *WebRARenewRequestOnApproveResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *WebRARenewRequestOnApproveResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *WebRARenewRequestOnApproveResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDn

`func (o *WebRARenewRequestOnApproveResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *WebRARenewRequestOnApproveResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *WebRARenewRequestOnApproveResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *WebRARenewRequestOnApproveResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRARenewRequestOnApproveResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARenewRequestOnApproveResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARenewRequestOnApproveResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARenewRequestOnApproveResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARenewRequestOnApproveResponse) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARenewRequestOnApproveResponse) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetExpirationDate

`func (o *WebRARenewRequestOnApproveResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *WebRARenewRequestOnApproveResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *WebRARenewRequestOnApproveResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *WebRARenewRequestOnApproveResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGlobalHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *WebRARenewRequestOnApproveResponse) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *WebRARenewRequestOnApproveResponse) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *WebRARenewRequestOnApproveResponse) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetHolderId

`func (o *WebRARenewRequestOnApproveResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *WebRARenewRequestOnApproveResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *WebRARenewRequestOnApproveResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetLabels

`func (o *WebRARenewRequestOnApproveResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRARenewRequestOnApproveResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRARenewRequestOnApproveResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRARenewRequestOnApproveResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRARenewRequestOnApproveResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRARenewRequestOnApproveResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLastModificationDate

`func (o *WebRARenewRequestOnApproveResponse) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *WebRARenewRequestOnApproveResponse) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *WebRARenewRequestOnApproveResponse) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetMetadata

`func (o *WebRARenewRequestOnApproveResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRARenewRequestOnApproveResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRARenewRequestOnApproveResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRARenewRequestOnApproveResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRARenewRequestOnApproveResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRARenewRequestOnApproveResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProfile

`func (o *WebRARenewRequestOnApproveResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARenewRequestOnApproveResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARenewRequestOnApproveResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProfileHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *WebRARenewRequestOnApproveResponse) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *WebRARenewRequestOnApproveResponse) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *WebRARenewRequestOnApproveResponse) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *WebRARenewRequestOnApproveResponse) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetRegistrationDate

`func (o *WebRARenewRequestOnApproveResponse) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *WebRARenewRequestOnApproveResponse) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *WebRARenewRequestOnApproveResponse) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetRemoveAt

`func (o *WebRARenewRequestOnApproveResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *WebRARenewRequestOnApproveResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *WebRARenewRequestOnApproveResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetRequester

`func (o *WebRARenewRequestOnApproveResponse) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *WebRARenewRequestOnApproveResponse) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *WebRARenewRequestOnApproveResponse) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *WebRARenewRequestOnApproveResponse) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *WebRARenewRequestOnApproveResponse) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *WebRARenewRequestOnApproveResponse) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetRequesterComment

`func (o *WebRARenewRequestOnApproveResponse) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRARenewRequestOnApproveResponse) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRARenewRequestOnApproveResponse) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRARenewRequestOnApproveResponse) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRARenewRequestOnApproveResponse) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRARenewRequestOnApproveResponse) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetStatus

`func (o *WebRARenewRequestOnApproveResponse) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebRARenewRequestOnApproveResponse) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebRARenewRequestOnApproveResponse) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *WebRARenewRequestOnApproveResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRARenewRequestOnApproveResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRARenewRequestOnApproveResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRARenewRequestOnApproveResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRARenewRequestOnApproveResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRARenewRequestOnApproveResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTriggerResults

`func (o *WebRARenewRequestOnApproveResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *WebRARenewRequestOnApproveResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *WebRARenewRequestOnApproveResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *WebRARenewRequestOnApproveResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *WebRARenewRequestOnApproveResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *WebRARenewRequestOnApproveResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


