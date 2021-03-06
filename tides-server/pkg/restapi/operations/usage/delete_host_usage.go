// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteHostUsageHandlerFunc turns a function with the right signature into a delete host usage handler
type DeleteHostUsageHandlerFunc func(DeleteHostUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHostUsageHandlerFunc) Handle(params DeleteHostUsageParams) middleware.Responder {
	return fn(params)
}

// DeleteHostUsageHandler interface for that can handle valid delete host usage params
type DeleteHostUsageHandler interface {
	Handle(DeleteHostUsageParams) middleware.Responder
}

// NewDeleteHostUsage creates a new http.Handler for the delete host usage operation
func NewDeleteHostUsage(ctx *middleware.Context, handler DeleteHostUsageHandler) *DeleteHostUsage {
	return &DeleteHostUsage{Context: ctx, Handler: handler}
}

/*DeleteHostUsage swagger:route DELETE /usage/deleteHost usage deleteHostUsage

delete datacenter usage info

*/
type DeleteHostUsage struct {
	Context *middleware.Context
	Handler DeleteHostUsageHandler
}

func (o *DeleteHostUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteHostUsageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteHostUsageBody delete host usage body
//
// swagger:model DeleteHostUsageBody
type DeleteHostUsageBody struct {

	// datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// host address
	HostAddress string `json:"hostAddress,omitempty"`
}

// Validate validates this delete host usage body
func (o *DeleteHostUsageBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostUsageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostUsageBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostUsageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteHostUsageOKBody delete host usage o k body
//
// swagger:model DeleteHostUsageOKBody
type DeleteHostUsageOKBody struct {

	// message
	// Enum: [success]
	Message string `json:"message,omitempty"`
}

// Validate validates this delete host usage o k body
func (o *DeleteHostUsageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteHostUsageOKBodyTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteHostUsageOKBodyTypeMessagePropEnum = append(deleteHostUsageOKBodyTypeMessagePropEnum, v)
	}
}

const (

	// DeleteHostUsageOKBodyMessageSuccess captures enum value "success"
	DeleteHostUsageOKBodyMessageSuccess string = "success"
)

// prop value enum
func (o *DeleteHostUsageOKBody) validateMessageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteHostUsageOKBodyTypeMessagePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteHostUsageOKBody) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(o.Message) { // not required
		return nil
	}

	// value enum
	if err := o.validateMessageEnum("deleteHostUsageOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostUsageOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
