// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLoadBalancersIDPoolsPoolIDParams creates a new GetLoadBalancersIDPoolsPoolIDParams object
// with the default values initialized.
func NewGetLoadBalancersIDPoolsPoolIDParams() *GetLoadBalancersIDPoolsPoolIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDParamsWithTimeout creates a new GetLoadBalancersIDPoolsPoolIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDPoolsPoolIDParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDPoolsPoolIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDParamsWithContext creates a new GetLoadBalancersIDPoolsPoolIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDPoolsPoolIDParamsWithContext(ctx context.Context) *GetLoadBalancersIDPoolsPoolIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDParamsWithHTTPClient creates a new GetLoadBalancersIDPoolsPoolIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDPoolsPoolIDParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDPoolsPoolIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDPoolsPoolIDParams contains all the parameters to send to the API endpoint
for the get load balancers ID pools pool ID operation typically these are written to a http.Request
*/
type GetLoadBalancersIDPoolsPoolIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*PoolID
	  The pool identifier

	*/
	PoolID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithContext(ctx context.Context) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithGeneration(generation int64) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithID(id string) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetID(id string) {
	o.ID = id
}

// WithPoolID adds the poolID to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithPoolID(poolID string) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithVersion adds the version to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) WithVersion(version string) *GetLoadBalancersIDPoolsPoolIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID pools pool ID params
func (o *GetLoadBalancersIDPoolsPoolIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDPoolsPoolIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param pool_id
	if err := r.SetPathParam("pool_id", o.PoolID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
