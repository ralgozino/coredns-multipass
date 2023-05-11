package multipass

import (
	"context"
	"net"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

type Multipass struct {
	Next plugin.Handler
}

func (mp Multipass) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	req := request.Request{W: w, Req: r}
	var log = clog.NewWithPlugin("multipass")
	vms, err := vmList()
	if err != nil {
		log.Errorf("Error while getting VM list: %s", err)
	}
	labels := dns.SplitDomainName(req.Name())
	if req.QType() == dns.TypeA && labels[0] != "" && len(vms[labels[0]]) > 0 {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true

		a := new(dns.A)
		a.Hdr = dns.RR_Header{Name: req.Name(), Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 3600}
		a.A = net.ParseIP(vms[labels[0]][0])
		m.Answer = []dns.RR{a}

		w.WriteMsg(m)
		return dns.RcodeSuccess, nil
	}

	// Call the next plugin in the chain
	return plugin.NextOrFailure(mp.Name(), mp.Next, ctx, w, r)
}

func (mp Multipass) Name() string { return "multipass" }
