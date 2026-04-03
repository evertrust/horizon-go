# WebRAMigrateRequestOnGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**NullableCertificate**](Certificate.md) | The certificate that was updated for this request. This is only available after the request has been approved | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in a migration | [optional] [default to false]
**Module** | [**Module**](Module.md) | The module of the certificate migrated. | 
**Profile** | **interface{}** | The target profile name | 
**Template** | [**WebRAMigrateRequestTemplateResponse**](WebRAMigrateRequestTemplateResponse.md) | The user-data that will be used to generate the certificate | 
**Workflow** | **string** | What this request will do. For a migration request, this is always &#x60;migrate&#x60; | 
**Id** | **string** | Object internal ID | 
**Approver** | Pointer to **NullableString** | The approver&#39;s principal identifier | [optional] 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**Contact** | Pointer to **NullableString** | The request&#39;s contact email | [optional] 
**Dn** | Pointer to **string** | Certificate&#39;s Distinguished Name | [optional] 
**ExpirationDate** | Pointer to **int64** | The date the request will expire. This is set by the system | [optional] 
**GlobalHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the Horizon database | [optional] 
**HolderId** | **string** | The computed holderID for this request. This is set by the system based on DN and SANs | 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The labels set in this request | [optional] 
**LastModificationDate** | **int64** | The date the request was last modified. This is set by the system | 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) | The metadata set in this request | [optional] 
**ProfileHolderIdCount** | Pointer to **NullableInt64** | The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile | [optional] 
**RegistrationDate** | **int64** | The date the request was created. This is set by the system | 
**RemoveAt** | **int64** | The date the requested will be deleted. This is set by the system | 
**Requester** | Pointer to **NullableString** | The requester&#39;s principal identifier | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Team** | Pointer to **NullableString** | The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them | [optional] 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) | The result of the execution of triggers on this request | [optional] 

## Methods

### NewWebRAMigrateRequestOnGetResponse

`func NewWebRAMigrateRequestOnGetResponse(module Module, profile interface{}, template WebRAMigrateRequestTemplateResponse, workflow string, id string, holderId string, lastModificationDate int64, registrationDate int64, removeAt int64, status RequestStatus, ) *WebRAMigrateRequestOnGetResponse`

NewWebRAMigrateRequestOnGetResponse instantiates a new WebRAMigrateRequestOnGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestOnGetResponseWithDefaults

`func NewWebRAMigrateRequestOnGetResponseWithDefaults() *WebRAMigrateRequestOnGetResponse`

NewWebRAMigrateRequestOnGetResponseWithDefaults instantiates a new WebRAMigrateRequestOnGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *WebRAMigrateRequestOnGetResponse) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WebRAMigrateRequestOnGetResponse) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WebRAMigrateRequestOnGetResponse) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *WebRAMigrateRequestOnGetResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *WebRAMigrateRequestOnGetResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *WebRAMigrateRequestOnGetResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDryRun

`func (o *WebRAMigrateRequestOnGetResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAMigrateRequestOnGetResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAMigrateRequestOnGetResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAMigrateRequestOnGetResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAMigrateRequestOnGetResponse) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAMigrateRequestOnGetResponse) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *WebRAMigrateRequestOnGetResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAMigrateRequestOnGetResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAMigrateRequestOnGetResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRAMigrateRequestOnGetResponse) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAMigrateRequestOnGetResponse) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAMigrateRequestOnGetResponse) SetProfile(v interface{})`

SetProfile sets Profile field to given value.


### SetProfileNil

`func (o *WebRAMigrateRequestOnGetResponse) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *WebRAMigrateRequestOnGetResponse) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetTemplate

`func (o *WebRAMigrateRequestOnGetResponse) GetTemplate() WebRAMigrateRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAMigrateRequestOnGetResponse) GetTemplateOk() (*WebRAMigrateRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAMigrateRequestOnGetResponse) SetTemplate(v WebRAMigrateRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRAMigrateRequestOnGetResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAMigrateRequestOnGetResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAMigrateRequestOnGetResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetId

`func (o *WebRAMigrateRequestOnGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRAMigrateRequestOnGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRAMigrateRequestOnGetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApprover

`func (o *WebRAMigrateRequestOnGetResponse) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *WebRAMigrateRequestOnGetResponse) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *WebRAMigrateRequestOnGetResponse) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *WebRAMigrateRequestOnGetResponse) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *WebRAMigrateRequestOnGetResponse) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *WebRAMigrateRequestOnGetResponse) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetApproverComment

