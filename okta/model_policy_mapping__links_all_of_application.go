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

// PolicyMappingLinksAllOfApplication struct for PolicyMappingLinksAllOfApplication
type PolicyMappingLinksAllOfApplication struct {
	Hints *HrefObjectHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Name *string `json:"name,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyMappingLinksAllOfApplication PolicyMappingLinksAllOfApplication

// NewPolicyMappingLinksAllOfApplication instantiates a new PolicyMappingLinksAllOfApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyMappingLinksAllOfApplication(href string) *PolicyMappingLinksAllOfApplication {
	this := PolicyMappingLinksAllOfApplication{}
	this.Href = href
	return &this
}

// NewPolicyMappingLinksAllOfApplicationWithDefaults instantiates a new PolicyMappingLinksAllOfApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyMappingLinksAllOfApplicationWithDefaults() *PolicyMappingLinksAllOfApplication {
	this := PolicyMappingLinksAllOfApplication{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *PolicyMappingLinksAllOfApplication) GetHints() HrefObjectHints {
	if o == nil || o.Hints == nil {
		var ret HrefObjectHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinksAllOfApplication) GetHintsOk() (*HrefObjectHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *PolicyMappingLinksAllOfApplication) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given HrefObjectHints and assigns it to the Hints field.
func (o *PolicyMappingLinksAllOfApplication) SetHints(v HrefObjectHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *PolicyMappingLinksAllOfApplication) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinksAllOfApplication) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *PolicyMappingLinksAllOfApplication) SetHref(v string) {
	o.Href = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyMappingLinksAllOfApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinksAllOfApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyMappingLinksAllOfApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyMappingLinksAllOfApplication) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyMappingLinksAllOfApplication) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinksAllOfApplication) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyMappingLinksAllOfApplication) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyMappingLinksAllOfApplication) SetType(v string) {
	o.Type = &v
}

func (o PolicyMappingLinksAllOfApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyMappingLinksAllOfApplication) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyMappingLinksAllOfApplication := _PolicyMappingLinksAllOfApplication{}

	err = json.Unmarshal(bytes, &varPolicyMappingLinksAllOfApplication)
	if err == nil {
		*o = PolicyMappingLinksAllOfApplication(varPolicyMappingLinksAllOfApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyMappingLinksAllOfApplication struct {
	value *PolicyMappingLinksAllOfApplication
	isSet bool
}

func (v NullablePolicyMappingLinksAllOfApplication) Get() *PolicyMappingLinksAllOfApplication {
	return v.value
}

func (v *NullablePolicyMappingLinksAllOfApplication) Set(val *PolicyMappingLinksAllOfApplication) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyMappingLinksAllOfApplication) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyMappingLinksAllOfApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyMappingLinksAllOfApplication(val *PolicyMappingLinksAllOfApplication) *NullablePolicyMappingLinksAllOfApplication {
	return &NullablePolicyMappingLinksAllOfApplication{value: val, isSet: true}
}

func (v NullablePolicyMappingLinksAllOfApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyMappingLinksAllOfApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}