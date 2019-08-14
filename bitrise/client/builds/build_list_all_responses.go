// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/charliekenney23/bitrisectl/var/work/bitrise/models"
)

// BuildListAllReader is a Reader for the BuildListAll structure.
type BuildListAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildListAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildListAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildListAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBuildListAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBuildListAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBuildListAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBuildListAllOK creates a BuildListAllOK with default headers values
func NewBuildListAllOK() *BuildListAllOK {
	return &BuildListAllOK{}
}

/*BuildListAllOK handles this case with default header values.

OK
*/
type BuildListAllOK struct {
	Payload *models.V0BuildListAllResponseModel
}

func (o *BuildListAllOK) Error() string {
	return fmt.Sprintf("[GET /builds][%d] buildListAllOK  %+v", 200, o.Payload)
}

func (o *BuildListAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildListAllResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildListAllBadRequest creates a BuildListAllBadRequest with default headers values
func NewBuildListAllBadRequest() *BuildListAllBadRequest {
	return &BuildListAllBadRequest{}
}

/*BuildListAllBadRequest handles this case with default header values.

Bad Request
*/
type BuildListAllBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildListAllBadRequest) Error() string {
	return fmt.Sprintf("[GET /builds][%d] buildListAllBadRequest  %+v", 400, o.Payload)
}

func (o *BuildListAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildListAllUnauthorized creates a BuildListAllUnauthorized with default headers values
func NewBuildListAllUnauthorized() *BuildListAllUnauthorized {
	return &BuildListAllUnauthorized{}
}

/*BuildListAllUnauthorized handles this case with default header values.

Unauthorized
*/
type BuildListAllUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildListAllUnauthorized) Error() string {
	return fmt.Sprintf("[GET /builds][%d] buildListAllUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildListAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildListAllNotFound creates a BuildListAllNotFound with default headers values
func NewBuildListAllNotFound() *BuildListAllNotFound {
	return &BuildListAllNotFound{}
}

/*BuildListAllNotFound handles this case with default header values.

Not Found
*/
type BuildListAllNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildListAllNotFound) Error() string {
	return fmt.Sprintf("[GET /builds][%d] buildListAllNotFound  %+v", 404, o.Payload)
}

func (o *BuildListAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildListAllInternalServerError creates a BuildListAllInternalServerError with default headers values
func NewBuildListAllInternalServerError() *BuildListAllInternalServerError {
	return &BuildListAllInternalServerError{}
}

/*BuildListAllInternalServerError handles this case with default header values.

Internal Server Error
*/
type BuildListAllInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildListAllInternalServerError) Error() string {
	return fmt.Sprintf("[GET /builds][%d] buildListAllInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildListAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
