# 我们的目标
* Restful

### 解析path中的param
第一个问题就是path是对应哪个handler, 我们现在的思路还是通过下面的判断
```go
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
很容易看出来这个pattern是写死的, 所以在这里要分析pattern, 假设path是```localhost/name/:name```,那么可能的选择就是

```
localhost/name/wang
localhost/name/liu
localhost/name/zhao
```

按照最原始的方法, 我们是需要对每个选择都建立一个HandleFunc的, 但是显然是无法做到的, 因为我们也不知道传过来的参数是什么, 看一看HandleFunction的源码, 最终:


ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler 
for the pattern that most closely matches the URL.

Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash). Longer patterns take precedence over shorter ones, so that if there are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will be called for paths beginning "/images/thumbnails/" and the former will receive requests for any other paths in the "/images/" subtree.

Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".

If a subtree has been registered and a request is received naming the subtree root without its trailing slash, ServeMux redirects that request to the subtree root (adding the trailing slash). This behavior can be overridden with a separate registration for the path without the trailing slash. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately.

可以看到, 我们注册的时候如果在path后面加入/, 就会将以这个path开头, 并且没有注册的path指向这个handler, 所以

我们需要注册localhost/name/, 然后在这里分析并且拿到名字


1. Query Parameters

2. Path Parameters





