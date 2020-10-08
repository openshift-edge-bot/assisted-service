// Code generated by go-swagger; DO NOT EDIT.

package managed_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListManagedDomainsHandlerFunc turns a function with the right signature into a list managed domains handler
type ListManagedDomainsHandlerFunc func(ListManagedDomainsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListManagedDomainsHandlerFunc) Handle(params ListManagedDomainsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListManagedDomainsHandler interface for that can handle valid list managed domains params
type ListManagedDomainsHandler interface {
	Handle(ListManagedDomainsParams, interface{}) middleware.Responder
}

// NewListManagedDomains creates a new http.Handler for the list managed domains operation
func NewListManagedDomains(ctx *middleware.Context, handler ListManagedDomainsHandler) *ListManagedDomains {
	return &ListManagedDomains{Context: ctx, Handler: handler}
}

/*ListManagedDomains swagger:route GET /domains managed_domains listManagedDomains

List of managed DNS domains.

*/
type ListManagedDomains struct {
	Context *middleware.Context
	Handler ListManagedDomainsHandler
}

func (o *ListManagedDomains) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListManagedDomainsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
