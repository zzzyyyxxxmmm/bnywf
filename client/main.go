package main

import (
	"bnywf"
)

func sayhello(c bnywf.Context) {
	c.String("hello world")
}

func sayok(c bnywf.Context) {
	c.String("ok")
}

func main() {

	b:=bnywf.NEW()

	b.GET("/hello",sayhello)
	b.GET("/ok",sayok)

	b.Start(":8080")
}
