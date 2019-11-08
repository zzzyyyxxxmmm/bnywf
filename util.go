package bnywf

import "strings"

func isRealPath(url string) []string{
	params := strings.Split( url,"/")
	realPath:=""
	s:=make([]string,2)
	for _, p:= range params{

		if len(p)==0{
			continue
		}

		if p[0]==':'{
			s[0]=realPath+"/"
			s[1]=p
			return s
		}
		realPath+="/"+p
	}
	s[0]=realPath
	s[1]=""
	return s
}
