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

// UserSchemaAttributeType the model 'UserSchemaAttributeType'
type UserSchemaAttributeType string

// List of UserSchemaAttributeType
const (
	USERSCHEMAATTRIBUTETYPE_ARRAY   UserSchemaAttributeType = "array"
	USERSCHEMAATTRIBUTETYPE_BOOLEAN UserSchemaAttributeType = "boolean"
	USERSCHEMAATTRIBUTETYPE_INTEGER UserSchemaAttributeType = "integer"
	USERSCHEMAATTRIBUTETYPE_NUMBER  UserSchemaAttributeType = "number"
	USERSCHEMAATTRIBUTETYPE_STRING  UserSchemaAttributeType = "string"
)

// All allowed values of UserSchemaAttributeType enum
var AllowedUserSchemaAttributeTypeEnumValues = []UserSchemaAttributeType{
	"array",
	"boolean",
	"integer",
	"number",
	"string",
}

func (v *UserSchemaAttributeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserSchemaAttributeType(value)
	for _, existing := range AllowedUserSchemaAttributeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserSchemaAttributeType", value)
}

// NewUserSchemaAttributeTypeFromValue returns a pointer to a valid UserSchemaAttributeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserSchemaAttributeTypeFromValue(v string) (*UserSchemaAttributeType, error) {
	ev := UserSchemaAttributeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserSchemaAttributeType: valid values are %v", v, AllowedUserSchemaAttributeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserSchemaAttributeType) IsValid() bool {
	for _, existing := range AllowedUserSchemaAttributeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserSchemaAttributeType value
func (v UserSchemaAttributeType) Ptr() *UserSchemaAttributeType {
	return &v
}

type NullableUserSchemaAttributeType struct {
	value *UserSchemaAttributeType
	isSet bool
}

func (v NullableUserSchemaAttributeType) Get() *UserSchemaAttributeType {
	return v.value
}

func (v *NullableUserSchemaAttributeType) Set(val *UserSchemaAttributeType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttributeType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttributeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttributeType(val *UserSchemaAttributeType) *NullableUserSchemaAttributeType {
	return &NullableUserSchemaAttributeType{value: val, isSet: true}
}

func (v NullableUserSchemaAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttributeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
