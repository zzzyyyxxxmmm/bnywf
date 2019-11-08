package bnywf

import (
	"net/http"
	"strconv"
	"strings"
)

type context struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
	Params string
}

type Context interface {
	String(code int ,s string)
	Param(s string) string
}

func (c *context) String(code int , s string){
	codestring:=strconv.Itoa(code)
	c.ResponseWriter.Header().Set("Status Code",codestring)
	c.ResponseWriter.Write([]byte(s))
}

func (c *context) Param(s string) string{

	//如果没有有param, 那直接报错就可以了
	if len(c.Params)==0{
	}

	runes := []rune(c.Request.URL.Path)
	// ... Convert back into a string from rune slice.
	subs := string(runes[strings.LastIndex(c.Request.URL.Path,"/")+1:])

	return subs
}
