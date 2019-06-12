// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetVolumeProfilesReader is a Reader for the GetVolumeProfiles structure.
type GetVolumeProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVolumeProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetVolumeProfilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVolumeProfilesOK creates a GetVolumeProfilesOK with default headers values
func NewGetVolumeProfilesOK() *GetVolumeProfilesOK {
	return &GetVolumeProfilesOK{}
}

/*GetVolumeProfilesOK handles this case with default header values.

dummy
*/
type GetVolumeProfilesOK struct {
	Payload *models.GetVolumeProfilesOKBody
}

func (o *GetVolumeProfilesOK) Error() string {
	return fmt.Sprintf("[GET /volume/profiles][%d] getVolumeProfilesOK  %+v", 200, o.Payload)
}

func (o *GetVolumeProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVolumeProfilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeProfilesInternalServerError creates a GetVolumeProfilesInternalServerError with default headers values
func NewGetVolumeProfilesInternalServerError() *GetVolumeProfilesInternalServerError {
	return &GetVolumeProfilesInternalServerError{}
}

/*GetVolumeProfilesInternalServerError handles this case with default header values.

error
*/
type GetVolumeProfilesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetVolumeProfilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volume/profiles][%d] getVolumeProfilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumeProfilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}