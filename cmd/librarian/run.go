package main

import "github.com/gari8/librarian/infra"

type IEdiNet interface {
	SearchCorporations(c infra.CorporationQuery)
	SearchDocuments(d infra.DocumentQuery)
}

type Set struct {
	IEdiNet
}

func NewSet(e IEdiNet) *Set {
	return &Set{e}
}

func Setup() *Set {
	h := infra.NewHttp()
	e := infra.NewEdiNet(h)
	return NewSet(e)
}

func (s *Set) Run() {
	//s.IEdiNet.SearchCorporations()
	s.IEdiNet.SearchDocuments(infra.DocumentQuery{})
}
