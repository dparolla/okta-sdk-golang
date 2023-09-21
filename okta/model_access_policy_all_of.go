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

// AccessPolicyAllOf struct for AccessPolicyAllOf
type AccessPolicyAllOf struct {
	Conditions           *PolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyAllOf AccessPolicyAllOf

// NewAccessPolicyAllOf instantiates a new AccessPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyAllOf() *AccessPolicyAllOf {
	this := AccessPolicyAllOf{}
	return &this
}

// NewAccessPolicyAllOfWithDefaults instantiates a new AccessPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyAllOfWithDefaults() *AccessPolicyAllOf {
	this := AccessPolicyAllOf{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *AccessPolicyAllOf) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

func (o AccessPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyAllOf := _AccessPolicyAllOf{}

	err = json.Unmarshal(bytes, &varAccessPolicyAllOf)
	if err == nil {
		*o = AccessPolicyAllOf(varAccessPolicyAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyAllOf struct {
	value *AccessPolicyAllOf
	isSet bool
}

func (v NullableAccessPolicyAllOf) Get() *AccessPolicyAllOf {
	return v.value
}

func (v *NullableAccessPolicyAllOf) Set(val *AccessPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyAllOf(val *AccessPolicyAllOf) *NullableAccessPolicyAllOf {
	return &NullableAccessPolicyAllOf{value: val, isSet: true}
}

func (v NullableAccessPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}