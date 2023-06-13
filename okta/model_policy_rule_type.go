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
	"fmt"
)

// PolicyRuleType the model 'PolicyRuleType'
type PolicyRuleType string

// List of PolicyRuleType
const (
	POLICYRULETYPE_ACCESS_POLICY      PolicyRuleType = "ACCESS_POLICY"
	POLICYRULETYPE_IDP_DISCOVERY      PolicyRuleType = "IDP_DISCOVERY"
	POLICYRULETYPE_MFA_ENROLL         PolicyRuleType = "MFA_ENROLL"
	POLICYRULETYPE_PASSWORD           PolicyRuleType = "PASSWORD"
	POLICYRULETYPE_PROFILE_ENROLLMENT PolicyRuleType = "PROFILE_ENROLLMENT"
	POLICYRULETYPE_RESOURCE_ACCESS    PolicyRuleType = "RESOURCE_ACCESS"
	POLICYRULETYPE_SIGN_ON            PolicyRuleType = "SIGN_ON"
)

// All allowed values of PolicyRuleType enum
var AllowedPolicyRuleTypeEnumValues = []PolicyRuleType{
	"ACCESS_POLICY",
	"IDP_DISCOVERY",
	"MFA_ENROLL",
	"PASSWORD",
	"PROFILE_ENROLLMENT",
	"RESOURCE_ACCESS",
	"SIGN_ON",
}

func (v *PolicyRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyRuleType(value)
	for _, existing := range AllowedPolicyRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyRuleType", value)
}

// NewPolicyRuleTypeFromValue returns a pointer to a valid PolicyRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyRuleTypeFromValue(v string) (*PolicyRuleType, error) {
	ev := PolicyRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyRuleType: valid values are %v", v, AllowedPolicyRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyRuleType) IsValid() bool {
	for _, existing := range AllowedPolicyRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyRuleType value
func (v PolicyRuleType) Ptr() *PolicyRuleType {
	return &v
}

type NullablePolicyRuleType struct {
	value *PolicyRuleType
	isSet bool
}

func (v NullablePolicyRuleType) Get() *PolicyRuleType {
	return v.value
}

func (v *NullablePolicyRuleType) Set(val *PolicyRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRuleType(val *PolicyRuleType) *NullablePolicyRuleType {
	return &NullablePolicyRuleType{value: val, isSet: true}
}

func (v NullablePolicyRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}