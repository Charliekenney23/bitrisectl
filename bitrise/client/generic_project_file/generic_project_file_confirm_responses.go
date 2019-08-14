// Code generated by go-swagger; DO NOT EDIT.

package generic_project_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/charliekenney23/bitrisectl/var/work/bitrise/models"
)

// GenericProjectFileConfirmReader is a Reader for the GenericProjectFileConfirm structure.
type GenericProjectFileConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenericProjectFileConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenericProjectFileConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGenericProjectFileConfirmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGenericProjectFileConfirmUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGenericProjectFileConfirmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGenericProjectFileConfirmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenericProjectFileConfirmOK creates a GenericProjectFileConfirmOK with default headers values
func NewGenericProjectFileConfirmOK() *GenericProjectFileConfirmOK {
	return &GenericProjectFileConfirmOK{}
}

/*GenericProjectFileConfirmOK handles this case with default header values.

OK
*/
type GenericProjectFileConfirmOK struct {
	Payload *models.V0ProjectFileStorageResponseModel
}

func (o *GenericProjectFileConfirmOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded][%d] genericProjectFileConfirmOK  %+v", 200, o.Payload)
}

func (o *GenericProjectFileConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileConfirmBadRequest creates a GenericProjectFileConfirmBadRequest with default headers values
func NewGenericProjectFileConfirmBadRequest() *GenericProjectFileConfirmBadRequest {
	return &GenericProjectFileConfirmBadRequest{}
}

/*GenericProjectFileConfirmBadRequest handles this case with default header values.

Bad Request
*/
type GenericProjectFileConfirmBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileConfirmBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded][%d] genericProjectFileConfirmBadRequest  %+v", 400, o.Payload)
}

func (o *GenericProjectFileConfirmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileConfirmUnauthorized creates a GenericProjectFileConfirmUnauthorized with default headers values
func NewGenericProjectFileConfirmUnauthorized() *GenericProjectFileConfirmUnauthorized {
	return &GenericProjectFileConfirmUnauthorized{}
}

/*GenericProjectFileConfirmUnauthorized handles this case with default header values.

Unauthorized
*/
type GenericProjectFileConfirmUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileConfirmUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded][%d] genericProjectFileConfirmUnauthorized  %+v", 401, o.Payload)
}

func (o *GenericProjectFileConfirmUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileConfirmNotFound creates a GenericProjectFileConfirmNotFound with default headers values
func NewGenericProjectFileConfirmNotFound() *GenericProjectFileConfirmNotFound {
	return &GenericProjectFileConfirmNotFound{}
}

/*GenericProjectFileConfirmNotFound handles this case with default header values.

Not Found
*/
type GenericProjectFileConfirmNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileConfirmNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded][%d] genericProjectFileConfirmNotFound  %+v", 404, o.Payload)
}

func (o *GenericProjectFileConfirmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileConfirmInternalServerError creates a GenericProjectFileConfirmInternalServerError with default headers values
func NewGenericProjectFileConfirmInternalServerError() *GenericProjectFileConfirmInternalServerError {
	return &GenericProjectFileConfirmInternalServerError{}
}

/*GenericProjectFileConfirmInternalServerError handles this case with default header values.

Internal Server Error
*/
type GenericProjectFileConfirmInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileConfirmInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded][%d] genericProjectFileConfirmInternalServerError  %+v", 500, o.Payload)
}

func (o *GenericProjectFileConfirmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
