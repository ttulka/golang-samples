package main

import (
  "net/http"
  "log"
)

func secureHeaders(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-XSS-Protection", "1; mode=block")
    w.Header().Set("X-Frame-Options", "deny")
    next.ServeHTTP(w, r)
  })
}

func logRequest(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
    next.ServeHTTP(w, r)
  })
}