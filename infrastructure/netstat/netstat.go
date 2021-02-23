package netstat

import "github.com/drael/GOnetstat"

type Netstat struct {
}

func NewNetstat() *Netstat {
	return &Netstat{}
}

func (n *Netstat) Tcp() []GOnetstat.Process {
	return GOnetstat.Tcp()
}

func (n *Netstat) Udp() []GOnetstat.Process {
	return GOnetstat.Udp()
}

func (n *Netstat) Tcp6() []GOnetstat.Process {
	return GOnetstat.Tcp6()
}

func (n *Netstat) Udp6() []GOnetstat.Process {
	return GOnetstat.Udp6()
}
