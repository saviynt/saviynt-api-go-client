# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CONNECTIONDESCRIPTION** | **string** |  | 
**CONNECTIONNAME** | **string** |  | 
**CONNECTIONTYPE** | **string** |  | 
**CREATEDBY** | Pointer to **string** |  | [optional] 
**CREATEDON** | **time.Time** |  | 
**STATUS** | **float32** |  | 
**UPDATEDBY** | Pointer to **string** |  | [optional] 
**UPDATEDON** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConnection

`func NewConnection(cONNECTIONDESCRIPTION string, cONNECTIONNAME string, cONNECTIONTYPE string, cREATEDON time.Time, sTATUS float32, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCONNECTIONDESCRIPTION

`func (o *Connection) GetCONNECTIONDESCRIPTION() string`

GetCONNECTIONDESCRIPTION returns the CONNECTIONDESCRIPTION field if non-nil, zero value otherwise.

### GetCONNECTIONDESCRIPTIONOk

`func (o *Connection) GetCONNECTIONDESCRIPTIONOk() (*string, bool)`

GetCONNECTIONDESCRIPTIONOk returns a tuple with the CONNECTIONDESCRIPTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONDESCRIPTION

`func (o *Connection) SetCONNECTIONDESCRIPTION(v string)`

SetCONNECTIONDESCRIPTION sets CONNECTIONDESCRIPTION field to given value.


### GetCONNECTIONNAME

`func (o *Connection) GetCONNECTIONNAME() string`

GetCONNECTIONNAME returns the CONNECTIONNAME field if non-nil, zero value otherwise.

### GetCONNECTIONNAMEOk

`func (o *Connection) GetCONNECTIONNAMEOk() (*string, bool)`

GetCONNECTIONNAMEOk returns a tuple with the CONNECTIONNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONNAME

`func (o *Connection) SetCONNECTIONNAME(v string)`

SetCONNECTIONNAME sets CONNECTIONNAME field to given value.


### GetCONNECTIONTYPE

`func (o *Connection) GetCONNECTIONTYPE() string`

GetCONNECTIONTYPE returns the CONNECTIONTYPE field if non-nil, zero value otherwise.

### GetCONNECTIONTYPEOk

`func (o *Connection) GetCONNECTIONTYPEOk() (*string, bool)`

GetCONNECTIONTYPEOk returns a tuple with the CONNECTIONTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONTYPE

`func (o *Connection) SetCONNECTIONTYPE(v string)`

SetCONNECTIONTYPE sets CONNECTIONTYPE field to given value.


### GetCREATEDBY

`func (o *Connection) GetCREATEDBY() string`

GetCREATEDBY returns the CREATEDBY field if non-nil, zero value otherwise.

### GetCREATEDBYOk

`func (o *Connection) GetCREATEDBYOk() (*string, bool)`

GetCREATEDBYOk returns a tuple with the CREATEDBY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEDBY

`func (o *Connection) SetCREATEDBY(v string)`

SetCREATEDBY sets CREATEDBY field to given value.

### HasCREATEDBY

`func (o *Connection) HasCREATEDBY() bool`

HasCREATEDBY returns a boolean if a field has been set.

### GetCREATEDON

`func (o *Connection) GetCREATEDON() time.Time`

GetCREATEDON returns the CREATEDON field if non-nil, zero value otherwise.

### GetCREATEDONOk

`func (o *Connection) GetCREATEDONOk() (*time.Time, bool)`

GetCREATEDONOk returns a tuple with the CREATEDON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEDON

`func (o *Connection) SetCREATEDON(v time.Time)`

SetCREATEDON sets CREATEDON field to given value.


### GetSTATUS

`func (o *Connection) GetSTATUS() float32`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *Connection) GetSTATUSOk() (*float32, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *Connection) SetSTATUS(v float32)`

SetSTATUS sets STATUS field to given value.


### GetUPDATEDBY

`func (o *Connection) GetUPDATEDBY() string`

GetUPDATEDBY returns the UPDATEDBY field if non-nil, zero value otherwise.

### GetUPDATEDBYOk

`func (o *Connection) GetUPDATEDBYOk() (*string, bool)`

GetUPDATEDBYOk returns a tuple with the UPDATEDBY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEDBY

`func (o *Connection) SetUPDATEDBY(v string)`

SetUPDATEDBY sets UPDATEDBY field to given value.

### HasUPDATEDBY

`func (o *Connection) HasUPDATEDBY() bool`

HasUPDATEDBY returns a boolean if a field has been set.

### GetUPDATEDON

`func (o *Connection) GetUPDATEDON() time.Time`

GetUPDATEDON returns the UPDATEDON field if non-nil, zero value otherwise.

### GetUPDATEDONOk

`func (o *Connection) GetUPDATEDONOk() (*time.Time, bool)`

GetUPDATEDONOk returns a tuple with the UPDATEDON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEDON

`func (o *Connection) SetUPDATEDON(v time.Time)`

SetUPDATEDON sets UPDATEDON field to given value.

### HasUPDATEDON

`func (o *Connection) HasUPDATEDON() bool`

HasUPDATEDON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


