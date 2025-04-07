# ADConnectionAttributesConnectionTimeoutConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryWait** | Pointer to **int32** | Wait time before retrying a failed connection. | [optional] 
**TokenRefreshMaxTryCount** | Pointer to **int32** | Maximum number of retries for token refresh. | [optional] 
**RetryFailureStatusCode** | Pointer to **int32** |  | [optional] 
**RetryWaitMaxValue** | Pointer to **int32** | Maximum wait time for retries. | [optional] 
**RetryCount** | Pointer to **int32** | Number of retry attempts allowed. | [optional] 
**ReadTimeout** | Pointer to **int32** | Read timeout duration (in seconds). | [optional] 
**ConnectionTimeout** | Pointer to **int32** | Connection timeout duration (in seconds). | [optional] 

## Methods

### NewADConnectionAttributesConnectionTimeoutConfig

`func NewADConnectionAttributesConnectionTimeoutConfig() *ADConnectionAttributesConnectionTimeoutConfig`

NewADConnectionAttributesConnectionTimeoutConfig instantiates a new ADConnectionAttributesConnectionTimeoutConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADConnectionAttributesConnectionTimeoutConfigWithDefaults

`func NewADConnectionAttributesConnectionTimeoutConfigWithDefaults() *ADConnectionAttributesConnectionTimeoutConfig`

NewADConnectionAttributesConnectionTimeoutConfigWithDefaults instantiates a new ADConnectionAttributesConnectionTimeoutConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryWait

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryWait() int32`

GetRetryWait returns the RetryWait field if non-nil, zero value otherwise.

### GetRetryWaitOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryWaitOk() (*int32, bool)`

GetRetryWaitOk returns a tuple with the RetryWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWait

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetRetryWait(v int32)`

SetRetryWait sets RetryWait field to given value.

### HasRetryWait

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasRetryWait() bool`

HasRetryWait returns a boolean if a field has been set.

### GetTokenRefreshMaxTryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetTokenRefreshMaxTryCount() int32`

GetTokenRefreshMaxTryCount returns the TokenRefreshMaxTryCount field if non-nil, zero value otherwise.

### GetTokenRefreshMaxTryCountOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetTokenRefreshMaxTryCountOk() (*int32, bool)`

GetTokenRefreshMaxTryCountOk returns a tuple with the TokenRefreshMaxTryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRefreshMaxTryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetTokenRefreshMaxTryCount(v int32)`

SetTokenRefreshMaxTryCount sets TokenRefreshMaxTryCount field to given value.

### HasTokenRefreshMaxTryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasTokenRefreshMaxTryCount() bool`

HasTokenRefreshMaxTryCount returns a boolean if a field has been set.

### GetRetryFailureStatusCode

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryFailureStatusCode() int32`

GetRetryFailureStatusCode returns the RetryFailureStatusCode field if non-nil, zero value otherwise.

### GetRetryFailureStatusCodeOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryFailureStatusCodeOk() (*int32, bool)`

GetRetryFailureStatusCodeOk returns a tuple with the RetryFailureStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailureStatusCode

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetRetryFailureStatusCode(v int32)`

SetRetryFailureStatusCode sets RetryFailureStatusCode field to given value.

### HasRetryFailureStatusCode

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasRetryFailureStatusCode() bool`

HasRetryFailureStatusCode returns a boolean if a field has been set.

### GetRetryWaitMaxValue

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryWaitMaxValue() int32`

GetRetryWaitMaxValue returns the RetryWaitMaxValue field if non-nil, zero value otherwise.

### GetRetryWaitMaxValueOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryWaitMaxValueOk() (*int32, bool)`

GetRetryWaitMaxValueOk returns a tuple with the RetryWaitMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWaitMaxValue

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetRetryWaitMaxValue(v int32)`

SetRetryWaitMaxValue sets RetryWaitMaxValue field to given value.

### HasRetryWaitMaxValue

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasRetryWaitMaxValue() bool`

HasRetryWaitMaxValue returns a boolean if a field has been set.

### GetRetryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetReadTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *ADConnectionAttributesConnectionTimeoutConfig) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *ADConnectionAttributesConnectionTimeoutConfig) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


