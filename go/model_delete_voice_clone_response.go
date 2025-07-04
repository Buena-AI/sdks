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

// checks if the DeleteVoiceCloneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVoiceCloneResponse{}

// DeleteVoiceCloneResponse struct for DeleteVoiceCloneResponse
type DeleteVoiceCloneResponse struct {
	Success *bool `json:"success,omitempty"`
	// Success message
	Message *string `json:"message,omitempty"`
	// The ID of the voice clone that was deleted
	DeletedId *string `json:"deleted_id,omitempty"`
}

// NewDeleteVoiceCloneResponse instantiates a new DeleteVoiceCloneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVoiceCloneResponse() *DeleteVoiceCloneResponse {
	this := DeleteVoiceCloneResponse{}
	return &this
}

// NewDeleteVoiceCloneResponseWithDefaults instantiates a new DeleteVoiceCloneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVoiceCloneResponseWithDefaults() *DeleteVoiceCloneResponse {
	this := DeleteVoiceCloneResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteVoiceCloneResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVoiceCloneResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteVoiceCloneResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteVoiceCloneResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteVoiceCloneResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVoiceCloneResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteVoiceCloneResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteVoiceCloneResponse) SetMessage(v string) {
	o.Message = &v
}

// GetDeletedId returns the DeletedId field value if set, zero value otherwise.
func (o *DeleteVoiceCloneResponse) GetDeletedId() string {
	if o == nil || IsNil(o.DeletedId) {
		var ret string
		return ret
	}
	return *o.DeletedId
}

// GetDeletedIdOk returns a tuple with the DeletedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVoiceCloneResponse) GetDeletedIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedId) {
		return nil, false
	}
	return o.DeletedId, true
}

// HasDeletedId returns a boolean if a field has been set.
func (o *DeleteVoiceCloneResponse) HasDeletedId() bool {
	if o != nil && !IsNil(o.DeletedId) {
		return true
	}

	return false
}

// SetDeletedId gets a reference to the given string and assigns it to the DeletedId field.
func (o *DeleteVoiceCloneResponse) SetDeletedId(v string) {
	o.DeletedId = &v
}

func (o DeleteVoiceCloneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVoiceCloneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.DeletedId) {
		toSerialize["deleted_id"] = o.DeletedId
	}
	return toSerialize, nil
}

type NullableDeleteVoiceCloneResponse struct {
	value *DeleteVoiceCloneResponse
	isSet bool
}

func (v NullableDeleteVoiceCloneResponse) Get() *DeleteVoiceCloneResponse {
	return v.value
}

func (v *NullableDeleteVoiceCloneResponse) Set(val *DeleteVoiceCloneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVoiceCloneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVoiceCloneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVoiceCloneResponse(val *DeleteVoiceCloneResponse) *NullableDeleteVoiceCloneResponse {
	return &NullableDeleteVoiceCloneResponse{value: val, isSet: true}
}

func (v NullableDeleteVoiceCloneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVoiceCloneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


