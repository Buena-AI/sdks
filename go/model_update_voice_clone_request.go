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

// checks if the UpdateVoiceCloneRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVoiceCloneRequest{}

// UpdateVoiceCloneRequest struct for UpdateVoiceCloneRequest
type UpdateVoiceCloneRequest struct {
	// Updated name for the voice clone
	Name *string `json:"name,omitempty"`
	// Updated description of the voice clone
	Description *string `json:"description,omitempty"`
	// Enable/disable the voice clone
	IsActive *bool `json:"isActive,omitempty"`
}

// NewUpdateVoiceCloneRequest instantiates a new UpdateVoiceCloneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVoiceCloneRequest() *UpdateVoiceCloneRequest {
	this := UpdateVoiceCloneRequest{}
	return &this
}

// NewUpdateVoiceCloneRequestWithDefaults instantiates a new UpdateVoiceCloneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVoiceCloneRequestWithDefaults() *UpdateVoiceCloneRequest {
	this := UpdateVoiceCloneRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVoiceCloneRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVoiceCloneRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVoiceCloneRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVoiceCloneRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVoiceCloneRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVoiceCloneRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateVoiceCloneRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVoiceCloneRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateVoiceCloneRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateVoiceCloneRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o UpdateVoiceCloneRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVoiceCloneRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	return toSerialize, nil
}

type NullableUpdateVoiceCloneRequest struct {
	value *UpdateVoiceCloneRequest
	isSet bool
}

func (v NullableUpdateVoiceCloneRequest) Get() *UpdateVoiceCloneRequest {
	return v.value
}

func (v *NullableUpdateVoiceCloneRequest) Set(val *UpdateVoiceCloneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVoiceCloneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVoiceCloneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVoiceCloneRequest(val *UpdateVoiceCloneRequest) *NullableUpdateVoiceCloneRequest {
	return &NullableUpdateVoiceCloneRequest{value: val, isSet: true}
}

func (v NullableUpdateVoiceCloneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVoiceCloneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


