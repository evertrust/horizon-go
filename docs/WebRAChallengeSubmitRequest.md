# WebRAChallengeSubmitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | **string** | The one-time challenge that authorizes this enrollment | 
**Profile** | **string** | The WebRA profile name the challenge was issued on | 
**Template** | [**WebRAChallengeSubmitRequestTemplate**](WebRAChallengeSubmitRequestTemplate.md) | The user-data that will be used to generate the certificate | 

## Methods

### NewWebRAChallengeSubmitRequest

`func NewWebRAChallengeSubmitRequest(challenge string, profile string, template WebRAChallengeSubmitRequestTemplate, ) *WebRAChallengeSubmitRequest`

NewWebRAChallengeSubmitRequest instantiates a new WebRAChallengeSubmitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAChallengeSubmitRequestWithDefaults

`func NewWebRAChallengeSubmitRequestWithDefaults() *WebRAChallengeSubmitRequest`

NewWebRAChallengeSubmitRequestWithDefaults instantiates a new WebRAChallengeSubmitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *WebRAChallengeSubmitRequest) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *WebRAChallengeSubmitRequest) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *WebRAChallengeSubmitRequest) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetProfile

`func (o *WebRAChallengeSubmitRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAChallengeSubmitRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAChallengeSubmitRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRAChallengeSubmitRequest) GetTemplate() WebRAChallengeSubmitRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAChallengeSubmitRequest) GetTemplateOk() (*WebRAChallengeSubmitRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAChallengeSubmitRequest) SetTemplate(v WebRAChallengeSubmitRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


