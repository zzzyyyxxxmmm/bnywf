package bnywf

import "testing"

var b=context{}

func TestContext_Param(t *testing.T) {
	b.Param("127.0.0.1/:name")
}
