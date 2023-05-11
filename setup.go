package multipass

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("multipass", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'multipass'
	if c.NextArg() {
		return plugin.Error("multipass", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Multipass{Next: next} // Set the Next field, so the plugin chaining works.
	})

	return nil
}
