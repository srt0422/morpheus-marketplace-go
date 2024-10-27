// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

// ProxyService contains methods and other services that help with interacting with
// the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProxyService] method instead.
type ProxyService struct {
	Options  []option.RequestOption
	Sessions *ProxySessionService
}

// NewProxyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProxyService(opts ...option.RequestOption) (r *ProxyService) {
	r = &ProxyService{}
	r.Options = opts
	r.Sessions = NewProxySessionService(opts...)
	return
}
