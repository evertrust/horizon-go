# ScepEnrollRequestOnApproveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Workflow** | **string** |  | 
**Dn** | Pointer to **string** | Certificate&#39;s Distinguished Name | [optional] 
**Template** | Pointer to [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password of the challenge. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**Id** | **string** | Object internal ID | 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Profile** | **string** | The associated profile name | 
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
**HolderId** | Pointer to **string** | The computed holderID for this request. This is set by the system based on DN and SANs | [optional] 
**GlobalHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the Horizon database | [optional] 
**ProfileHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile | [optional] 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The labels set in this request | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) | The metadata set in this request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewScepEnrollRequestOnApproveResponse

`func NewScepEnrollRequestOnApproveResponse(module string, workflow string, id string, status RequestStatus, profile string, registrationDate int64, lastModificationDate int64, removeAt int64, ) *ScepEnrollRequestOnApproveResponse`

NewScepEnrollRequestOnApproveResponse instantiates a new ScepEnrollRequestOnApproveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestOnApproveResponseWithDefaults

`func NewScepEnrollRequestOnApproveResponseWithDefaults() *ScepEnrollRequestOnApproveResponse`

NewScepEnrollRequestOnApproveResponseWithDefaults instantiates a new ScepEnrollRequestOnApproveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ScepEnrollRequestOnApproveResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepEnrollRequestOnApproveResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepEnrollRequestOnApproveResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *ScepEnrollRequestOnApproveResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ScepEnrollRequestOnApproveResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ScepEnrollRequestOnApproveResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetDn

`func (o *ScepEnrollRequestOnApproveResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ScepEnrollRequestOnApproveResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ScepEnrollRequestOnApproveResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ScepEnrollRequestOnApproveResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetTemplate

`func (o *ScepEnrollRequestOnApproveResponse) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ScepEnrollRequestOnApproveResponse) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ScepEnrollRequestOnApproveResponse) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ScepEnrollRequestOnApproveResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetPassword

`func (o *ScepEnrollRequestOnApproveResponse) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ScepEnrollRequestOnApproveResponse) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ScepEnrollRequestOnApproveResponse) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ScepEnrollRequestOnApproveResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ScepEnrollRequestOnApproveResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ScepEnrollRequestOnApproveResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetId

`func (o *ScepEnrollRequestOnApproveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScepEnrollRequestOnApproveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScepEnrollRequestOnApproveResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ScepEnrollRequestOnApproveResponse) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScepEnrollRequestOnApproveResponse) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScepEnrollRequestOnApproveResponse) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetProfile

`func (o *ScepEnrollRequestOnApproveResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScepEnrollRequestOnApproveResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScepEnrollRequestOnApproveResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRequester

`func (o *ScepEnrollRequestOnApproveResponse) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ScepEnrollRequestOnApproveResponse) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ScepEnrollRequestOnApproveResponse) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ScepEnrollRequestOnApproveResponse) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *ScepEnrollRequestOnApproveResponse) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *ScepEnrollRequestOnApproveResponse) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetTeam

`func (o *ScepEnrollRequestOnApproveResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ScepEnrollRequestOnApproveResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ScepEnrollRequestOnApproveResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ScepEnrollRequestOnApproveResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ScepEnrollRequestOnApproveResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ScepEnrollRequestOnApproveResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetApprover

`func (o *ScepEnrollRequestOnApproveResponse) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ScepEnrollRequestOnApproveResponse) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ScepEnrollRequestOnApproveResponse) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *ScepEnrollRequestOnApproveResponse) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *ScepEnrollRequestOnApproveResponse) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *ScepEnrollRequestOnApproveResponse) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetContact

`func (o *ScepEnrollRequestOnApproveResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ScepEnrollRequestOnApproveResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ScepEnrollRequestOnApproveResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ScepEnrollRequestOnApproveResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ScepEnrollRequestOnApproveResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ScepEnrollRequestOnApproveResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetRequesterComment

`func (o *ScepEnrollRequestOnApproveResponse) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *ScepEnrollRequestOnApproveResponse) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *ScepEnrollRequestOnApproveResponse) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *ScepEnrollRequestOnApproveResponse) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *ScepEnrollRequestOnApproveResponse) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *ScepEnrollRequestOnApproveResponse) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetApproverComment

`func (o *ScepEnrollRequestOnApproveResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *ScepEnrollRequestOnApproveResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *ScepEnrollRequestOnApproveResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *ScepEnrollRequestOnApproveResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *ScepEnrollRequestOnApproveResponse) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *ScepEnrollRequestOnApproveResponse) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetRegistrationDate

`func (o *ScepEnrollRequestOnApproveResponse) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *ScepEnrollRequestOnApproveResponse) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *ScepEnrollRequestOnApproveResponse) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetLastModificationDate

`func (o *ScepEnrollRequestOnApproveResponse) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *ScepEnrollRequestOnApproveResponse) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *ScepEnrollRequestOnApproveResponse) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetExpirationDate

`func (o *ScepEnrollRequestOnApproveResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ScepEnrollRequestOnApproveResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ScepEnrollRequestOnApproveResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ScepEnrollRequestOnApproveResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRemoveAt

`func (o *ScepEnrollRequestOnApproveResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *ScepEnrollRequestOnApproveResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *ScepEnrollRequestOnApproveResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetTriggerResults

`func (o *ScepEnrollRequestOnApproveResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *ScepEnrollRequestOnApproveResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *ScepEnrollRequestOnApproveResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *ScepEnrollRequestOnApproveResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *ScepEnrollRequestOnApproveResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *ScepEnrollRequestOnApproveResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil
### GetHolderId

`func (o *ScepEnrollRequestOnApproveResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *ScepEnrollRequestOnApproveResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *ScepEnrollRequestOnApproveResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.

### HasHolderId

`func (o *ScepEnrollRequestOnApproveResponse) HasHolderId() bool`

HasHolderId returns a boolean if a field has been set.

### GetGlobalHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *ScepEnrollRequestOnApproveResponse) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *ScepEnrollRequestOnApproveResponse) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *ScepEnrollRequestOnApproveResponse) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetProfileHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *ScepEnrollRequestOnApproveResponse) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *ScepEnrollRequestOnApproveResponse) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *ScepEnrollRequestOnApproveResponse) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *ScepEnrollRequestOnApproveResponse) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetLabels

`func (o *ScepEnrollRequestOnApproveResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScepEnrollRequestOnApproveResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScepEnrollRequestOnApproveResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ScepEnrollRequestOnApproveResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ScepEnrollRequestOnApproveResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ScepEnrollRequestOnApproveResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *ScepEnrollRequestOnApproveResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ScepEnrollRequestOnApproveResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ScepEnrollRequestOnApproveResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ScepEnrollRequestOnApproveResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ScepEnrollRequestOnApproveResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ScepEnrollRequestOnApproveResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDryRun

`func (o *ScepEnrollRequestOnApproveResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScepEnrollRequestOnApproveResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScepEnrollRequestOnApproveResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScepEnrollRequestOnApproveResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *ScepEnrollRequestOnApproveResponse) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *ScepEnrollRequestOnApproveResponse) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


