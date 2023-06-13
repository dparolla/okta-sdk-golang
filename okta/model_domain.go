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

// Domain struct for Domain
type Domain struct {
	BrandId               *string                      `json:"brandId,omitempty"`
	CertificateSourceType *DomainCertificateSourceType `json:"certificateSourceType,omitempty"`
	DnsRecords            []DNSRecord                  `json:"dnsRecords,omitempty"`
	Domain                *string                      `json:"domain,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	PublicCertificate     *DomainCertificateMetadata   `json:"publicCertificate,omitempty"`
	ValidationStatus      *DomainValidationStatus      `json:"validationStatus,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _Domain Domain

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *Domain) GetBrandId() string {
	if o == nil || o.BrandId == nil {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetBrandIdOk() (*string, bool) {
	if o == nil || o.BrandId == nil {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *Domain) HasBrandId() bool {
	if o != nil && o.BrandId != nil {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *Domain) SetBrandId(v string) {
	o.BrandId = &v
}

// GetCertificateSourceType returns the CertificateSourceType field value if set, zero value otherwise.
func (o *Domain) GetCertificateSourceType() DomainCertificateSourceType {
	if o == nil || o.CertificateSourceType == nil {
		var ret DomainCertificateSourceType
		return ret
	}
	return *o.CertificateSourceType
}

// GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCertificateSourceTypeOk() (*DomainCertificateSourceType, bool) {
	if o == nil || o.CertificateSourceType == nil {
		return nil, false
	}
	return o.CertificateSourceType, true
}

// HasCertificateSourceType returns a boolean if a field has been set.
func (o *Domain) HasCertificateSourceType() bool {
	if o != nil && o.CertificateSourceType != nil {
		return true
	}

	return false
}

// SetCertificateSourceType gets a reference to the given DomainCertificateSourceType and assigns it to the CertificateSourceType field.
func (o *Domain) SetCertificateSourceType(v DomainCertificateSourceType) {
	o.CertificateSourceType = &v
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *Domain) GetDnsRecords() []DNSRecord {
	if o == nil || o.DnsRecords == nil {
		var ret []DNSRecord
		return ret
	}
	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsRecordsOk() ([]DNSRecord, bool) {
	if o == nil || o.DnsRecords == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *Domain) HasDnsRecords() bool {
	if o != nil && o.DnsRecords != nil {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given []DNSRecord and assigns it to the DnsRecords field.
func (o *Domain) SetDnsRecords(v []DNSRecord) {
	o.DnsRecords = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Domain) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Domain) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Domain) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Domain) SetId(v string) {
	o.Id = &v
}

// GetPublicCertificate returns the PublicCertificate field value if set, zero value otherwise.
func (o *Domain) GetPublicCertificate() DomainCertificateMetadata {
	if o == nil || o.PublicCertificate == nil {
		var ret DomainCertificateMetadata
		return ret
	}
	return *o.PublicCertificate
}

// GetPublicCertificateOk returns a tuple with the PublicCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPublicCertificateOk() (*DomainCertificateMetadata, bool) {
	if o == nil || o.PublicCertificate == nil {
		return nil, false
	}
	return o.PublicCertificate, true
}

// HasPublicCertificate returns a boolean if a field has been set.
func (o *Domain) HasPublicCertificate() bool {
	if o != nil && o.PublicCertificate != nil {
		return true
	}

	return false
}

// SetPublicCertificate gets a reference to the given DomainCertificateMetadata and assigns it to the PublicCertificate field.
func (o *Domain) SetPublicCertificate(v DomainCertificateMetadata) {
	o.PublicCertificate = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *Domain) GetValidationStatus() DomainValidationStatus {
	if o == nil || o.ValidationStatus == nil {
		var ret DomainValidationStatus
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetValidationStatusOk() (*DomainValidationStatus, bool) {
	if o == nil || o.ValidationStatus == nil {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *Domain) HasValidationStatus() bool {
	if o != nil && o.ValidationStatus != nil {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given DomainValidationStatus and assigns it to the ValidationStatus field.
func (o *Domain) SetValidationStatus(v DomainValidationStatus) {
	o.ValidationStatus = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BrandId != nil {
		toSerialize["brandId"] = o.BrandId
	}
	if o.CertificateSourceType != nil {
		toSerialize["certificateSourceType"] = o.CertificateSourceType
	}
	if o.DnsRecords != nil {
		toSerialize["dnsRecords"] = o.DnsRecords
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PublicCertificate != nil {
		toSerialize["publicCertificate"] = o.PublicCertificate
	}
	if o.ValidationStatus != nil {
		toSerialize["validationStatus"] = o.ValidationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Domain) UnmarshalJSON(bytes []byte) (err error) {
	varDomain := _Domain{}

	err = json.Unmarshal(bytes, &varDomain)
	if err == nil {
		*o = Domain(varDomain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "brandId")
		delete(additionalProperties, "certificateSourceType")
		delete(additionalProperties, "dnsRecords")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "id")
		delete(additionalProperties, "publicCertificate")
		delete(additionalProperties, "validationStatus")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}