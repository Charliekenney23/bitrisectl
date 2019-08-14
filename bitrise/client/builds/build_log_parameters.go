// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBuildLogParams creates a new BuildLogParams object
// with the default values initialized.
func NewBuildLogParams() *BuildLogParams {
	var ()
	return &BuildLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildLogParamsWithTimeout creates a new BuildLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildLogParamsWithTimeout(timeout time.Duration) *BuildLogParams {
	var ()
	return &BuildLogParams{

		timeout: timeout,
	}
}

// NewBuildLogParamsWithContext creates a new BuildLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildLogParamsWithContext(ctx context.Context) *BuildLogParams {
	var ()
	return &BuildLogParams{

		Context: ctx,
	}
}

// NewBuildLogParamsWithHTTPClient creates a new BuildLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildLogParamsWithHTTPClient(client *http.Client) *BuildLogParams {
	var ()
	return &BuildLogParams{
		HTTPClient: client,
	}
}

/*BuildLogParams contains all the parameters to send to the API endpoint
for the build log operation typically these are written to a http.Request
*/
type BuildLogParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*BuildSlug
	  Build slug

	*/
	BuildSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build log params
func (o *BuildLogParams) WithTimeout(timeout time.Duration) *BuildLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build log params
func (o *BuildLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build log params
func (o *BuildLogParams) WithContext(ctx context.Context) *BuildLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build log params
func (o *BuildLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build log params
func (o *BuildLogParams) WithHTTPClient(client *http.Client) *BuildLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build log params
func (o *BuildLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build log params
func (o *BuildLogParams) WithAppSlug(appSlug string) *BuildLogParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build log params
func (o *BuildLogParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildSlug adds the buildSlug to the build log params
func (o *BuildLogParams) WithBuildSlug(buildSlug string) *BuildLogParams {
	o.SetBuildSlug(buildSlug)
	return o
}

// SetBuildSlug adds the buildSlug to the build log params
func (o *BuildLogParams) SetBuildSlug(buildSlug string) {
	o.BuildSlug = buildSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param build-slug
	if err := r.SetPathParam("build-slug", o.BuildSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
