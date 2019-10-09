// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteClusterRoleBindingParams creates a new DeleteClusterRoleBindingParams object
// with the default values initialized.
func NewDeleteClusterRoleBindingParams() *DeleteClusterRoleBindingParams {
	var ()
	return &DeleteClusterRoleBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterRoleBindingParamsWithTimeout creates a new DeleteClusterRoleBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterRoleBindingParamsWithTimeout(timeout time.Duration) *DeleteClusterRoleBindingParams {
	var ()
	return &DeleteClusterRoleBindingParams{

		timeout: timeout,
	}
}

// NewDeleteClusterRoleBindingParamsWithContext creates a new DeleteClusterRoleBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterRoleBindingParamsWithContext(ctx context.Context) *DeleteClusterRoleBindingParams {
	var ()
	return &DeleteClusterRoleBindingParams{

		Context: ctx,
	}
}

// NewDeleteClusterRoleBindingParamsWithHTTPClient creates a new DeleteClusterRoleBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterRoleBindingParamsWithHTTPClient(client *http.Client) *DeleteClusterRoleBindingParams {
	var ()
	return &DeleteClusterRoleBindingParams{
		HTTPClient: client,
	}
}

/*DeleteClusterRoleBindingParams contains all the parameters to send to the API endpoint
for the delete cluster role binding operation typically these are written to a http.Request
*/
type DeleteClusterRoleBindingParams struct {

	/*BindingID*/
	BindingID string
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*ProjectID*/
	ProjectID string
	/*RoleID*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithTimeout(timeout time.Duration) *DeleteClusterRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithContext(ctx context.Context) *DeleteClusterRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithHTTPClient(client *http.Client) *DeleteClusterRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBindingID adds the bindingID to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithBindingID(bindingID string) *DeleteClusterRoleBindingParams {
	o.SetBindingID(bindingID)
	return o
}

// SetBindingID adds the bindingId to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetBindingID(bindingID string) {
	o.BindingID = bindingID
}

// WithClusterID adds the clusterID to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithClusterID(clusterID string) *DeleteClusterRoleBindingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithDc(dc string) *DeleteClusterRoleBindingParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithProjectID(projectID string) *DeleteClusterRoleBindingParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) WithRoleID(roleID string) *DeleteClusterRoleBindingParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the delete cluster role binding params
func (o *DeleteClusterRoleBindingParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param binding_id
	if err := r.SetPathParam("binding_id", o.BindingID); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}