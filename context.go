package bnywf

import (
	"net/http"
	"strconv"
)

type context struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
}

type Context interface {
	String(code int ,s string)
}

func (c *context) String(code int , s string){
	codestring:=strconv.Itoa(code)
	c.ResponseWriter.Header().Set("Status Code",codestring)
	c.ResponseWriter.Write([]byte(s))
}