`func (o *WebRAMigrateRequestOnGetResponse) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRAMigrateRequestOnGetResponse) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRAMigrateRequestOnGetResponse) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRAMigrateRequestOnGetResponse) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRAMigrateRequestOnGetResponse) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRAMigrateRequestOnGetResponse) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetContact

`func (o *WebRAMigrateRequestOnGetResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *WebRAMigrateRequestOnGetResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *WebRAMigrateRequestOnGetResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *WebRAMigrateRequestOnGetResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *WebRAMigrateRequestOnGetResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *WebRAMigrateRequestOnGetResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDn

`func (o *WebRAMigrateRequestOnGetResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *WebRAMigrateRequestOnGetResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *WebRAMigrateRequestOnGetResponse) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *WebRAMigrateRequestOnGetResponse) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetExpirationDate

`func (o *WebRAMigrateRequestOnGetResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *WebRAMigrateRequestOnGetResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *WebRAMigrateRequestOnGetResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *WebRAMigrateRequestOnGetResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGlobalHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *WebRAMigrateRequestOnGetResponse) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *WebRAMigrateRequestOnGetResponse) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *WebRAMigrateRequestOnGetResponse) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetHolderId

`func (o *WebRAMigrateRequestOnGetResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *WebRAMigrateRequestOnGetResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *WebRAMigrateRequestOnGetResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetLabels

`func (o *WebRAMigrateRequestOnGetResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAMigrateRequestOnGetResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAMigrateRequestOnGetResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAMigrateRequestOnGetResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAMigrateRequestOnGetResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAMigrateRequestOnGetResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLastModificationDate

`func (o *WebRAMigrateRequestOnGetResponse) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *WebRAMigrateRequestOnGetResponse) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *WebRAMigrateRequestOnGetResponse) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetMetadata

`func (o *WebRAMigrateRequestOnGetResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAMigrateRequestOnGetResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAMigrateRequestOnGetResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAMigrateRequestOnGetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAMigrateRequestOnGetResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAMigrateRequestOnGetResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProfileHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *WebRAMigrateRequestOnGetResponse) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *WebRAMigrateRequestOnGetResponse) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *WebRAMigrateRequestOnGetResponse) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *WebRAMigrateRequestOnGetResponse) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetRegistrationDate

`func (o *WebRAMigrateRequestOnGetResponse) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *WebRAMigrateRequestOnGetResponse) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *WebRAMigrateRequestOnGetResponse) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetRemoveAt

`func (o *WebRAMigrateRequestOnGetResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *WebRAMigrateRequestOnGetResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *WebRAMigrateRequestOnGetResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetRequester

`func (o *WebRAMigrateRequestOnGetResponse) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *WebRAMigrateRequestOnGetResponse) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *WebRAMigrateRequestOnGetResponse) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *WebRAMigrateRequestOnGetResponse) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *WebRAMigrateRequestOnGetResponse) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *WebRAMigrateRequestOnGetResponse) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetRequesterComment

`func (o *WebRAMigrateRequestOnGetResponse) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRAMigrateRequestOnGetResponse) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRAMigrateRequestOnGetResponse) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRAMigrateRequestOnGetResponse) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRAMigrateRequestOnGetResponse) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRAMigrateRequestOnGetResponse) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetStatus

`func (o *WebRAMigrateRequestOnGetResponse) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebRAMigrateRequestOnGetResponse) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebRAMigrateRequestOnGetResponse) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *WebRAMigrateRequestOnGetResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAMigrateRequestOnGetResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAMigrateRequestOnGetResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAMigrateRequestOnGetResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAMigrateRequestOnGetResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAMigrateRequestOnGetResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTriggerResults

`func (o *WebRAMigrateRequestOnGetResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *WebRAMigrateRequestOnGetResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *WebRAMigrateRequestOnGetResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *WebRAMigrateRequestOnGetResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *WebRAMigrateRequestOnGetResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *WebRAMigrateRequestOnGetResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


