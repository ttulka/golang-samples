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

func recoverPanic(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Create a deferred function (which will always be run in the event of a panic as Go unwinds the stack).
    defer func() {
      // Use the builtin recover function to check if there has been a panic or not. If there has...
      if err := recover(); err != nil {
        // Set a "Connection: close" header on the response.
        w.Header().Set("Connection", "close")
        // Call the app.serverError helper method to return a 500 Internal Server response.
        http.Error(w, "Internal Server Error", 500)
      }
    }()
    next.ServeHTTP(w, r)
  })
}