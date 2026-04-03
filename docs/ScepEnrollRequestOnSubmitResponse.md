# ScepEnrollRequestOnSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The DN of the challenge | [optional] 
**Module** | **string** | The module that will be used to process this request. For a SCEP request, this is always &#x60;scep&#x60; | 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password of the challenge. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**Template** | Pointer to [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**Workflow** | **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | 
**Id** | **string** | Object internal ID | 
**Approver** | Pointer to **NullableString** | The approver&#39;s principal identifier | [optional] 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**Contact** | Pointer to **NullableString** | The request&#39;s contact email | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**ExpirationDate** | Pointer to **int64** | The date the request will expire. This is set by the system | [optional] 
**GlobalHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the Horizon database | [optional] 
**HolderId** | Pointer to **string** | The computed holderID for this request. This is set by the system based on DN and SANs | [optional] 
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

### NewScepEnrollRequestOnSubmitResponse

`func NewScepEnrollRequestOnSubmitResponse(module string, workflow string, id string, lastModificationDate int64, profile string, registrationDate int64, removeAt int64, status RequestStatus, ) *ScepEnrollRequestOnSubmitResponse`

NewScepEnrollRequestOnSubmitResponse instantiates a new ScepEnrollRequestOnSubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestOnSubmitResponseWithDefaults

`func NewScepEnrollRequestOnSubmitResponseWithDefaults() *ScepEnrollRequestOnSubmitResponse`

NewScepEnrollRequestOnSubmitResponseWithDefaults instantiates a new ScepEnrollRequestOnSubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *ScepEnrollRequestOnSubmitResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ScepEnrollRequestOnSubmitResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ScepEnrollRequestOnSubmitResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetModule

`func (o *ScepEnrollRequestOnSubmitResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepEnrollRequestOnSubmitResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetPassword

`func (o *ScepEnrollRequestOnSubmitResponse) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ScepEnrollRequestOnSubmitResponse) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ScepEnrollRequestOnSubmitResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTemplate

`func (o *ScepEnrollRequestOnSubmitResponse) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ScepEnrollRequestOnSubmitResponse) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ScepEnrollRequestOnSubmitResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *ScepEnrollRequestOnSubmitResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ScepEnrollRequestOnSubmitResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetId

`func (o *ScepEnrollRequestOnSubmitResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScepEnrollRequestOnSubmitResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApprover

`func (o *ScepEnrollRequestOnSubmitResponse) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ScepEnrollRequestOnSubmitResponse) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *ScepEnrollRequestOnSubmitResponse) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetApproverComment

`func (o *ScepEnrollRequestOnSubmitResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *ScepEnrollRequestOnSubmitResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *ScepEnrollRequestOnSubmitResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetContact

`func (o *ScepEnrollRequestOnSubmitResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ScepEnrollRequestOnSubmitResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ScepEnrollRequestOnSubmitResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDryRun

`func (o *ScepEnrollRequestOnSubmitResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScepEnrollRequestOnSubmitResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScepEnrollRequestOnSubmitResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetExpirationDate

`func (o *ScepEnrollRequestOnSubmitResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ScepEnrollRequestOnSubmitResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ScepEnrollRequestOnSubmitResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGlobalHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetHolderId

`func (o *ScepEnrollRequestOnSubmitResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *ScepEnrollRequestOnSubmitResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.

### HasHolderId

`func (o *ScepEnrollRequestOnSubmitResponse) HasHolderId() bool`

HasHolderId returns a boolean if a field has been set.

### GetLabels

`func (o *ScepEnrollRequestOnSubmitResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScepEnrollRequestOnSubmitResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ScepEnrollRequestOnSubmitResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLastModificationDate

`func (o *ScepEnrollRequestOnSubmitResponse) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *ScepEnrollRequestOnSubmitResponse) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetMetadata

`func (o *ScepEnrollRequestOnSubmitResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ScepEnrollRequestOnSubmitResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ScepEnrollRequestOnSubmitResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProfile

`func (o *ScepEnrollRequestOnSubmitResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScepEnrollRequestOnSubmitResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProfileHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *ScepEnrollRequestOnSubmitResponse) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetRegistrationDate

`func (o *ScepEnrollRequestOnSubmitResponse) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *ScepEnrollRequestOnSubmitResponse) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetRemoveAt

`func (o *ScepEnrollRequestOnSubmitResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *ScepEnrollRequestOnSubmitResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetRequester

`func (o *ScepEnrollRequestOnSubmitResponse) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ScepEnrollRequestOnSubmitResponse) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ScepEnrollRequestOnSubmitResponse) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetRequesterComment

`func (o *ScepEnrollRequestOnSubmitResponse) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *ScepEnrollRequestOnSubmitResponse) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *ScepEnrollRequestOnSubmitResponse) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetStatus

`func (o *ScepEnrollRequestOnSubmitResponse) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScepEnrollRequestOnSubmitResponse) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *ScepEnrollRequestOnSubmitResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ScepEnrollRequestOnSubmitResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ScepEnrollRequestOnSubmitResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTriggerResults

`func (o *ScepEnrollRequestOnSubmitResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *ScepEnrollRequestOnSubmitResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *ScepEnrollRequestOnSubmitResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *ScepEnrollRequestOnSubmitResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *ScepEnrollRequestOnSubmitResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *ScepEnrollRequestOnSubmitResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


