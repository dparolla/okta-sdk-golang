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

// KeyTrustLevelOSMode Represents the attestation strength used by the Chrome Verified Access API
type KeyTrustLevelOSMode string

// List of KeyTrustLevelOSMode
const (
	KEYTRUSTLEVELOSMODE_DEVELOPER_MODE KeyTrustLevelOSMode = "CHROME_OS_DEVELOPER_MODE"
	KEYTRUSTLEVELOSMODE_VERIFIED_MODE  KeyTrustLevelOSMode = "CHROME_OS_VERIFIED_MODE"
)

// All allowed values of KeyTrustLevelOSMode enum
var AllowedKeyTrustLevelOSModeEnumValues = []KeyTrustLevelOSMode{
	"CHROME_OS_DEVELOPER_MODE",
	"CHROME_OS_VERIFIED_MODE",
}

func (v *KeyTrustLevelOSMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyTrustLevelOSMode(value)
	for _, existing := range AllowedKeyTrustLevelOSModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyTrustLevelOSMode", value)
}

// NewKeyTrustLevelOSModeFromValue returns a pointer to a valid KeyTrustLevelOSMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyTrustLevelOSModeFromValue(v string) (*KeyTrustLevelOSMode, error) {
	ev := KeyTrustLevelOSMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyTrustLevelOSMode: valid values are %v", v, AllowedKeyTrustLevelOSModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyTrustLevelOSMode) IsValid() bool {
	for _, existing := range AllowedKeyTrustLevelOSModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyTrustLevelOSMode value
func (v KeyTrustLevelOSMode) Ptr() *KeyTrustLevelOSMode {
	return &v
}

type NullableKeyTrustLevelOSMode struct {
	value *KeyTrustLevelOSMode
	isSet bool
}

func (v NullableKeyTrustLevelOSMode) Get() *KeyTrustLevelOSMode {
	return v.value
}

func (v *NullableKeyTrustLevelOSMode) Set(val *KeyTrustLevelOSMode) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyTrustLevelOSMode) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyTrustLevelOSMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyTrustLevelOSMode(val *KeyTrustLevelOSMode) *NullableKeyTrustLevelOSMode {
	return &NullableKeyTrustLevelOSMode{value: val, isSet: true}
}

func (v NullableKeyTrustLevelOSMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyTrustLevelOSMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
