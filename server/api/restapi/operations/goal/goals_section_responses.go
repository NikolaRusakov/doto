// Code generated by go-swagger; DO NOT EDIT.

package goal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rusakov/doto/server/api/models"
)

// GoalsSectionOKCode is the HTTP code returned for type GoalsSectionOK
const GoalsSectionOKCode int = 200

/*GoalsSectionOK OK

swagger:response goalsSectionOK
*/
type GoalsSectionOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Goal `json:"body,omitempty"`
}

// NewGoalsSectionOK creates GoalsSectionOK with default headers values
func NewGoalsSectionOK() *GoalsSectionOK {

	return &GoalsSectionOK{}
}

// WithPayload adds the payload to the goals section o k response
func (o *GoalsSectionOK) WithPayload(payload []*models.Goal) *GoalsSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the goals section o k response
func (o *GoalsSectionOK) SetPayload(payload []*models.Goal) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GoalsSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Goal, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}