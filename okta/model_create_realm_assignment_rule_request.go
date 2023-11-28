/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// CreateRealmAssignmentRuleRequest struct for CreateRealmAssignmentRuleRequest
type CreateRealmAssignmentRuleRequest struct {
	Actions              *Actions    `json:"actions,omitempty"`
	Conditions           *Conditions `json:"conditions,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	Priority             *int32      `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRealmAssignmentRuleRequest CreateRealmAssignmentRuleRequest

// NewCreateRealmAssignmentRuleRequest instantiates a new CreateRealmAssignmentRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRealmAssignmentRuleRequest() *CreateRealmAssignmentRuleRequest {
	this := CreateRealmAssignmentRuleRequest{}
	return &this
}

// NewCreateRealmAssignmentRuleRequestWithDefaults instantiates a new CreateRealmAssignmentRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRealmAssignmentRuleRequestWithDefaults() *CreateRealmAssignmentRuleRequest {
	this := CreateRealmAssignmentRuleRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CreateRealmAssignmentRuleRequest) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRealmAssignmentRuleRequest) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CreateRealmAssignmentRuleRequest) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *CreateRealmAssignmentRuleRequest) SetActions(v Actions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *CreateRealmAssignmentRuleRequest) GetConditions() Conditions {
	if o == nil || o.Conditions == nil {
		var ret Conditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRealmAssignmentRuleRequest) GetConditionsOk() (*Conditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CreateRealmAssignmentRuleRequest) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given Conditions and assigns it to the Conditions field.
func (o *CreateRealmAssignmentRuleRequest) SetConditions(v Conditions) {
	o.Conditions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRealmAssignmentRuleRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRealmAssignmentRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRealmAssignmentRuleRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRealmAssignmentRuleRequest) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateRealmAssignmentRuleRequest) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRealmAssignmentRuleRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateRealmAssignmentRuleRequest) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *CreateRealmAssignmentRuleRequest) SetPriority(v int32) {
	o.Priority = &v
}

func (o CreateRealmAssignmentRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateRealmAssignmentRuleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateRealmAssignmentRuleRequest := _CreateRealmAssignmentRuleRequest{}

	err = json.Unmarshal(bytes, &varCreateRealmAssignmentRuleRequest)
	if err == nil {
		*o = CreateRealmAssignmentRuleRequest(varCreateRealmAssignmentRuleRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateRealmAssignmentRuleRequest struct {
	value *CreateRealmAssignmentRuleRequest
	isSet bool
}

func (v NullableCreateRealmAssignmentRuleRequest) Get() *CreateRealmAssignmentRuleRequest {
	return v.value
}

func (v *NullableCreateRealmAssignmentRuleRequest) Set(val *CreateRealmAssignmentRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRealmAssignmentRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRealmAssignmentRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRealmAssignmentRuleRequest(val *CreateRealmAssignmentRuleRequest) *NullableCreateRealmAssignmentRuleRequest {
	return &NullableCreateRealmAssignmentRuleRequest{value: val, isSet: true}
}

func (v NullableCreateRealmAssignmentRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRealmAssignmentRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
