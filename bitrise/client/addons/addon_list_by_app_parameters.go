// Code generated by go-swagger; DO NOT EDIT.

package addons

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

// NewAddonListByAppParams creates a new AddonListByAppParams object
// with the default values initialized.
func NewAddonListByAppParams() *AddonListByAppParams {
	var ()
	return &AddonListByAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddonListByAppParamsWithTimeout creates a new AddonListByAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddonListByAppParamsWithTimeout(timeout time.Duration) *AddonListByAppParams {
	var ()
	return &AddonListByAppParams{

		timeout: timeout,
	}
}

// NewAddonListByAppParamsWithContext creates a new AddonListByAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddonListByAppParamsWithContext(ctx context.Context) *AddonListByAppParams {
	var ()
	return &AddonListByAppParams{

		Context: ctx,
	}
}

// NewAddonListByAppParamsWithHTTPClient creates a new AddonListByAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddonListByAppParamsWithHTTPClient(client *http.Client) *AddonListByAppParams {
	var ()
	return &AddonListByAppParams{
		HTTPClient: client,
	}
}

/*AddonListByAppParams contains all the parameters to send to the API endpoint
for the addon list by app operation typically these are written to a http.Request
*/
type AddonListByAppParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the addon list by app params
func (o *AddonListByAppParams) WithTimeout(timeout time.Duration) *AddonListByAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the addon list by app params
func (o *AddonListByAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the addon list by app params
func (o *AddonListByAppParams) WithContext(ctx context.Context) *AddonListByAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the addon list by app params
func (o *AddonListByAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the addon list by app params
func (o *AddonListByAppParams) WithHTTPClient(client *http.Client) *AddonListByAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the addon list by app params
func (o *AddonListByAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the addon list by app params
func (o *AddonListByAppParams) WithAppSlug(appSlug string) *AddonListByAppParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the addon list by app params
func (o *AddonListByAppParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AddonListByAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
