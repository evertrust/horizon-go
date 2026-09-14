# Challenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**Problem**](Problem.md) |  | [optional] 
**Status** | [**ChallengeStatus**](ChallengeStatus.md) |  | 
**Token** | **string** |  | 
**Type** | [**AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | 
**Validated** | Pointer to **int64** |  | [optional] 

## Methods

### NewChallenge

`func NewChallenge(status ChallengeStatus, token string, type_ AcmeAuthorizationType, ) *Challenge`

NewChallenge instantiates a new Challenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeWithDefaults

`func NewChallengeWithDefaults() *Challenge`

NewChallengeWithDefaults instantiates a new Challenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Challenge) GetError() Problem`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Challenge) GetErrorOk() (*Problem, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Challenge) SetError(v Problem)`

SetError sets Error field to given value.

### HasError

`func (o *Challenge) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *Challenge) GetStatus() ChallengeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Challenge) GetStatusOk() (*ChallengeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Challenge) SetStatus(v ChallengeStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *Challenge) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Challenge) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Challenge) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *Challenge) GetType() AcmeAuthorizationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Challenge) GetTypeOk() (*AcmeAuthorizationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Challenge) SetType(v AcmeAuthorizationType)`

SetType sets Type field to given value.


### GetValidated

`func (o *Challenge) GetValidated() int64`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *Challenge) GetValidatedOk() (*int64, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *Challenge) SetValidated(v int64)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *Challenge) HasValidated() bool`

HasValidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


