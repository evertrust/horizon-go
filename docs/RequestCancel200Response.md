# RequestCancel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Workflow** | **string** |  | 
**Template** | [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | 
**Pkcs12** | Pointer to [**NullableSecretString**](SecretString.md) | The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode | [optional] 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password of the challenge. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**Certificate** | [**NullableCertificate**](Certificate.md) | The certificate that was generated for this request. | 
**Id** | **string** | Object internal ID | 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Profile** | **string** | The associated profile name | 
**Dn** | Pointer to **string** | Certificate&#39;s Distinguished Name | [optional] 
**Requester** | Pointer to **NullableString** | The requester&#39;s principal identifier | [optional] 
**Team** | Pointer to **NullableString** | The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them | [optional] 
**Approver** | Pointer to **NullableString** | The approver&#39;s principal identifier | [optional] 
**Contact** | Pointer to **NullableString** | The request&#39;s contact email | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**RegistrationDate** | **int64** | The date the request was created. This is set by the system | 
**LastModificationDate** | **int64** | The date the request was last modified. This is set by the system | 
**ExpirationDate** | Pointer to **int64** | The date the request will expire. This is set by the system | [optional] 
**RemoveAt** | **int64** | The date the requested will be deleted. This is set by the system | 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) | The result of the execution of triggers on this request | [optional] 
**HolderId** | **string** | The computed holderID for this request. This is set by the system based on DN and SANs | 
**GlobalHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the Horizon database | [optional] 
**ProfileHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile | [optional] 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The labels set in this request | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) | The metadata set in this request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewRequestCancel200Response

`func NewRequestCancel200Response(module string, workflow string, template ScepEnrollRequestTemplate, certificate NullableCertificate, id string, status RequestStatus, profile string, registrationDate int64, lastModificationDate int64, removeAt int64, holderId string, ) *RequestCancel200Response`

NewRequestCancel200Response instantiates a new RequestCancel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCancel200ResponseWithDefaults

`func NewRequestCancel200ResponseWithDefaults() *RequestCancel200Response`

NewRequestCancel200ResponseWithDefaults instantiates a new RequestCancel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *RequestCancel200Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestCancel200Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestCancel200Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *RequestCancel200Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestCancel200Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestCancel200Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *RequestCancel200Response) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *RequestCancel200Response) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *RequestCancel200Response) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.


### GetPkcs12

`func (o *RequestCancel200Response) GetPkcs12() SecretString`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *RequestCancel200Response) GetPkcs12Ok() (*SecretString, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *RequestCancel200Response) SetPkcs12(v SecretString)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *RequestCancel200Response) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### SetPkcs12Nil

`func (o *RequestCancel200Response) SetPkcs12Nil(b bool)`

 SetPkcs12Nil sets the value for Pkcs12 to be an explicit nil

### UnsetPkcs12
`func (o *RequestCancel200Response) UnsetPkcs12()`

UnsetPkcs12 ensures that no value is present for Pkcs12, not even an explicit nil
### GetPassword

`func (o *RequestCancel200Response) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestCancel200Response) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestCancel200Response) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestCancel200Response) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RequestCancel200Response) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RequestCancel200Response) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCertificate

`func (o *RequestCancel200Response) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RequestCancel200Response) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RequestCancel200Response) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *RequestCancel200Response) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RequestCancel200Response) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetId

`func (o *RequestCancel200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestCancel200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestCancel200Response) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *RequestCancel200Response) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestCancel200Response) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestCancel200Response) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetProfile

`func (o *RequestCancel200Response) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RequestCancel200Response) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RequestCancel200Response) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetDn

`func (o *RequestCancel200Response) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *RequestCancel200Response) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *RequestCancel200Response) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *RequestCancel200Response) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRequester

`func (o *RequestCancel200Response) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *RequestCancel200Response) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *RequestCancel200Response) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *RequestCancel200Response) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *RequestCancel200Response) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *RequestCancel200Response) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetTeam

`func (o *RequestCancel200Response) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *RequestCancel200Response) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *RequestCancel200Response) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *RequestCancel200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *RequestCancel200Response) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *RequestCancel200Response) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetApprover

`func (o *RequestCancel200Response) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *RequestCancel200Response) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *RequestCancel200Response) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *RequestCancel200Response) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *RequestCancel200Response) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *RequestCancel200Response) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetContact

`func (o *RequestCancel200Response) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RequestCancel200Response) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RequestCancel200Response) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *RequestCancel200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *RequestCancel200Response) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *RequestCancel200Response) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetRequesterComment

`func (o *RequestCancel200Response) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *RequestCancel200Response) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *RequestCancel200Response) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *RequestCancel200Response) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *RequestCancel200Response) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *RequestCancel200Response) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetApproverComment

`func (o *RequestCancel200Response) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *RequestCancel200Response) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *RequestCancel200Response) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *RequestCancel200Response) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *RequestCancel200Response) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *RequestCancel200Response) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetRegistrationDate

`func (o *RequestCancel200Response) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *RequestCancel200Response) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *RequestCancel200Response) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetLastModificationDate

`func (o *RequestCancel200Response) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *RequestCancel200Response) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *RequestCancel200Response) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetExpirationDate

`func (o *RequestCancel200Response) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RequestCancel200Response) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RequestCancel200Response) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RequestCancel200Response) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRemoveAt

`func (o *RequestCancel200Response) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *RequestCancel200Response) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *RequestCancel200Response) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetTriggerResults

`func (o *RequestCancel200Response) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *RequestCancel200Response) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *RequestCancel200Response) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *RequestCancel200Response) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *RequestCancel200Response) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *RequestCancel200Response) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil
### GetHolderId

`func (o *RequestCancel200Response) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *RequestCancel200Response) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *RequestCancel200Response) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetGlobalHolderIdCount

`func (o *RequestCancel200Response) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *RequestCancel200Response) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *RequestCancel200Response) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *RequestCancel200Response) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *RequestCancel200Response) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *RequestCancel200Response) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetProfileHolderIdCount

`func (o *RequestCancel200Response) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *RequestCancel200Response) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *RequestCancel200Response) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *RequestCancel200Response) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *RequestCancel200Response) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *RequestCancel200Response) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetLabels

`func (o *RequestCancel200Response) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RequestCancel200Response) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RequestCancel200Response) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RequestCancel200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *RequestCancel200Response) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RequestCancel200Response) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *RequestCancel200Response) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestCancel200Response) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestCancel200Response) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestCancel200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RequestCancel200Response) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RequestCancel200Response) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDryRun

`func (o *RequestCancel200Response) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RequestCancel200Response) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RequestCancel200Response) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RequestCancel200Response) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *RequestCancel200Response) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *RequestCancel200Response) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


