// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2UpdateHostInstallerArgsHandlerFunc turns a function with the right signature into a v2 update host installer args handler
type V2UpdateHostInstallerArgsHandlerFunc func(V2UpdateHostInstallerArgsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2UpdateHostInstallerArgsHandlerFunc) Handle(params V2UpdateHostInstallerArgsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2UpdateHostInstallerArgsHandler interface for that can handle valid v2 update host installer args params
type V2UpdateHostInstallerArgsHandler interface {
	Handle(V2UpdateHostInstallerArgsParams, interface{}) middleware.Responder
}

// NewV2UpdateHostInstallerArgs creates a new http.Handler for the v2 update host installer args operation
func NewV2UpdateHostInstallerArgs(ctx *middleware.Context, handler V2UpdateHostInstallerArgsHandler) *V2UpdateHostInstallerArgs {
	return &V2UpdateHostInstallerArgs{Context: ctx, Handler: handler}
}

/* V2UpdateHostInstallerArgs swagger:route PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args installer v2UpdateHostInstallerArgs

Updates a host's installer arguments.

*/
type V2UpdateHostInstallerArgs struct {
	Context *middleware.Context
	Handler V2UpdateHostInstallerArgsHandler
}

func (o *V2UpdateHostInstallerArgs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2UpdateHostInstallerArgsParams()
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
