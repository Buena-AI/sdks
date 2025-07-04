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
	"time"
)

// checks if the VoiceCloneResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceCloneResponseData{}

// VoiceCloneResponseData struct for VoiceCloneResponseData
type VoiceCloneResponseData struct {
	// Unique voice clone identifier
	VoiceId *string `json:"voiceId,omitempty"`
	// User-defined name for the voice clone
	Name *string `json:"name,omitempty"`
	// Description of the voice clone
	Description *string `json:"description,omitempty"`
	// Number of audio samples uploaded for training
	SampleCount *int32 `json:"sampleCount,omitempty"`
	// When the voice clone was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewVoiceCloneResponseData instantiates a new VoiceCloneResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceCloneResponseData() *VoiceCloneResponseData {
	this := VoiceCloneResponseData{}
	return &this
}

// NewVoiceCloneResponseDataWithDefaults instantiates a new VoiceCloneResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceCloneResponseDataWithDefaults() *VoiceCloneResponseData {
	this := VoiceCloneResponseData{}
	return &this
}

// GetVoiceId returns the VoiceId field value if set, zero value otherwise.
func (o *VoiceCloneResponseData) GetVoiceId() string {
	if o == nil || IsNil(o.VoiceId) {
		var ret string
		return ret
	}
	return *o.VoiceId
}

// GetVoiceIdOk returns a tuple with the VoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCloneResponseData) GetVoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.VoiceId) {
		return nil, false
	}
	return o.VoiceId, true
}

// HasVoiceId returns a boolean if a field has been set.
func (o *VoiceCloneResponseData) HasVoiceId() bool {
	if o != nil && !IsNil(o.VoiceId) {
		return true
	}

	return false
}

// SetVoiceId gets a reference to the given string and assigns it to the VoiceId field.
func (o *VoiceCloneResponseData) SetVoiceId(v string) {
	o.VoiceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VoiceCloneResponseData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCloneResponseData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VoiceCloneResponseData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VoiceCloneResponseData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VoiceCloneResponseData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCloneResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VoiceCloneResponseData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VoiceCloneResponseData) SetDescription(v string) {
	o.Description = &v
}

// GetSampleCount returns the SampleCount field value if set, zero value otherwise.
func (o *VoiceCloneResponseData) GetSampleCount() int32 {
	if o == nil || IsNil(o.SampleCount) {
		var ret int32
		return ret
	}
	return *o.SampleCount
}

// GetSampleCountOk returns a tuple with the SampleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCloneResponseData) GetSampleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleCount) {
		return nil, false
	}
	return o.SampleCount, true
}

// HasSampleCount returns a boolean if a field has been set.
func (o *VoiceCloneResponseData) HasSampleCount() bool {
	if o != nil && !IsNil(o.SampleCount) {
		return true
	}

	return false
}

// SetSampleCount gets a reference to the given int32 and assigns it to the SampleCount field.
func (o *VoiceCloneResponseData) SetSampleCount(v int32) {
	o.SampleCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VoiceCloneResponseData) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCloneResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VoiceCloneResponseData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VoiceCloneResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VoiceCloneResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceCloneResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoiceId) {
		toSerialize["voiceId"] = o.VoiceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SampleCount) {
		toSerialize["sampleCount"] = o.SampleCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableVoiceCloneResponseData struct {
	value *VoiceCloneResponseData
	isSet bool
}

func (v NullableVoiceCloneResponseData) Get() *VoiceCloneResponseData {
	return v.value
}

func (v *NullableVoiceCloneResponseData) Set(val *VoiceCloneResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceCloneResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceCloneResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceCloneResponseData(val *VoiceCloneResponseData) *NullableVoiceCloneResponseData {
	return &NullableVoiceCloneResponseData{value: val, isSet: true}
}

func (v NullableVoiceCloneResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceCloneResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


