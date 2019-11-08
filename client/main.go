package main

import (
	"bnywf"
	"net/http"
)

func sayhello(c bnywf.Context) {
	c.String(http.StatusOK, "hello world")
}

func sayok(c bnywf.Context) {
	c.String(http.StatusAccepted, "ok")
}

func main() {

	b:=bnywf.NEW()

	b.GET("/hello",sayhello)
	b.GET("/ok",sayok)

	b.Start(":8080")
}
