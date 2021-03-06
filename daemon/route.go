package daemon

import (
	"errors"
	"github.com/zhgwenming/vrouter/Godeps/_workspace/src/github.com/docker/libcontainer/netlink"
	"net"
	"strings"
)

type Route struct {
	target string
	gw     string
}

func NewRoute(target, gw string) *Route {
	return &Route{target: target, gw: gw}
}

func ParseRoute(str string) (*Route, error) {
	r := strings.Split(str, ":")
	if len(r) != 2 {
		return nil, errors.New("Wrong route format")
	}
	return &Route{target: r[0], gw: r[1]}, nil
}

func (r *Route) AddRoute(iface *net.Interface) error {
	return netlink.AddRoute(r.target, "", r.gw, iface.Name)
}

func (r *Route) String() string {
	return r.target + ":" + r.gw
}
