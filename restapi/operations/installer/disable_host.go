// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DisableHostHandlerFunc turns a function with the right signature into a disable host handler
type DisableHostHandlerFunc func(DisableHostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DisableHostHandlerFunc) Handle(params DisableHostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DisableHostHandler interface for that can handle valid disable host params
type DisableHostHandler interface {
	Handle(DisableHostParams, interface{}) middleware.Responder
}

// NewDisableHost creates a new http.Handler for the disable host operation
func NewDisableHost(ctx *middleware.Context, handler DisableHostHandler) *DisableHost {
	return &DisableHost{Context: ctx, Handler: handler}
}

/* DisableHost swagger:route DELETE /v1/clusters/{cluster_id}/hosts/{host_id}/actions/enable installer disableHost

Disables a host for inclusion in the cluster.

*/
type DisableHost struct {
	Context *middleware.Context
	Handler DisableHostHandler
}

func (o *DisableHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDisableHostParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
