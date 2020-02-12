package main

import (
  "net/http"
  "net/http/httptest"
	"net/http/httputil"
  "net/url"
  "os"
  "log"
  "fmt"
  "strings"
)

func main() {
  if len(os.Args[1:]) <= 0 {
    log.Fatalln("Usage: <host>")
  }
  host, err := url.Parse(os.Args[1:][0])
  if err != nil {
    log.Fatalln("Wrong host")
  }
  
  frontendProxy := httptest.NewServer(NewSingleHostReverseProxy(host))
	defer frontendProxy.Close()
  
  fmt.Printf("Proxy URL: %v\n", frontendProxy.URL)
  
  for {}
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
    
    req.Host = target.Host
    
    fmt.Printf("Request headers: %v\n", req.Header)
    fmt.Printf("Request method: %v\n", req.Method)
    fmt.Printf("Request Url: %v\n", req.URL)
    fmt.Printf("Request host: %v\n", req.Host)
    fmt.Printf("Request remote addr: %v\n", req.RemoteAddr)
    fmt.Printf("Request URI: %v\n", req.RequestURI)
	}
	return &httputil.ReverseProxy{Director: director}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}