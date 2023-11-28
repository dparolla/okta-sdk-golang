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
	"reflect"
	"strings"
)

// InboundProvisioningApplicationFeature struct for InboundProvisioningApplicationFeature
type InboundProvisioningApplicationFeature struct {
	ApplicationFeature
	AdditionalProperties map[string]interface{}
}

type _InboundProvisioningApplicationFeature InboundProvisioningApplicationFeature

// NewInboundProvisioningApplicationFeature instantiates a new InboundProvisioningApplicationFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundProvisioningApplicationFeature() *InboundProvisioningApplicationFeature {
	this := InboundProvisioningApplicationFeature{}
	return &this
}

// NewInboundProvisioningApplicationFeatureWithDefaults instantiates a new InboundProvisioningApplicationFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundProvisioningApplicationFeatureWithDefaults() *InboundProvisioningApplicationFeature {
	this := InboundProvisioningApplicationFeature{}
	return &this
}

func (o InboundProvisioningApplicationFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedApplicationFeature, errApplicationFeature := json.Marshal(o.ApplicationFeature)
	if errApplicationFeature != nil {
		return []byte{}, errApplicationFeature
	}
	errApplicationFeature = json.Unmarshal([]byte(serializedApplicationFeature), &toSerialize)
	if errApplicationFeature != nil {
		return []byte{}, errApplicationFeature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InboundProvisioningApplicationFeature) UnmarshalJSON(bytes []byte) (err error) {
	type InboundProvisioningApplicationFeatureWithoutEmbeddedStruct struct{}

	varInboundProvisioningApplicationFeatureWithoutEmbeddedStruct := InboundProvisioningApplicationFeatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInboundProvisioningApplicationFeatureWithoutEmbeddedStruct)
	if err == nil {
		varInboundProvisioningApplicationFeature := _InboundProvisioningApplicationFeature{}
		*o = InboundProvisioningApplicationFeature(varInboundProvisioningApplicationFeature)
	} else {
		return err
	}

	varInboundProvisioningApplicationFeature := _InboundProvisioningApplicationFeature{}

	err = json.Unmarshal(bytes, &varInboundProvisioningApplicationFeature)
	if err == nil {
		o.ApplicationFeature = varInboundProvisioningApplicationFeature.ApplicationFeature
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {

		// remove fields from embedded structs
		reflectApplicationFeature := reflect.ValueOf(o.ApplicationFeature)
		for i := 0; i < reflectApplicationFeature.Type().NumField(); i++ {
			t := reflectApplicationFeature.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInboundProvisioningApplicationFeature struct {
	value *InboundProvisioningApplicationFeature
	isSet bool
}

func (v NullableInboundProvisioningApplicationFeature) Get() *InboundProvisioningApplicationFeature {
	return v.value
}

func (v *NullableInboundProvisioningApplicationFeature) Set(val *InboundProvisioningApplicationFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundProvisioningApplicationFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundProvisioningApplicationFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundProvisioningApplicationFeature(val *InboundProvisioningApplicationFeature) *NullableInboundProvisioningApplicationFeature {
	return &NullableInboundProvisioningApplicationFeature{value: val, isSet: true}
}

func (v NullableInboundProvisioningApplicationFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundProvisioningApplicationFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
