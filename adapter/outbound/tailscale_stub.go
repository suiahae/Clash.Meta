//go:build !with_gvisor || no_tailscale

package outbound

import (
	"context"
	"fmt"
	"net"

	"github.com/metacubex/tailscale/ipn/ipnstate"
)

type Tailscale struct {
	*Base
}

type TailscaleOption struct {
	BasicOption
	Name       string `proxy:"name"`
	Hostname   string `proxy:"hostname,omitempty"`
	AuthKey    string `proxy:"auth-key,omitempty"`
	ControlURL string `proxy:"control-url,omitempty"`
	StateDir   string `proxy:"state-dir,omitempty"`
	Ephemeral  bool   `proxy:"ephemeral,omitempty"`
	UDP        bool   `proxy:"udp,omitempty"`

	AcceptRoutes           *bool  `proxy:"accept-routes,omitempty"`
	ExitNode               string `proxy:"exit-node,omitempty"`
	ExitNodeAllowLANAccess *bool  `proxy:"exit-node-allow-lan-access,omitempty"`
}

func NewTailscale(option TailscaleOption) (*Tailscale, error) {
	return nil, fmt.Errorf("tailscale support is disabled by \"no_tailscale\" build tag or not include \"with_gvisor\" build tag")
}

func (t *Tailscale) TailscaleLastError() string {
	return ""
}

func (t *Tailscale) TailscaleStatus(ctx context.Context) (*ipnstate.Status, error) {
	return nil, errTailscaleUnsupported
}

func (t *Tailscale) TailscaleStartLogin(ctx context.Context) error {
	return errTailscaleUnsupported
}

func (t *Tailscale) TailscaleLogout(ctx context.Context) error {
	return errTailscaleUnsupported
}

func (t *Tailscale) TailscaleDial(ctx context.Context, network, address string) (net.Conn, error) {
	return nil, errTailscaleUnsupported
}

var errTailscaleUnsupported = fmt.Errorf("tailscale support is disabled by \"no_tailscale\" build tag or not include \"with_gvisor\" build tag")
