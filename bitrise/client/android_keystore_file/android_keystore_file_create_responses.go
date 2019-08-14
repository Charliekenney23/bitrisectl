// Code generated by go-swagger; DO NOT EDIT.

package android_keystore_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/charliekenney23/bitrisectl/var/work/bitrise/models"
)

// AndroidKeystoreFileCreateReader is a Reader for the AndroidKeystoreFileCreate structure.
type AndroidKeystoreFileCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AndroidKeystoreFileCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAndroidKeystoreFileCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAndroidKeystoreFileCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAndroidKeystoreFileCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAndroidKeystoreFileCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAndroidKeystoreFileCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAndroidKeystoreFileCreateCreated creates a AndroidKeystoreFileCreateCreated with default headers values
func NewAndroidKeystoreFileCreateCreated() *AndroidKeystoreFileCreateCreated {
	return &AndroidKeystoreFileCreateCreated{}
}

/*AndroidKeystoreFileCreateCreated handles this case with default header values.

Created
*/
type AndroidKeystoreFileCreateCreated struct {
	Payload *models.V0ProjectFileStorageResponseModel
}

func (o *AndroidKeystoreFileCreateCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileCreateCreated  %+v", 201, o.Payload)
}

func (o *AndroidKeystoreFileCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileCreateBadRequest creates a AndroidKeystoreFileCreateBadRequest with default headers values
func NewAndroidKeystoreFileCreateBadRequest() *AndroidKeystoreFileCreateBadRequest {
	return &AndroidKeystoreFileCreateBadRequest{}
}

/*AndroidKeystoreFileCreateBadRequest handles this case with default header values.

Bad Request
*/
type AndroidKeystoreFileCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AndroidKeystoreFileCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AndroidKeystoreFileCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileCreateUnauthorized creates a AndroidKeystoreFileCreateUnauthorized with default headers values
func NewAndroidKeystoreFileCreateUnauthorized() *AndroidKeystoreFileCreateUnauthorized {
	return &AndroidKeystoreFileCreateUnauthorized{}
}

/*AndroidKeystoreFileCreateUnauthorized handles this case with default header values.

Unauthorized
*/
type AndroidKeystoreFileCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AndroidKeystoreFileCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AndroidKeystoreFileCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileCreateNotFound creates a AndroidKeystoreFileCreateNotFound with default headers values
func NewAndroidKeystoreFileCreateNotFound() *AndroidKeystoreFileCreateNotFound {
	return &AndroidKeystoreFileCreateNotFound{}
}

/*AndroidKeystoreFileCreateNotFound handles this case with default header values.

Not Found
*/
type AndroidKeystoreFileCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AndroidKeystoreFileCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileCreateNotFound  %+v", 404, o.Payload)
}

func (o *AndroidKeystoreFileCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileCreateInternalServerError creates a AndroidKeystoreFileCreateInternalServerError with default headers values
func NewAndroidKeystoreFileCreateInternalServerError() *AndroidKeystoreFileCreateInternalServerError {
	return &AndroidKeystoreFileCreateInternalServerError{}
}

/*AndroidKeystoreFileCreateInternalServerError handles this case with default header values.

Internal Server Error
*/
type AndroidKeystoreFileCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AndroidKeystoreFileCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *AndroidKeystoreFileCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
