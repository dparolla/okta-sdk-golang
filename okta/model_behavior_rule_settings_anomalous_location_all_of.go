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

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// BehaviorRuleSettingsAnomalousLocationAllOf struct for BehaviorRuleSettingsAnomalousLocationAllOf
type BehaviorRuleSettingsAnomalousLocationAllOf struct {
	Granularity LocationGranularity `json:"granularity"`
	// Required when `granularity` is `LAT_LONG`. Radius from the provided coordinates in kilometers.
	RadiusKilometers     *int32 `json:"radiusKilometers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousLocationAllOf BehaviorRuleSettingsAnomalousLocationAllOf

// NewBehaviorRuleSettingsAnomalousLocationAllOf instantiates a new BehaviorRuleSettingsAnomalousLocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousLocationAllOf(granularity LocationGranularity) *BehaviorRuleSettingsAnomalousLocationAllOf {
	this := BehaviorRuleSettingsAnomalousLocationAllOf{}
	this.Granularity = granularity
	return &this
}

// NewBehaviorRuleSettingsAnomalousLocationAllOfWithDefaults instantiates a new BehaviorRuleSettingsAnomalousLocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousLocationAllOfWithDefaults() *BehaviorRuleSettingsAnomalousLocationAllOf {
	this := BehaviorRuleSettingsAnomalousLocationAllOf{}
	return &this
}

// GetGranularity returns the Granularity field value
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetGranularity() LocationGranularity {
	if o == nil {
		var ret LocationGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetGranularityOk() (*LocationGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) SetGranularity(v LocationGranularity) {
	o.Granularity = v
}

// GetRadiusKilometers returns the RadiusKilometers field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetRadiusKilometers() int32 {
	if o == nil || o.RadiusKilometers == nil {
		var ret int32
		return ret
	}
	return *o.RadiusKilometers
}

// GetRadiusKilometersOk returns a tuple with the RadiusKilometers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetRadiusKilometersOk() (*int32, bool) {
	if o == nil || o.RadiusKilometers == nil {
		return nil, false
	}
	return o.RadiusKilometers, true
}

// HasRadiusKilometers returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) HasRadiusKilometers() bool {
	if o != nil && o.RadiusKilometers != nil {
		return true
	}

	return false
}

// SetRadiusKilometers gets a reference to the given int32 and assigns it to the RadiusKilometers field.
func (o *BehaviorRuleSettingsAnomalousLocationAllOf) SetRadiusKilometers(v int32) {
	o.RadiusKilometers = &v
}

func (o BehaviorRuleSettingsAnomalousLocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.RadiusKilometers != nil {
		toSerialize["radiusKilometers"] = o.RadiusKilometers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleSettingsAnomalousLocationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleSettingsAnomalousLocationAllOf := _BehaviorRuleSettingsAnomalousLocationAllOf{}

	err = json.Unmarshal(bytes, &varBehaviorRuleSettingsAnomalousLocationAllOf)
	if err == nil {
		*o = BehaviorRuleSettingsAnomalousLocationAllOf(varBehaviorRuleSettingsAnomalousLocationAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "radiusKilometers")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBehaviorRuleSettingsAnomalousLocationAllOf struct {
	value *BehaviorRuleSettingsAnomalousLocationAllOf
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousLocationAllOf) Get() *BehaviorRuleSettingsAnomalousLocationAllOf {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousLocationAllOf) Set(val *BehaviorRuleSettingsAnomalousLocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousLocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousLocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousLocationAllOf(val *BehaviorRuleSettingsAnomalousLocationAllOf) *NullableBehaviorRuleSettingsAnomalousLocationAllOf {
	return &NullableBehaviorRuleSettingsAnomalousLocationAllOf{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousLocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousLocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}