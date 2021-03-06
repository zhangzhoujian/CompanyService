// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "CompanyService/openapi/gen/companyservice/models"
)

// GetAPIV1CompaniesCompanyIDOKCode is the HTTP code returned for type GetAPIV1CompaniesCompanyIDOK
const GetAPIV1CompaniesCompanyIDOKCode int = 200

/*GetAPIV1CompaniesCompanyIDOK Status 200

swagger:response getApiV1CompaniesCompanyIdOK
*/
type GetAPIV1CompaniesCompanyIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Company `json:"body,omitempty"`
}

// NewGetAPIV1CompaniesCompanyIDOK creates GetAPIV1CompaniesCompanyIDOK with default headers values
func NewGetAPIV1CompaniesCompanyIDOK() *GetAPIV1CompaniesCompanyIDOK {

	return &GetAPIV1CompaniesCompanyIDOK{}
}

// WithPayload adds the payload to the get Api v1 companies company Id o k response
func (o *GetAPIV1CompaniesCompanyIDOK) WithPayload(payload *models.Company) *GetAPIV1CompaniesCompanyIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api v1 companies company Id o k response
func (o *GetAPIV1CompaniesCompanyIDOK) SetPayload(payload *models.Company) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIV1CompaniesCompanyIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIV1CompaniesCompanyIDBadRequestCode is the HTTP code returned for type GetAPIV1CompaniesCompanyIDBadRequest
const GetAPIV1CompaniesCompanyIDBadRequestCode int = 400

/*GetAPIV1CompaniesCompanyIDBadRequest Status 400

swagger:response getApiV1CompaniesCompanyIdBadRequest
*/
type GetAPIV1CompaniesCompanyIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIV1CompaniesCompanyIDBadRequest creates GetAPIV1CompaniesCompanyIDBadRequest with default headers values
func NewGetAPIV1CompaniesCompanyIDBadRequest() *GetAPIV1CompaniesCompanyIDBadRequest {

	return &GetAPIV1CompaniesCompanyIDBadRequest{}
}

// WithPayload adds the payload to the get Api v1 companies company Id bad request response
func (o *GetAPIV1CompaniesCompanyIDBadRequest) WithPayload(payload *models.Error) *GetAPIV1CompaniesCompanyIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api v1 companies company Id bad request response
func (o *GetAPIV1CompaniesCompanyIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIV1CompaniesCompanyIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIV1CompaniesCompanyIDNotFoundCode is the HTTP code returned for type GetAPIV1CompaniesCompanyIDNotFound
const GetAPIV1CompaniesCompanyIDNotFoundCode int = 404

/*GetAPIV1CompaniesCompanyIDNotFound Status 404

swagger:response getApiV1CompaniesCompanyIdNotFound
*/
type GetAPIV1CompaniesCompanyIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIV1CompaniesCompanyIDNotFound creates GetAPIV1CompaniesCompanyIDNotFound with default headers values
func NewGetAPIV1CompaniesCompanyIDNotFound() *GetAPIV1CompaniesCompanyIDNotFound {

	return &GetAPIV1CompaniesCompanyIDNotFound{}
}

// WithPayload adds the payload to the get Api v1 companies company Id not found response
func (o *GetAPIV1CompaniesCompanyIDNotFound) WithPayload(payload *models.Error) *GetAPIV1CompaniesCompanyIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api v1 companies company Id not found response
func (o *GetAPIV1CompaniesCompanyIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIV1CompaniesCompanyIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIV1CompaniesCompanyIDInternalServerErrorCode is the HTTP code returned for type GetAPIV1CompaniesCompanyIDInternalServerError
const GetAPIV1CompaniesCompanyIDInternalServerErrorCode int = 500

/*GetAPIV1CompaniesCompanyIDInternalServerError Status 500

swagger:response getApiV1CompaniesCompanyIdInternalServerError
*/
type GetAPIV1CompaniesCompanyIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIV1CompaniesCompanyIDInternalServerError creates GetAPIV1CompaniesCompanyIDInternalServerError with default headers values
func NewGetAPIV1CompaniesCompanyIDInternalServerError() *GetAPIV1CompaniesCompanyIDInternalServerError {

	return &GetAPIV1CompaniesCompanyIDInternalServerError{}
}

// WithPayload adds the payload to the get Api v1 companies company Id internal server error response
func (o *GetAPIV1CompaniesCompanyIDInternalServerError) WithPayload(payload *models.Error) *GetAPIV1CompaniesCompanyIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api v1 companies company Id internal server error response
func (o *GetAPIV1CompaniesCompanyIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIV1CompaniesCompanyIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
