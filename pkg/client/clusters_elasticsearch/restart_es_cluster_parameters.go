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

// NewRestartEsClusterParams creates a new RestartEsClusterParams object
// with the default values initialized.
func NewRestartEsClusterParams() *RestartEsClusterParams {
	var (
		cancelPendingDefault     = bool(false)
		groupAttributeDefault    = string("__zone__")
		restoreSnapshotDefault   = bool(true)
		shardInitWaitTimeDefault = int64(600)
		skipSnapshotDefault      = bool(true)
	)
	return &RestartEsClusterParams{
		CancelPending:     &cancelPendingDefault,
		GroupAttribute:    &groupAttributeDefault,
		RestoreSnapshot:   &restoreSnapshotDefault,
		ShardInitWaitTime: &shardInitWaitTimeDefault,
		SkipSnapshot:      &skipSnapshotDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRestartEsClusterParamsWithTimeout creates a new RestartEsClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestartEsClusterParamsWithTimeout(timeout time.Duration) *RestartEsClusterParams {
	var (
		cancelPendingDefault     = bool(false)
		groupAttributeDefault    = string("__zone__")
		restoreSnapshotDefault   = bool(true)
		shardInitWaitTimeDefault = int64(600)
		skipSnapshotDefault      = bool(true)
	)
	return &RestartEsClusterParams{
		CancelPending:     &cancelPendingDefault,
		GroupAttribute:    &groupAttributeDefault,
		RestoreSnapshot:   &restoreSnapshotDefault,
		ShardInitWaitTime: &shardInitWaitTimeDefault,
		SkipSnapshot:      &skipSnapshotDefault,

		timeout: timeout,
	}
}

// NewRestartEsClusterParamsWithContext creates a new RestartEsClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestartEsClusterParamsWithContext(ctx context.Context) *RestartEsClusterParams {
	var (
		cancelPendingDefault     = bool(false)
		groupAttributeDefault    = string("__zone__")
		restoreSnapshotDefault   = bool(true)
		shardInitWaitTimeDefault = int64(600)
		skipSnapshotDefault      = bool(true)
	)
	return &RestartEsClusterParams{
		CancelPending:     &cancelPendingDefault,
		GroupAttribute:    &groupAttributeDefault,
		RestoreSnapshot:   &restoreSnapshotDefault,
		ShardInitWaitTime: &shardInitWaitTimeDefault,
		SkipSnapshot:      &skipSnapshotDefault,

		Context: ctx,
	}
}

// NewRestartEsClusterParamsWithHTTPClient creates a new RestartEsClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestartEsClusterParamsWithHTTPClient(client *http.Client) *RestartEsClusterParams {
	var (
		cancelPendingDefault     = bool(false)
		groupAttributeDefault    = string("__zone__")
		restoreSnapshotDefault   = bool(true)
		shardInitWaitTimeDefault = int64(600)
		skipSnapshotDefault      = bool(true)
	)
	return &RestartEsClusterParams{
		CancelPending:     &cancelPendingDefault,
		GroupAttribute:    &groupAttributeDefault,
		RestoreSnapshot:   &restoreSnapshotDefault,
		ShardInitWaitTime: &shardInitWaitTimeDefault,
		SkipSnapshot:      &skipSnapshotDefault,
		HTTPClient:        client,
	}
}

