# WebRAImportRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElementResponse**](CertificateContactEmailElementResponse.md) | The contact email for this certificate | [optional] 
**DiscoveryData** | Pointer to [**HostDiscoveryData**](HostDiscoveryData.md) | The host discovery data associated with the certificate (discovery metadata) | [optional] 
**DiscoveryInfo** | Pointer to [**NullableDiscoveryInfo**](DiscoveryInfo.md) | Information about the discovery of this certificate | [optional] 
**Labels** | Pointer to [**[]RequestLabelElementResponse**](RequestLabelElementResponse.md) | The labels for this certificate | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElementResponse**](CertificateMetadataElementResponse.md) | The technical metadata for this certificate | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElementResponse**](CertificateOwnerElementResponse.md) | The owner for this certificate | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElementResponse**](CertificateTeamElementResponse.md) | The team for this certificate | [optional] 
**ThirdPartyData** | Pointer to [**[]ThirdPartyItem**](ThirdPartyItem.md) | The third party data associated with the certificate | [optional] 

## Methods

### NewWebRAImportRequestTemplateResponse

`func NewWebRAImportRequestTemplateResponse() *WebRAImportRequestTemplateResponse`

NewWebRAImportRequestTemplateResponse instantiates a new WebRAImportRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestTemplateResponseWithDefaults

`func NewWebRAImportRequestTemplateResponseWithDefaults() *WebRAImportRequestTemplateResponse`

NewWebRAImportRequestTemplateResponseWithDefaults instantiates a new WebRAImportRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *WebRAImportRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAImportRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAImportRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAImportRequestTemplateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAImportRequestTemplateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAImportRequestTemplateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetDiscoveryData

`func (o *WebRAImportRequestTemplateResponse) GetDiscoveryData() HostDiscoveryData`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *WebRAImportRequestTemplateResponse) GetDiscoveryDataOk() (*HostDiscoveryData, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *WebRAImportRequestTemplateResponse) SetDiscoveryData(v HostDiscoveryData)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *WebRAImportRequestTemplateResponse) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### GetDiscoveryInfo

`func (o *WebRAImportRequestTemplateResponse) GetDiscoveryInfo() DiscoveryInfo`

GetDiscoveryInfo returns the DiscoveryInfo field if non-nil, zero value otherwise.

### GetDiscoveryInfoOk

`func (o *WebRAImportRequestTemplateResponse) GetDiscoveryInfoOk() (*DiscoveryInfo, bool)`

GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryInfo

`func (o *WebRAImportRequestTemplateResponse) SetDiscoveryInfo(v DiscoveryInfo)`

SetDiscoveryInfo sets DiscoveryInfo field to given value.

### HasDiscoveryInfo

`func (o *WebRAImportRequestTemplateResponse) HasDiscoveryInfo() bool`

HasDiscoveryInfo returns a boolean if a field has been set.

### SetDiscoveryInfoNil

`func (o *WebRAImportRequestTemplateResponse) SetDiscoveryInfoNil(b bool)`

 SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil

### UnsetDiscoveryInfo
`func (o *WebRAImportRequestTemplateResponse) UnsetDiscoveryInfo()`

UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
### GetLabels

`func (o *WebRAImportRequestTemplateResponse) GetLabels() []RequestLabelElementResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAImportRequestTemplateResponse) GetLabelsOk() (*[]RequestLabelElementResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAImportRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAImportRequestTemplateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAImportRequestTemplateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAImportRequestTemplateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAImportRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAImportRequestTemplateResponse) GetMetadataOk() (*[]CertificateMetadataElementResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAImportRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAImportRequestTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAImportRequestTemplateResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAImportRequestTemplateResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAImportRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAImportRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAImportRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAImportRequestTemplateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAImportRequestTemplateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAImportRequestTemplateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTeam

`func (o *WebRAImportRequestTemplateResponse) GetTeam() CertificateTeamElementResponse`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAImportRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAImportRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAImportRequestTemplateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAImportRequestTemplateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAImportRequestTemplateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetThirdPartyData

`func (o *WebRAImportRequestTemplateResponse) GetThirdPartyData() []ThirdPartyItem`

GetThirdPartyData returns the ThirdPartyData field if non-nil, zero value otherwise.

### GetThirdPartyDataOk

`func (o *WebRAImportRequestTemplateResponse) GetThirdPartyDataOk() (*[]ThirdPartyItem, bool)`

GetThirdPartyDataOk returns a tuple with the ThirdPartyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyData

`func (o *WebRAImportRequestTemplateResponse) SetThirdPartyData(v []ThirdPartyItem)`

SetThirdPartyData sets ThirdPartyData field to given value.

### HasThirdPartyData

`func (o *WebRAImportRequestTemplateResponse) HasThirdPartyData() bool`

HasThirdPartyData returns a boolean if a field has been set.

### SetThirdPartyDataNil

`func (o *WebRAImportRequestTemplateResponse) SetThirdPartyDataNil(b bool)`

 SetThirdPartyDataNil sets the value for ThirdPartyData to be an explicit nil

### UnsetThirdPartyData
`func (o *WebRAImportRequestTemplateResponse) UnsetThirdPartyData()`

UnsetThirdPartyData ensures that no value is present for ThirdPartyData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


