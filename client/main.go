package main

import (
	"bnywf"
	"net/http"
	"fmt"
)

func sayhello(c bnywf.Context) {
	c.String(http.StatusOK, "hello world")
}

func sayok(c bnywf.Context) {
	c.String(http.StatusAccepted, "ok")
}

func sayName(c bnywf.Context)  {
	fmt.Println("sayname")
	name:=c.Param("name")
	c.String(http.StatusOK, "hello"+name)
}

func main() {

	b:=bnywf.NEW()

	b.GET("/hello",sayhello)
	b.GET("/ok",sayok)
	b.GET("/name/:name",sayName)

	b.Start(":8080")
}
