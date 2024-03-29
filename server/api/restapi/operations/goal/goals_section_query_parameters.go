// Code generated by go-swagger; DO NOT EDIT.

package goal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGoalsSectionQueryParams creates a new GoalsSectionQueryParams object
// no default values defined in spec.
func NewGoalsSectionQueryParams() GoalsSectionQueryParams {

	return GoalsSectionQueryParams{}
}

// GoalsSectionQueryParams contains all the bound params for the goals section query operation
// typically these are obtained from a http.Request
//
// swagger:parameters goals-section-query
type GoalsSectionQueryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	  Collection Format: csv
	*/
	Sections []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGoalsSectionQueryParams() beforehand.
func (o *GoalsSectionQueryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qSections, qhkSections, _ := qs.GetOK("sections")
	if err := o.bindSections(qSections, qhkSections, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSections binds and validates array parameter Sections from query.
//
// Arrays are parsed according to CollectionFormat: "csv" (defaults to "csv" when empty).
func (o *GoalsSectionQueryParams) bindSections(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sections", "query")
	}

	var qvSections string
	if len(rawData) > 0 {
		qvSections = rawData[len(rawData)-1]
	}

	// CollectionFormat: csv
	sectionsIC := swag.SplitByFormat(qvSections, "csv")

	if len(sectionsIC) == 0 {
		return errors.Required("sections", "query")
	}

	var sectionsIR []string
	for i, sectionsIV := range sectionsIC {
		sectionsI := sectionsIV

		if err := validate.Enum(fmt.Sprintf("%s.%v", "sections", i), "query", sectionsI, []interface{}{"productivity", "todo", "enhancement", "proficiency"}); err != nil {
			return err
		}

		sectionsIR = append(sectionsIR, sectionsI)
	}

	o.Sections = sectionsIR

	return nil
}
