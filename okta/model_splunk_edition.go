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
	"fmt"
)

// SplunkEdition Edition of the Splunk Cloud instance
type SplunkEdition string

// List of SplunkEdition
const (
	SPLUNKEDITION_AWS          SplunkEdition = "aws"
	SPLUNKEDITION_AWS_GOVCLOUD SplunkEdition = "aws_govcloud"
	SPLUNKEDITION_GCP          SplunkEdition = "gcp"
)

// All allowed values of SplunkEdition enum
var AllowedSplunkEditionEnumValues = []SplunkEdition{
	"aws",
	"aws_govcloud",
	"gcp",
}

func (v *SplunkEdition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SplunkEdition(value)
	for _, existing := range AllowedSplunkEditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SplunkEdition", value)
}

// NewSplunkEditionFromValue returns a pointer to a valid SplunkEdition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSplunkEditionFromValue(v string) (*SplunkEdition, error) {
	ev := SplunkEdition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SplunkEdition: valid values are %v", v, AllowedSplunkEditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SplunkEdition) IsValid() bool {
	for _, existing := range AllowedSplunkEditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplunkEdition value
func (v SplunkEdition) Ptr() *SplunkEdition {
	return &v
}

type NullableSplunkEdition struct {
	value *SplunkEdition
	isSet bool
}

func (v NullableSplunkEdition) Get() *SplunkEdition {
	return v.value
}

func (v *NullableSplunkEdition) Set(val *SplunkEdition) {
	v.value = val
	v.isSet = true
}

func (v NullableSplunkEdition) IsSet() bool {
	return v.isSet
}

func (v *NullableSplunkEdition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplunkEdition(val *SplunkEdition) *NullableSplunkEdition {
	return &NullableSplunkEdition{value: val, isSet: true}
}

func (v NullableSplunkEdition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplunkEdition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}