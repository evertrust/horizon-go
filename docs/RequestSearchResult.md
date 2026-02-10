# RequestSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the request | 
**Module** | Pointer to [**Module**](Module.md) |  | [optional] 
**Workflow** | Pointer to [**Workflow**](Workflow.md) |  | [optional] 
**Status** | Pointer to [**RequestStatus**](RequestStatus.md) |  | [optional] 
**Profile** | Pointer to **string** | Any profile configured for a protocol in Horizon | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**CertificateId** | Pointer to **NullableString** | The id of the certificate in the request | [optional] 
**Certificate** | Pointer to [**NullableCertificate**](Certificate.md) | The certificate associated with the request | [optional] 
**Dn** | Pointer to **string** | Associated certificate&#39;s Distinguished Name | [optional] 
**HolderId** | Pointer to **string** | The computed holderID for this request. This is set by the system based on DN and SANs | [optional] 
**Permissions** | [**RequestPermissions**](RequestPermissions.md) | The permissions of the principal on this request. | 

## Methods

### NewRequestSearchResult

`func NewRequestSearchResult(id string, permissions RequestPermissions, ) *RequestSearchResult`

NewRequestSearchResult instantiates a new RequestSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSearchResultWithDefaults

`func NewRequestSearchResultWithDefaults() *RequestSearchResult`

NewRequestSearchResultWithDefaults instantiates a new RequestSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestSearchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestSearchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestSearchResult) SetId(v string)`

SetId sets Id field to given value.


### GetModule

`func (o *RequestSearchResult) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestSearchResult) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestSearchResult) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *RequestSearchResult) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetWorkflow

`func (o *RequestSearchResult) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestSearchResult) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestSearchResult) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *RequestSearchResult) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetStatus

`func (o *RequestSearchResult) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestSearchResult) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestSearchResult) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestSearchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProfile

`func (o *RequestSearchResult) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RequestSearchResult) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RequestSearchResult) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *RequestSearchResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRequesterComment

`func (o *RequestSearchResult) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *RequestSearchResult) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *RequestSearchResult) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *RequestSearchResult) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *RequestSearchResult) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *RequestSearchResult) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetCertificateId

`func (o *RequestSearchResult) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *RequestSearchResult) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *RequestSearchResult) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *RequestSearchResult) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *RequestSearchResult) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *RequestSearchResult) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificate

`func (o *RequestSearchResult) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RequestSearchResult) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RequestSearchResult) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RequestSearchResult) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *RequestSearchResult) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RequestSearchResult) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDn

`func (o *RequestSearchResult) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *RequestSearchResult) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *RequestSearchResult) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *RequestSearchResult) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHolderId

`func (o *RequestSearchResult) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *RequestSearchResult) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *RequestSearchResult) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.

### HasHolderId

`func (o *RequestSearchResult) HasHolderId() bool`

HasHolderId returns a boolean if a field has been set.

### GetPermissions

`func (o *RequestSearchResult) GetPermissions() RequestPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RequestSearchResult) GetPermissionsOk() (*RequestPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RequestSearchResult) SetPermissions(v RequestPermissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


