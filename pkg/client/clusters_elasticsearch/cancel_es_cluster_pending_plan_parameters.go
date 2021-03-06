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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCancelEsClusterPendingPlanParams creates a new CancelEsClusterPendingPlanParams object
// with the default values initialized.
func NewCancelEsClusterPendingPlanParams() *CancelEsClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelEsClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelEsClusterPendingPlanParamsWithTimeout creates a new CancelEsClusterPendingPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelEsClusterPendingPlanParamsWithTimeout(timeout time.Duration) *CancelEsClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelEsClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		timeout: timeout,
	}
}

// NewCancelEsClusterPendingPlanParamsWithContext creates a new CancelEsClusterPendingPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelEsClusterPendingPlanParamsWithContext(ctx context.Context) *CancelEsClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelEsClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		Context: ctx,
	}
}

// NewCancelEsClusterPendingPlanParamsWithHTTPClient creates a new CancelEsClusterPendingPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelEsClusterPendingPlanParamsWithHTTPClient(client *http.Client) *CancelEsClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelEsClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,
		HTTPClient:    client,
	}
}

/*CancelEsClusterPendingPlanParams contains all the parameters to send to the API endpoint
for the cancel es cluster pending plan operation typically these are written to a http.Request
*/
type CancelEsClusterPendingPlanParams struct {

	/*ClusterID
	  The Elasticsearch cluster identifier.

	*/
	ClusterID string
	/*ForceDelete
	  "When `true`, deletes the pending plan instead of attempting a graceful cancellation. The default is `false`.

	*/
	ForceDelete *bool
	/*IgnoreMissing
	  When `true`, returns successfully, even when plans are pending. The default is `false`.

	*/
	IgnoreMissing *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithTimeout(timeout time.Duration) *CancelEsClusterPendingPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithContext(ctx context.Context) *CancelEsClusterPendingPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithHTTPClient(client *http.Client) *CancelEsClusterPendingPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithClusterID(clusterID string) *CancelEsClusterPendingPlanParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceDelete adds the forceDelete to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithForceDelete(forceDelete *bool) *CancelEsClusterPendingPlanParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithIgnoreMissing adds the ignoreMissing to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) WithIgnoreMissing(ignoreMissing *bool) *CancelEsClusterPendingPlanParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the cancel es cluster pending plan params
func (o *CancelEsClusterPendingPlanParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WriteToRequest writes these params to a swagger request
func (o *CancelEsClusterPendingPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ForceDelete != nil {

		// query param force_delete
		var qrForceDelete bool
		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {
			if err := r.SetQueryParam("force_delete", qForceDelete); err != nil {
				return err
			}
		}

	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
