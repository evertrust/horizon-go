# RequestApprove200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**NullableCertificate**](Certificate.md) | The certificate that was generated for this request. | 
**Module** | **string** |  | 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password of the challenge. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**Pkcs12** | Pointer to [**NullableSecretString**](SecretString.md) | The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode | [optional] 
**Template** | [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | 
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

### NewRequestApprove200Response

`func NewRequestApprove200Response(certificate NullableCertificate, module string, template ScepEnrollRequestTemplate, workflow string, id string, holderId string, lastModificationDate int64, profile string, registrationDate int64, removeAt int64, status RequestStatus, ) *RequestApprove200Response`

NewRequestApprove200Response instantiates a new RequestApprove200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApprove200ResponseWithDefaults

`func NewRequestApprove200ResponseWithDefaults() *RequestApprove200Response`

NewRequestApprove200ResponseWithDefaults instantiates a new RequestApprove200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *RequestApprove200Response) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RequestApprove200Response) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RequestApprove200Response) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *RequestApprove200Response) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RequestApprove200Response) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetModule

`func (o *RequestApprove200Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestApprove200Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestApprove200Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetPassword

`func (o *RequestApprove200Response) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestApprove200Response) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestApprove200Response) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestApprove200Response) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RequestApprove200Response) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RequestApprove200Response) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPkcs12

`func (o *RequestApprove200Response) GetPkcs12() SecretString`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *RequestApprove200Response) GetPkcs12Ok() (*SecretString, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *RequestApprove200Response) SetPkcs12(v SecretString)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *RequestApprove200Response) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### SetPkcs12Nil

`func (o *RequestApprove200Response) SetPkcs12Nil(b bool)`

 SetPkcs12Nil sets the value for Pkcs12 to be an explicit nil

### UnsetPkcs12
`func (o *RequestApprove200Response) UnsetPkcs12()`

UnsetPkcs12 ensures that no value is present for Pkcs12, not even an explicit nil
### GetTemplate

`func (o *RequestApprove200Response) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *RequestApprove200Response) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *RequestApprove200Response) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *RequestApprove200Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestApprove200Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestApprove200Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetId

`func (o *RequestApprove200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestApprove200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestApprove200Response) SetId(v string)`

SetId sets Id field to given value.


### GetApprover

`func (o *RequestApprove200Response) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *RequestApprove200Response) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *RequestApprove200Response) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *RequestApprove200Response) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *RequestApprove200Response) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *RequestApprove200Response) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetApproverComment

`func (o *RequestApprove200Response) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *RequestApprove200Response) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *RequestApprove200Response) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *RequestApprove200Response) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *RequestApprove200Response) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *RequestApprove200Response) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetContact

`func (o *RequestApprove200Response) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RequestApprove200Response) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RequestApprove200Response) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *RequestApprove200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *RequestApprove200Response) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *RequestApprove200Response) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDn

`func (o *RequestApprove200Response) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *RequestApprove200Response) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *RequestApprove200Response) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *RequestApprove200Response) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetDryRun

`func (o *RequestApprove200Response) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RequestApprove200Response) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RequestApprove200Response) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RequestApprove200Response) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *RequestApprove200Response) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *RequestApprove200Response) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetExpirationDate

`func (o *RequestApprove200Response) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RequestApprove200Response) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RequestApprove200Response) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RequestApprove200Response) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGlobalHolderIdCount

`func (o *RequestApprove200Response) GetGlobalHolderIdCount() int64`

GetGlobalHolderIdCount returns the GlobalHolderIdCount field if non-nil, zero value otherwise.

### GetGlobalHolderIdCountOk

`func (o *RequestApprove200Response) GetGlobalHolderIdCountOk() (*int64, bool)`

GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHolderIdCount

`func (o *RequestApprove200Response) SetGlobalHolderIdCount(v int64)`

SetGlobalHolderIdCount sets GlobalHolderIdCount field to given value.

### HasGlobalHolderIdCount

`func (o *RequestApprove200Response) HasGlobalHolderIdCount() bool`

HasGlobalHolderIdCount returns a boolean if a field has been set.

### SetGlobalHolderIdCountNil

`func (o *RequestApprove200Response) SetGlobalHolderIdCountNil(b bool)`

 SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil

### UnsetGlobalHolderIdCount
`func (o *RequestApprove200Response) UnsetGlobalHolderIdCount()`

UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
### GetHolderId

`func (o *RequestApprove200Response) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *RequestApprove200Response) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *RequestApprove200Response) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetLabels

`func (o *RequestApprove200Response) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RequestApprove200Response) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RequestApprove200Response) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RequestApprove200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *RequestApprove200Response) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RequestApprove200Response) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLastModificationDate

`func (o *RequestApprove200Response) GetLastModificationDate() int64`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *RequestApprove200Response) GetLastModificationDateOk() (*int64, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *RequestApprove200Response) SetLastModificationDate(v int64)`

SetLastModificationDate sets LastModificationDate field to given value.


### GetMetadata

`func (o *RequestApprove200Response) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestApprove200Response) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestApprove200Response) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestApprove200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RequestApprove200Response) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RequestApprove200Response) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProfile

`func (o *RequestApprove200Response) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RequestApprove200Response) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RequestApprove200Response) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProfileHolderIdCount

`func (o *RequestApprove200Response) GetProfileHolderIdCount() int64`

GetProfileHolderIdCount returns the ProfileHolderIdCount field if non-nil, zero value otherwise.

### GetProfileHolderIdCountOk

`func (o *RequestApprove200Response) GetProfileHolderIdCountOk() (*int64, bool)`

GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileHolderIdCount

`func (o *RequestApprove200Response) SetProfileHolderIdCount(v int64)`

SetProfileHolderIdCount sets ProfileHolderIdCount field to given value.

### HasProfileHolderIdCount

`func (o *RequestApprove200Response) HasProfileHolderIdCount() bool`

HasProfileHolderIdCount returns a boolean if a field has been set.

### SetProfileHolderIdCountNil

`func (o *RequestApprove200Response) SetProfileHolderIdCountNil(b bool)`

 SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil

### UnsetProfileHolderIdCount
`func (o *RequestApprove200Response) UnsetProfileHolderIdCount()`

UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
### GetRegistrationDate

`func (o *RequestApprove200Response) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *RequestApprove200Response) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *RequestApprove200Response) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetRemoveAt

`func (o *RequestApprove200Response) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *RequestApprove200Response) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *RequestApprove200Response) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.


### GetRequester

`func (o *RequestApprove200Response) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *RequestApprove200Response) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *RequestApprove200Response) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *RequestApprove200Response) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *RequestApprove200Response) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *RequestApprove200Response) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetRequesterComment

`func (o *RequestApprove200Response) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *RequestApprove200Response) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *RequestApprove200Response) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *RequestApprove200Response) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *RequestApprove200Response) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *RequestApprove200Response) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetStatus

`func (o *RequestApprove200Response) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestApprove200Response) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestApprove200Response) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *RequestApprove200Response) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *RequestApprove200Response) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *RequestApprove200Response) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *RequestApprove200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *RequestApprove200Response) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *RequestApprove200Response) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTriggerResults

`func (o *RequestApprove200Response) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *RequestApprove200Response) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *RequestApprove200Response) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *RequestApprove200Response) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *RequestApprove200Response) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *RequestApprove200Response) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


