# WebRARevokeRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevocationReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 

## Methods

### NewWebRARevokeRequestTemplate

`func NewWebRARevokeRequestTemplate() *WebRARevokeRequestTemplate`

NewWebRARevokeRequestTemplate instantiates a new WebRARevokeRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARevokeRequestTemplateWithDefaults

`func NewWebRARevokeRequestTemplateWithDefaults() *WebRARevokeRequestTemplate`

NewWebRARevokeRequestTemplateWithDefaults instantiates a new WebRARevokeRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevocationReason

`func (o *WebRARevokeRequestTemplate) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *WebRARevokeRequestTemplate) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *WebRARevokeRequestTemplate) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *WebRARevokeRequestTemplate) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *WebRARevokeRequestTemplate) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *WebRARevokeRequestTemplate) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


