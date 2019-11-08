package bnywf

import "net/http"

type context struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
}

type Context interface {
	String(s string)
}

func (c *context) String(s string){
	c.ResponseWriter.Write([]byte(s))
}
