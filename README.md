# 我们的目标
* Restful

# 实现最基本的一个server
```go
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
```

# 实现Restful

## 能够分别处理Get, POST请求
```go
package main

import (
	"bnywf"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {

	b:=bnywf.NEW()

	b.GET("/bar",sayhello)

	b.Start(":8080")
}
```


```go
package bnywf

import "net/http"

type bnywf struct {
}

func NEW() *bnywf {
	return &bnywf{}
}

func (b *bnywf)Start(addr string) {
	http.ListenAndServe(addr,nil)
}

func (b *bnywf) GET(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}
```


## 封装一个context
Context里包含了需要的所有读写方法, 所以context应该持有request和response

```go
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
```

```go
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
```

```go
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
```


