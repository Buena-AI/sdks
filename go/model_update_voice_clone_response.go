/*
Buena.ai API v2

The most powerful LinkedIn automation and lead management API for modern sales teams and developers.

API version: 2.0.0
Contact: support@buena.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buena

import (
	"encoding/json"
)

// checks if the UpdateVoiceCloneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVoiceCloneResponse{}

// UpdateVoiceCloneResponse struct for UpdateVoiceCloneResponse
type UpdateVoiceCloneResponse struct {
	Success *bool `json:"success,omitempty"`
	// Success message
	Message *string `json:"message,omitempty"`
	Data *UpdateVoiceCloneResponseData `json:"data,omitempty"`
}

// NewUpdateVoiceCloneResponse instantiates a new UpdateVoiceCloneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVoiceCloneResponse() *UpdateVoiceCloneResponse {
	this := UpdateVoiceCloneResponse{}
	return &this
}

// NewUpdateVoiceCloneResponseWithDefaults instantiates a new UpdateVoiceCloneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVoiceCloneResponseWithDefaults() *UpdateVoiceCloneResponse {
	this := UpdateVoiceCloneResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateVoiceCloneResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdateVoiceCloneResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateVoiceCloneResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateVoiceCloneResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateVoiceCloneResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateVoiceCloneResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateVoiceCloneResponse) GetData() UpdateVoiceCloneResponseData {
	if o == nil || IsNil(o.Data) {
		var ret UpdateVoiceCloneResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneResponse) GetDataOk() (*UpdateVoiceCloneResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateVoiceCloneResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateVoiceCloneResponseData and assigns it to the Data field.
func (o *UpdateVoiceCloneResponse) SetData(v UpdateVoiceCloneResponseData) {
	o.Data = &v
}

func (o UpdateVoiceCloneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVoiceCloneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateVoiceCloneResponse struct {
	value *UpdateVoiceCloneResponse
	isSet bool
}

func (v NullableUpdateVoiceCloneResponse) Get() *UpdateVoiceCloneResponse {
	return v.value
}

func (v *NullableUpdateVoiceCloneResponse) Set(val *UpdateVoiceCloneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVoiceCloneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVoiceCloneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVoiceCloneResponse(val *UpdateVoiceCloneResponse) *NullableUpdateVoiceCloneResponse {
	return &NullableUpdateVoiceCloneResponse{value: val, isSet: true}
}

func (v NullableUpdateVoiceCloneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVoiceCloneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