/*RestartEsClusterParams contains all the parameters to send to the API endpoint
for the restart es cluster operation typically these are written to a http.Request
*/
type RestartEsClusterParams struct {

	/*CancelPending
	  When `true`, cancels the pending plans, then restarts the cluster.

	*/
	CancelPending *bool
	/*ClusterID
	  The Elasticsearch cluster identifier.

	*/
	ClusterID string
	/*GroupAttribute
	  Specifies the properties that divide the instances into groups. To restart all of the instances, use '\_\_all\_\_'. To restart the instances by logical zone, use '\_\_zone\_\_'. To restart one instance at a time, use '\_\_name\_\_', or use a comma-separated list of instance attributes.

	*/
	GroupAttribute *string
	/*RestoreSnapshot
	  When `true` and restoring from a shutdown, restores the cluster from the last available snapshot.

	*/
	RestoreSnapshot *bool
	/*ShardInitWaitTime
	  The time, in seconds, to wait for shards that show no progress of initializing, before rolling the next group (default: 10 minutes)

	*/
	ShardInitWaitTime *int64
	/*SkipSnapshot
	  When `true`, does not capture a snapshot before restarting the cluster.

	*/
	SkipSnapshot *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restart es cluster params
func (o *RestartEsClusterParams) WithTimeout(timeout time.Duration) *RestartEsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart es cluster params
func (o *RestartEsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart es cluster params
func (o *RestartEsClusterParams) WithContext(ctx context.Context) *RestartEsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart es cluster params
func (o *RestartEsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart es cluster params
func (o *RestartEsClusterParams) WithHTTPClient(client *http.Client) *RestartEsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart es cluster params
func (o *RestartEsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCancelPending adds the cancelPending to the restart es cluster params
func (o *RestartEsClusterParams) WithCancelPending(cancelPending *bool) *RestartEsClusterParams {
	o.SetCancelPending(cancelPending)
	return o
}

// SetCancelPending adds the cancelPending to the restart es cluster params
func (o *RestartEsClusterParams) SetCancelPending(cancelPending *bool) {
	o.CancelPending = cancelPending
}

// WithClusterID adds the clusterID to the restart es cluster params
func (o *RestartEsClusterParams) WithClusterID(clusterID string) *RestartEsClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the restart es cluster params
func (o *RestartEsClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithGroupAttribute adds the groupAttribute to the restart es cluster params
func (o *RestartEsClusterParams) WithGroupAttribute(groupAttribute *string) *RestartEsClusterParams {
	o.SetGroupAttribute(groupAttribute)
	return o
}

// SetGroupAttribute adds the groupAttribute to the restart es cluster params
func (o *RestartEsClusterParams) SetGroupAttribute(groupAttribute *string) {
	o.GroupAttribute = groupAttribute
}

// WithRestoreSnapshot adds the restoreSnapshot to the restart es cluster params
func (o *RestartEsClusterParams) WithRestoreSnapshot(restoreSnapshot *bool) *RestartEsClusterParams {
	o.SetRestoreSnapshot(restoreSnapshot)
	return o
}

// SetRestoreSnapshot adds the restoreSnapshot to the restart es cluster params
func (o *RestartEsClusterParams) SetRestoreSnapshot(restoreSnapshot *bool) {
	o.RestoreSnapshot = restoreSnapshot
}

// WithShardInitWaitTime adds the shardInitWaitTime to the restart es cluster params
func (o *RestartEsClusterParams) WithShardInitWaitTime(shardInitWaitTime *int64) *RestartEsClusterParams {
	o.SetShardInitWaitTime(shardInitWaitTime)
	return o
}

// SetShardInitWaitTime adds the shardInitWaitTime to the restart es cluster params
func (o *RestartEsClusterParams) SetShardInitWaitTime(shardInitWaitTime *int64) {
	o.ShardInitWaitTime = shardInitWaitTime
}

// WithSkipSnapshot adds the skipSnapshot to the restart es cluster params
func (o *RestartEsClusterParams) WithSkipSnapshot(skipSnapshot *bool) *RestartEsClusterParams {
	o.SetSkipSnapshot(skipSnapshot)
	return o
}

// SetSkipSnapshot adds the skipSnapshot to the restart es cluster params
func (o *RestartEsClusterParams) SetSkipSnapshot(skipSnapshot *bool) {
	o.SkipSnapshot = skipSnapshot
}

// WriteToRequest writes these params to a swagger request
func (o *RestartEsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CancelPending != nil {

		// query param cancel_pending
		var qrCancelPending bool
		if o.CancelPending != nil {
			qrCancelPending = *o.CancelPending
		}
		qCancelPending := swag.FormatBool(qrCancelPending)
		if qCancelPending != "" {
			if err := r.SetQueryParam("cancel_pending", qCancelPending); err != nil {
				return err
			}
		}

	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.GroupAttribute != nil {

		// query param group_attribute
		var qrGroupAttribute string
		if o.GroupAttribute != nil {
			qrGroupAttribute = *o.GroupAttribute
		}
		qGroupAttribute := qrGroupAttribute
		if qGroupAttribute != "" {
			if err := r.SetQueryParam("group_attribute", qGroupAttribute); err != nil {
				return err
			}
		}

	}

	if o.RestoreSnapshot != nil {

		// query param restore_snapshot
		var qrRestoreSnapshot bool
		if o.RestoreSnapshot != nil {
			qrRestoreSnapshot = *o.RestoreSnapshot
		}
		qRestoreSnapshot := swag.FormatBool(qrRestoreSnapshot)
		if qRestoreSnapshot != "" {
			if err := r.SetQueryParam("restore_snapshot", qRestoreSnapshot); err != nil {
				return err
			}
		}

	}

	if o.ShardInitWaitTime != nil {

		// query param shard_init_wait_time
		var qrShardInitWaitTime int64
		if o.ShardInitWaitTime != nil {
			qrShardInitWaitTime = *o.ShardInitWaitTime
		}
		qShardInitWaitTime := swag.FormatInt64(qrShardInitWaitTime)
		if qShardInitWaitTime != "" {
			if err := r.SetQueryParam("shard_init_wait_time", qShardInitWaitTime); err != nil {
				return err
			}
		}

	}

	if o.SkipSnapshot != nil {

		// query param skip_snapshot
		var qrSkipSnapshot bool
		if o.SkipSnapshot != nil {
			qrSkipSnapshot = *o.SkipSnapshot
		}
		qSkipSnapshot := swag.FormatBool(qrSkipSnapshot)
		if qSkipSnapshot != "" {
			if err := r.SetQueryParam("skip_snapshot", qSkipSnapshot); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
