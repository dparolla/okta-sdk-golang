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

// LogAuthenticationProvider the model 'LogAuthenticationProvider'
type LogAuthenticationProvider string

// List of LogAuthenticationProvider
const (
	LOGAUTHENTICATIONPROVIDER_ACTIVE_DIRECTORY             LogAuthenticationProvider = "ACTIVE_DIRECTORY"
	LOGAUTHENTICATIONPROVIDER_FACTOR_PROVIDER              LogAuthenticationProvider = "FACTOR_PROVIDER"
	LOGAUTHENTICATIONPROVIDER_FEDERATION                   LogAuthenticationProvider = "FEDERATION"
	LOGAUTHENTICATIONPROVIDER_LDAP                         LogAuthenticationProvider = "LDAP"
	LOGAUTHENTICATIONPROVIDER_OKTA_AUTHENTICATION_PROVIDER LogAuthenticationProvider = "OKTA_AUTHENTICATION_PROVIDER"
	LOGAUTHENTICATIONPROVIDER_SOCIAL                       LogAuthenticationProvider = "SOCIAL"
)

// All allowed values of LogAuthenticationProvider enum
var AllowedLogAuthenticationProviderEnumValues = []LogAuthenticationProvider{
	"ACTIVE_DIRECTORY",
	"FACTOR_PROVIDER",
	"FEDERATION",
	"LDAP",
	"OKTA_AUTHENTICATION_PROVIDER",
	"SOCIAL",
}

func (v *LogAuthenticationProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogAuthenticationProvider(value)
	for _, existing := range AllowedLogAuthenticationProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogAuthenticationProvider", value)
}

// NewLogAuthenticationProviderFromValue returns a pointer to a valid LogAuthenticationProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogAuthenticationProviderFromValue(v string) (*LogAuthenticationProvider, error) {
	ev := LogAuthenticationProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogAuthenticationProvider: valid values are %v", v, AllowedLogAuthenticationProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogAuthenticationProvider) IsValid() bool {
	for _, existing := range AllowedLogAuthenticationProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogAuthenticationProvider value
func (v LogAuthenticationProvider) Ptr() *LogAuthenticationProvider {
	return &v
}

type NullableLogAuthenticationProvider struct {
	value *LogAuthenticationProvider
	isSet bool
}

func (v NullableLogAuthenticationProvider) Get() *LogAuthenticationProvider {
	return v.value
}

func (v *NullableLogAuthenticationProvider) Set(val *LogAuthenticationProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableLogAuthenticationProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableLogAuthenticationProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogAuthenticationProvider(val *LogAuthenticationProvider) *NullableLogAuthenticationProvider {
	return &NullableLogAuthenticationProvider{value: val, isSet: true}
}

func (v NullableLogAuthenticationProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogAuthenticationProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
