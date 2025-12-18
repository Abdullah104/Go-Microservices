# DataProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | **int32** | The ID for this product | 
**Name** | **string** |  | 
**Price** | Pointer to **float32** |  | [optional] 
**Sku** | **string** |  | 

## Methods

### NewDataProduct

`func NewDataProduct(id int32, name string, sku string, ) *DataProduct`

NewDataProduct instantiates a new DataProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProductWithDefaults

`func NewDataProductWithDefaults() *DataProduct`

NewDataProductWithDefaults instantiates a new DataProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DataProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DataProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataProduct) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DataProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataProduct) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *DataProduct) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DataProduct) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DataProduct) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DataProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSku

`func (o *DataProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *DataProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *DataProduct) SetSku(v string)`

SetSku sets Sku field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


