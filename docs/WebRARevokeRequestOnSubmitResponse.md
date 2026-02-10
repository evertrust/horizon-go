# WebRARevokeRequestOnSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | [**Module**](Module.md) | The module of the certificate revoked. | 
**Workflow** | **string** | What this request will do. For a revocation request, this is always &#x60;revoke&#x60; | 
**Template** | [**WebRARevokeRequestTemplate**](WebRARevokeRequestTemplate.md) | The user-data that was used to revoke the certificate | 
**Certificate** | Pointer to [**NullableCertificate**](Certificate.md) | The certificate that was revoked for this request. This is only available after the request has been approved | [optional] 
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

### NewWebRARevokeRequestOnSubmitResponse

`func NewWebRARevokeRequestOnSubmitResponse(module Module, workflow string, template WebRARevokeRequestTemplate, id string, status RequestStatus, profile string, registrationDate int64, lastModificationDate int64, removeAt int64, holderId string, ) *WebRARevokeRequestOnSubmitResponse`

NewWebRARevokeRequestOnSubmitResponse instantiates a new WebRARevokeRequestOnSubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARevokeRequestOnSubmitResponseWithDefaults

`func NewWebRARevokeRequestOnSubmitResponseWithDefaults() *WebRARevokeRequestOnSubmitResponse`

NewWebRARevokeRequestOnSubmitResponseWithDefaults instantiates a new WebRARevokeRequestOnSubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRARevokeRequestOnSubmitResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARevokeRequestOnSubmitResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *WebRARevokeRequestOnSubmitResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARevokeRequestOnSubmitResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *WebRARevokeRequestOnSubmitResponse) GetTemplate() WebRARevokeRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetTemplateOk() (*WebRARevokeRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARevokeRequestOnSubmitResponse) SetTemplate(v WebRARevokeRequestTemplate)`

SetTemplate sets Template field to given value.


### GetCertificate

`func (o *WebRARevokeRequestOnSubmitResponse) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WebRARevokeRequestOnSubmitResponse) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WebRARevokeRequestOnSubmitResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetId

`func (o *WebRARevokeRequestOnSubmitResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRARevokeRequestOnSubmitResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *WebRARevokeRequestOnSubmitResponse) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebRARevokeRequestOnSubmitResponse) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetProfile

`func (o *WebRARevokeRequestOnSubmitResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARevokeRequestOnSubmitResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetDn

`func (o *WebRARevokeRequestOnSubmitResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *WebRARevokeRequestOnSubmitResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *WebRARevokeRequestOnSubmitResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRequester

`func (o *WebRARevokeRequestOnSubmitResponse) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *WebRARevokeRequestOnSubmitResponse) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *WebRARevokeRequestOnSubmitResponse) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetTeam

`func (o *WebRARevokeRequestOnSubmitResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRARevokeRequestOnSubmitResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRARevokeRequestOnSubmitResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetApprover

`func (o *WebRARevokeRequestOnSubmitResponse) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *WebRARevokeRequestOnSubmitResponse) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *WebRARevokeRequestOnSubmitResponse) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetContact

`func (o *WebRARevokeRequestOnSubmitResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *WebRARevokeRequestOnSubmitResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *WebRARevokeRequestOnSubmitResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetRequesterComment

`func (o *WebRARevokeRequestOnSubmitResponse) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRARevokeRequestOnSubmitResponse) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRARevokeRequestOnSubmitResponse) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetApproverComment

`func (o *WebRARevokeRequestOnSubmitResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRARevokeRequestOnSubmitResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRARevokeRequestOnSubmitResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetRegistrationDate

`func (o *WebRARevokeRequestOnSubmitResponse) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *WebRARevokeRequestOnSubmitResponse) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetLastModificationDate

`func (o *WebRARevokeRequestOnSubmitResponse) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *WebRARevokeRequestOnSubmitResponse) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetExpirationDate

`func (o *WebRARevokeRequestOnSubmitResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *WebRARevokeRequestOnSubmitResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *WebRARevokeRequestOnSubmitResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRemoveAt

`func (o *WebRARevokeRequestOnSubmitResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *WebRARevokeRequestOnSubmitResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetTriggerResults

`func (o *WebRARevokeRequestOnSubmitResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *WebRARevokeRequestOnSubmitResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *WebRARevokeRequestOnSubmitResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil
### GetHolderId

`func (o *WebRARevokeRequestOnSubmitResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *WebRARevokeRequestOnSubmitResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetGlobalHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetProfileHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *WebRARevokeRequestOnSubmitResponse) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetLabels

`func (o *WebRARevokeRequestOnSubmitResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRARevokeRequestOnSubmitResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRARevokeRequestOnSubmitResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRARevokeRequestOnSubmitResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRARevokeRequestOnSubmitResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRARevokeRequestOnSubmitResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDryRun

`func (o *WebRARevokeRequestOnSubmitResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARevokeRequestOnSubmitResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARevokeRequestOnSubmitResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARevokeRequestOnSubmitResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARevokeRequestOnSubmitResponse) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARevokeRequestOnSubmitResponse) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


