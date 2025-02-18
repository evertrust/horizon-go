# ThirdPartyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** | The third party connector name on which this certificate is synchronized | 
**Id** | **string** | The Id of this certificate on the third party | 
**Fingerprint** | Pointer to **NullableString** | The fingerprint of this certificate on the third party | [optional] 
**PushDate** | Pointer to **NullableInt64** | The date when the certificate was pushed to this third party | [optional] 
**RemoveDate** | Pointer to **NullableInt64** | The date when the certificate was removed from this third party (in case of revocation) | [optional] 

## Methods

### NewThirdPartyItem

`func NewThirdPartyItem(connector string, id string, ) *ThirdPartyItem`

NewThirdPartyItem instantiates a new ThirdPartyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyItemWithDefaults

`func NewThirdPartyItemWithDefaults() *ThirdPartyItem`

NewThirdPartyItemWithDefaults instantiates a new ThirdPartyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ThirdPartyItem) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ThirdPartyItem) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ThirdPartyItem) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetId

`func (o *ThirdPartyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyItem) SetId(v string)`

SetId sets Id field to given value.


### GetFingerprint

`func (o *ThirdPartyItem) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ThirdPartyItem) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ThirdPartyItem) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ThirdPartyItem) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ThirdPartyItem) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ThirdPartyItem) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPushDate

`func (o *ThirdPartyItem) GetPushDate() int64`

GetPushDate returns the PushDate field if non-nil, zero value otherwise.

### GetPushDateOk

`func (o *ThirdPartyItem) GetPushDateOk() (*int64, bool)`

GetPushDateOk returns a tuple with the PushDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDate

`func (o *ThirdPartyItem) SetPushDate(v int64)`

SetPushDate sets PushDate field to given value.

### HasPushDate

`func (o *ThirdPartyItem) HasPushDate() bool`

HasPushDate returns a boolean if a field has been set.

### SetPushDateNil

`func (o *ThirdPartyItem) SetPushDateNil(b bool)`

 SetPushDateNil sets the value for PushDate to be an explicit nil

### UnsetPushDate
`func (o *ThirdPartyItem) UnsetPushDate()`

UnsetPushDate ensures that no value is present for PushDate, not even an explicit nil
### GetRemoveDate

`func (o *ThirdPartyItem) GetRemoveDate() int64`

GetRemoveDate returns the RemoveDate field if non-nil, zero value otherwise.

### GetRemoveDateOk

`func (o *ThirdPartyItem) GetRemoveDateOk() (*int64, bool)`

GetRemoveDateOk returns a tuple with the RemoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDate

`func (o *ThirdPartyItem) SetRemoveDate(v int64)`

SetRemoveDate sets RemoveDate field to given value.

### HasRemoveDate

`func (o *ThirdPartyItem) HasRemoveDate() bool`

HasRemoveDate returns a boolean if a field has been set.

### SetRemoveDateNil

`func (o *ThirdPartyItem) SetRemoveDateNil(b bool)`

 SetRemoveDateNil sets the value for RemoveDate to be an explicit nil

### UnsetRemoveDate
`func (o *ThirdPartyItem) UnsetRemoveDate()`

UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


