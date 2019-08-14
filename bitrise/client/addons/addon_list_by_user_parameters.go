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

// NewAddonListByUserParams creates a new AddonListByUserParams object
// with the default values initialized.
func NewAddonListByUserParams() *AddonListByUserParams {
	var ()
	return &AddonListByUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddonListByUserParamsWithTimeout creates a new AddonListByUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddonListByUserParamsWithTimeout(timeout time.Duration) *AddonListByUserParams {
	var ()
	return &AddonListByUserParams{

		timeout: timeout,
	}
}

// NewAddonListByUserParamsWithContext creates a new AddonListByUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddonListByUserParamsWithContext(ctx context.Context) *AddonListByUserParams {
	var ()
	return &AddonListByUserParams{

		Context: ctx,
	}
}

// NewAddonListByUserParamsWithHTTPClient creates a new AddonListByUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddonListByUserParamsWithHTTPClient(client *http.Client) *AddonListByUserParams {
	var ()
	return &AddonListByUserParams{
		HTTPClient: client,
	}
}

/*AddonListByUserParams contains all the parameters to send to the API endpoint
for the addon list by user operation typically these are written to a http.Request
*/
type AddonListByUserParams struct {

	/*UserSlug
	  User slug

	*/
	UserSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the addon list by user params
func (o *AddonListByUserParams) WithTimeout(timeout time.Duration) *AddonListByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the addon list by user params
func (o *AddonListByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the addon list by user params
func (o *AddonListByUserParams) WithContext(ctx context.Context) *AddonListByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the addon list by user params
func (o *AddonListByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the addon list by user params
func (o *AddonListByUserParams) WithHTTPClient(client *http.Client) *AddonListByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the addon list by user params
func (o *AddonListByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserSlug adds the userSlug to the addon list by user params
func (o *AddonListByUserParams) WithUserSlug(userSlug string) *AddonListByUserParams {
	o.SetUserSlug(userSlug)
	return o
}

// SetUserSlug adds the userSlug to the addon list by user params
func (o *AddonListByUserParams) SetUserSlug(userSlug string) {
	o.UserSlug = userSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AddonListByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user-slug
	if err := r.SetPathParam("user-slug", o.UserSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
