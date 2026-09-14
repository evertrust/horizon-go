# X509SanElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | SAN type | 
**Values** | **[]string** | The requested SAN values for this type | 

## Methods

### NewX509SanElementResponse

`func NewX509SanElementResponse(type_ string, values []string, ) *X509SanElementResponse`

NewX509SanElementResponse instantiates a new X509SanElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509SanElementResponseWithDefaults

`func NewX509SanElementResponseWithDefaults() *X509SanElementResponse`

NewX509SanElementResponseWithDefaults instantiates a new X509SanElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *X509SanElementResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *X509SanElementResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *X509SanElementResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *X509SanElementResponse) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *X509SanElementResponse) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *X509SanElementResponse) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


