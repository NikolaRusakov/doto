// Code generated by go-swagger; DO NOT EDIT.

package goal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rusakov/doto/server/api/models"
)

// GoalsSectionQueryOKCode is the HTTP code returned for type GoalsSectionQueryOK
const GoalsSectionQueryOKCode int = 200

/*GoalsSectionQueryOK OK

swagger:response goalsSectionQueryOK
*/
type GoalsSectionQueryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Goal `json:"body,omitempty"`
}

// NewGoalsSectionQueryOK creates GoalsSectionQueryOK with default headers values
func NewGoalsSectionQueryOK() *GoalsSectionQueryOK {

	return &GoalsSectionQueryOK{}
}

// WithPayload adds the payload to the goals section query o k response
func (o *GoalsSectionQueryOK) WithPayload(payload []*models.Goal) *GoalsSectionQueryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the goals section query o k response
func (o *GoalsSectionQueryOK) SetPayload(payload []*models.Goal) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GoalsSectionQueryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Goal, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
