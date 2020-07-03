// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"tides-server/pkg/models"
)

// AddResourceHandlerFunc turns a function with the right signature into a add resource handler
type AddResourceHandlerFunc func(AddResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddResourceHandlerFunc) Handle(params AddResourceParams) middleware.Responder {
	return fn(params)
}

// AddResourceHandler interface for that can handle valid add resource params
type AddResourceHandler interface {
	Handle(AddResourceParams) middleware.Responder
}

// NewAddResource creates a new http.Handler for the add resource operation
func NewAddResource(ctx *middleware.Context, handler AddResourceHandler) *AddResource {
	return &AddResource{Context: ctx, Handler: handler}
}

/*AddResource swagger:route POST /resource/add resource addResource

AddResource add resource API

*/
type AddResource struct {
	Context *middleware.Context
	Handler AddResourceHandler
}

func (o *AddResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddResourceBody add resource body
//
// swagger:model AddResourceBody
type AddResourceBody struct {

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// datacenters
	Datacenters string `json:"datacenters,omitempty"`

	// host address
	HostAddress string `json:"hostAddress,omitempty"`

	// is resource pool
	IsResourcePool bool `json:"isResourcePool,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// policy
	Policy int64 `json:"policy,omitempty"`

	// resources
	Resources []string `json:"resources"`

	// username
	Username string `json:"username,omitempty"`

	// vmtype
	// Enum: [vsphere kvm hyper-v]
	Vmtype string `json:"vmtype,omitempty"`
}

// Validate validates this add resource body
func (o *AddResourceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVmtype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addResourceBodyTypeVmtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vsphere","kvm","hyper-v"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addResourceBodyTypeVmtypePropEnum = append(addResourceBodyTypeVmtypePropEnum, v)
	}
}

const (

	// AddResourceBodyVmtypeVsphere captures enum value "vsphere"
	AddResourceBodyVmtypeVsphere string = "vsphere"

	// AddResourceBodyVmtypeKvm captures enum value "kvm"
	AddResourceBodyVmtypeKvm string = "kvm"

	// AddResourceBodyVmtypeHyperv captures enum value "hyper-v"
	AddResourceBodyVmtypeHyperv string = "hyper-v"
)

// prop value enum
func (o *AddResourceBody) validateVmtypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addResourceBodyTypeVmtypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddResourceBody) validateVmtype(formats strfmt.Registry) error {

	if swag.IsZero(o.Vmtype) { // not required
		return nil
	}

	// value enum
	if err := o.validateVmtypeEnum("reqBody"+"."+"vmtype", "body", o.Vmtype); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddResourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddResourceBody) UnmarshalBinary(b []byte) error {
	var res AddResourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddResourceNotFoundBody add resource not found body
//
// swagger:model AddResourceNotFoundBody
type AddResourceNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add resource not found body
func (o *AddResourceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddResourceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddResourceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AddResourceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddResourceOKBody add resource o k body
//
// swagger:model AddResourceOKBody
type AddResourceOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// results
	Results []*models.ResourceAddItem `json:"results"`
}

// Validate validates this add resource o k body
func (o *AddResourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddResourceOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addResourceOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddResourceOKBody) UnmarshalBinary(b []byte) error {
	var res AddResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}