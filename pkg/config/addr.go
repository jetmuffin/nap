package config

import (
	"net/url"
)

type Addr struct {
	addr *url.URL
}

func (a *Addr) UnmarshalText(text []byte) error {
	var err error
	a.addr, err = url.Parse(string(text))
	return err
}

func (a *Addr) Address() *url.URL {
	return a.addr
}
