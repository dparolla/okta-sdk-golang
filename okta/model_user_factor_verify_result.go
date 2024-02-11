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

// UserFactorVerifyResult Result of a Factor verification
type UserFactorVerifyResult string

// List of UserFactorVerifyResult
const (
	USERFACTORVERIFYRESULT_CHALLENGE            UserFactorVerifyResult = "CHALLENGE"
	USERFACTORVERIFYRESULT_ERROR                UserFactorVerifyResult = "ERROR"
	USERFACTORVERIFYRESULT_EXPIRED              UserFactorVerifyResult = "EXPIRED"
	USERFACTORVERIFYRESULT_FAILED               UserFactorVerifyResult = "FAILED"
	USERFACTORVERIFYRESULT_PASSCODE_REPLAYED    UserFactorVerifyResult = "PASSCODE_REPLAYED"
	USERFACTORVERIFYRESULT_REJECTED             UserFactorVerifyResult = "REJECTED"
	USERFACTORVERIFYRESULT_SUCCESS              UserFactorVerifyResult = "SUCCESS"
	USERFACTORVERIFYRESULT_TIMEOUT              UserFactorVerifyResult = "TIMEOUT"
	USERFACTORVERIFYRESULT_TIME_WINDOW_EXCEEDED UserFactorVerifyResult = "TIME_WINDOW_EXCEEDED"
	USERFACTORVERIFYRESULT_WAITING              UserFactorVerifyResult = "WAITING"
)

// All allowed values of UserFactorVerifyResult enum
var AllowedUserFactorVerifyResultEnumValues = []UserFactorVerifyResult{
	"CHALLENGE",
	"ERROR",
	"EXPIRED",
	"FAILED",
	"PASSCODE_REPLAYED",
	"REJECTED",
	"SUCCESS",
	"TIMEOUT",
	"TIME_WINDOW_EXCEEDED",
	"WAITING",
}

func (v *UserFactorVerifyResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserFactorVerifyResult(value)
	for _, existing := range AllowedUserFactorVerifyResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserFactorVerifyResult", value)
}

// NewUserFactorVerifyResultFromValue returns a pointer to a valid UserFactorVerifyResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserFactorVerifyResultFromValue(v string) (*UserFactorVerifyResult, error) {
	ev := UserFactorVerifyResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserFactorVerifyResult: valid values are %v", v, AllowedUserFactorVerifyResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserFactorVerifyResult) IsValid() bool {
	for _, existing := range AllowedUserFactorVerifyResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserFactorVerifyResult value
func (v UserFactorVerifyResult) Ptr() *UserFactorVerifyResult {
	return &v
}

type NullableUserFactorVerifyResult struct {
	value *UserFactorVerifyResult
	isSet bool
}

func (v NullableUserFactorVerifyResult) Get() *UserFactorVerifyResult {
	return v.value
}

func (v *NullableUserFactorVerifyResult) Set(val *UserFactorVerifyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorVerifyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorVerifyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorVerifyResult(val *UserFactorVerifyResult) *NullableUserFactorVerifyResult {
	return &NullableUserFactorVerifyResult{value: val, isSet: true}
}

func (v NullableUserFactorVerifyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorVerifyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}