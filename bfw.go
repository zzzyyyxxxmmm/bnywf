package bnywf

import (
	"net/http"
)

type (
	bnywf struct {
	}

	HandlerFunc func(Context)
)

func NEW() *bnywf {
	return &bnywf{}
}

func (b *bnywf) Start(addr string) {
	http.ListenAndServe(addr, nil)
}

func (b *bnywf) GET(pattern string, h HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c:=&context{
			ResponseWriter:w,
			Request:r,
		}
		h(c)
	})
}
