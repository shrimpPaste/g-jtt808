package v2013

import "strings"

type M0102 struct {
	RawToken string `jtt808:"0,string"`
}

func (m M0102) Token() string {
	return strings.Trim(m.RawToken, "\u0000")
}

func (m *M0102) SetToken(t string) *M0102 {
	m.RawToken = t
	return m
}
