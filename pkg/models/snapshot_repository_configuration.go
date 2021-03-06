// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotRepositoryConfiguration The configuration for an Elasticsearch snapshot repository.
//
// swagger:model SnapshotRepositoryConfiguration
type SnapshotRepositoryConfiguration struct {

	// Elasticsearch repository configuration settings JSON. See [Elasticsearch documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html) for more information
	// Required: true
	Settings interface{} `json:"settings"`

	// Repository type, (Currently supported: 's3')
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this snapshot repository configuration
func (m *SnapshotRepositoryConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotRepositoryConfiguration) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotRepositoryConfiguration) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotRepositoryConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotRepositoryConfiguration) UnmarshalBinary(b []byte) error {
	var res SnapshotRepositoryConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
